### E-Commerce simulation in Pure Python
Intended to practice Design Patterns right in Python

We’ll build a simple E-Commerce System that allows users to manage products, orders, and notifications. This project will cover the most commonly used design patterns in Python.


## Features:
# Product Management:
- Use the __Factory Pattern__ to create different types of products.
- Use the __Decorator Pattern__ to add dynamic behavior (e.g., gift wrapping, discounts).

# Shopping Cart:
- Use the __Singleton Pattern__ to ensure only one cart exists per user.

# Checkout Process:
- Use the __Facade Pattern__ to simplify the checkout process.

# Order Management:
- Use the __Command Pattern__ to handle order actions (e.g., place, cancel).

# Notifications:
- Use the __Observer Pattern__ to notify users of order status changes.

# Pricing Strategies:
Use the __Strategy Pattern__ to apply different pricing strategies.


## Phases

# Phase 1: Creational Design Patterns
Factory Pattern:

Create a ProductFactory to generate different types of products (e.g., Book, Electronics, Clothing).

Each product type will have its own class with specific attributes.

Singleton Pattern:

Create a ShoppingCart class that acts as a singleton. Only one instance of the shopping cart should exist for a user.

Builder Pattern:

Create an OrderBuilder to construct complex Order objects step by step (e.g., add products, set shipping address, apply discounts).

# Phase 2: Structural Design Patterns
Adapter Pattern:

Create an adapter to integrate a third-party payment gateway (e.g., PayPal) with your e-commerce system.

Decorator Pattern:

Add dynamic behavior to products (e.g., add gift wrapping or discount decorators).

Facade Pattern:

Create a CheckoutFacade to simplify the checkout process (e.g., validate cart, process payment, send confirmation).

# Phase 3: Behavioral Design Patterns
Observer Pattern:

Implement a notification system where users are notified when their order status changes (e.g., "Order Shipped", "Order Delivered").

Strategy Pattern:

Implement different pricing strategies (e.g., RegularPricing, DiscountedPricing, SeasonalPricing).

Command Pattern:

Implement a command-based system for order actions (e.g., PlaceOrderCommand, CancelOrderCommand).