package main

import (
	"testing"
)

func TestIsContain(t *testing.T) {
	elf1 := Assigment{
		start: 1,
		end:   10,
	}

	elf2 := Assigment{
		start: 2,
		end:   5,
	}

	value := IsContain(elf1, elf2)

	if value == false {
		t.Errorf("1: Expected value to be %v and received %v", true, value)
	}

	elf3 := Assigment{
		start: 3,
		end:   5,
	}
	elf4 := Assigment{
		start: 4,
		end:   6,
	}

	value2 := IsContain(elf3, elf4)

	if value2 {
		t.Errorf("2; Expected value to be %v and received %v", false, value)
	}
}

func TestComputeLazyElf(t *testing.T) {
	var elfs [][2]Assigment
	elfs = append(elfs,
		[2]Assigment{NewAssigment(2, 4), NewAssigment(6, 8)},
		[2]Assigment{NewAssigment(2, 3), NewAssigment(4, 5)},
		[2]Assigment{NewAssigment(5, 7), NewAssigment(7, 9)},
		[2]Assigment{NewAssigment(2, 8), NewAssigment(3, 7)},
		[2]Assigment{NewAssigment(6, 6), NewAssigment(4, 6)},
		[2]Assigment{NewAssigment(2, 6), NewAssigment(4, 8)},
	)

    value := ComputeLazyElf(elfs);

    if value != 4 {
        t.Errorf("Compute: expected(%d), received(%d)", 4, value)
    }
	
}
