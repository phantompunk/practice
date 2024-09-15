class MyQueue(object):

    def __init__(self):
        self.queue = []
        self.stack = []
        

    def push(self, x):
        """
        :type x: int
        :rtype: None
        """
        while self.queue:
            self.stack.append(self.queue.pop())
        self.stack.append(x)
        while self.stack:
            self.queue.append(self.stack.pop())
        

    def pop(self):
        """
        :rtype: int
        """
        if not self.queue:
            return
        return self.queue.pop()
        

    def peek(self):
        """
        :rtype: int
        """
        if not self.queue:
            return
        return self.queue[-1]
        

    def empty(self):
        """
        :rtype: bool
        """
        return not bool(self.queue)
        


# Your MyQueue object will be instantiated and called as such:
# obj = MyQueue()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.peek()
# param_4 = obj.empty()
