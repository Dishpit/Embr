---
sidebar_position: 6
---

# Arrays

As of v0.8.0, Omega has functional arrays.

## Using Arrays

Arrays in Omega have a variety of built-in functions accessible to them.

Arrays in Omega are immutable, so any modifications to arrays return new copies of the original array.

Currently, arrays can only be used through live interpretation, and are unable to be stored as variables due to a lack of an array data type.

### Syntax

The basic syntax for using arrays in Omega is as follows:

```omega
[1, 2, 5 * 3, 2 + 4, "hello, Omega!"][2]; // this returns: 15
```
