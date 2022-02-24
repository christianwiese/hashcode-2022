package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var days = []int64{0}

func main() {
	input := ParseInput(os.Args[1])

	fmt.Println("WORKING ON: ", os.Args[1])

	fmt.Println(input.numProjects, input.numContributors)
	//for _, c := range input.contributors {
	//	fmt.Println(c)
	//}
	//for _, p := range input.projects {
	//	fmt.Println(p)
	//}

	projects := input.projects
	contributors := input.contributors

	//sort.SliceStable(projects, func(i, j int) bool {
	//	return projects[i].score*projects[j].numDays*projects[j].numRoles > projects[j].score*projects[i].numDays*projects[i].numRoles
	//})
	sort.SliceStable(projects, func(i, j int) bool {
		return projects[j].numRoles > projects[i].numRoles
	})
	//sort.SliceStable(projects, func(i, j int) bool {
	//	return projects[i].numDays < projects[j].numDays
	//})

	res := Result{}

	for len(days) > 0 {
		fmt.Println(days[0])
		sort.SliceStable(days, func(i, j int) bool {
			return days[i] < days[j]
		})
		for _, p := range projects {
			if p.done {
				continue
			}
			names := FindContributors(p, contributors, days[0])
			if len(names) < len(p.roles) {
				continue
			}
			p.done = true
			days = append(days, days[0]+p.numDays)
			res = append(res, &Assignment{projectName: p.name, contributors: names})
		}
		days = days[1:]
	}

	for _, p := range res {
		fmt.Println(p)
	}

	Dump(res, os.Args[1])
}

func FindContributors(p *Project, cc []*Contributor, day int64) []string {
	names := []string{}
	for _, role := range p.roles {
		found := false
		for _, c := range cc {
			if c.occupiedUntil > day {
				continue
			}
			if c.skills[role.name] == role.level {
				names = append(names, c.name)
				c.occupiedUntil = day + p.numDays
				c.skills[role.name]++
				found = true
				break
			}
		}
		if found {
			continue
		}
		minLevel := int64(100000000)
		minID := -1
		for id, c := range cc {
			if c.occupiedUntil > day {
				continue
			}
			if c.skills[role.name] > role.level && c.skills[role.name] < minLevel {
				minLevel = c.skills[role.name]
				minID = id
			}
		}
		if minID >= 0 {
			cc[minID].occupiedUntil = day + p.numDays
			names = append(names, cc[minID].name)
		}
	}
	return names
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

func ParseInt(v string) int64 {
	x, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	return int64(x)
}

type Input struct {
	numContributors int64
	numProjects     int64
	contributors    []*Contributor
	projects        []*Project
}

type Contributor struct {
	name          string
	numSkills     int64
	skills        Skills
	occupiedUntil int64
}

type Skills map[string]int64

type Project struct {
	name       string
	numDays    int64
	score      int64
	bestBefore int64
	numRoles   int64
	roles      []*ProjectRole
	done       bool
}

type ProjectRole struct {
	name  string
	level int64
}

type Result []*Assignment

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

	for i := int64(0); i < input.numContributors; i++ {
		c := &Contributor{
			skills: Skills{},
		}

		scanner.Scan()
		c.name = scanner.Text()
		scanner.Scan()
		c.numSkills = ParseInt(scanner.Text())

		for j := int64(0); j < c.numSkills; j++ {
			scanner.Scan()
			skillName := scanner.Text()
			scanner.Scan()
			c.skills[skillName] = ParseInt(scanner.Text())
		}
		input.contributors = append(input.contributors, c)
	}

	for i := int64(0); i < input.numProjects; i++ {
		p := &Project{}

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

		for j := int64(0); j < p.numRoles; j++ {
			scanner.Scan()
			skillName := scanner.Text()
			scanner.Scan()
			skillLevel := ParseInt(scanner.Text())
			p.roles = append(p.roles, &ProjectRole{
				name:  skillName,
				level: skillLevel,
			})
		}
		input.projects = append(input.projects, p)
	}

	return input
}

func Score(input *Input, res []*Result) int {

	return 0
}
