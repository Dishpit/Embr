#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "compiler.h"
#include "memory.h"
#include "common.h"
#include "chunk.h"
#include "debug.h"
#include "object.h"
#include "vm.h"

static void repl() {
  char line[1024];
  for (;;) {
    printf(">> ");

    if (!fgets(line, sizeof(line), stdin)) {
      printf("\n");
      break;
    }

    interpret(line);
  }
}

static char* readFile(const char* path) {
  FILE* file = fopen(path, "rb");
  if (file == NULL) {
    fprintf(stderr, "Could not open file \"%s\".\n", path);
    exit(74);
  }

  fseek(file, 0L, SEEK_END);
  size_t fileSize = ftell(file);
  rewind(file);

  char* buffer = (char*)malloc(fileSize + 1);
  if (buffer == NULL) {
    fprintf(stderr, "Not enough memory to read \"%s\".\n", path);
    exit(74);
  }

  size_t bytesRead = fread(buffer, sizeof(char), fileSize, file);
  if (bytesRead < fileSize) {
    fprintf(stderr, "Could not read file \"%s\".\n", path);
    exit(74);
  }
  
  buffer[bytesRead] = '\0';

  fclose(file);
  return buffer;
}

static void runFile(const char* path) {
  // restrict files to those that contain a .o extension
  const char* extension = strrchr(path, '.');
  if (extension == NULL || strcmp(extension, ".omg") != 0) {
    fprintf(stderr, "Error: File must be an Omega code file (.omg extension).\n");
    exit(74);
  }
  
  char* source = readFile(path);
  InterpretResult result = interpret(source);
  free(source);

  if (result == INTERPRET_COMPILE_ERROR) exit(65);
  if (result == INTERPRET_RUNTIME_ERROR) exit(70);
}

void loadStandardLibrary() {
  const char *stlPath = "./stl/";
  const char *mathFile = "math.omg";
  char filePath[256];
  snprintf(filePath, sizeof(filePath), "%s%s", stlPath, mathFile);

  FILE *file = fopen(filePath, "r");  // Open as a text file
  if (!file) {
    fprintf(stderr, "Failed to open standard library file: %s\n", filePath);
    return;
  }

  fseek(file, 0, SEEK_END);
  size_t fileSize = ftell(file);
  fseek(file, 0, SEEK_SET);

  char *source = (char *)malloc(fileSize + 1);
  if (!source) {
    fprintf(stderr, "Failed to allocate memory for standard library.\n");
    fclose(file);
    return;
  }

  size_t index = 0;
  int c;
  while ((c = fgetc(file)) != EOF) {
    if (c == '\n' || c == '\r') {
      source[index++] = ' ';
    } else {
      source[index++] = (char)c;
    }
  }
  source[index] = '\0';

  fclose(file);

  InterpretResult result = interpret(source);
  free(source);

  if (result == INTERPRET_COMPILE_ERROR) exit(65);
  if (result == INTERPRET_RUNTIME_ERROR) exit(70);
}



int main(int argc, const char* argv[]) {
  initVM();

  loadStandardLibrary();

  if (argc == 1) {
    repl();
  } else if (argc == 2) {
    runFile(argv[1]);
  } else {
    fprintf(stderr, "Usage: omega [path]\n");
    exit(64);
  }

  freeVM();
  return 0;
}
