

class SmartPath:
    """
    A class that wraps a string path in a way that it can be used as an argument to most functions that expect a path-like object.
    
    It can be used to add a layer of abstraction between pure string paths
    and path-like objects. For example, it can be used to allow a function to 
    accept both string paths and path-like objects.
    
    :param path: The path to be wrapped.
    :type path: str or os.PathLike
    """
    def __init__(self, path):
        self.path = path

    def __fspath__(self):
        super().__fspath__()

    def __truediv__(self, other):
        return super().__truediv__(other)
    
    def __getitem__(self, key):
        return super().__getitem__()
    
    def __contains__(self, item):
        return super().__contains__(item)
    
    def __call__(self, *args, **kwargs):
        return super().__call__(*args, **kwargs)


    def __str__(self):
        return self.path
    
    def __repr__(self):
        return f"SmartPath<{self.path!r}>"
