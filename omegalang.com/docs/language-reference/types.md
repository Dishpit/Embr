---
sidebar_position: 2
---

# Types

As of v0.10.0, Omega is dynamically typed.

However, we're planning to shift toward static typing in the near future.

Function return types are annotated with a `@` after the parameters of a function are declared, like so: `fn my_func(x, y) @int {};`.

Every function has a return type.

| Type | Keyword | Usage |
|------|---------|-------|
| Void | `@void` | A special return type for functions that don't return anything. |
| Integer | `@int`  | Regular integers, such as `420`, `69`, or `8675309` |
| Boolean | `@bool`  | Boolean values: `true` or `false`. |
| String | `@str`  | Strings of text, encapsulated in double-quotations, such as `"Hello, Omega!"` or `"Zaphod Beeblebrox"`. |
