CC = gcc
SRC = src/*.c
JSON_SRC = src/json/*.c
MATH_LINKER = -lm
BIN = bin

run:
	$(CC) $(SRC) $(JSON_SRC) -o $(BIN)/trlog $(MATH_LINKER)
	./$(BIN)/trlog

clean:
	rm $(BIN)/trlog