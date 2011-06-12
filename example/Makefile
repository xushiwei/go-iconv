include $(GOROOT)/src/Make.inc

all:
	$(GC) convstr.go; $(LD) -o convstr.exe convstr.$(O)
	$(GC) reader.go; $(LD) -o reader.exe reader.$(O)
	$(GC) writer.go; $(LD) -o writer.exe writer.$(O)

clean:
	rm -f convstr reader writer *.$(O) *.out

