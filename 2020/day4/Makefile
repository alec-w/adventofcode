CC=g++
CFLAGS=-Wall
LDFLAGS=-Wall

.PHONY: all
all: part1 part2

.PHONY: clean
clean:
	$(RM) *.o part1 part2

part1: part1.o
	$(CC) -o $@ $< $(LDFLAGS)

part2: part2.o
	$(CC) -o $@ $< $(LDFLAGS)

%.o: %.cpp
	$(CC) -c -o $@ $< $(CFLAGS)

