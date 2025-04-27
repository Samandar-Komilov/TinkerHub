from abc import ABC, abstractmethod

from ecommercedp.src.choices import ProductType

"""
    Factory Pattern: ProductFactory
"""

class Product(ABC):
    def __init__(self, name, price):
        self.name = name
        self.price = price
    
    @abstractmethod
    def __str__(self):
        pass


class Book(Product):
    def __init__(self, name, price, author):
        super().__init__(name, price)
        self.author = author

    def __str__(self):
        return f"Book: {self.name} - ${self.price}"
    

class Electronics(Product):
    def __init__(self, name, price, warranty):
        super().__init__(name, price)
        self.warranty = warranty

    def __str__(self):
        return f"Electronics: {self.name} - ${self.price}"
    

class Clothing(Product):
    def __init__(self, name, price, size):
        super().__init__(name, price)
        self.size = size

    def __str__(self):
        return f"Clothing: {self.name} - ${self.price}"
    


class ProductFactory:
    @staticmethod
    def create_product(product_type, *args):
        if product_type == ProductType.BOOK:
            return Book(*args)
        elif product_type == ProductType.ELECTRONICS:
            return Electronics(*args)
        elif product_type == ProductType.CLOTHING:
            return Clothing(*args)
        else:
            raise ValueError("Invalid product type!")
        
