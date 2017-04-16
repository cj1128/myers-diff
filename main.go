/*
* @Author: CJ Ting
* @Date:   2017-04-08 19:12:05
* @Last Modified by:   CJ Ting
* @Last Modified time: 2017-04-16 11:28:28
 */

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Operation uint

const INSERT Operation = 1
const DELETE Operation = 2
const MOVE Operation = 3

func (op Operation) String() string {
	switch op {
	case INSERT:
		return "INS"
	case DELETE:
		return "DEL"
	case MOVE:
		return "MOV"
	default:
		return "UNKNOWN"
	}
}

var colors = map[Operation]string{
	INSERT: "\033[32m",
	DELETE: "\033[31m",
	MOVE:   "\033[39m",
}

// 支持使用负数作为索引
type IntArray []int

func (a IntArray) get(i int) int {
	if i < 0 {
		return a[len(a)+i]
	}
	return a[i]
}

func (a IntArray) set(i, v int) {
	if i < 0 {
		a[len(a)+i] = v
	} else {
		a[i] = v
	}
}

func (a IntArray) clone() IntArray {
	return append(IntArray{}, a...)
}

func (a IntArray) String() string {
	k := len(a) / 2
	var result []string
	for i := -k; i <= k; i++ {
		result = append(result, strconv.Itoa(a.get(i)))
	}
	return strings.Join(result, " ")
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: myers-diff [src] [dst]")
		os.Exit(1)
	}
	src, err := getFileLines(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	dst, err := getFileLines(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	generateDiff(src, dst)
}

func generateDiff(src, dst []string) {
	script := shortestEditScript(src, dst)
	srcIndex, dstIndex := 0, 0
	for _, op := range script {
		switch op {
		case INSERT:
			fmt.Println(colors[op] + "+" + dst[dstIndex])
			dstIndex += 1
		case MOVE:
			fmt.Println(colors[op] + " " + src[srcIndex])
			srcIndex += 1
			dstIndex += 1
		case DELETE:
			fmt.Println(colors[op] + "-" + src[srcIndex])
			srcIndex += 1
		}
	}
}

// 生成最短的编辑脚本
func shortestEditScript(src, dst []string) []Operation {
	n := len(src)
	m := len(dst)
	max := n + m
	v := make(IntArray, 2*max+1)
	var trace []IntArray
	var x, y int

loop:
	for d := 0; d <= max; d++ {
		trace = append(trace, v.clone())
		for k := -d; k <= d; k += 2 {
			if k == -d {
				x = v.get(k + 1)
			} else if k != d && v.get(k-1) < v.get(k+1) {
				x = v.get(k + 1)
			} else {
				x = v.get(k-1) + 1
			}

			y = x - k

			for x < n && y < m && src[x] == dst[y] {
				x, y = x+1, y+1
			}

			v.set(k, x)

			if x == n && y == m {
				break loop
			}
		}
	}

	// this is for debug
	// for d := len(trace) - 1; d > 0; d-- {
	// 	fmt.Println(d, ":", trace[d])
	// }

	// 反向回溯
	var script []Operation

	x = n
	y = m
	var k, prevK, prevX, prevY int

	for d := len(trace) - 1; d > 0; d-- {
		k = x - y
		v := trace[d]
		if k == -d || (k != d && v.get(k-1) < v.get(k+1)) {
			prevK = k + 1
		} else {
			prevK = k - 1
		}
		prevX = v.get(prevK)
		prevY = prevX - prevK

		// fmt.Printf("d: %d, k: %d, x: %d, y: %d, prevX: %d, prevY: %d\n", d, k, x, y, prevX, prevY)

		for x > prevX && y > prevY {
			script = prepend(script, MOVE)
			x -= 1
			y -= 1
		}

		if x == prevX {
			script = prepend(script, INSERT)
		} else {
			script = prepend(script, DELETE)
		}

		x, y = prevX, prevY
	}

	return script
}

func prepend(s []Operation, op Operation) []Operation {
	return append([]Operation{op}, s...)
}

func getFileLines(p string) ([]string, error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
