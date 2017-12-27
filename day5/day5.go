package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"bytes"
	"encoding/hex"
)

func main() {
	input := "wtnhxymk"

	buf := bytes.Buffer{}
	for i := 0; buf.Len() != 8; i++ {
		md5Sum := md5.Sum([]byte(input + strconv.Itoa(i)))
		if encoded := hex.EncodeToString(md5Sum[:]); encoded[:5] == "00000" {
			buf.WriteRune(rune(encoded[5]))
		}
	}
	fmt.Println("Part 1:", buf.String())

	password := [8]rune{}
	for i, found := 0, 0; found != 8; i++ {
		md5Sum := md5.Sum([]byte(input + strconv.Itoa(i)))
		if encoded := hex.EncodeToString(md5Sum[:]); encoded[:5] == "00000" {
			pos, err := strconv.Atoi(string(encoded[5]))
			if err != nil || pos > 7 || password[pos] != 0 {
				continue
			}
			password[pos] = rune(encoded[6])
			found++
		}
	}
	fmt.Println("Part 2:", string(password[:]))
}
