"""
    Singleton Pattern: ShoppingCart
"""

class ShoppingCart:
    _instance = None

    def __new__(cls):
        if cls._instance is None:
            cls._instance = super().__new__(cls)
            cls._instance.items = []
        return cls._instance

    def add_item(self, product):
        self.items.append(product)

    def remove_item(self, product):
        self.items.remove(product)

    def __str__(self):
        cart_items = [item.__str__() for item in self.items]
        return f"ShoppingCart: {cart_items}"