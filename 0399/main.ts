import { runTests } from "../utils/utils.ts"

/**
 * This was a nice one. It's modeled as a weighted graph problem, which isn't obvious from the first couple of
 * examples, since those only ask about pairs that already appear directly in `equations`. It only turns into a
 * real graph problem once a query asks for the ratio between two variables that aren't directly connected.
 *
 * Preprocessing:
 * - The key insight is that each equation `Ai / Bi = values[i]` is a weighted, directed edge from `Ai` to
 *   `Bi`. Since division is invertible, we add the reverse edge `Bi -> Ai` at the same time, weighted
 *   `1 / values[i]`, so both `a/b` and `b/a` are a single lookup away without having to search twice.
 * - For `a/b = x` and `b/c = y` this builds:
 *
 *      /<-- 1/x ---\ /<-- 1/y ---\
 *     a ---- x ---> b ---- y ---> c
 *
 * Query resolution:
 * - If either variable in a query never showed up in any equation, the answer can't be determined and we
 *   return `-1` right away without searching.
 * - Otherwise, answering `Cj / Dj` becomes a path-finding problem: multiplying the edge weights along any
 *   path from `Cj` to `Dj` gives the answer, since each hop is itself a ratio and the chain telescopes back
 *   into a single division. For `a/b = x` and `b/c = y`, finding `a/c` works out to:
 *
 *      b = c*y
 *      a/b = a/(c*y) = x
 *      a/c = x*y
 *
 *   which is exactly what multiplying the edge weights along the path a -> b -> c gives us directly.
 * - The DFS from `source` walks every unvisited neighbor, multiplying the running product by each edge's
 *   weight, and stops as soon as it reaches `dest`. The `visited` set keeps it from doubling back over an
 *   edge already used on the current path.
 * - There is an advantage here: every time the DFS reaches a node that isn't already directly connected to
 *   `source`, it adds a new edge from `source` to that node with the product computed so far. So resolving
 *   `a/c` for the first time also wires up a direct `a <-> c` edge:
 *
 *      a --- x --> b --- y --> c
 *       \-------- x*y ------->/
 *        \<---- 1/(x*y) ----/
 *
 *   This isn't limited to the final destination — every intermediate node the DFS passes through on the way
 *   gets the same shortcut back to `source`. That means later queries sharing any of those variables with an
 *   earlier query become an O(1) lookup instead of a fresh traversal, which pays off when queries repeat or
 *   overlap on the same variables.
 */

class Graph<T> {
    private adjacencyList: Map<T, Map<T, number>>;

    constructor() {
        this.adjacencyList = new Map<T, Map<T, number>>();
    }

    addVertex(vertex: T): void {
        if (!this.adjacencyList.has(vertex)) {
            this.adjacencyList.set(vertex, new Map<T, number>());
        }
    }

    hasVertex(vertex: T): boolean {
        return this.adjacencyList.has(vertex);
    }

    addEdge(source: T, destination: T, weight: number): void {
        this.addVertex(source);
        this.addVertex(destination);

        this.adjacencyList.get(source)!.set(destination, weight);
        this.adjacencyList.get(destination)!.set(source, 1 / weight);
    }

    getNeighbors(vertex: T): Map<T, number> {
        return this.adjacencyList.get(vertex) || new Map<T, number>();
    }

    isConnected(source: T, destination: T): boolean {
        return this.adjacencyList.get(source)?.has(destination) ?? false;
    }

    printGraph(): void {
        for (const [vertex, edges] of this.adjacencyList.entries()) {
            const edgeStrings = [...edges.entries()].map(([node, weight]) => `${vertex} -> ${node} (w: ${weight})`);
            console.log(`${vertex}: [ ${edgeStrings.join(', ')} ]`);
        }
    }
}


function calcEquation(equations: string[][], values: number[], queries: string[][]): number[] {
    const graph = new Graph<string>()
    for (let i = 0; i < equations.length; i++) {
        graph.addEdge(equations[i][0], equations[i][1], values[i])
    }
    // graph.printGraph()
    const results: number[] = new Array(queries.length).fill(-1.0)

    for (let i = 0; i < queries.length; i++) {
        const source = queries[i][0]
        const dest = queries[i][1]
        if (!graph.hasVertex(source) || !graph.hasVertex(dest)) {
            continue
        }

        const visited = new Set<string>()
        function dfs(node: string, result: number) {
            if (node === dest) {
                results[i] = Number(result.toFixed(5))
                return
            }
            if (!graph.isConnected(source, node)) {
                graph.addEdge(source, node, result)
            }

            visited.add(node);

            for (const n of graph.getNeighbors(node)) {
                if (!visited.has(n[0])) {
                    dfs(n[0], result * n[1]);
                }
            }
        }

        dfs(source, 1.0)
    }

    return results
}

runTests(calcEquation, [
    { input: [[["a", "b"], ["b", "c"]], [2.0, 3.0], [["a", "c"], ["b", "a"], ["a", "e"], ["a", "a"], ["x", "x"]]], expected: [6.00000, 0.50000, -1.00000, 1.00000, -1.00000] },
    { input: [[["a", "b"], ["b", "c"], ["bc", "cd"]], [1.5, 2.5, 5.0], [["a", "c"], ["c", "b"], ["bc", "cd"], ["cd", "bc"]]], expected: [3.75000, 0.40000, 5.00000, 0.20000] },
    { input: [[["a", "b"]], [0.5], [["a", "b"], ["b", "a"], ["a", "c"], ["x", "y"]]], expected: [0.50000, 2.00000, -1.00000, -1.00000] },
    { input: [[["a", "b"], ["c", "b"], ["d", "b"], ["w", "x"], ["y", "x"], ["z", "x"], ["w", "d"]], [2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0], [["a", "c"], ["b", "c"], ["a", "e"], ["a", "a"], ["x", "x"], ["a", "z"]]], expected: [0.66667, 0.33333, -1.00000, 1.00000, 1.00000, 0.04464] },
    { input: [[["x1", "x2"], ["x2", "x3"], ["x3", "x4"], ["x4", "x5"]], [3.0, 4.0, 5.0, 6.0], [["x1", "x5"], ["x5", "x2"], ["x2", "x4"], ["x2", "x2"], ["x2", "x9"], ["x9", "x9"]]], expected: [360.00000, 0.00833, 20.00000, 1.00000, -1.00000, -1.00000] },
    { input: [[["ab", "cd"], ["a", "c"]], [4.0, 2.0], [["b", "d"]]], expected: [-1.00000] },
    { input: [[["a", "b"], ["c", "b"], ["d", "b"], ["w", "x"], ["y", "x"], ["z", "x"], ["w", "d"]], [2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0], [["a", "c"], ["b", "c"], ["a", "e"], ["a", "a"], ["x", "x"], ["a", "z"]]], expected: [0.66667, 0.33333, -1.00000, 1.00000, 1.00000, 0.04464] },
    { input: [[["a", "b"], ["b", "c"], ["a", "c"]], [2.0, 3.0, 6.0], [["a", "c"], ["b", "a"], ["a", "e"], ["a", "a"], ["x", "x"]]], expected: [6.00000, 0.50000, -1.00000, 1.00000, -1.00000] },
])
