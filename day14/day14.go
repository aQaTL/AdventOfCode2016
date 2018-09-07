package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	input := []byte("jlmsuwbz")

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []byte) int64 {
	foundKeys := 0
	keyIdx := int64(0)
	for ; foundKeys != 64; keyIdx++ {
		data := strconv.AppendInt(input, keyIdx, 10)
		md5Sum := md5.Sum(data)

		if IsKey(md5Sum, keyIdx, input) {
			foundKeys++
		}
	}
	keyIdx--
	return keyIdx
}

func part2(input []byte) int64 {
	foundKeys := 0
	keyIdx := int64(0)

	hashes := make([][16]byte, 0, 100000)

	for ; foundKeys != 64; keyIdx++ {
		data := []byte(string(input) + strconv.Itoa(int(keyIdx)))
		md5Sum := md5.Sum(data)

		hexBuf := make([]byte, hex.EncodedLen(len(md5Sum)))
		hex.Encode(hexBuf, md5Sum[:])

		for i := 0; i < 2016; i++ {
			sum := md5.Sum(hexBuf)
			hex.Encode(hexBuf, sum[:])
		}

		hex.Decode(md5Sum[:], hexBuf)
		hashes = append(hashes, md5Sum)

		if keyIdx >= 1000 {
			if IsKeyStretched(hashes, keyIdx-1000) {
				foundKeys++
			}
		}
	}
	keyIdx -= 1001
	return keyIdx
}

func ContainsTriplet(hash *[16]byte) (bool, byte) {
	if c := hash[0] >> 4; c == hash[0]<<4>>4 && c == hash[1]>>4 {
		return true, c
	} else if c := hash[15] >> 4; c == hash[15]<<4>>4 && c == hash[14]<<4>>4 {
		return true, c
	} else {
		for i := 1; i < len(hash)-1; i++ {
			if c := hash[i] >> 4; c == hash[i]<<4>>4 {
				if c == hash[i-1]<<4>>4 || c == hash[i+1]>>4 {
					return true, c
				}
			}
		}
	}
	return false, 0
}

func IsKey(hash [16]byte, idx int64, data []byte) bool {
	containsTriplet, triplet := ContainsTriplet(&hash)
	if !containsTriplet {
		return false
	}

	for i := int64(0); i < 1000; i++ {
		md5Sum := md5.Sum(strconv.AppendInt(data, idx+i+1, 10))

		if c := md5Sum[0] >> 4; c == triplet && c == md5Sum[0]<<4>>4 && md5Sum[0] == md5Sum[1] {
			if c == md5Sum[2]>>4 {
				return true
			}
		} else if c := md5Sum[14] >> 4; c == triplet && c == md5Sum[14]<<4>>4 && md5Sum[14] == md5Sum[15] {
			if c == md5Sum[13]<<4>>4 {
				return true
			}
		} else {
			for j := 1; j < len(md5Sum)-2; j++ {
				if c := md5Sum[j] >> 4; c == triplet && c == md5Sum[j]<<4>>4 && md5Sum[j] == md5Sum[j+1] {
					if c == md5Sum[j-1]<<4>>4 || c == md5Sum[j+2]>>4 {
						return true
					}
				}
			}
		}

	}

	return false
}

func IsKeyStretched(hashes [][16]byte, idx int64) bool {
	containsTriplet, triplet := ContainsTriplet(&hashes[idx])
	if !containsTriplet {
		return false
	}

	for i := idx + 1; i <= idx+1000; i++ {
		hash := hashes[i]
		if c := hash[0] >> 4; c == triplet && c == hash[0]<<4>>4 && hash[0] == hash[1] {
			if c == hash[2]>>4 {
				return true
			}
		} else if c := hash[14] >> 4; c == triplet && c == hash[14]<<4>>4 && hash[14] == hash[15] {
			if c == hash[13]<<4>>4 {
				return true
			}
		} else {
			for j := 1; j < len(hash)-2; j++ {
				if c := hash[j] >> 4; c == triplet && c == hash[j]<<4>>4 && hash[j] == hash[j+1] {
					if c == hash[j-1]<<4>>4 || c == hash[j+2]>>4 {
						return true
					}
				}
			}
		}
	}
	return false
}
