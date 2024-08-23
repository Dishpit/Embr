---
sidebar_position: 8
---

# Arrays

Arrays have been introduced to Embr in v0.11.0.

## Introduction

Arrays in Embr are versatile and powerful data structures that allow you to store multiple values in a single variable. Arrays can contain elements of any data type, including numbers, strings, booleans, other arrays, and even functions!

## Creating Arrays

To create an array in Embr, use square brackets `[]` and separate elements with commas. Arrays can hold any type of data.

```embr
var empty_arr = [];
var number_arr = [1, 2, 3, 4, 5];
var mixed_arr = [1, "two", true, 3.14, [5, 6]];
```

## Accessing Array Elements

You can access elements in an array using the index operator `[]`. Array indices are zero-based, meaning the first element is at index `0`.

```embr
var my_arr = [10, 20, 30, 40, 50];
var first_element = my_arr[0]; // 10
var third_element = my_arr[2]; // 30
```

## Modifying Array Elements

You can modify elements in an array by assigning a new value to a specific index.

```embr
var my_arr = [10, 20, 30, 40, 50];
my_arr[0] = 15; // my_arr is now [15, 20, 30, 40, 50]
```

## Examples

### Example 1: Basic Array Operations

```embr
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

```embr
var nested_array = [[1, 2], [3, 4], [5, 6]];
out nested_array; // Outputs: [[1, 2], [3, 4], [5, 6]]

var first_nested = nested_array[0];
out first_nested; // Outputs: [1, 2]

append(first_nested, 3);
out nested_array; // Outputs: [[1, 2, 3], [3, 4], [5, 6]]
```

### Example 3: Using Array Functions in Expressions

```embr
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
  out "Hello, Embr!";
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
