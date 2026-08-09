function rob(nums: number[]): number {
    let prev = 0;
    let curr = 0;

    for (const num of nums) {
        const next = Math.max(curr, prev + num);
        prev = curr;
        curr = next;
    }

    return curr;
}

console.log(rob([1, 2, 3, 1]) === 4);
console.log(rob([2, 7, 9, 3, 1]) === 12);