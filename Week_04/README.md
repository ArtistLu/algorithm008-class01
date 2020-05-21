### DFS & BFS

#### DFS 递归模板

```python
visited = set() 

def dfs(node, visited):
    if node in visited: # terminator
    	# already visited 
    	return 

	visited.add(node) 

	# process current node here. 
	...
	for next_node in node.children(): 
		if next_node not in visited: 
			dfs(next_node, visited)
```
#### DFS 非递归模板

```python
def DFS(self, tree): 

	if tree.root is None: 
		return [] 

	visited, stack = [], [tree.root]

	while stack: 
		node = stack.pop() 
		visited.add(node)

		process (node) 
		nodes = generate_related_nodes(node) 
		stack.push(nodes) 

	# other processing work 
	...
```

#### BFS 模板

```python
def BFS(graph, start, end):
    visited = set()
	queue = [] 
	queue.append([start]) 

	while queue: 
		node = queue.pop() 
		visited.add(node)

		process(node) 
		nodes = generate_related_nodes(node) 
		queue.push(nodes)

	# other processing work 
	...
```
### 贪心算法
- 贪心：当下做局部最优判断
- 回溯：能够回退
- 动态规划：最优判断+回退
### 二分查找
- 目标函数的单调性	
- 存在上下界
- 能够通过索引访问
#### 二分查找模板
```python
left, right = 0, len(array) - 1 
while left <= right: 
	  mid = (left + right) / 2 
	  if array[mid] == target: 
		    # find the target!! 
		    break or return result 
	  elif array[mid] < target: 
		    left = mid + 1 
	  else: 
		    right = mid - 1
```
### 每日练习

#### day1

- [x] [day1-1](https://leetcode-cn.com/problems/climbing-stairs/)
- [x] [二叉树前序遍历](https://leetcode-cn.com/problems/binary-tree-preorder-traversal/)
- [x] [二叉树层序遍历](https://leetcode-cn.com/problems/binary-tree-level-order-traversal/)
- [x] [二叉树后序遍历](https://leetcode-cn.com/problems/binary-tree-postorder-traversal/)
- [x] [二叉树中序遍历](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)
- [x] [341. 扁平化嵌套列表迭代器](https://leetcode-cn.com/problems/flatten-nested-list-iterator/)

#### day2

- [x] [433. 最小基因变化](https://leetcode-cn.com/problems/minimum-genetic-mutation/)

#### day3
- [x] [529. 扫雷游戏](https://leetcode-cn.com/problems/minesweeper/description/)

#### day4
- [x] [169. 多数元素](https://leetcode-cn.com/problems/majority-element/description/)

#### day5
- [x] [面试题49. 丑数](https://leetcode-cn.com/problems/chou-shu-lcof/)

#### day6
- [x] [51. N皇后](https://leetcode-cn.com/problems/n-queens/)

### day29
- [x] [102. 二叉树的层序遍历](https://leetcode-cn.com/problems/binary-tree-level-order-traversal/#/description)

### day30
- [x] [69. x 的平方根](https://leetcode-cn.com/problems/sqrtx/)
- [ ] [153. 寻找旋转排序数组中的最小值](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/)

### day31
- [x] [455. 分发饼干](https://leetcode-cn.com/problems/assign-cookies/description/)

### day32
- [x] [33. 搜索旋转排序数组](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/)

### day33
- [x] [45. 跳跃游戏 II](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/)
  
### day34
- [x] [70. 爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/)
  
### day35
- [x] [142. 环形链表 II](https://leetcode.com/problems/linked-list-cycle-ii/)
  
### day36
- [x] [面试题62. 圆圈中最后剩下的数字 **复习**](https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/)

### day37
- [x] [860. 柠檬水找零 **复习**](https://leetcode-cn.com/problems/lemonade-change/description/)

### day38
- [x] [127. 单词接龙](https://leetcode-cn.com/problems/word-ladder/description/)

### day39
- [x] [51. N皇后](https://leetcode-cn.com/problems/n-queens/)