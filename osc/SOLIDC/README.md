# SOLID patterns in C

Everyone believes SOLID patterns are only for Object Oriented Programming languages. But it is not true. You are not required to be depend on OOP all the time. SOLID is actually a set of principles that can be applied to any programming language. 

The reason we are implementing them in C is, in this language we have the most control (and the ability to shoot our own foot) and least abstractions, which allows us to see every possible way to mess things up, see the issues and solve them.

## The SOLID principles

Recall that, the 5 principles we discussed here are:

- Single Responsibility Principle
- Open/Closed Principle
- Liskov Substitution Principle
- Interface Segregation Principle
- Dependency Inversion Principle

## The Application: Transaction Logging System

We will build a transaction logging system in C to demonstrate the SOLID principles. The requirements are hostile and contradictory, which requires to change the system frequently. And we gradually implement the features, hit the pain points, and apply the SOLID principles to solve them.