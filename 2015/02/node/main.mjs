import { readFile } from "node:fs/promises"
import { join } from "node:path"

async function getInputs() {
    const inputs = await readFile(join("..", "inputs.txt"), { encoding: "utf-8" })
    return inputs.split("\n").slice(0, -1)
}

/**
 * @param {Array<string>} inputs
 * @param {function(number, number, number)} calculateFn
 * @returns {number}
*/
function getTotalWrappingPapper(inputs, calculateFn) {
    let totalWrappingPapper = 0
    for (const metric of inputs) {
        const [l, w, h] = metric.split("x")

        totalWrappingPapper += calculateFn(parseInt(l), parseInt(w), parseInt(h))
    }
    return totalWrappingPapper
}

/**
 * @param {number} l
 * @param {number} w
 * @param {number} h
 * @returns {number}
*/
function calculateRectangleWrappingPapper(l, w, h) {
    const baseArea = l * w
    const sideArea = w * h
    const frontArea = l * h

    const wrappingPapper = 2 * baseArea + 2 * sideArea + 2 * frontArea

    return wrappingPapper + Math.min(baseArea, sideArea, frontArea)
}

/**
 * @param {number} l
 * @param {number} w
 * @param {number} h
 * @returns {number}
*/
function calculateRibbon(l, w, h) {
    const basePerimeter = 2 * (l + w)
    const sidePerimeter = 2 * (w + h)
    const frontPerimeter = 2 * (l + h)

    const bow = l * w * h

    return bow + Math.min(basePerimeter, sidePerimeter, frontPerimeter)
}


const inputs = await getInputs()
const totalWrappingPapper = getTotalWrappingPapper(inputs, calculateRectangleWrappingPapper)
const totalRibbon = getTotalWrappingPapper(inputs, calculateRibbon)

console.log(`Total wrapping papper: ${totalWrappingPapper}`)
console.log(`Total ribbon: ${totalRibbon}`)
