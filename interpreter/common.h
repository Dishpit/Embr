#ifndef omega_common_h
#define omega_common_h

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>
// DEV FLAGS
#define NAN_BOXING // there's a chance your CPU architecture may not play nice with this flag, so turn it off if your Omega code seems to be running abnormally slow
// #define DEBUG_PRINT_CODE
// #define DEBUG_TRACE_EXECUTION

// #define DEBUG_STRESS_GC
// #define DEBUG_LOG_GC

#define UINT8_COUNT (UINT8_MAX + 1)

#endif
