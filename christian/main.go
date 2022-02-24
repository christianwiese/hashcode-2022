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

	fmt.Println(input.numProjects, input.numContributors)
	for _, c := range input.contributors {
		fmt.Println(c)
	}
	for _, p := range input.projects {
		fmt.Println(p)
	}
	res := Result{}

	Dump(res, os.Args[1])
}

func Dump(o Result, file string) {
	output := fmt.Sprintf("./%s.out", strings.TrimSuffix(file, ".in"))
	f, _ := os.Create(output)
	defer f.Close()

	w := bufio.NewWriter(f)

	num := len(o)
	w.WriteString(fmt.Sprintf("%d", num))
	w.WriteString("\n")
	for _, a := range o {
		w.WriteString(fmt.Sprintf("%s\n", a.projectName))
		w.WriteString(fmt.Sprintf("%s\n", strings.Join(a.contributors, " ")))
	}
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
	numContributors int
	numProjects     int
	contributors    []*Contributor
	projects        []*Project
}

type Contributor struct {
	name      string
	numSkills int
	skills    Skills
}

type Skills map[string]int

type Project struct {
	name       string
	numDays    int
	score      int
	bestBefore int
	numRoles   int
	roles      Skills
}

type Result []Assignment

type Assignment struct {
	projectName  string
	contributors []string
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

	scanner.Scan()
	input.numContributors = ParseInt(scanner.Text())
	scanner.Scan()
	input.numProjects = ParseInt(scanner.Text())

	for i := 0; i < input.numContributors; i++ {
		c := &Contributor{
			skills: Skills{},
		}

		scanner.Scan()
		c.name = scanner.Text()
		scanner.Scan()
		c.numSkills = ParseInt(scanner.Text())

		for j := 0; j < c.numSkills; j++ {
			scanner.Scan()
			skillName := scanner.Text()
			scanner.Scan()
			c.skills[skillName] = ParseInt(scanner.Text())
		}
		input.contributors = append(input.contributors, c)
	}

	for i := 0; i < input.numProjects; i++ {
		p := &Project{
			roles: Skills{},
		}

		scanner.Scan()
		p.name = scanner.Text()
		scanner.Scan()
		p.numDays = ParseInt(scanner.Text())
		scanner.Scan()
		p.score = ParseInt(scanner.Text())
		scanner.Scan()
		p.bestBefore = ParseInt(scanner.Text())
		scanner.Scan()
		p.numRoles = ParseInt(scanner.Text())

		for j := 0; j < p.numRoles; j++ {
			scanner.Scan()
			skillName := scanner.Text()
			scanner.Scan()
			p.roles[skillName] = ParseInt(scanner.Text())
		}
		input.projects = append(input.projects, p)
	}

	return input
}

func Score(input *Input, res []*Result) int {

	return 0
}
