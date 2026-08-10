# 💡 LeetCode

[![Dashboard](https://img.shields.io/badge/Dashboard-coding--challenges-blue?style=for-the-badge)](https://github.com/LorranSutter/coding-challenges) <!-- SUMMARY:START -->[![Solved Challenges](https://img.shields.io/badge/Solved%20Challenges-130-brightgreen?style=for-the-badge&logo=leetcode&logoColor=white)](https://leetcode.com/)<!-- SUMMARY:END -->

This repository contains my solutions for [LeetCode](https://leetcode.com/) problems.

LeetCode is a platform featuring coding challenges designed to improve algorithmic thinking and prepare for technical interviews.

## 🛠️ Setup

Ensure you have [Go](https://go.dev/) installed (version 1.24+ is recommended).

Some problems are solved in TypeScript instead. Ensure you have [Node.js](https://nodejs.org/) installed (version 23.6+ is recommended) — Node runs `.ts` files directly, so no `ts-node`, `tsx`, or `package.json` is needed.

No additional installation or virtual environments are needed.

## ✨ Creating a New Problem

To create a new problem structure, use the `create_problem.sh` script:

```bash
./create_problem.sh <number>       # Go (default)
./create_problem.sh <number> --ts  # TypeScript
```

Example:
```bash
./create_problem.sh 1
./create_problem.sh 2 --ts
```

This will create a folder structure (e.g. `0001/`) with either `main.go` or `main.ts` as a starter template for the solution.

## 🚀 Running Solutions

You can run solutions directly:

```bash
go run ./0001/main.go
node ./0002/main.ts
```

Replace the folder number with the specific problem you want to execute.

## 🔄 Updating Progress Summary

To update the progress summary in this README after solving new problems, run the `generate_readme.py` script:

```bash
python3 generate_readme.py
```
