"""
    Builder Pattern: Order Builder
"""

class Order:
    def __init__(self):
        self.products = []
        self.shipping_address = None
        self.discount = 0

    def __str__(self):
        return f"Order: {[p.display() for p in self.products]}, Shipping: {self.shipping_address}, Discount: {self.discount}%"
    

class OrderBuilder:
    def __init__(self):
        self.order = Order()

    def add_order(self, order):
        self.products.append(order)
        return self
    
    def set_shipping_address(self, shipping_address):
        self.shipping_address = shipping_address
        return self
    
    def apply_discount(self, discount):
        self.discount = discount
        return self
    
    def build(self):
        return self.order
    

""" Testing:
order = (
    OrderBuilder()
    .add_product(book)
    .add_product(clothing)
    .set_shipping_address("123 Main St")
    .apply_discount(10)
    .build()
)

Builds the Order step-by-step, but hiding all of the complexity from us.
"""