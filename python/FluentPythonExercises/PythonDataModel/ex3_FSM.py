"""

Usages:
    - FSM(states=iterable("state1", "state2", "state3"), current_state=None)
    - FSM("state2") - switch to state2
    - FSM["state2"] = <something> specify state
    - 

"""

class FSM:
    def __init__(self):
        self.state = 0

    def __setattr__(self, name, value):
        pass

    def __getattr__(self, name):
        pass

    def __call__(self, *args, **kwargs):
        pass

    def __enter__(self):
        pass

    def __exit__(self):
        pass

    def __await__(self):
        pass