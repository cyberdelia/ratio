package ratio

import (
	"bytes"
	"crypto/rand"
	"io"
	"io/ioutil"
	"strings"
	"testing"
	"time"
)

func TestWriter(t *testing.T) {
	buf := new(bytes.Buffer)
	rw := RateLimitedWriter(buf, 2, time.Millisecond)
	defer rw.Close()
	io.Copy(rw, strings.NewReader("aloha"))
	if buf.String() != "aloha" {
		t.Fatalf("'%s' don't match '%s'", buf.String(), "aloha")
	}
}

func TestReader(t *testing.T) {
	buf := new(bytes.Buffer)
	rr := RateLimitedReader(strings.NewReader("aloha"), 2, time.Millisecond)
	defer rr.Close()
	io.Copy(buf, rr)
	if buf.String() != "aloha" {
		t.Fatalf("'%s' don't match '%s'", buf.String(), "aloha")
	}
}

func BenchmarkWriter(b *testing.B) {
	b.StopTimer()
	buf := make([]byte, 2*MB)
	b.SetBytes(2 * MB)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rw := RateLimitedWriter(ioutil.Discard, 1*MB, time.Second)
		rw.Write(buf)
		rw.Close()
	}
}

func BenchmarkReader(b *testing.B) {
	b.StopTimer()
	buf := make([]byte, 2*MB)
	b.SetBytes(2 * MB)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rw := RateLimitedReader(rand.Reader, 1*MB, time.Second)
		rw.Read(buf)
		rw.Close()
	}
}
