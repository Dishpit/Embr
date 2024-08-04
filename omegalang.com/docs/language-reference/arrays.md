---
sidebar_position: 8
---

# Arrays

Arrays have been introduced to Omega in v0.11.0.

## Introduction

Arrays in Omega are versatile and powerful data structures that allow you to store multiple values in a single variable. Arrays can contain elements of any data type, including numbers, strings, booleans, other arrays, and even functions! This documentation provides a comprehensive guide on how to use arrays in Omega, including array creation, access, manipulation, and built-in array functions.

## Creating Arrays

To create an array in Omega, use square brackets `[]` and separate elements with commas. Arrays can hold any type of data.

```omega
var empty_arr = [];
var number_arr = [1, 2, 3, 4, 5];
var mixeed_arr = [1, "two", true, 3.14, [5, 6]];
```

## Accessing Array Elements

You can access elements in an array using the index operator `[]`. Array indices are zero-based, meaning the first element is at index `0`.

```omega
var my_arr = [10, 20, 30, 40, 50];
var first_element = my_arr[0]; // 10
var third_element = my_arr[2]; // 30
```

## Modifying Array Elements

You can modify elements in an array by assigning a new value to a specific index.

```omega
var my_arr = [10, 20, 30, 40, 50];
my_arr[0] = 15; // my_arr is now [15, 20, 30, 40, 50]
```

## Built-in Array Functions

Omega provides a variety of built-in functions for manipulating arrays:

### prepend()

The `prepend()` function adds an element to the beginning of an array.

#### Syntax:

```omega
prepend(array, value);
```

#### Example:

```omega
var my_arr = [2, 3, 4];
prepend(my_arr, 1); // my_arr is now [1, 2, 3, 4]
```

### append()

The `append()` function adds an element to the end of an array.

#### Syntax:

```omega
append(array, value);
```

#### Example:

```omega
var my_arr = [1, 2, 3];
append(my_arr, 4); // my_arr is now [1, 2, 3, 4]
```

### head()

The `head()` function removes and returns the first element of an array.

#### Syntax:

```omega
var first_el = head(array);
```

#### Example:

```omega
var my_arr = [1, 2, 3, 4];
var first_el = head(my_arr); // first_el is 1, my_arr is now [2, 3, 4]
```

### tail()

The `tail()` function removes and returns the last element of an array.

#### Syntax:

```omega
var last_el = tail(array);
```

#### Example:

```omega
var my_arr = [1, 2, 3, 4];
var last_el = tail(my_arr); // last_el is 4, my_arr is now [1, 2, 3]
```

### rest()

The `rest()` function returns a new array containing all elements except the first one.

#### Syntax:

```omega
var new_arr = rest(arr);
```

#### Example:

```omega
var my_arr = [1, 2, 3, 4];
var new_arr = rest(my_arr); // new_arr is [2, 3, 4]
```

## Printing Arrays

You can print arrays using the `out` statement. Omega's `out` statement prints arrays in a readable format, showing all elements.

```omega
var my_arr = [1, 2, 3, 4, 5];
out my_arr; // outputs: [1, 2, 3, 4, 5]
```

## Further Examples

### Example 1: Basic Array Operations

```omega
var my_array = [1, 2, 3];
out my_array; // Outputs: [1, 2, 3]

append(my_array, 4);
out my_array; // Outputs: [1, 2, 3, 4]

prepend(my_array, 0);
out my_array; // Outputs: [0, 1, 2, 3, 4]

var first = head(my_array);
out first; // Outputs: 0
out my_array; // Outputs: [1, 2, 3, 4]

var last = tail(my_array);
out last; // Outputs: 4
out my_array; // Outputs: [1, 2, 3]

var rest_array = rest(my_array);
out rest_array; // Outputs: [2, 3]
```

### Example 2: Nested Arrays

```omega
var nested_array = [[1, 2], [3, 4], [5, 6]];
out nested_array; // Outputs: [[1, 2], [3, 4], [5, 6]]

var first_nested = nested_array[0];
out first_nested; // Outputs: [1, 2]

append(first_nested, 3);
out nested_array; // Outputs: [[1, 2, 3], [3, 4], [5, 6]]
```

### Example 3: Using Array Functions in Expressions

```omega
var my_array = [10, 20, 30];
append(my_array, head([1, 2, 3]));
out my_array; // Outputs: [10, 20, 30, 1]

var another_array = rest(my_array);
out another_array; // Outputs: [20, 30, 1]
```

### Example 4: Storing Functions in an Array and Calling Them

```
// Define some functions
fn say_hello() {
  out "Hello, Omega!";
}

fn add(a, b) {
  return a + b;
}

fn multiply(a, b) {
  return a * b;
}

// Store functions in an array
var function_array = [say_hello, add, multiply];

// Call the functions using array indices

// Call say_hello function
function_array[0](); // Outputs: Hello, World!

// Call add function with arguments
var sum = function_array[1](5, 3);
out sum; // Outputs: 8

// Call multiply function with arguments
var product = function_array[2](4, 2);
out product; // Outputs: 8
```