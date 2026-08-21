# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

This is a personal LeetCode solutions repository. Each problem lives in its own numbered folder (4-digit, zero-padded — e.g. `0001/`, `2996/`) at the repo root, matching the problem's LeetCode number.

## Commands

- Create a new problem folder: `./create_problem.sh <number>` (Go, default), `./create_problem.sh <number> --ts` (TypeScript), or `./create_problem.sh <number> --py` (Python). If the folder already exists, it adds the requested language's file to it (e.g. adding `--py` to a folder that already has `main.go`); it still refuses to overwrite an existing file for that language.
- Run a Go solution: `go run ./<folder>/main.go` (e.g. `go run ./0001/main.go`)
- Run a TypeScript solution: `node ./<folder>/main.ts` (Node 23.6+ runs `.ts` directly — no ts-node/tsx/package.json)
- Run a Python solution: `python3 ./<folder>/main.py`
- Regenerate the solved-count badge in README.md: `python3 generate_readme.py` (scans root for `\d{4}` folders and rewrites the block between `<!-- SUMMARY:START -->` / `<!-- SUMMARY:END -->`)

There is no test suite, linter, or build step — correctness is verified by each solution file's own `main()`/top-level call into the shared per-language test runner (see below), printed to stdout.

## Solution file conventions

Each `main.go` / `main.ts` / `main.py` is self-contained and executable directly:

- The solution function/method is defined at the top (in Go/TypeScript, a plain function; in Python, a `Solution` class matching LeetCode's submission format).
- `main()` (Go) or top-level statements (TS/Python) pass the sample cases from the LeetCode problem statement, as input/expected-output pairs, into the shared test runner (`utils.RunTests` / `runTests` / `run_tests`) — there are no `_test.go` files or testing frameworks involved.
- Each solution imports the shared `utils` package/module for its language: Go via `leetcode/utils` (module name is `leetcode`, see `go.mod`), TypeScript via a relative import of `utils/utils.ts`, Python via `sys.path` manipulation to import `utils/utils.py` (see the scaffold template in `create_problem.sh` for the exact boilerplate).

## Shared utilities (`utils/`)

Per-language helper files, one per language, used across solutions:

- `utils/utils.go` (`package utils`, imported as `leetcode/utils`): `ListNode`, `TreeNode` structs and constructors `MakeList`, `MakeBinaryTree`, `MakeBinaryTreeFromLevelOrder`; `PrintList`, `PrintTree`, `PrintMatrix` for debugging output; `EqualSlices` (order-independent), `DeepEqualSlices` (order-dependent), `EqualListNodes`, `DeepEqualMatrix` for standalone comparisons; and `RunTests[R any](cases []TestCase[R])`, the test runner — each `TestCase` carries `Input` (for reporting), `Got` (the result of calling the solution, computed inline by the caller since Go generics can't unify arbitrary argument lists), and `Expected`, compared via `reflect.DeepEqual`.
- `utils/utils.ts`: `runTests(fn, cases)`, where each case is `{ input: Args, expected: R }` (`input` is the positional-argument tuple passed to `fn`), compared via Node's `util.isDeepStrictEqual`.
- `utils/utils.py`: `run_tests(fn, cases)`, where each case is `{"input": [...], "expected": ...}` (`input` is unpacked as positional args to `fn`), compared via `!=`.

All three runners follow the same contract: print one success line if every case passes, otherwise print only the failing cases with their input, expected output, and actual output.

## Adding a new solution

1. Run `./create_problem.sh <number>`, `./create_problem.sh <number> --ts`, or `./create_problem.sh <number> --py` to scaffold the folder — the generated file already wires up the language's test runner with a placeholder case.
2. Implement the solution, replacing the placeholder function/method.
3. Replace the placeholder test case(s) with the problem's actual sample input/expected-output pairs.
4. Run `python3 generate_readme.py` to update the solved-count badge.
