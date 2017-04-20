package counterwriter

import "io"

type CounterWriter struct {
	writer io.Writer
	count  int64
}

func (w *CounterWriter) Write(p []byte) (n int, err error) {
	n, err = w.writer.Write(p)
	w.count += int64(n)
	return
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &CounterWriter{
		writer: w,
		count:  0,
	}
	return cw, &cw.count
}
