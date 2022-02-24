package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := ParseInput(os.Args[1])

	fmt.Println("WORKING ON: ", os.Args[1])

	fmt.Println(input)
	res := Result{}

	Dump(res, os.Args[1])
}

func Dump(o Result, file string) {
	output := fmt.Sprintf("./%s.out", strings.TrimSuffix(file, ".in"))
	f, _ := os.Create(output)
	defer f.Close()

	w := bufio.NewWriter(f)

	//num := len(o)
	//w.WriteString(fmt.Sprintf("%d", num))
	//w.WriteString("\n")
	//for interID, inter := range o {
	//	w.WriteString(fmt.Sprintf("%d\n", interID))
	//	numStreets := len(inter)
	//	w.WriteString(fmt.Sprintf("%d\n", numStreets))
	//	for name, duration := range inter {
	//		w.WriteString(fmt.Sprintf("%s %d\n", name, duration))
	//	}
	//}
	w.Flush()
}

func ParseInt(v string) int {
	x, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	return x
}

type Input struct {
}

type Result struct {
}

func ParseInput(path string) *Input {

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	input := &Input{}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	//scanner.Scan()
	//input.duration = ParseInt(scanner.Text())
	//scanner.Scan()
	//input.intersections = ParseInt(scanner.Text())
	//scanner.Scan()
	//input.numOfStreets = ParseInt(scanner.Text())
	//scanner.Scan()
	//input.numOfCars = ParseInt(scanner.Text())
	//scanner.Scan()
	//input.bonusPoints = ParseInt(scanner.Text())

	return input
}

func Score(input *Input, res []*Result) int {

	return 0
}
