package processor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockReader struct {
	closed bool
}

func (m *mockReader) Read(p []byte) (n int, err error) { return 0, nil }
func (m *mockReader) Close() error {
	m.closed = true
	return nil
}

const maxReadBytes = 1024 * 30

func TestLogWriterEasyLength(t *testing.T) {
	r := &mockReader{}

	w := NewLogWriter(r)
	n, err := w.Write(longBytes((maxReadBytes / 2) + 1))

	assert.Nil(t, err)
	assert.Equal(t, (maxReadBytes/2)+1, n)
	assert.Equal(t, (maxReadBytes/2)+1, len(w.Output()))
	assert.False(t, r.closed)
}

func TestLogWriterMaxLength(t *testing.T) {
	r := &mockReader{}

	w := NewLogWriter(r)
	n, err := w.Write(longBytes(maxReadBytes + 100))

	assert.Nil(t, err)
	assert.Equal(t, maxReadBytes+100, n)
	assert.Equal(t, maxReadBytes, len(w.Output()))
	assert.True(t, r.closed)
}

func longBytes(amount int) []byte {
	var bs = make([]byte, amount)
	for i := 0; i < amount; i++ {
		bs[i] = 'x'
	}
	return bs
}
