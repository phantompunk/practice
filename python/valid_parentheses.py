class Solution(object):
    def isValid(self, s):
        """
        :type s: str
        :rtype: bool
        """
        stack = []
        for char in s:
            if stack and char in ["}","]",")"]:
                if stack[-1]+char in "(){}[]":
                    stack.pop()
            else:
                stack.append(char)
        return not stack 
