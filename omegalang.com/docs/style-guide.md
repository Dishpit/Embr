---
sidebar_position: 1
---

# Style Guide

By sticking to an established style guide, programmers don't need to formulate ad hoc style rules, nor do they need to debate with other programmers about which style rules should be used, which saves time, communication overhead, and mental energy.

Humans comprehend information through pattern matching. By ensuring that all Omega code has similar formatting, less mental effort is required to comprehend a new project, lowering the barrier to entry for new developers.

Thus, there's a very large benefit to using a community-consistent format.

## The default Omega style

The Omega Style Guide defines the default Omega style. Everything in this style guide, whether or not it uses language such as "must" or the imperative mood such as "insert a space ..." or "break the line after ...", refers to the default style.

This should not be interpreted as forbidding developers from following a non-default style, or forbidding tools from adding any particular configuration options, though the original creator of Omega may forcefully extract your kneecaps as a sacrifice to the gods if he catches you drifting from the default style.

## Formatting Conventions

### Indentation and line width

- Use spaces, not tabs
- Each level of indentation must be 2 spaces (that is, all indentation outside of string literals and comments must be a multiple of 2).
- The maximum width for a line is 80 characters.

### Block indent

Prefer block indent over visual indent:

```omega
function_call(
  foo,
  bar,
);
```
This makes for smaller diffs and less rightward drift.

### Trailing commas

In comma-separated listts of any kind, do not use a trailing comma when followed by a newline:

```omega
function_call(
  argument,
  another_argument
);
```

### Blank lines

Separate items and statements by either zero or one blank lines.

```omega
fn foo() @int {
  var x = ...;

  var y = ...;
  var z = ...;
};

fn bar() @void {};
fn baz() @void {};
```
