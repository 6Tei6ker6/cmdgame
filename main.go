package main

import (
	"fmt"
	"os/exec"
	"strconv"
)

const (
	Cols = 100
	Lines = 42 // исправить на 40 и поправить в коде, костыль в 42 сделан для непостоянной генерации
)

var console_lines string = "lines=" + strconv.Itoa(Lines)
var console_cols string = "cols=" + strconv.Itoa(Cols)

func main() {
	//генерация поля в консоли в заданных параметрах 
	cmd := exec.Command("cmd", "/c", "mode", "con", console_cols, console_lines)
	cmd.Run()
	gamemap := make([]string, 0, Cols*(Lines - 2))
	for x := Cols*(Lines - 2); x > 0; x--{
		switch x % 2{
			case 0:
				gamemap = append(gamemap, "@") //перенос поля в слайс для более удобного взаимодействия с ним
			default:
				gamemap = append(gamemap, " ")
		}
	}

	for _, x := range gamemap{
		fmt.Printf(x)
	}
}