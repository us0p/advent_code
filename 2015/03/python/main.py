from os.path import join
from enum import Enum


class Direction(Enum):
    UP = "^"
    RIGHT = ">"
    DOWN = "v"
    LEFT = "<"


def calculate_visited_houses(
        house_grid: list[list[bool]],
        position: dict,
        secondary_position: dict
) -> int:
    first_row_len = len(house_grid[position["y"]])

    match character:
        case Direction.UP.value:
            if (position["y"] - 1 < 0):
                new_line = [False for x in range(first_row_len)]
                house_grid.insert(0, new_line)
                position["y"] += 1
                secondary_position["y"] += 1
            position["y"] -= 1

            return check_new_house(house_grid, position)
        case Direction.RIGHT.value:
            if (position["x"] + 1 >= first_row_len):
                for row in house_grid:
                    row.append(False)
            position["x"] += 1

            return check_new_house(house_grid, position)
        case Direction.DOWN.value:
            if (position["y"] + 1 >= len(house_grid)):
                new_line = [False for x in range(first_row_len)]
                house_grid.append(new_line)
            position["y"] += 1

            return check_new_house(house_grid, position)
        case Direction.LEFT.value:
            if (position["x"] - 1 < 0):
                for row in house_grid:
                    row.insert(0, False)
                position["x"] += 1
                secondary_position["x"] += 1
            position["x"] -= 1

            return check_new_house(house_grid, position)

    return 0


def check_new_house(house_grid: list[list[bool]], position: dict) -> int:
    delivered_presents = 0
    if (not house_grid[position["y"]][position["x"]]):
        delivered_presents = 1
    house_grid[position["y"]][position["x"]] = True
    return delivered_presents


if __name__ == "__main__":
    file_name = join("..", "inputs.txt")

    position = {"x": 4, "y": 4}
    robot_position = {"x": 4, "y": 4}
    house_grid = [[False for x in range(10)] for y in range(10)]
    house_grid[position["x"]][position["y"]] = True
    count = 1
    is_santa_turn = True
    with open(file_name) as f:
        while True:
            character = f.read(1)
            if (character == ""):
                break

            if (is_santa_turn):
                count += calculate_visited_houses(
                    house_grid,
                    position,
                    robot_position
                )
                is_santa_turn = False
                continue

            count += calculate_visited_houses(
                house_grid,
                robot_position,
                position
            )
            is_santa_turn = True

    print(count)
