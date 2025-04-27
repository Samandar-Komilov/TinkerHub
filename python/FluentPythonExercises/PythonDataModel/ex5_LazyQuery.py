

class LazyQuery:
    def __init__(self):
        self.query = 0

    def __getattr__(self, key):
        pass

    def __call__(self, *args, **kwargs):
        pass

    def __bool__(self):
        pass

    def __getitem__(self, key):
        pass

    def __contains__(self):
        pass

    def __iter__(self):
        pass