#include <stdlib.h>

#include "chunk.h"
#include "memory.h"
#include "vm.h"

void initChunk(Chunk* chunk) {
  chunk->count = 0;
  chunk->capacity = 0;
  chunk->code = NULL;
  chunk->lines = NULL;
  initValueArray(&chunk->constants);
}

void freeChunk(Chunk* chunk) {
  FREE_ARRAY(uint8_t, chunk->code, chunk->capacity);
  FREE_ARRAY(LineRun, chunk->lines, chunk->capacity);
  freeValueArray(&chunk->constants);
  initChunk(chunk);
}

void writeChunk(Chunk* chunk, uint8_t byte, int line) {
  if (chunk->capacity < chunk->count + 1) {
    int oldCapacity = chunk->capacity;
    chunk->capacity = GROW_CAPACITY(oldCapacity);
    chunk->code = GROW_ARRAY(uint8_t, chunk->code, oldCapacity, chunk->capacity);
    chunk->lines = GROW_ARRAY(LineRun, chunk->lines, oldCapacity, chunk->capacity);
  } 

  chunk->code[chunk->count] = byte;
  chunk->count++;

  if (chunk->count == 1 || chunk->lines[chunk->count - 2].line != line) {
    if (chunk->count > 1 && chunk->lines[chunk->count - 2].line == line) {
      chunk->lines[chunk->count - 2].count++;
    } else {
      LineRun run;
      run.line = line;
      run.count = 1;
      chunk->lines[chunk->count - 1] = run;
    }
  } else {
    chunk->lines[chunk->count - 2].count++;
  }
}

int addConstant(Chunk* chunk, Value value) {
  push(value);
  writeValueArray(&chunk->constants, value);
  pop();
  return chunk->constants.count - 1;
}

int getLine(Chunk* chunk, int instruction) {
  int offset = 0;
  for (int i = 0; i < chunk->count; i++) {
    if (instruction < offset + chunk->lines[i].count) {
      return chunk->lines[i].line;
    }
    offset += chunk->lines[i].count;
  }
  return -1; // if this ever happens, something very very bad occurred. but this shouldn't ever happen anyway, so i'm sure it'll be fine.
}
