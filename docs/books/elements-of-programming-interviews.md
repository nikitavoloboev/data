---
title: Elements of programming interviews
---

# [Elements of programming interviews](http://goodreads.com/book/show/16253895)

## Chapter 1 - strategies for a great interview

- Once you have an algorithm, it is important to present it in a clear manner. Your solution will be much simpler if you take advantage of libraries such as Java Collections or C++ Boost.
- Important that you use the language you are most comfortable with.
- Master the libraries, especially the data structures.
  - Do not waste time and lose credibility trying to remember how to pass an explicit comparator to a BST constructor.
  - Remember that a hash function should use exactly those fields which are used in the equality check. A comparison function should be transitive.
- Focus on the top-level algorithm:
  - It's OK to use functions that you will implement later. This will let you focus on the main part of the algorithm, will penalize you less if you don't complete the algorithm. (Hash, equals, and compare functions are good candidates for deferred implementation.) Specify that you will handle main algorithm first, then comer cases. Add TODO comments for portions that you want to come back to.
- Manage the whiteboard:
  - You will likely use more of the board than you expect, so start at the top-left comer.
  - Make use of functions skip implementing anything that's trivial (e.g., finding the maximum of an array) or standard (e.g., a thread pool).
  - Have a convention for identifiers, e.g., i , j , k for array indices, A , B , C for arrays, u , v , w for vectors, s for a String, sb for a StringBuilder, etc.
- Assume valid inputs.
- Test for comer cases.
  - For many problems, your general idea may work for most valid inputs but there may be pathological valid inputs where your algorithm (or your implementation of it) fails.
    - For example, your binary search code may crash if the input is an empty array; or you may do arithmetic without considering the possibility of overflow.
  - It is important to systematically consider these possibilities. If there is time, write unit tests.
  - Small, extreme, or random inputs make for good stimuli. Don't forget to add code for checking the result. Occasionally, the code to handle obscure corner cases may be too complicated to implement in an interview setting. If so, you should mention to the interviewer that you are aware of these problems, and could address them if required.
- Syntax
  - Interviewers rarely penalize you for small syntax errors since modern IDE excel at handling these details.
    - However, lots of bad syntax may result in the impression that you have limited coding experience. Once you are done writing your program, make a pass through it to fix any obvious syntax errors before claiming you are done.
- Know your interviewers & the company.
-
