import unittest

from ecommercedp.src.cart import ShoppingCart


class TestCartSingleton(unittest.TestCase):
    def setUp(self):
        self.cart = ShoppingCart()

    def test_cart_singleton(self):
        new_cart = ShoppingCart()
        print("> Cart:", id(self.cart), id(new_cart))
        self.assertEqual(id(self.cart), id(new_cart))