import os
from typing import Callable


def getInputs():
    metrics = None
    with open(path) as inputs:
        metrics = inputs.read()
        metrics = metrics.split("\n")[:-1]
    return metrics


def calculateRectangleWrappingPapper(
        length: int,
        width: int,
        height: int
) -> int:
    baseArea = length * width
    sideArea = width * height
    frontArea = length * height

    wrappingPapper = 2 * baseArea + 2 * sideArea + 2 * frontArea

    return wrappingPapper + min(baseArea, sideArea, frontArea)


def calculateRibbon(
        length: int,
        width: int,
        height: int
) -> int:
    basePerimeter = 2 * (length + width)
    sidePerimeter = 2 * (width + height)
    frontPerimeter = 2 * (length + height)

    bow = length * width * height

    return bow + min(basePerimeter, sidePerimeter, frontPerimeter)


def getTotalPapperSquareFeet(
        metrics: list[str],
        calculateFn: Callable[[int, int, int], int]
) -> int:
    totalPapperSquareFeet = 0

    for metric in metrics:
        metricList = metric.split("x")

        totalPapperSquareFeet += calculateFn(
            int(metricList[0]),
            int(metricList[1]),
            int(metricList[2])
        )

    return totalPapperSquareFeet


if __name__ == "__main__":
    path = os.path.join("..", "inputs.txt")
    metrics = getInputs()
    totalPapperSquareFeet = getTotalPapperSquareFeet(
        metrics,
        calculateRectangleWrappingPapper
    )
    totalRibbonSquareFeet = getTotalPapperSquareFeet(
        metrics,
        calculateRibbon
    )

    print(f'Total wrapping papper: {totalPapperSquareFeet}')
    print(f'Total ribbon: {totalRibbonSquareFeet}')
