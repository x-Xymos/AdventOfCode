package main

import (
	"fmt"
	"strconv"
)

func main() {
	passwordCount := 0
	for i := 168630; i < 718099; i++ {
		intArray := strconv.Itoa(i)
		if func() bool {
			var pair bool
			for x := 0; x < len(intArray); x++ {
				if x+1 < len(intArray) {
					if intArray[x+1] < intArray[x] {
						return false
					}
					if x == 0 {
						if intArray[x] == intArray[x+1] && intArray[x] != intArray[x+2] {
							pair = true
						}
					} else if x != len(intArray) {
						if intArray[x-1] != intArray[x+1] {
							if intArray[x] == intArray[x-1] {
								if x-2 >= 0 {
									if intArray[x-1] != intArray[x-2] {
										pair = true
									}
								} else {
									pair = true
								}
							}
							if intArray[x] == intArray[x+1] {
								if x+2 < len(intArray) {
									if intArray[x+1] != intArray[x+2] {
										pair = true
									}
								} else {
									pair = true
								}
							}
						}
					} else if intArray[x] == intArray[x-1] && intArray[x-1] != intArray[x-2] {
						pair = true
					}
				}
			}
			return pair
		}() {
			passwordCount++
		}
	}
	fmt.Println(passwordCount)
}
