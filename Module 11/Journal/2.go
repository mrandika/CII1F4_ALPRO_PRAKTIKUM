package main

import "fmt"

type tag struct {
	topik string
	count int
}

type arrTopic []string
type arrTag []tag

func main() {
	var topik string
	var dataTopik arrTopic
	var dataTag arrTag

	fmt.Scan(&topik)

	for topik != "." {
		dataTopik = append(dataTopik, topik)

		fillArrayTag(topik, &dataTag)

		fmt.Scan(&topik)
	}

	trendingIndex := trendingSearch(dataTag)

	fmt.Println(dataTag[trendingIndex].topik)
}

func checkTopic(topik string, dataTag arrTag) bool {
	var found bool

	found = false

	i := 0

	for i < len(dataTag) && found == false {
		if topik == dataTag[i].topik {
			found = true
		} else {
			i++
		}
	}

	return found
}

func fillArrayTag(topik string, dataTag *arrTag) {
	var tag tag

	if checkTopic(topik, *dataTag) == false {
		tag.topik = topik
		tag.count = 1

		*dataTag = append(*dataTag, tag)
	} else {
		for i, v := range *dataTag {
			if topik == v.topik {
				(*dataTag)[i].count++
			}
		}
	}
}

func trendingSearch(dataTag arrTag) int {
	index := 0
	highest := 0

	for i, v := range dataTag {
		if v.count > highest {
			index = i
			highest = v.count
		}
	}

	return index
}