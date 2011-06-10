//
// iconv.go
//
package iconv

// #include <iconv.h>
// #include <errno.h>
import "C"

import (
	"os"
	"io"
	"unsafe"
	"bytes"
)

var EILSEQ = os.Errno(int(C.EILSEQ))
var E2BIG = os.Errno(int(C.E2BIG))

const DefaultBufSize = 4096

type Iconv struct {
	inbuf []byte
	outbuf []byte
	pointer C.iconv_t
	output io.Writer
	n int // inbuf[0:n] is valid
	autoSync bool
}

func Open(tocode string, fromcode string) (*Iconv, os.Error) {
	return OpenWith(tocode, fromcode, os.Stdout, 0, true)
}

func OpenWith(tocode string, fromcode string, output io.Writer, bufSize int, autoSync bool) (*Iconv, os.Error) {
	ret, err := C.iconv_open(C.CString(tocode), C.CString(fromcode))
	if err != nil {
		return nil, err
	}
	if bufSize < 16 { bufSize = DefaultBufSize }
	var inbuf []byte
	if !autoSync {
		inbuf = make([]byte, bufSize)
	}
	outbuf := make([]byte, bufSize)
	return &Iconv{inbuf, outbuf, ret, output, 0, autoSync}, nil
}

func (cd *Iconv) Close() os.Error {
	err1 := cd.Sync()
	_, err := C.iconv_close(cd.pointer)
	if err1 != nil { return err1 }
	return err
}

func (cd *Iconv) Output(w io.Writer) {
	cd.Sync()
	cd.output = w
}

func (cd *Iconv) AutoSync(b bool) {
	cd.autoSync = b
	if !b && cd.inbuf == nil {
		cd.inbuf = make([]byte, len(cd.outbuf))
	}
}

func (cd *Iconv) Sync() os.Error {

	if cd.n == 0 { return nil }
	
	inleft, err := iconv(cd.pointer, cd.output, cd.inbuf, cd.n, cd.outbuf)
	if inleft > 0 {
		copy(cd.inbuf, cd.inbuf[cd.n-inleft:cd.n])
	}
	cd.n = inleft
	return err
}

func (cd *Iconv) Write(b []byte) (n int, err os.Error) {

	if cd.autoSync {
		var inleft int
		inleft, err = iconv(cd.pointer, cd.output, b, len(b), cd.outbuf)
		n = len(b) - inleft
		return
	}
	for {
		n1 := copy(cd.inbuf[cd.n:], b)
		if n1 == 0 {
			if len(b) > 0 { return n, EILSEQ }
			break
		}
		cd.n += n1
		n += n1
		if cd.n == len(cd.inbuf) {
			err = cd.Sync()
			if err != nil && err != os.EINVAL { return }
		}
		if len(b) == n1 { break }
		b = b[n1:]
	}
	return n, nil
}

func (cd *Iconv) WriteString(b string) (n int, err os.Error) {

	if cd.autoSync {
		var inleft int
		inleft, err = iconv(cd.pointer, cd.output, []byte(b), len(b), cd.outbuf)
		n = len(b) - inleft
		return
	}
	for {
		n1 := copy(cd.inbuf[cd.n:], b)
		if n1 == 0 {
			if len(b) > 0 { return n, EILSEQ }
			break
		}
		cd.n += n1
		n += n1
		if cd.n == len(cd.inbuf) {
			err = cd.Sync()
			if err != nil && err != os.EINVAL { return }
		}
		if len(b) == n1 { break }
		b = b[n1:]
	}
	return n, nil
}

func (cd *Iconv) Conv(b []byte) (out []byte, err os.Error) {

	inbytes := C.size_t(len(b))
	inptr := &b[0]

	outbuf := cd.outbuf
	outbytes := C.size_t(len(outbuf))
	outptr := &outbuf[0]
	_, err = C.iconv(cd.pointer,
		(**C.char)(unsafe.Pointer(&inptr)), &inbytes,
		(**C.char)(unsafe.Pointer(&outptr)), &outbytes)

	out = outbuf[:len(outbuf)-int(outbytes)]
	if err == nil && err != E2BIG { return }

	w := bytes.NewBuffer(nil)
	w.Write(out)

	n := int(inbytes)
	_, err = iconv(cd.pointer, w, b[len(b)-n:], n, outbuf)
	out = w.Bytes()
	return
}

func (cd *Iconv) ConvString(s string) string {
	s1, _ := cd.Conv([]byte(s))
	return string(s1)
}

func iconv(cd C.iconv_t, w io.Writer, inbuf []byte, in int, outbuf []byte) (inleft int, err os.Error) {

	inbytes := C.size_t(in)
	inptr := &inbuf[0]

	for inbytes > 0 {
		outbytes := C.size_t(len(outbuf))
		outptr := &outbuf[0]
		_, err = C.iconv(cd,
			(**C.char)(unsafe.Pointer(&inptr)), &inbytes,
			(**C.char)(unsafe.Pointer(&outptr)), &outbytes)
		w.Write(outbuf[:len(outbuf)-int(outbytes)])
		if err != nil && err != E2BIG {
			return int(inbytes), err
		}
	}

	return 0, nil
}

