# from sense_hat import SenseHat


on = [255, 255, 255]
off = [0, 0, 0]


class Grid:
    def __init__(self, value):
        self.grid = []
        for i in range(0, 8):
            row = []
            for j in range(0, 8):
                row.append(off)
            self.grid.append(row)
        self.set_value(value)

    def set_digit(self, digit, offset_top, offset_left):
        assert 0 <= digit < 10
        lights = set()
        if digit in [2, 3, 5, 6, 7, 8, 9, 0]:  # top
            lights.add((0, 0))
            lights.add((0, 1))
            lights.add((0, 2))
        if digit in [4, 5, 6, 8, 9, 0]:  # top left
            lights.add((0, 0))
            lights.add((1, 0))
            lights.add((2, 0))
        if digit in [1, 2, 3, 4, 7, 8, 9, 0]:  # top right
            lights.add((0, 2))
            lights.add((1, 2))
            lights.add((2, 2))
        if digit in [2, 3, 4, 5, 6, 8, 9, 0]:  # middle
            lights.add((2, 0))
            lights.add((2, 1))
            lights.add((2, 2))
        if digit in [2, 6, 8, 0]:  # bottom left
            lights.add((2, 0))
            lights.add((3, 0))
            lights.add((4, 0))
        if digit in [1, 3, 4, 5, 6, 7, 8, 9, 0]:  # bottom right
            lights.add((2, 2))
            lights.add((3, 2))
            lights.add((4, 2))
        if digit in [2, 3, 5, 6, 8, 0]:  # bottom
            lights.add((4, 0))
            lights.add((4, 1))
            lights.add((4, 2))

        for light in lights:
            self.grid[offset_top + light[0]][offset_left + light[1]] = on

    def set_left(self, digit):
        self.set_digit(digit, 2, 1)

    def set_right(self, digit):
        self.set_digit(digit, 2, 5)

    def set_value(self, value):
        tens = int(value/10)
        self.set_left(tens)
        self.set_right(value - tens*10)


v = Grid(79)
print(v)
