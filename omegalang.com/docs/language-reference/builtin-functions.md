---
sidebar_position: 7
---

# Builtin Functions

As of v0.8.0, Omega has a variety of builtin functions.

## len(x)

Returns the length of either the string or the array that's passed into it.

```omega
len("Hello, world!"); // 13
len([1, 2, 3, "Hello, Omega!", 420, 69]); // 6
```

## first(x)

Returns the first element of an array.

```omega
first([420, 69, 8675309]); // 420
```

## last(x)

Returns the last element of an array.

```omega
last([420, 69, 8675309]); // 8675309
```

## tail(x)

Returns a new copy of the original array, containing everything except for the first element.

```omega
tail([420, 69, 8675309]); // [69, 8675309]
```

## push(x, y)

Returns a new copy of the original array, containing the added value appended onto the end.

```omega
push([69, 8675309], "Zaphod Beeblebrox"); // [69, 8675309, Zaphod Beeblebrox]
```

## out(x)

Outputs the contents to the terminal. Note that `out` also returns a void object.

```omega
out("Hello World!") // Hello World\nvoid.
```
