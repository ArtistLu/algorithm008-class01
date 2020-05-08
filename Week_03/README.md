### 递归
```python
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
```go
func recursion(level int, param int) {
  // terminator  递归终结条件
	if level > MAX_LEVEL {
		//process result
		return
	}

	//process current logic 处理当前层逻辑
	process(level, param)

	//drill down 下探到下一层
 	recursion(level + 1, neqParsm)
	//restore current status 清理当前层
}
```
#### 三个思维要点
1. 不要人肉递归（最大误区）
2. 找最近最简方法，将其拆解成可重复解决的问题（重复子问题）
3. 数学归纳法思维

### 分治、回溯
```python
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
#### 分治可以理解成一种特殊的递归

### 每日练习

#### day1
- [x] [day1-1](https://leetcode-cn.com/problems/two-sum/)
- [ ] [day1-2](https://leetcode-cn.com/problems/3sum/)
- [ ] [day1-3](https://leetcode-cn.com/problems/get-kth-magic-number-lcci/)
#### day2
- [x] [day2-1](https://leetcode-cn.com/problems/add-strings/)
- [x] [day2-2](https://leetcode-cn.com/problems/number-of-burgers-with-no-waste-of-ingredients/)
- [ ] [day2-3](https://leetcode-cn.com/problems/spiral-matrix/)

#### day3
- [x] [day3-1](https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/)
- [x] [day3-2](https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/)
- [x] [day3-3](https://leetcode-cn.com/problems/er-cha-shu-de-zui-jin-gong-gong-zu-xian-lcof/)

#### day4

- [x] [day4-1](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/submissions/)
- [x] [day4-2](https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/)
#### day5

- [x] [day5-1](https://leetcode-cn.com/problems/number-of-islands/)
  
#### day6
- [x] [day6-1](https://leetcode-cn.com/problems/3sum/)
#### day7
- [x] [day7-1](https://leetcode-cn.com/problems/valid-anagram/description/)






