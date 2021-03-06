# 问题
1.  new和make的区别
    
    new和make都在堆上分配内存，但是它们的行为不同，适用于不同的类型。
    
    new(T) 为每个新的类型T分配一片内存，初始化为 0 并且返回类型为*T的内存地址：这种方法 返回一个指向类型为 T，值为 0 的地址的指针，它适用于值类型如数组和结构体；它相当于 &T{}。
    
    make(T) 返回一个类型为 T 的初始值，是三个引用类型本身，它只适用于3种内建的引用类型：slice、map 和 channel。
    
    换言之，new 函数分配内存，make 函数初始化；
2.  解决hash冲突的方法
    1.  开放定址法（线性探测再散列，二次探测再散列，伪随机探测再散列）
    2.  再哈希法
    3.  链地址法
    4.  建立一个公共溢出区
3.  tcp与udp的区别
    https://blog.csdn.net/qq_34624951/article/details/82669444
4. b+树相比于b树的查询优势：
   
   b树（balance tree）和b+树应用在数据库索引，可以认为是m叉的多路平衡查找树，但是从理论上讲，二叉树查找速度和比较次数都是最小的，为什么不用二叉树呢？ 
   
   因为我们要考虑磁盘IO的影响，它相对于内存来说是很慢的。数据库索引是存储在磁盘上的，当数据量大时，就不能把整个索引全部加载到内存了，只能逐一加载每一个磁盘页（对应索引树的节点）。所以我们要减少IO次数，对于树来说，IO次数就是树的高度，而“矮胖”就是b树的特征之一，它的每个节点最多包含m个孩子，m称为b树的阶，m的大小取决于磁盘页的大小
   1.   b+树的中间节点不保存数据，所以磁盘页能容纳更多节点元素，更“矮胖”；
   2.   b+树查询必须查找到叶子节点，b树只要匹配到即可不用管元素位置，因此b+树查找更稳定（并不慢）；
   3.   对于范围查找来说，b+树只需遍历叶子节点链表即可，b树却需要重复地中序遍历