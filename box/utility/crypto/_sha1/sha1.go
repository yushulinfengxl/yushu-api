package _sha1

import "encoding/binary"

// 常量初始化（SHA-1 初始哈希值）
var h0 uint32 = 0x67452301
var h1 uint32 = 0xEFCDAB89
var h2 uint32 = 0x98BADCFE
var h3 uint32 = 0x10325476
var h4 uint32 = 0xC3D2E1F0

// 循环左移
func leftRotate(value uint32, bits uint) uint32 {
	return (value << bits) | (value >> (32 - bits))
}

// 消息填充，返回填充后的字节切片
func padMessage(data []byte) []byte {
	originalLen := uint64(len(data) * 8)

	// 填充消息尾部的 "1" 位
	data = append(data, 0x80)

	// 填充 "0" 位，直到消息长度满足 512 位的倍数
	for len(data)%64 != 56 {
		data = append(data, 0x00)
	}

	// 在末尾添加消息的长度，64 位表示
	lengthBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(lengthBytes, originalLen)
	data = append(data, lengthBytes...)

	return data
}

// Sum SHA-1 主函数
func Sum(data []byte) [20]byte {
	// 消息填充
	paddedData := padMessage(data)

	// 初始化哈希值
	h0, h1, h2, h3, h4 := h0, h1, h2, h3, h4

	// 遍历每个 512 位块
	for i := 0; i < len(paddedData); i += 64 {
		block := paddedData[i : i+64]

		// 消息扩展，将块扩展为 80 个 32 位字
		var w [80]uint32
		for j := 0; j < 16; j++ {
			w[j] = binary.BigEndian.Uint32(block[j*4 : (j+1)*4])
		}
		for j := 16; j < 80; j++ {
			w[j] = leftRotate(w[j-3]^w[j-8]^w[j-14]^w[j-16], 1)
		}

		// 初始化哈希值
		a, b, c, d, e := h0, h1, h2, h3, h4

		// 80 次迭代
		for j := 0; j < 80; j++ {
			var f, k uint32
			switch {
			case j < 20:
				f = (b & c) | ((^b) & d)
				k = 0x5A827999
			case j < 40:
				f = b ^ c ^ d
				k = 0x6ED9EBA1
			case j < 60:
				f = (b & c) | (b & d) | (c & d)
				k = 0x8F1BBCDC
			default:
				f = b ^ c ^ d
				k = 0xCA62C1D6
			}

			temp := leftRotate(a, 5) + f + e + k + w[j]
			e = d
			d = c
			c = leftRotate(b, 30)
			b = a
			a = temp
		}

		// 更新哈希值
		h0 += a
		h1 += b
		h2 += c
		h3 += d
		h4 += e
	}

	// 将最终哈希值转换为字节数组
	var hash [20]byte
	binary.BigEndian.PutUint32(hash[0:4], h0)
	binary.BigEndian.PutUint32(hash[4:8], h1)
	binary.BigEndian.PutUint32(hash[8:12], h2)
	binary.BigEndian.PutUint32(hash[12:16], h3)
	binary.BigEndian.PutUint32(hash[16:20], h4)
	return hash
}
