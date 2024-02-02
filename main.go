package main

import (
	"fmt"
	"hash/fnv"
)

type BloomFilter struct {
	bitArray []bool
	hashList []BloomFilterHash
}

type BloomFilterHash func(s string) uint32

func NewBloomFilter(size uint32) *BloomFilter {
	return &BloomFilter{
		bitArray: make([]bool, size),
	}
}

func (b *BloomFilter) RegisterHash(h ...BloomFilterHash) {
	b.hashList = append(b.hashList, h...)
}

func (b *BloomFilter) Add(s string) {
	bitLen := uint32(len(b.bitArray))
	for _, h := range b.hashList {
		index := h(s) % bitLen
		b.bitArray[index] = true
	}
}

func (b *BloomFilter) Contains(s string) bool {
	bitLen := uint32(len(b.bitArray))
	for _, h := range b.hashList {
		index := h(s) % bitLen
		if !b.bitArray[index] {
			return false
		}
	}
	return true
}

func hash1(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func hash2(s string) uint32 {
	h := fnv.New32()
	h.Write([]byte(s))
	return h.Sum32()
}

func main() {
	var data = "www.baidu.com"

	bloom := NewBloomFilter(1 << 16)
	bloom.RegisterHash(hash1, hash2)

	bloom.Add(data)

	fmt.Println(bloom.Contains(data))
}
