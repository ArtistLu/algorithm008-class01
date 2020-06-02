### 字典树 trie

#### 字典树数据结构
- 字典树，即 Trie 树，又称单词查找树或键树，是一种树形结构。典型应用于统计和排序大量的字符串（但不仅限于字符串）所以经常被搜索引擎系统用于文本词频统计。
- 优点：最大限度的减少无谓的字符串比较，查询效率比hash表高
#### 字典树核心思想
- 空间换时间
- 利用字符串公共前缀降低查询时间开销，以达到提高效率目的
#### 字典树的基本性质
- 节点本身不存完整单词
- 从根结点到某一结点，路径上经过的结点连接起来，为该结点对应的字符串
- 每个结点的所有子结点路径代表的字符都不相同

### trie 树模板
```py
class Trie(object):
  
	def __init__(self): 
		self.root = {} 
		self.end_of_word = "#" 
 
	def insert(self, word): 
		node = self.root 
		for char in word: 
			node = node.setdefault(char, {}) 
		node[self.end_of_word] = self.end_of_word 
 
	def search(self, word): 
		node = self.root 
		for char in word: 
			if char not in node: 
				return False 
			node = node[char] 
		return self.end_of_word in node 
 
	def startsWith(self, prefix): 
		node = self.root 
		for char in prefix: 
			if char not in node: 
				return False 
			node = node[char] 
		return True
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

### day40
- [x] [429. N叉树的层序遍历](https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/)

### day41
- [x] [169. 多数元素](https://leetcode-cn.com/problems/majority-element/description/)

### day42
- [x] [47. 全排列 II](https://leetcode-cn.com/problems/permutations-ii/)

### day43
- [x] [198. 打家劫舍](https://leetcode-cn.com/problems/house-robber/)

### day44
- [x] [322. 零钱兑换](https://leetcode-cn.com/problems/coin-change/)

### day45
- [x] [120. 三角形最小路径和](https://leetcode-cn.com/problems/triangle/)

### day46
- [x] [77. 组合](https://leetcode-cn.com/problems/combinations/)

### day47
- [x] [312. 戳气球](https://leetcode-cn.com/problems/burst-balloons/)

### day48
- [x] [221. 最大正方形](https://leetcode-cn.com/problems/maximal-square/)

  ### day49
- [x] [621. 任务调度器](https://leetcode-cn.com/problems/task-scheduler/)

### day50
- [x] [94. 二叉树的中序遍历](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)
