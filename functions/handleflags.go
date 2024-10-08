package reloaded

import (
	"log"
	"strconv"
)

// this function handle all the flags passed
func HandleTheFlags(line []string) []string {
	for i := 0; i < len(line); i++ {
		if line[i] == "(cap)" {
			if i-1 >= 0 {
				_, b := Cap(line, i)
				if !b {
					log.Println("The flag doesn't work correctly so it would be handled as a string !")
					i++
				} else {
					line = append(line[:i], line[i+1:]...)
					i = 0
				}
			}
		} else if line[i] == "(up)" {
			if i-1 >= 0 {
				_, b := Up(line, i)
				if !b {
					log.Println("The flag doesn't work correctly so it would be handled as a string !")
					i++
				} else {
					line = append(line[:i], line[i+1:]...)
					i = 0
				}
			}
		} else if line[i] == "(low)" {
			if i-1 >= 0 {
				_, b := Low(line, i)
				if !b {
					log.Println("The flag doesn't work correctly so it would be handled as a string !")
					i++
				} else {
					line = append(line[:i], line[i+1:]...)
					i = 0
				}
			}
		} else if line[i] == "(cap," {
			if i+1 < len(line) && line[i+1] != "" {
				n, _ := strconv.Atoi(line[i+1][:len(line[i+1])-1])
				if n == 0 || n > i {
					log.Println("The flag doesn't work correctly so it would be handled as a string !")
					i++
				} else {
					CapNumbers(line, i, n)
					line = append(line[:i], line[i+2:]...)
					i = 0
				}
			}
		} else if line[i] == "(up," {
			if i+1 < len(line) && line[i+1] != "" {
				n, _ := strconv.Atoi(line[i+1][:len(line[i+1])-1])
				if n == 0 || n > i {
					log.Println("The flag doesn't work correctly so it would be handled as a string !")
					i++
				} else {
					UpNumbers(line, i, n)
					line = append(line[:i], line[i+2:]...)
					i = 0
				}
			}
		} else if line[i] == "(low," {
			if i+1 < len(line) && line[i+1] != "" {
				n, _ := strconv.Atoi(line[i+1][:len(line[i+1])-1])
				if n == 0 || n > i {
					log.Println("The (low,) flag doesn't work correctly so it would be handled as a string !")
					i++
				} else {
					LowNumbers(line, i, n)
					line = append(line[:i], line[i+2:]...)
					i = 0
				}
			}
		} else if line[i] == "(hex)" {
			if i-1 >= 0 {
				_, b := HexToDecimal(line, i)
				if !b {
					log.Println("The flag (hex) doesn't work correctly so it would be handled as a string !")
					i++
				} else {
					line = append(line[:i], line[i+1:]...)
					i = 0
				}
			}
		} else if line[i] == "(bin)" {
			if i-1 >= 0 {
				_, b := BinToDecimal(line, i)
				if !b {
					log.Println("The (bin) flag doesn't work correctly so it would be handled as a string !")
					i++
				} else {
					line = append(line[:i], line[i+1:]...)
					i = 0
				}
			}
		}
	}
	return line
}
