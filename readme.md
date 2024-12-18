# Advent of code 2024
[Link](https://adventofcode.com/2024)

# Notes
- The goal is to end up with code that 1. can be easily understood (at least if you have read the problem) and 2. can evolve efficiently, that is with mostly extensions and only few changes. Refactorings should be easy (or unnecessary).
- put one and only one layer of abstraction, indicated by extensive use of `type PagesUpdates [][]int` meaning that you relabel built-in types, but no structs
- Have methods on the types, a method changes only data on its own object. Here "data" refers to the potentially "big" data. For small data it is acceptable (or even desirable) not to use reference semantics. This results for big data in a object-oriented style (and for small data in a functional style). Here you could dive deep into a stack vs. heap discussion.
- Name your variables according to the problem description, that is to what you are actually modelling and that is not terms of the technical solution.
- No matter how immature the code and project phase seems, write tests.
- Programming experience differs strongly from script-like languages such as Python: There is more low-level code, but there are less ambiguities and, moreover, with static code analysis, from time to time you will experience successful runs of newly written code on first attempt.
