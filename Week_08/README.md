### 位运算

#### xor 异或
异或操作的一些特点
- x^0=x
- x^1s=~x //1s=~0
- x^(~x)=1s
- x^x=0
- c=a^b -> a^c=b b^c=a
- a^b^c=(a^b)^c=a^(b^c)

#### 指定位置的位运算
1. 将x最右边的n位清零: x & (~0<<n)
2. 获取x的第n位值（0或1）:(x>>n) & 1
3. 获取x的第n位幂值 :x&(1<<n)
4. 将n位置为1:x|(1<<n)
5. 将n位置为0:x&(~(1<<n))
6. x最高位至n位（包含）置0:x&((1<<n)-1)
#### 实战位运算要求
- 判断奇偶性
    - x % 2 == 1 ----> (x&1) == 1
    - x % 2 == 0 ----> (x&1) == 0
- x >> 1 -----------> x / 2
- x=x&(x-1)清零最低位的1
- x&-x 得到最低位的1
- x&~x == 0
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

#### day29
- [x] [102. 二叉树的层序遍历](https://leetcode-cn.com/problems/binary-tree-level-order-traversal/#/description)

#### day30
- [x] [69. x 的平方根](https://leetcode-cn.com/problems/sqrtx/)
- [ ] [153. 寻找旋转排序数组中的最小值](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/)

#### day31
- [x] [455. 分发饼干](https://leetcode-cn.com/problems/assign-cookies/description/)

#### day32
- [x] [33. 搜索旋转排序数组](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/)

#### day33
- [x] [45. 跳跃游戏 II](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/)
  
#### day34
- [x] [70. 爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/)
  
#### day35
- [x] [142. 环形链表 II](https://leetcode.com/problems/linked-list-cycle-ii/)
  
#### day36
- [x] [面试题62. 圆圈中最后剩下的数字 **复习**](https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/)

#### day37
- [x] [860. 柠檬水找零 **复习**](https://leetcode-cn.com/problems/lemonade-change/description/)

#### day38
- [x] [127. 单词接龙](https://leetcode-cn.com/problems/word-ladder/description/)

#### day39
- [x] [51. N皇后](https://leetcode-cn.com/problems/n-queens/)

#### day40
- [x] [429. N叉树的层序遍历](https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/)

#### day41
- [x] [169. 多数元素](https://leetcode-cn.com/problems/majority-element/description/)

#### day42
- [x] [47. 全排列 II](https://leetcode-cn.com/problems/permutations-ii/)

#### day43
- [x] [198. 打家劫舍](https://leetcode-cn.com/problems/house-robber/)

#### day44
- [x] [322. 零钱兑换](https://leetcode-cn.com/problems/coin-change/)

#### day45
- [x] [120. 三角形最小路径和](https://leetcode-cn.com/problems/triangle/)

#### day46
- [x] [77. 组合](https://leetcode-cn.com/problems/combinations/)

#### day47
- [x] [312. 戳气球](https://leetcode-cn.com/problems/burst-balloons/)

#### day48
- [x] [221. 最大正方形](https://leetcode-cn.com/problems/maximal-square/)

  #### day49
- [x] [621. 任务调度器](https://leetcode-cn.com/problems/task-scheduler/)

#### day50
- [x] [94. 二叉树的中序遍历](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)

#### day51
- [x] [547. 朋友圈](https://leetcode-cn.com/problems/friend-circles/)


#### day52
- [x] [51. N皇后](https://leetcode-cn.com/problems/n-queens/)

#### day53
- [x] [130. 被围绕的区域](https://leetcode-cn.com/problems/surrounded-regions/)

#### day54
- [x] [36. 有效的数独](https://leetcode-cn.com/problems/valid-sudoku/description/)

#### day55
- [x] [217. 存在重复元素](https://leetcode-cn.com/problems/contains-duplicate/)
- [x] [350. 两个数组的交集 II](https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/)

#### day56
- [x] [127. 单词接龙](https://leetcode-cn.com/problems/word-ladder/)

#### day57
- [x] [46. 全排列](https://leetcode-cn.com/problems/permutations/)

#### day58
- [x] [191. 位1的个数](https://leetcode-cn.com/problems/number-of-1-bits/)

#### day59
- [x] [1122. 数组的相对排序](https://leetcode-cn.com/problems/relative-sort-array/)