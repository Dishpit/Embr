---
sidebar_position: 2
---

# Types

As of v0.10.0, Embr is dynamically typed.

However, we're planning to shift toward static typing in the near future.

Function return types are annotated with a `@` after the parameters of a function are declared, like so: `fn my_func(x, y) @int {};`.

Every function has a return type. Note that despite this, specifying the return type for a function is completely optional as they are also implicitly determined. Because of this, valid Embr functions can also look like: `fn my_func(x, y) {};`.

| Type | Keyword | Usage |
|------|---------|-------|
| Void | `@void` | A special return type for functions that don't return anything. |
| Integer | `@int`  | Regular integers, such as `420`, `69`, or `8675309` |
| Boolean | `@bool`  | Boolean values: `true` or `false`. |
| String | `@str`  | Strings of text, encapsulated in double-quotations, such as `"Hello, Embr!"` or `"Zaphod Beeblebrox"`. |
