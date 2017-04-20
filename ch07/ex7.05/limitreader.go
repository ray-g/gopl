package limitreader

import "io"

// https://golang.org/src/io/io.go
type LimitedReader struct {
	reader io.Reader
	limit  int64
}

func (r *LimitedReader) Read(p []byte) (n int, err error) {
	if r.limit <= 0 {
		return 0, io.EOF
	}

	if int64(len(p)) > r.limit {
		p = p[0:r.limit]
	}

	n, err = r.reader.Read(p)
	r.limit -= int64(n)
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r, n}
}
