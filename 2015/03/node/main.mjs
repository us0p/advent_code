import { open } from "node:fs/promises"
import { join } from "node:path"

/** @typedef Coordinates {Object}
 * @property {number} x
 * @property {number} y
 */

const DIRECTIONS = {
    "UP": "^",
    "RIGHT": ">",
    "DOWN": "v",
    "LEFT": "<"
}

/**
 * @param {Coordinates} coordinates
 * @param {Coordinates} secondaryCoordinates
 * @param {Array<Array<boolean>>} grid
 * @param {string} CHAR
 * @returns {number}
 */
function calculateNewHouses(coordinates, secondaryCoordinates, grid, CHAR) {
    switch (CHAR) {
        case DIRECTIONS.UP:
            if (coordinates.y - 1 < 0) {
                grid.unshift(new Array(grid[0].length).fill(false))
                coordinates.y += 1
                secondaryCoordinates.y += 1
            }
            coordinates.y -= 1
            return checkHouse(grid, coordinates)
        case DIRECTIONS.RIGHT:
            if (coordinates.x + 1 >= grid[0].length) {
                for (const row of grid) {
                    row.push(false);
                }
            }
            coordinates.x += 1
            return checkHouse(grid, coordinates)
        case DIRECTIONS.DOWN:
            if (coordinates.y + 1 >= grid.length) {
                grid.push(new Array(grid[0].length).fill(false))
            }
            coordinates.y += 1
            return checkHouse(grid, coordinates)
        case DIRECTIONS.LEFT:
            if (coordinates.x - 1 < 0) {
                for (const row of grid) {
                    row.unshift(false)
                }
                coordinates.x += 1
                secondaryCoordinates.x += 1
            }
            coordinates.x -= 1
            return checkHouse(grid, coordinates)
        default:
            return 0
    }
}

/**
 * @param {Array<Array<boolean>>} houseGrid
 * @param {Coordinates} position
 * @returns {number}
 */
function checkHouse(grid, position) {
    let deliveredGifts = 0
    if (!grid[position.y][position.x]) {
        deliveredGifts = 1
    }
    grid[position.y][position.x] = true
    return deliveredGifts
}

const filePath = join("..", "inputs.txt")

const f = await open(filePath)

const buff = Buffer.alloc(1)

const houseGrid = new Array(10)

for (let i = 0; i < houseGrid.length; i++) {
    houseGrid[i] = new Array(10).fill(false)

    if (i == 4) {
        houseGrid[i][4] = true
    }
}

/** @type {Coordinates} */
const santaPosition = { x: 4, y: 4 }

/** @type {Coordinates} */
const robotPosition = { x: 4, y: 4 }

let count = 1

let isSantaTurn = true

while ((await f.read(buff, 0, 1)).bytesRead > 0) {
    const CHAR = buff.toString("utf8")
    if (isSantaTurn) {
        count += calculateNewHouses(santaPosition, robotPosition, houseGrid, CHAR)
        isSantaTurn = false
        continue
    }

    count += calculateNewHouses(robotPosition, santaPosition, houseGrid, CHAR)
    isSantaTurn = true
}

console.log(count)

await f.close()
