# Cracking the coding interview
## Chapter 1 - arrays and strings
- hash table is a data structure that maps keys to values for highly efficient lookup. 
	- In a very simple implementation of a hash table, the hash table has an underlying array and a hash function.
	- When you want to insert an object and its key, the hash function maps the key to an integer, which indicates the index in the array. The object is then stored at that index. 
- we can implement the hash table with a binary search tree.
	- We can then guarantee an 0(log n) lookup time, since we can keep the tree balanced.
	- Additionally, we may use less space, since a large array no longer needs to be allocated in the very beginning.
- ArrayList (dynamical resizing array)
	- An ArrayList, or a dynamically resizing array, is an array that resizes itself as needed while still providing 0(1) access. 
	- A typical implementation is that when the array is full, the array doubles in size.
		- Each doubling takes 0(n) time, but happens so rarely that its amortized time is still O(1).
