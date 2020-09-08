# Trees

# What is a Tree?
> A tree is a data structure similar to a linked list but instead of each node pointing simply to the next node in a linear fashion, each node points to a number of nodes. Tree is an example of a non-linear data structure. A tree structure is a way of representing the hierarchical nature of a structure in a graphical form.

# Glossary

- The root of a tree is the node with no parents. There can be at most one root node in a tree(node A in the above example).
- An edge refers to the link from parent to child (all links in the figure).
- A node with no children is called leaf node(E, J, K, H and I).
- Children of same parent are called sibling (B,C,D are siblings of A, and E, F are siblings of B).
- A node p is an ancestor of node q if there exists a path from root to q and p appears on the path. The node q is called a descendant of p. For example, A, C and G are the ancestors of K.
- The set of all nodes at a given depth is called the level of the tree(B, C and D are the same level). The root node is at level zero.
- The depth of a node is the length of the path from the root to the node(depth of G is 2, A-C-G).
- The height of a node is the length of the path from that node to deepest node. The height of a tree is the length of the path from the root to the deepest node in the tree. A(rooted) tree with only one node(the root) has a height of zero. In the previous example, the height of B is 2(B-F-J).
- Height of the tree is the maximum height among all the nodes in the tree and depth of the tree is the maximum depth among all the nodes in the tree. For a given tree, depth and height returns the same value. But for individual nodes we may get different results.
- The size of a node is the number of descendants it has including itself(the size of the subtree C is 3).
- If every node in a tree has only one child(except leaf nodes)then we call such trees skew trees.  If every node has only left child then we call them left skew trees. Similarly, if every node has only right child then we call them right skew trees.

# Binary Trees

A tree is called binary tree if each node has zero child, one child or two children. Empty tree is also a valid binary tree. We can visualize a binary tree as consisting of a root and two disjoint binary trees, called the left and right subtrees of the root.

# Types of Binary Trees

- Strict Binary Tree: A binary tree is called strict binary tree if each node has exactly two children or no children.
- Full Binary Tree: A binary tree is called full binary tree if each node has exactly two children and all leaf node are the same level.
- Complete Binary Tree: Before defining the complete binary tree, let us assume that the height of the binary tree is h. In complete binary trees, if we give numbering for the nodes by starting at the root (let us say the root node has 1)then we get a complete sequence from 1 to the number of nodes in the tree.

# Structure of Binary Trees
> Now let us define structure of the binary tree. For simplicity, assume that the data of the nodes are integers. One way to represent a node(which contains data) is to have two links which point to left and right children along with data fields as shown below.

# Operations on Binary Trees

Basic Operations

- Inserting an element into a tree
- Deleting an element from a tree
- Searching for an element
- Traversing the tree

Auxiliary Operations

- Finding the size of the tree
- Finding the height of the tree
- Finding the level which has maximum sum
- Finding the least common ancestor(LCA) for a given pair of nodes, and many more

*Applications of Binary Trees*

- Expression trees are used in compilers
- Huffman coding trees that are used in data compression algorithms
- Binary Search Tree(BST), which supports search, insertion and deletion on a collection of items in O(logn)(average)
- Priority Queue(PQ), which supports search and deletion of minimum (or maximum) on a collection of items in logarithmic time(in worst case).

# Binary Tree Traversals
> In order to process trees, we need a mechanism for traversing them, and that forms the subject of this section.
The process visiting all nodes of a tree is called tree traversal. Each node is processed only but it may be visited more than once. 

Traversal Possibilities

Starting at the root of a binary tree, there are three main steps that can be performed and the order in which they are performed defines the traversal type.

These steps are: performing an action on the current node(referred to as "visiting" the node and denoted with "D"), traversing to the left child node(denoted with "L"), and traversing to the right child node (denoted with "R"). This process can be easily described through recursion.

Based on the above definition there are 6 possibilities:

- 1、LDR: Process left subtree, process the current node data and then process right subtree
- 2、LRD: Process left subtree, process right subtree and then process the current node data
- 3、DLR: Process the current node data, process left subtree and then process right subtree
- 4、DRL: Process the current node data, process right subtree and then process left subtree
- 5、RDL: Process right subtree, process the current node data and then process left subtree
- 6、RLD: Process right subtree, process left subtree and then process the current node data

# Classifying the Traversals

- PreOrder(DLR) Traversal
- InOrder(LDR) Traversal
- PostOrder(LRD) Traversal

There is another traversal method which does not depend on the above orders and it is:
> Level Order Traversal: This method is inspired from Breadth First Traversal(BFS of Graph algorithms).

