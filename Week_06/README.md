
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
  