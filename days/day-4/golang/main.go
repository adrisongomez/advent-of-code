package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Assigment struct {
	start uint8
	end   uint8
}

func NewAssigment(start, end uint8) Assigment {
	return Assigment{
		start,
		end,
	}
}

func ComputeLazyElf(elfs [][2]Assigment) uint32 {
	var counter uint32
	for _, elf := range elfs {
		if IsOverlapped(elf[0], elf[1]) || IsOverlapped(elf[1], elf[0]) {
			fmt.Printf("Match %v %v\n", elf[0], elf[1])
			counter += 1
		}
	}
	return counter
}

func IsOverlapped(firstElf Assigment, secondElf Assigment) bool {
    if IsContain(firstElf, secondElf)  {
        return true
    }

    leftSide :=  firstElf.start <= secondElf.start && firstElf.end >= secondElf.start
    rightSide := firstElf.start <= secondElf.end && firstElf.end >= secondElf.end
    return  leftSide || rightSide
}

func IsContain(firstElf Assigment, secondElf Assigment) bool {
	return (firstElf.start <= secondElf.start && firstElf.end >= secondElf.end)
}

func FormStringToAssigment(line string) Assigment {
	ranges := strings.Split(line, "-")
	start, _ := strconv.ParseUint(ranges[0], 10, 8)
	end, _ := strconv.ParseUint(ranges[1], 10, 8)
	return NewAssigment(uint8(start), uint8(end))
}

func main() {
	value, _ := os.LookupEnv("PROJECT_ROOT")
	path := fmt.Sprintf("%s/days/day-4/input.txt", value)
	file, err := os.ReadFile(path)

	file_string := string(file)

	lines := strings.Split(file_string, "\n")

	var elfs [][2]Assigment

	if err != nil {
		fmt.Println("Error reading file")
		log.Fatal(err)
	}

	for _, line := range lines {
		parts := strings.Split(line, ",")
        if len(parts) != 2 {
            continue
        }
		elf1 := FormStringToAssigment(parts[0])
		elf2 := FormStringToAssigment(parts[1])

		elfs = append(elfs, [2]Assigment{elf1, elf2})
	}

	fmt.Println(ComputeLazyElf(elfs))

}
