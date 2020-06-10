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

### 并查集

#### 适用场景
- 组团配对问题
- Group or not？

#### 基本操作

- makeSet(s) : 建立一个新的并查集，其中包含s个单元素集合
- unionSet(x, y): 把 x 和 y 所在的集合合并，要求 x 和 y 所在的集合不想交，如果相交则不合并。
- find(x) : 找到x所在集合的代表，该操作也可用于判断两个元素是否在同一个集合，只要将他们各自代表比较下就可以了。

#### 代码模板

```py
# Python 

def init(p): 
	# for i = 0 .. n: p[i] = i; 
	p = [i for i in range(n)] 
 
def union(self, p, i, j): 
	p1 = self.parent(p, i) 
	p2 = self.parent(p, j) 
	p[p1] = p2 
 
def parent(self, p, i): 
	root = i
	while p[root] != root: 
		root = p[root] 
	while p[i] != i: # 路径压缩 ?
		x = i; i = p[i]; p[x] = root 
	return root
```

### 高级搜索

#### 目录

- 剪枝
- 双向 BFS
- 启发式搜索（A*）

#### 初级搜索

1. 朴素搜索
2. 优化方式：不重复（fibonacci）、剪枝（生成括号问题）
3. 搜索方向：
   - DFS
   - BFS
   - 双向搜索，启发式搜索
