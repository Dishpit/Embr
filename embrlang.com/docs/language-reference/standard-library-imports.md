---
sidebar_position: 9
---

# Standard Library & Imports

The Embr Standard Library has been introduced to Embr as of v0.12.0.

## Introduction

The Embr STL is prepackeged with every new release of Embr, in a separate directory called `/stl`. The entire STL is written in Embr.

As of v0.12.0, Embr contains a very basic Math library containing two functions: `power` and `sqrt`, which return the resulting power of a number, and the square root of a number, respectively.

You can import files from the Embr STL in one of two ways: the explicit `import` keyword followed by the name of the library you'd like to import, or use the syntactic sugar `#` followed by the name of the library you'd like to import. You can import your own custom Embr files in the same method. Note that file importing only needs the name of the file, not the extension.

When importing, Embr will attempt to search through the STL first. Because of this, you won't be able to import your own files if they share the same name as a file that's in the STL.

## Example

SomeOtherFile.mbr:
```embr
fn hello_world() {
  out "Hello, Embr!";
}
```
main.mbr:
```embr
import Math;
# SomeOtherFile;

var my_num = 42;

out Math.power(my_num, 2); // -> outputs 1764
out Math.sqrt(my_num); // -> outputs 6.48074
hello_world(); // -> outputs "Hello, Embr!"
```

## STL

### Math

```embr
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
