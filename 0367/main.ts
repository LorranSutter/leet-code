function isPerfectSquare1(num: number): boolean {
    for (let i = 1; ; i++) {
        let square = i * i
        if (square > num) {
            return false
        }
        if (square == num) {
            return true
        }
    }
};

function isPerfectSquare2(num: number): boolean {
    let left = 1
    let middle = 0
    let right = num
    let square = 0

    while (left <= right) {
        middle = Math.floor(left + (right - left) / 2)
        square = middle * middle
        if (square == num) {
            return true
        }
        if (square < num) {
            left = middle + 1
        } else {
            right = middle - 1
        }
    }

    return false
};

console.log(isPerfectSquare1(16) === true)
console.log(isPerfectSquare1(14) === false)
console.log(isPerfectSquare1(1) === true)
console.log(isPerfectSquare2(16) === true)
console.log(isPerfectSquare2(14) === false)
console.log(isPerfectSquare2(1) === true)
