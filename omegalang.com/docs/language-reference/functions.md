---
sidebar_position: 4
---

# Functions

Functions are central to structuring code in Omega, allowing you to define reusable blocks of code.

In Omega, functions are first-class citizens. This means they can be treated in the same way as any other data type!

## Defining Functions

In Omega, functions are defined using the `fn` keyword. This is followed by the function name, a pair of parentheses `()` indicating parameters, the function's return type denoted with `@<type>`, and the function body enclosed in braces `{}`.

## Returning

All functions in Omega have a return type. If a function doesn't return anything, then its return type is void.

## Calling Functions

Functions can be called by just referring to the name of the function, followed by a pair of parentheses that contain any applicable arguments. For example: `my_func(420, 69);`

### Syntax

The basic syntax for defining a function in Omega is as follows:

```omega
fn my_func() @void {
  // function body
}
```

### Example

Here is a simple example of a function in Omega:

```omega
fn add_ints(x, y) @int {
  return x + y;
};

fn main() @void {
  add_ints(420, 69);
};
```
