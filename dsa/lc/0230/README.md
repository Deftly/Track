# 230 - Kth Smallest Element in a BST
Given the `root` of a binary search tree, and an integer `k`, return the <code>k<sup>th</sup></code> smallest value(1-indexed) of all the values of the nodes in the tree.

**Example 1:**
![ex1](./assets/kthtree1.jpg)
```
Input: root = [3,1,4,null,2], k = 1
Output: 1
```

**Example 2:**
![ex2](./assets/kthtree2.jpg)
```
Input: root = [5,3,6,2,4,null,null,1], k = 3
Output: 3
```

**Constraints:**
- The number of nodes in the tree is `n`.
- <code>1 <= k <= 10<sup>4</sup></code>
- <code>0 <= Node.val <= 10<sup>4</sup></code>

**Follow up:** If the BST is modified often(i.e. we can do insert and delete operations) and you need to find the kth smallest frequently, how would you optimize?
