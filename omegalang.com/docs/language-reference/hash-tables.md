---
sidebar_position: 10
---

# Hash Tables

Hashes have been introduced to Omega as of v0.16.0.

## Introduction

In Omega, a hash table (also known as a dictionary, dict, hashmap, or simply a hash) is a data structure that allows you to store key-value pairs. This data type is particularly useful for scenarios where you need to associate values with unique keys for quick lookup, insertion, and deletion.

## Syntax

To create a hash table in Omega, you can use the following syntax:

SomeOtherFile.omg:
```omega
var my_hash = {"key1": "value1", "key2": "value2"};
```

Values can be of any data type, and you can add new key-value pairs to an existing hash table using dot notation or bracket notation.

## Example Usage

Here's an example of how to use hash tables in Omega.

```omega
var my_hash = {"hash table!": "first hash key!"};

my_hash.dog = "Golden Retriever";
my_hash.apples = ["Golden Delicious", "Granny Smith"];

out my_hash["hash table!"]; // -> "first hash key!"
out my_hash.apples; // -> ["Golden Delicious", "Granny Smith"]
out my_hash.apples[0]; // -> "Golden Delicious"

out my_hash; // -> {"hash table!": "first hash key!", dog: "Golden Retriever", apples: ["Golden Delicious", "Granny Smith"]}

remove(my_hash, "dog");

out my_hash; // -> {"hash table!": "first hash key!", apples: ["Golden Delicious", "Granny Smith"]}
```
