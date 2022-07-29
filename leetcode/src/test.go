/*
* @Author: Ivy
* @Date: 2022/5/13 2:40 PM
**/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println()
	fmt.Printf("%v", strings.Join(os.Args[:], " "))
	fmt.Println()
}

func dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()] += 1
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func dup2() {

}
