package main

import "fmt"

func main() {
	reversibleString := reverse("Hiroto")
	fmt.Println(reversibleString)
}

func reverse(s string) string {
	b := []byte(s)
	bLen := len(b)
	bCopy := make([]byte, bLen)
	copy(bCopy, b)
	for i := 0; i < bLen; i++ {
		bCopy[i] = b[bLen-i-1]
	}
	return string(bCopy)
}
