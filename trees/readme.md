Binary Search Tree

Perfect Binary Tree
All the leaf nodes are full. There is no node with only one child.
1. Number of nodes at each level doubles as we move down the tree.
2. Number of nodes on the last level is equal to the sum of the
   number of nodes on the all the other levels +1.
3. At each level we can say number of nodes
   Level 0 = 2^0 = 1 - 1 = 0
   Level 1 = 2^1 = 2 - 1 = 1
   Level 2 = 2^2 = 4 - 1 = 3
   Level 3 = 2^3 = 8 - 1 = 7
   Level n = 2^(n-1) = n

   # of nodes = 2^h -1 
   log nodes = height
   log 100 = 2
   10^2 = 100

Suppose n = 16:
16 ➔ 8 ➔ 4 ➔ 2 ➔ 1

How many divisions by 2? 4 times.
So,
log₂(16) = 4

✅ Because 2⁴ = 16.

📈 Visual intuition:
If you had 1 million (10⁶) nodes:

log₂(1,000,000) ≈ 20

2pow20 = 1,000,000

Only about 20 steps needed to find any element!

Way faster than O(n) (linear search).

Full Binary Tree
All the nodes have either 0 or 2 children. There is no node with only one child.

Complete Binary Tree
All the leaf nodes are at the same level. All the nodes are as far left as possible. The last level may not be full.



