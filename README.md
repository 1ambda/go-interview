# go-interview

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

# Credits

- http://bigocheatsheet.com/
