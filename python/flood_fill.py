from typing import List


class Solution:
    def floodFill(
        self, image: List[List[int]], sr: int, sc: int, color: int
    ) -> List[List[int]]:
        if image[sr][sc] == color:
            return image

        start_color = image[sr][sc]
        queue = [(sr, sc)]
        while queue:
            x, y = queue.pop(0)

            # print(f"Step: ({x},{y})\n")
            if x >= 0 and y >= 0:
                image[x][y] = color
            if x + 1 < len(image):
                if image[x + 1][y] == start_color:
                    queue.append((x + 1, y))
            if x - 1 >= 0:
                if image[x - 1][y] == start_color:
                    queue.append((x - 1, y))
            if y + 1 < len(image[0]):
                if image[x][y + 1] == start_color:
                    queue.append((x, y + 1))
            if y - 1 < len(image):
                if image[x][y - 1] == start_color:
                    queue.append((x, y - 1))
            # print(f"Step: ({x},{y})\n", image[0], "\n", image[1], "\n", image[2])

        return image
