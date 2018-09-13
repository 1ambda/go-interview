# go-interview

## Articles

- Leetcode Patterns
    * [BST](https://medium.com/leetcode-patterns/leetcode-pattern-0-iterative-traversals-on-trees-d373568eb0ec)
    * [DFS + BFS Part 1](https://medium.com/leetcode-patterns/leetcode-pattern-1-bfs-dfs-25-of-the-problems-part-1-519450a84353)
    * [DFS + BFS Part 2](https://medium.com/leetcode-patterns/leetcode-pattern-2-dfs-bfs-25-of-the-problems-part-2-a5b269597f52)
    * [Sliding Windows for String](https://medium.com/leetcode-patterns/leetcode-pattern-2-sliding-windows-for-strings-e19af105316b)
    * [Backtracking](https://medium.com/leetcode-patterns/leetcode-pattern-3-backtracking-5d9e5a03dc26)

## Sorting Algorithms

- [x] Bubble Sort
- [x] Selection Sort
- [x] Insertion Sort
    * [Selection Sort vs Insertion Sort](https://stackoverflow.com/questions/15799034/insertion-sort-vs-selection-sort)

Both bubble sort and insertion sort have `Ω(n)` best case time complexity while selection sort provides `Ω(n^2)`.

- [x] Quick Sort
    * Hoare partition schema is more efficient due to less swap than Lomuto's partition scheme.
    * In case of (reverse) sorted can perform the worse case. To avoid this, Use [Pivot Selection Strategy](https://stackoverflow.com/a/7561147/3435720)
    * [Why is quick sort considered to be better than merge sort?](https://www.quora.com/Why-is-quicksort-considered-to-be-better-than-merge-sort)
        1. quick sort is in-place sort which means doesn't require additional space
        2. the `O(n2)` worst case of quick sort can easily be avoided using pivot strategy discussed above.
        3. qucik sort exhibits good cache locality than merge sort.
- [x] Merge Sort
    * merge sort is efficient for data doesn't fit into memory and when it's expensive get to an item (single linked list, file)
    * [Why is merge sort preferred over quick sort for sorting linked lists](https://stackoverflow.com/questions/5222730/why-is-merge-sort-preferred-over-quick-sort-for-sorting-linked-lists)
    * stable sort
    
## Binary Search Tree and Heap

- [x] Binary Search Tree
    * BST provides `O(logN)` operations: **Access**, **Search**, **Insert**, **Delete**
    * [Heap vs Binary Search Tree](https://stackoverflow.com/questions/6147242/heap-vs-binary-search-tree-bst)
- [ ] Self-balanced BST
- [ ] Heap
    
## Graph Algorithms

- [ ] DFS
- [ ] BFS
- [ ] Minimum Spanning Tree
- [ ] Prim's Algorithm
- [ ] Kruskal's Algorithm
- [ ] Dijkstra Algorithm 
- [ ] Bellman-Ford Algorithm
- [ ] Floyd Algorithm

# Credits

- http://bigocheatsheet.com/
