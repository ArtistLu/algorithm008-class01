
#### 递归复习

递归模板
```py
# Python
def recursion(level, param1, param2, ...): 
    # recursion terminator 
    if level > MAX_LEVEL: 
	   process_result 
	   return 

    # process logic in current level 
    process(level, data...) 

    # drill down 
    self.recursion(level + 1, p1, ...) 

    # reverse the current level status if needed

```
分治模板
```py
# Python
def divide_conquer(problem, param1, param2, ...): 
  # recursion terminator 
  if problem is None: 
	print_result 
	return 

  # prepare data 
  data = prepare_data(problem) 
  subproblems = split_problem(problem, data) 

  # conquer subproblems 
  subresult1 = self.divide_conquer(subproblems[0], p1, ...) 
  subresult2 = self.divide_conquer(subproblems[1], p1, ...) 
  subresult3 = self.divide_conquer(subproblems[2], p1, ...) 
  …

  # process and generate the final result 
  result = process_result(subresult1, subresult2, subresult3, …)
	
  # revert the current level states
```
#### 感触
1. 不要人肉递归（低效）
2. 找到最近简单的方法，将其拆解成可重复解决的问题
3. 数学归纳法（n=1,n=2....）

- 本质：寻找重复性->计算机指令集
- 如果还是想不清除，画递归树
  

### 动态规划

#### 关键点

- 动态规划和递归或分治没有本质区别（关键看有无最有子结构）
- 共性：找到重复子问题
- 差异性：最优子结构，中途可以淘汰次优解
- simplifying a complicated problem by breaking it down into simpler sub-problems


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