

class Config:
    def __init__(self):
        return self
    
    def __getattr__(self, key):
        pass

    def __setattr__(self, name, value):
        pass

    def __delattr__(self, name):
        pass

    def __getitem__(self, key):
        pass

    def __setitem__(self, key, value):
        pass

    def __delitem__(self, key):
        pass

    def __dir__(self):
        pass

    def __contains__(self):
        pass