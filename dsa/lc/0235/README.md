# 235 - Lowest Common Ancestor of a Binary Search Tree
Given a binary search tree(BST), find the lowest common ancestor(LCA) node of two given nodes in the BST.

The LCA of two nodes `p` and `q` in a tree is defined as the lowest node in `T` that has both `p` and `q` as descendants(We allow a node to be a descendant of itself).

**Example 1:**
![ex1](./assets/ex1.png)
```
Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
Output: 6
Explanation: The LCA of nodes 2 and 8 is 6.
```

**Example 2:**
![ex2](./assets/ex2.png)
```
Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
Output: 2
Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant of itself according to the LCA definition.
```

**Constraints:**
- The number of nodes in the tree is in the range <code>[2, 10<sup>5</sup>]</code>
- <code>-10<sup>9</sup> <= Node.val <= 10<sup>9</sup></code>
- All `Node.val` are unique.
- `p != q`
- `p` and `q` will exist in the BST.
