package bloom_filter

/*
布隆过滤器

可能会把一些没有的判断成有
但有的一定会判断成有
=> 判断成没有的那么一定是没有; 判断成有的不一定有 (有准确率)
 */

type BloomFilter struct {
	m int  // 布隆过滤器的长度（如比特数组的大小）
	k int  // 哈希的次数
	cnt int  // 已经过滤元素的数量
	array []bool
}

func NewBloomFilter() *BloomFilter {
	return &BloomFilter{m: 1024, k: 3, array: make([]bool, 1024, 1024)}
}

func (bf *BloomFilter) Put(s string) {
	h1 := bf.SDBMHash(s)
	h2 := bf.BKDRHash(s)
	//h3 := bf.RSHash(s)
	h4 := bf.JSHash(s)

	//log.Printf("%d, %d, %d, %d\n", h1, h2, h3, h4)

	bf.array[h1] = true
	bf.array[h2] = true
	//bf.array[h3] = true
	bf.array[h4] = true

	bf.cnt += 1

	// 无法扩容... 因为没有保存字符串. 所以需要预先就预测好空间
	// TODO: 处理扩容逻辑, 保证精确度
	//if float32(bf.m) * math.Ln2 * math.Ln2 / float32(bf.cnt) < 0.8 {
	//	bf.expand()
	//}
}

func (bf *BloomFilter) Get(s string) bool {
	h1 := bf.SDBMHash(s)
	h2 := bf.BKDRHash(s)
	//h3 := bf.RSHash(s)
	h4 := bf.JSHash(s)
	if bf.array[h1] && bf.array[h2] && bf.array[h4] {
		return true
	}
	return false
}

// 哈希算法参考资料: https://developer.aliyun.com/article/252773
func (bf BloomFilter) SDBMHash(s string) int {
	h := 0
	for i := 0; i < len(s); i++ {
		h = int(s[i]) + (h << 6) + (h << 16) - h
	}
	return (h& 0x7FFFFFFF) % bf.m
}

func (bf BloomFilter) BKDRHash(s string) int {
	h := 0
	seed := 131
	for i := 0; i < len(s); i++ {
		h = seed * h + int(s[i])
	}
	return (h& 0x7FFFFFFF) % bf.m
}

func (bf BloomFilter) RSHash(s string) int {
	h := 0
	b := 378551
	a := 63689
	for i := 0; i < len(s); i++ {
		h = h * a + int(s[i])
		a = a * b
	}
	return (h& 0x7FFFFFFF) % bf.m
}

func (bf BloomFilter) JSHash(s string) int {
	var h uint =  1315423911
	for i := 0; i < len(s); i++ {
		h ^= (h << 5) + uint(s[i]) + (h >> 2)
	}
	return (int(h)& 0x7FFFFFFF) % bf.m
}