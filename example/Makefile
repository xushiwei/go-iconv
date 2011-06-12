include $(GOROOT)/src/Make.inc

all:
	$(GC) convstr.go; $(LD) -o convstr convstr.$(O)
	$(GC) reader.go; $(LD) -o reader reader.$(O)
	$(GC) writer.go; $(LD) -o writer writer.$(O)

clean:
	rm -f convstr reader writer *.$(O) *.out

