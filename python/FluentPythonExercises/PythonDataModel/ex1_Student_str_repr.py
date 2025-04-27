"""
######## EX 1 ########

Example of the difference between __str__ and __repr__

The __str__ method should return a string that is a nice human-readable representation of the object, while the __repr__ method should return a string that is a valid Python expression that could be used to recreate the object.

In this example, the __str__ method returns a simple string with the name of the student, while the __repr__ method returns a string that includes both the name and age of the student.
"""

class Student:
    def __init__(self, name, age):
        self.name = name
        self.age = age

    def __str__(self):
        return self.name
    
    def __repr__(self):
        return f"Name: {self.name}, Age: {self.age}"

s1 = Student("Hi", 12)

print(s1, s1.__str__(), s1.__repr__())
