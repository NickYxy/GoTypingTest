package wal

import (
	"errors"
	"fmt"
	lru "github.com/hashicorp/golang-lru/v2"
	"io"
	"os"
	"sync"
)

type ChunkType = byte
type SegmentID = uint32

const (
	ChunkTypeFull ChunkType = iota
	ChunkTypeFirst
	ChunkTypeMiddle
	ChunkTypeLast
)

var (
	ErrClosed     = errors.New("the segment file is closed")
	ErrInvalidCRC = errors.New("invalid crc, the data may be corrupted")
)

const (
	// 7 Bytes
	// Checksum Length Type
	//    4      2     1
	chunkHeaderSize = 7

	blockSize = 32 * KB

	fileModePerm = 0644
)

// Segment represents a single segment file in WAL.
// The segment file is append-only, and the data is written in blocks.
// Each block is 32 KB, and the data is written in chunks.
type segment struct {
	id                 SegmentID
	fd                 *os.File
	currentBlockNumber uint32
	currentBlockSize   uint32
	closed             bool
	cache              *lru.Cache[uint64, []byte]
	header             []byte
	blockPool          sync.Pool
}

// segmentReader is used to iterate all the data from the segment file.
// You can call Next to get the next chunk data,
// and io.EOF will be returned when there is no data.
type segmentReader struct {
	segment     *segment
	blockNumber uint32
	chunkOffset int64
}

// block and chunk header, saved in pool
type blockAndHeader struct {
	block  []byte
	header []byte
}

// ChunkPosition represents the position of a chunk in a segment file.
// Used to read the data from the segment file.
type ChunkPosition struct {
	SegmentId   SegmentID
	BlockNumber uint32
	ChunkOffset int64
	ChunkSize   uint32
}

func openSegmentFile(dirPath, extName string, id uint32, cache *lru.Cache[uint64, []byte]) (*segment, error) {
	fd, err := os.OpenFile(
		SegmentFileName(dirPath, extName, id),
		os.O_CREATE|os.O_RDWR|os.O_APPEND,
		fileModePerm,
	)

	if err != nil {
		return nil, err
	}

	// set the current block number and block size
	offset, err := fd.Seek(0, io.SeekEnd)
	if err != nil {
		panic(any(fmt.Errorf("seek to the end of segment file %d%s failed: %v", id, extName, err)))
	}

	return &segment{
		id:                 id,
		fd:                 fd,
		currentBlockNumber: uint32(offset / blockSize),
		currentBlockSize:   uint32(offset % blockSize),
		cache:              cache,
		header:             make([]byte, chunkHeaderSize),
		blockPool:          sync.Pool{New: newBlockAndHeader},
	}, nil
}

func newBlockAndHeader() interface{} {
	return &blockAndHeader{
		block:  make([]byte, blockSize),
		header: make([]byte, chunkHeaderSize),
	}
}

// NewReader creates a new segment reader.
// You can call Next to get the next chunk data,
// and io.EOF will be returned when there is no data.
func (seg *segment) NewReader() *segmentReader {
	return &segmentReader{
		segment:     seg,
		blockNumber: 0,
		chunkOffset: 0,
	}
}

// Sync flushes the segment file to disk
func (seg *segment) Sync() error {
	if seg.closed {
		return nil
	}
	return seg.fd.Sync()
}

// Remove removes the segment file.
func (seg *segment) Remove() error {
	if !seg.closed {
		seg.closed = true
		_ = seg.fd.Close()
	}
	return os.Remove(seg.fd.Name())
}

// Close closes the segment file.
func (seg *segment) Close() error {
	if seg.closed {
		return nil
	}
	seg.closed = true
	return seg.fd.Close()
}

// Size returns the size of the segment file.
func (seg *segment) Size() int64 {
	return int64(seg.currentBlockNumber*blockSize + seg.currentBlockSize)
}

// Write writes the data to the segment file.
// The data will be written in chunks, and the chunk has four types:
// chunkTypeFull, chunkTypeFirst, chunkTypeMiddle, chunkTypeLast.
//
// Each chunk has a header, and the header contains the length, type and checksum.
// And the payload of the chunk is the real data you want to write.
func (seg *segment) Write(data []byte) (*ChunkPosition, error) {
	if seg.closed {
		return nil, ErrClosed
	}

	// The left block space is not enough for a chunk header
	if seg.currentBlockSize+chunkHeaderSize >= blockSize {
		// padding if necessary
		if seg.currentBlockSize < blockSize {
			padding := make([]byte, blockSize-seg.currentBlockSize)
			if _, err := seg.fd.Write(padding); err != nil {
				return nil, err
			}
		}

		// A new block, clear the current block size.
		seg.currentBlockNumber += 1
		seg.currentBlockSize = 0
	}

	// the start position (for read operation)
	position := &ChunkPosition{
		SegmentId:   seg.id,
		BlockNumber: seg.currentBlockNumber,
		ChunkOffset: int64(seg.currentBlockSize),
	}

	dataSize := uint32(len(data))
	// The entire chunk can fit into the block
	if seg.currentBlockSize+dataSize+chunkHeaderSize <= blockSize {
		err := seg.WriteInternal(data, ChunkTypeFull)
		if err != nil {
			return nil, err
		}
		position.ChunkSize = dataSize + chunkHeaderSize
		return position, nil
	}

	// If the size of the data exceeds the size of the block,
	// the data should be written to the block in batches.
}