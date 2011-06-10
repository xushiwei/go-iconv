include $(GOROOT)/src/Make.inc

TARG=xushiwei.com/iconv

CGOFILES=iconv.go
ifneq ($(GOOS),linux)
CGO_LDFLAGS=-liconv
endif

include $(GOROOT)/src/Make.pkg

