---
sidebar_position: 6
---

# Builtin Functions

As of v0.10.0, Omega has a couple of builtin functions.

## clock()

Returns the length of time elapsed since the interpreter has started.

```omega
out clock();
```

## out

Outputs the evaluation of the remainder of the line into the console.

```omega
out "Hello, Omega!"; // Hello, Omega!
```

## out(x)

Outputs the contents to the terminal. Note that `out` also returns a void object.

```omega
out("Hello World!") // Hello World\nvoid.
```
