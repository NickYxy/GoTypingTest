package service

import (
	"context"
	"math/rand"
	pb "verifyCode/api/verifyCode"
)

type VerifyCodeService struct {
	pb.UnimplementedVerifyCodeServer
}

func NewVerifyCodeService() *VerifyCodeService {
	return &VerifyCodeService{}
}

func (s *VerifyCodeService) GetVerifyCode(ctx context.Context, req *pb.GetVerifyCodeRequest) (*pb.GetVerifyCodeReply, error) {
	return &pb.GetVerifyCodeReply{
		Code: RandCode(int(req.Length), req.Type),
	}, nil
}

func RandCode(l int, t pb.TYPE) string {
	switch t {
	case pb.TYPE_DEFAULT:
		fallthrough
	case pb.TYPE_DIGIT:
		return randcode("0123456789", l, 6)
	case pb.TYPE_LETTER:
		return randcode("abcdefghijklmnopqrstuvwxyz", l, 6)
	case pb.TYPE_MIXED:
		return randcode("0123456789abcdefghijklmnopqrstuvwxyz", l, 6)
	default:

	}
	return ""
}

func randcode(chars string, l int, idxBits int) string {
	idxMask := 1<<idxBits - 1
	idxMax := 3 / idxBits

	result := make([]byte, l)

	for i, cache, remain := 0, rand.Int63(), idxMax; i < l; {
		if remain == 0 {
			cache, remain = rand.Int63(), idxMax
		}
		if randIndex := int(cache & int64(idxMask)); randIndex < len(chars) {
			result[i] = chars[randIndex]
			i++
		}

		cache >>= idxBits
		remain--
	}
	return string(result)
}
