package main

import (
	"crypto/md5"
	"strconv"
	"fmt"
	"encoding/hex"
)

func main() {
	input := []byte("jlmsuwbz")

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

	fmt.Println(keyIdx)

	foundKeys = 0
	keyIdx = 0

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

		if IsKeyStretched(md5Sum, keyIdx, input) {
			fmt.Println(string(data))
			fmt.Println(keyIdx)
			foundKeys++
		}
	}
	keyIdx--

	fmt.Println(keyIdx)
}

func IsKey(checkedHash [16]byte, idx int64, data []byte) bool {
	containsTriplet := false
	var triplet byte

	if c := checkedHash[0] >> 4; c == checkedHash[0]<<4>>4 && c == checkedHash[1]>>4 {
		triplet = c
		containsTriplet = true
	} else if c := checkedHash[15] >> 4; c == checkedHash[15]<<4>>4 && c == checkedHash[14]<<4>>4 {
		triplet = c
		containsTriplet = true
	} else {
		for i := 1; i < len(checkedHash)-1; i++ {
			if c := checkedHash[i] >> 4; c == checkedHash[i]<<4>>4 {
				if c == checkedHash[i-1]<<4>>4 || c == checkedHash[i+1]>>4 {
					triplet = c
					containsTriplet = true
					break
				}
			}
		}
	}

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

func IsKeyStretched(checkedHash [16]byte, idx int64, data []byte) bool {
	containsTriplet := false
	var triplet byte

	if c := checkedHash[0] >> 4; c == checkedHash[0]<<4>>4 && c == checkedHash[1]>>4 {
		triplet = c
		containsTriplet = true
	} else if c := checkedHash[15] >> 4; c == checkedHash[15]<<4>>4 && c == checkedHash[14]<<4>>4 {
		triplet = c
		containsTriplet = true
	} else {
		for i := 1; i < len(checkedHash)-1; i++ {
			if c := checkedHash[i] >> 4; c == checkedHash[i]<<4>>4 {
				if c == checkedHash[i-1]<<4>>4 || c == checkedHash[i+1]>>4 {
					triplet = c
					containsTriplet = true
					break
				}
			}
		}
	}

	if !containsTriplet {
		return false
	}

	for i := int64(0); i < 1000; i++ {
		md5Sum := md5.Sum(strconv.AppendInt(data, idx+i+1, 10))
		hexBuf := make([]byte, hex.EncodedLen(len(md5Sum)))
		hex.Encode(hexBuf, md5Sum[:])

		for j := 0; j < 2016; j++ {
			sum := md5.Sum(hexBuf)
			hex.Encode(hexBuf, sum[:])
		}

		hex.Decode(md5Sum[:], hexBuf)

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