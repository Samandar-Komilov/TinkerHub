CC = gcc
INCLUDE = -Isrc/include
SRC = src/*.c
STRICT_FLAGS = -Wall -Wextra -Wpedantic -Werror
JSON_SRC = src/json/*.c
MATH_LINKER = -lm
BIN = bin

run:
	$(CC) $(STRICT_FLAGS) $(INCLUDE) $(SRC) $(JSON_SRC) -o $(BIN)/trlog $(MATH_LINKER)
	./$(BIN)/trlog

clean:
	rm $(BIN)/trlog