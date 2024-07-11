---
sidebar_position: 2
---

# Types

Omega has a robust static type system.

In Omega, types are denoted with three-character keywords. Currently, Omega only has one type: integers.

Omega has two versions of each type: the value type, and the return type.

Value types are typically used for things like variables, such as: `int my_var = 42069;`. Return types are a bit different. They're annotated with a `@` after the parameters of a function are declared, like so: `fn my_func(x, y) @int {};`.

Every function has a return type, just like how every variable has a value type.

| Type | Keyword | Usage |
|------|---------|-------|
| Integer | `int` or `@int`  | Regular integers, such as `420`, `69`, or `8675309` |
