import os
from typing import List


def compute_calories(args: List[List[int]]) -> int:
    values = []
    for calories in args:
        values.append(sum(calories, 0))
    values.sort(reverse=True)
    return sum(values[0:3], 0)

if __name__ == "__main__":
    input_path = os.environ.get("PROJECT_ROOT", "") + "/days/day-1/input.txt"
    with open(input_path) as f:
        lines = f.readlines()
        payload: List[List[int]] = []
        new_list = []
        for line in lines:
            if line == "\n":
                new_list = []
                payload.append(new_list)
                continue

            new_list.append(int(line))

        value = compute_calories(payload)
        print("The most carried calories:", value)
