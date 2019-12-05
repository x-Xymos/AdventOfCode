package main

import (
	"fmt"
	"log"
	"strconv"
)

type Instruction struct {
	input  string
	opcode int
	mode1  int
	mode2  int
	param1 int
	param2 int
}

func parseOpcode(input int) Instruction {
	instructionAsString := strconv.Itoa(input)

	instruction := Instruction{mode1: 0, mode2: 0, input: instructionAsString}

	for len(instructionAsString) < 4 {
		instructionAsString = "0" + instructionAsString
	}
	instruction.opcode, _ = strconv.Atoi(string(instructionAsString[2:]))
	instruction.mode1, _ = strconv.Atoi(string(instructionAsString[1]))
	instruction.mode2, _ = strconv.Atoi(string(instructionAsString[0]))

	return instruction
}

func (instruction *Instruction) getParams(input *[]int, i *int) {
	if instruction.mode1 == 1 {
		instruction.param1 = (*input)[*i+1]
	} else {
		instruction.param1 = (*input)[(*input)[*i+1]]
	}

	if instruction.mode2 == 1 {
		instruction.param2 = (*input)[*i+2]
	} else {
		instruction.param2 = (*input)[(*input)[*i+2]]
	}
}

func calculateInput(input []int) {
	i := 0
	for {
		instruction := parseOpcode(input[i])
		defer func() {
			if r := recover(); r != nil {
				log.Println(instruction, input)
			}
		}()
		switch instruction.opcode {
		//add
		case 01:
			instruction.getParams(&input, &i)

			input[input[i+3]] = instruction.param1 + instruction.param2
			fmt.Println(instruction, input[i:4+i])
			i += 4
		//multiply
		case 02:
			instruction.getParams(&input, &i)

			input[input[i+3]] = instruction.param1 * instruction.param2
			fmt.Println(instruction, input[i:4+i])
			i += 4
		//input
		case 03:
			input[input[i+1]] = 5
			fmt.Println(instruction, input[i:2+i])
			i += 2
		//output
		case 04:
			fmt.Println(instruction, input[i:2+i])
			fmt.Println("Test code", input[input[i+1]])
			i += 2
		//jump if true
		case 05:
			instruction.getParams(&input, &i)

			if instruction.param1 != 0 {
				i = instruction.param2
			} else {
				i += 3
			}
			fmt.Println(instruction, input[i:3+i])
		//jump if false
		case 06:
			instruction.getParams(&input, &i)

			if instruction.param1 == 0 {
				i = instruction.param2
			} else {
				i += 3
			}
			fmt.Println(instruction, input[i:3+i])
		//less than
		case 07:
			instruction.getParams(&input, &i)

			if instruction.param1 < instruction.param2 {
				input[input[i+3]] = 1
			} else {
				input[input[i+3]] = 0
			}
			i += 4
			fmt.Println(instruction, input[i:4+i])
		//equals
		case 8:
			instruction.getParams(&input, &i)

			if instruction.param1 == instruction.param2 {
				input[input[i+3]] = 1
			} else {
				input[input[i+3]] = 0
			}
			i += 4
			fmt.Println(instruction, input[i:4+i])
		case 99:
			fmt.Println("Program end")
			return
		default:
		}
	}
}

func main() {
	//                                                                  1101
	var input = []int{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1002, 114, 19, 224, 1001, 224, -646, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 7, 224, 1, 223, 224, 223, 1101, 40, 62, 225, 1101, 60, 38, 225, 1101, 30, 29, 225, 2, 195, 148, 224, 1001, 224, -40, 224, 4, 224, 1002, 223, 8, 223, 101, 2, 224, 224, 1, 224, 223, 223, 1001, 143, 40, 224, 101, -125, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 3, 224, 1, 224, 223, 223, 101, 29, 139, 224, 1001, 224, -99, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 2, 224, 1, 224, 223, 223, 1101, 14, 34, 225, 102, 57, 39, 224, 101, -3420, 224, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 7, 224, 1, 223, 224, 223, 1101, 70, 40, 225, 1102, 85, 69, 225, 1102, 94, 5, 225, 1, 36, 43, 224, 101, -92, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 1, 224, 224, 1, 224, 223, 223, 1102, 94, 24, 224, 1001, 224, -2256, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 1, 224, 1, 223, 224, 223, 1102, 8, 13, 225, 1101, 36, 65, 224, 1001, 224, -101, 224, 4, 224, 102, 8, 223, 223, 101, 3, 224, 224, 1, 223, 224, 223, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 8, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 329, 1001, 223, 1, 223, 1108, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 344, 101, 1, 223, 223, 1108, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 359, 101, 1, 223, 223, 107, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 374, 101, 1, 223, 223, 1107, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 389, 101, 1, 223, 223, 107, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 404, 101, 1, 223, 223, 1008, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 419, 101, 1, 223, 223, 108, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 434, 101, 1, 223, 223, 1108, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 449, 101, 1, 223, 223, 1008, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 464, 1001, 223, 1, 223, 108, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 479, 101, 1, 223, 223, 7, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 494, 1001, 223, 1, 223, 8, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 509, 101, 1, 223, 223, 107, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 524, 1001, 223, 1, 223, 7, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 539, 1001, 223, 1, 223, 1007, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 554, 1001, 223, 1, 223, 8, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 569, 101, 1, 223, 223, 7, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 584, 1001, 223, 1, 223, 1008, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 599, 101, 1, 223, 223, 1007, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 614, 101, 1, 223, 223, 1107, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 629, 101, 1, 223, 223, 1107, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 644, 101, 1, 223, 223, 1007, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 659, 1001, 223, 1, 223, 108, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 674, 101, 1, 223, 223, 4, 223, 99, 226}
	// var input = []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
	// 	1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
	// 	999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	//var input = []int{1002, 4, 3, 4, 33}
	calculateInput(append([]int(nil), input...))

}
