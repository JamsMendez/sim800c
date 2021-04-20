package sim800c

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func indexOf(array []string, value string) bool {
	for i := 0; i < len(array); i++ {
		v := array[i]
		if strings.EqualFold(v, value) {
			return true
		}
	}
	return false
}

func isIMEI(array []string, value string) bool {
	for i := 0; i < len(array); i++ {
		v := array[i]
		if strings.EqualFold(v, value) {
			return true
		}
	}
	return false
}

func isIP(array []string) bool {
	for i := 0; i < len(array); i++ {
		element := array[i]
		num := strings.Split(element, ".")
		if len(num) == 4 {
			return true
		}
	}
	return false
}

func printInputCmd(cmd string) {
	s := strings.ReplaceAll(cmd, sNL, "")
	fmt.Println(color.YellowString("EXEC: %s", s))
}

func printOutputCmd(lines []string) {
	s := strings.Join(lines, ",")
	fmt.Println(color.GreenString("> %s\n", s))
}

func printErrCmd(err error) {
	fmt.Println(color.RedString("%v\n", err))
}
