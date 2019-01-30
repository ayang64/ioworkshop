DOTS		:= $(wildcard *.dot)
SVGS		:= $(DOTS:.dot=.svg)
PNGS		:= $(DOTS:.dot=.png)
IMGDIR	:= ../../../static/postimages/advent-2018/interfaces-and-reflect/

.PHONY : clean

%.png : %.dot
	dot -Tpng -o$@ $<

%.svg : %.dot
	dot -Tsvg -o$@ $<

# all: $(SVGS) $(PNGS)
all: $(SVGS)

clean:
	rm -f $(SVGS) $(PNGS)
