package bloom_filter

import (
	"fmt"
	"github.com/howeyc/crc16"
)

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
	idx := bf.getHashIdx(s)
	fmt.Println(idx)

	for _, i := range idx {
		bf.array[i] = true
	}

	bf.cnt += 1

	// 无法扩容... 因为没有保存字符串. 所以需要预先就预测好空间
	// TODO: 处理扩容逻辑, 保证精确度
	//if float32(bf.m) * math.Ln2 * math.Ln2 / float32(bf.cnt) < 0.8 {
	//	bf.expand()
	//}
}

func (bf *BloomFilter) Get(s string) bool {
	idx := bf.getHashIdx(s)

	bit := true
	for _, i := range idx {
		bit = bit && bf.array[i]
	}
	return bit
}

func (bf BloomFilter) getHashIdx(s string) []int {
	funcs := []func(data []byte)uint16 {crc16.ChecksumCCITT, crc16.ChecksumCCITTFalse,
		crc16.ChecksumIBM, crc16.ChecksumMBus, crc16.ChecksumSCSI}

	idx := make([]int, 5)
	for i, fu := range funcs {
		idx[i] = int(fu([]byte(s))) % bf.m
	}
	return idx
}

/*
这些hash算法都不好用... 垃圾. 我直接调用 crc16库了

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
 */
