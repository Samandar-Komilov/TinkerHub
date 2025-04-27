import unittest

from ecommercedp.src.choices import ProductType
from ecommercedp.src.product import ProductFactory


class TestProductFactory(unittest.TestCase):
    def test_create_book(self):
        book = ProductFactory.create_product(ProductType.BOOK, "Start With Why", 50.00, "Simon Sinek")
        print("> Book:", book)
        self.assertEqual(book.__str__(), f"Book: {book.name} - ${book.price}")

    def test_create_electronics(self):
        electronics = ProductFactory.create_product(ProductType.ELECTRONICS, "Amplifier", 1.75, 1)
        print("> Electronics:", electronics)
        self.assertEqual(electronics.__str__(), f"Electronics: {electronics.name} - ${electronics.price}")

    def test_create_clothing(self):
        clothing = ProductFactory.create_product(ProductType.CLOTHING, "Trousers", 25.00, "L")
        print("> Clothing:", clothing)
        self.assertEqual(clothing.__str__(), f"Clothing: {clothing.name} - ${clothing.price}")