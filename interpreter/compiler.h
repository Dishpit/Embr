#ifndef omega_compiler_h
#define omega_compiler_h

#include "object.h"
#include "vm.h"

ObjFunction* compile(const char* source);
void markCompilerRoots();

#endif