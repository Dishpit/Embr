---
sidebar_position: 4
---

# Functions

Functions are central to structuring code in Omega, allowing you to define reusable blocks of code.

In Omega, functions are first-class citizens. This means they can be treated in the same way as any other data type!

## Defining Functions

In Omega, functions are defined using the `fn` keyword. This is followed by the function name, a pair of parentheses `()` indicating parameters, the function's return type denoted with `@<type>`, and the function body enclosed in braces `{}`.

### Syntax

The basic syntax for defining a function in Omega is as follows:

```omega
fn my_func() @int {
  // function body
}
```

### Example

Here is a simple example of a function in Omega:

```omega
fn add_nums(x, y) @int {
  return x + y;
};
```
