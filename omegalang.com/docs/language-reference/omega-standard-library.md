---
sidebar_position: 9
---

# Omega Standard Library

The Omega Standard Library has been introduced to Omega as of v0.12.0.

## Introduction

The Omega STL is prepackeged with every new release of Omega, in a separate directory called `/stl`. The entire STL is written in Omega.

As of v0.12.0, Omega contains a very basi Math library containing two functions: `power` and `sqrt`, which return the resulting power of a number, and the square root of a number, respectively.

__⚠ With the current version, the entire contents of `/stl` is loaded into the virtual machine at runtime. There are plans to have this patched ASAP in order to save on memory usage. ⚠__

## Examples

```omega
var my_num = 42;

out Math.power(my_num, 2); // -> outputs 1764
out Math.sqrt(my_num); // -> outputs 6.48074
```

## STL

### Math

```omega
class Math {
  power(base, exponent) {
    var result = 1;
    for (var i = 0; i < exponent; i = i + 1) {
      result = result * base;
    }
    return result;
  }

  sqrt(x) {
    if (x < 0) {
      return "Error: Negative argument for sqrt";
    }
    var guess = x / 2.0;
    var epsilon = 0.00001;
    var difference = guess * guess - x;
    while (difference > epsilon or difference < -epsilon) {
      guess = (guess + x / guess) / 2.0;
      difference = guess * guess - x;
    }
    return guess;
  }
}

var Math = Math();
```
