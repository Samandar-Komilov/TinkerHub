

class ORMModel:
    def __prepare__(name, bases, namespace, **kwargs):
        pass

    def __init_subclass__(cls, **kwargs):
        pass

    def __class_getitem__(cls, item):
        pass

    def __getattr__(self, name):
        pass

    def __setattr__(self, name, value):
        pass

    def __dir__(self):
        pass