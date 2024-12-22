package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type farm struct {
	ants_number int
	rooms       map[string][]int
	start       map[string][]int
	end         map[string][]int
	links       map[string][]string
}

func main() {
	var myFarm farm
	myFarm.Read("test.txt")
	BFS(myFarm)
	fmt.Println("number of ants is : ", myFarm.ants_number)
	fmt.Println("rooms are : ", myFarm.rooms)
	fmt.Println("start is : ", myFarm.start)
	fmt.Println("end is : ", myFarm.end)
	fmt.Println("links are : ", myFarm.links)
	fmt.Println("adjacent is : ", Graph(myFarm))
}

func (myFarm *farm) Read(filename string) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Println("error reading", err)
	}
	content := strings.Split(string(bytes), "\n")

	myFarm.rooms = make(map[string][]int)
	myFarm.start = make(map[string][]int)
	myFarm.end = make(map[string][]int)
	myFarm.links = make(map[string][]string)

	var st, en int
	number, err := strconv.Atoi(content[0])
	if err != nil {
		log.Println("couldn't convert", err)
	}
	myFarm.ants_number = number

	for index := range content {
		if strings.TrimSpace(content[index]) == "##start" {
			st++
			if index+1 <= len(content)-1 {
				split := strings.Split(strings.TrimSpace(content[index+1]), " ")
				x, err := strconv.Atoi(split[1])
				y, err2 := strconv.Atoi(split[2])
				if err == nil && err2 == nil {
					myFarm.start[split[0]] = []int{x, y}
				}

			}

		} else if strings.TrimSpace(content[index]) == "##end" {
			en++
			if index+1 <= len(content)-1 {
				split := strings.Split(strings.TrimSpace(content[index+1]), " ")
				x, err := strconv.Atoi(split[1])
				y, err2 := strconv.Atoi(split[2])
				if err == nil && err2 == nil {
					myFarm.end[split[0]] = []int{x, y}
				}

			}
		} else if strings.Contains(content[index], "-") {
			split := strings.Split(strings.TrimSpace(content[index]), "-")
			if len(split) == 2 {
				myFarm.links[split[0]] = append(myFarm.links[split[0]], split[1])
			}
		} else if strings.Count(content[index], " ") == 2 {
			split := strings.Split(strings.TrimSpace(content[index]), " ")
			if len(split) == 3 {
				x, err := strconv.Atoi(split[1])
				y, err2 := strconv.Atoi(split[2])
				if err == nil || err2 == nil {
					myFarm.rooms[split[0]] = []int{x, y}
				}
			}
		} else if (strings.HasPrefix(strings.TrimSpace(content[index]), "#") || strings.HasPrefix(strings.TrimSpace(content[index]), "L")) && (strings.TrimSpace(content[index]) != "##start" && strings.TrimSpace(content[index]) != "##end") {
			continue
		}
	}
	if en != 1 || st != 1 {
		log.Println("rooms setup is incorrect", err)
	}
}

func Graph(farm farm) map[string][]string {
	adjacent := make(map[string][]string)
	for room := range farm.rooms {
		adjacent[room] = []string{}
	}
	for room, links := range farm.links {
		for _, link := range links {
			adjacent[room] = append(adjacent[room], link)
			adjacent[link] = append(adjacent[link], room)

		}
	}

	return adjacent
}

func BFS(myFarm farm) {
	adjacent := Graph(myFarm)
	var Queue []string
	var endd string
	start := myFarm.start
	end := myFarm.end

	for key := range start {
		for _, adj := range adjacent[key] {
			Visited := make(map[string]bool)
			Parents := make(map[string]string)

			Queue = append(Queue, adj)
			Visited[adj] = true

			for key := range end {
				endd = key
			}

			for len(Queue) > 0 {
				current := Queue[0]
				Queue = Queue[1:]
				if current == endd {
					Queue = []string{}
					break
				}

				for _, link := range adjacent[current] {
					if !Visited[link] {
						Queue = append(Queue, link)
						Visited[link] = true
						Parents[link] = current
					}
				}
			}

			if !Visited[endd] {
				fmt.Printf("\n No path found to end room \n")
				return
			}

			path := []string{endd}
			current := endd

			for Parents[current] != "" {
				current = Parents[current]
				path = append([]string{current}, path...)
			}

			fmt.Printf("\nFinal path from start to end: %v\n", path)


		}
	}
}
