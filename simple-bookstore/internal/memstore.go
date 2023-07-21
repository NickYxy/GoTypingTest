package internal

import (
	mystore "simple-bookstore/store"
	factory "simple-bookstore/store/factory"
	"sync"
)

func init() {
	factory.Register("mem", &MemStore{
		books: make(map[string]*mystore.Book),
	})
}

type MemStore struct {
	sync.RWMutex
	books map[string]*mystore.Book
}

func (m *MemStore) Create(book *mystore.Book) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.books[book.Id]; ok {
		return mystore.ErrExit
	}

	newBook := *book
	m.books[book.Id] = &newBook

	return nil
}

func (m *MemStore) Update(book *mystore.Book) error {
	m.Lock()
	defer m.Unlock()

	oldBook, ok := m.books[book.Id]
	if !ok {
		return mystore.ErrNotFound
	}

	newBook := *oldBook
	if book.Name != "" {
		newBook.Name = book.Name
	}

	if book.Authors != nil {
		newBook.Authors = book.Authors
	}

	if book.Press != "" {
		newBook.Press = book.Press
	}

	m.books[book.Id] = &newBook

	return nil
}

func (m *MemStore) Get(id string) (mystore.Book, error) {
	m.RLock()
	defer m.RUnlock()

	t, ok := m.books[id]
	if ok {
		return *t, nil
	}

	return mystore.Book{}, mystore.ErrNotFound
}

func (m *MemStore) GetAll() ([]mystore.Book, error) {
	m.RLock()
	defer m.RUnlock()

	allBooks := make([]mystore.Book, 0, len(m.books))

	for _, book := range m.books {
		allBooks = append(allBooks, *book)
	}

	return allBooks, nil
}

func (m *MemStore) Delete(id string) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.books[id]; !ok {
		return mystore.ErrNotFound
	}

	delete(m.books, id)
	return nil
}
