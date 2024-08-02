#ifndef omega_chunk_h
#define omega_chunk_h

#include "common.h"

typedef enum {
  OP_RETURN,
} OpCode;

typedef struct {
  int count;
  int capacity;
  uint8_t* code;
} Chunk;

void initChunk(Chunk* chunk);

#endif
