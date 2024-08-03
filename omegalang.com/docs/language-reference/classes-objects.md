---
sidebar_position: 7
---

# Classes and Objects

As of v0.10.0, Omega has classes.

## Class Declaration

In Omega, a class is declared using the `class` keyword followed by the class name and a block of code containing methods. A class serves as a blueprint for creating objects (instants).

## Syntax

```omega
class ClassName {
  // Method definitions
}
```

## Example:

```omega
class Animal {
  speak() {
    out "Animal sound";
  }
}
```

# Instantiating Objects

To create an object from a class, call the class name followed by parentheses. This invokes the class's initializer if it has one.

## Syntax:

```omega
var objectName = ClassName();
```

## Example:

```omega
var dog = Animal();
dog.speak(); // outputs: Animal sound
```

# Methods

Methods are functions defined within a class. They can be called on objects instantiated from the class.

## Syntax:

```omega
class ClassName {
  methodName() {
    // method body
  }
}
```

## Example:

```omega
class Dog {
  bark() {
    out "Woof!";
  }
}

var my_dog = Dog();
my_dog.bark(); // outputs: Woof!
```

# Initializer (Constructor)

An initializer is a special method named `init` that gets called when an object is created. It initializes the object's state.

## Syntax:

```omega
class ClassName {
  init() {
    // initialization code
  }
}
```

## Example:

```omega
class Person {
  init(name) {
    this.name = name;
  }

  greet() {
    out "Hello, " + this.name;
  }
}

var alice = Person("Alice");
alice.greet(); // outputs: Hello, Alice
```

# Inheritance and Super

Omega supports single inheritance, allowing a class to inherit methods and properties from another class using the `<` symbol. The `super` keyword is used to call methods from the superclass.

## Syntax:

```omega
class SubClass < SuperClass {
  // subclass methods
}
```

## Example:

```omega
class Vehicle {
  start() {
    out "Vehicle starting";
  }
}

class Car < Vehicle {
  start() {
    super.start();
    out "Car starting";
  }
}

var my_car = Car();
my_car.start(); 
// Outputs: 
// Vehicle starting
// Car starting
```
