package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dc := 0
	seriesName := getString("Enter name of series", "")
	if seriesName == "" {
		fmt.Println("No series name entered; exiting!")
		os.Exit(1)
	}
	seasons := getInt("  Number of seasons", 1)
	if seasons < 1 || seasons > 30 {
		fmt.Println("Number of seasons seems invalid?")
		os.Exit(2)
	}
	discs := getInt("  Discs per season", 7)
	if seasons < 1 || seasons > 12 {
		fmt.Println("Number of discs seems invalid?")
		os.Exit(3)
	}

	os.Mkdir(seriesName, os.ModePerm)
	dc++
	os.Chdir(seriesName)

	for s := 1; s <= seasons; s++ {
		sf := folderName("s", s, seasons)
		os.Mkdir(sf, os.ModePerm)
		dc++
		os.Chdir(sf)

		for d := 1; d <= discs; d++ {
			df := folderName("d", d, discs)
			os.Mkdir(df, os.ModePerm)
			dc++
		}
		os.Chdir("..")
	}
	os.Chdir("..")

	fmt.Printf("%d folders created!\n", dc)
}

func folderName(prefix string, num, maxnum int) string {
	sl := len(strconv.Itoa(maxnum))
	rv := strconv.Itoa(num)
	for len(rv) < sl {
		rv = "0" + rv
	}
	return prefix + rv
}

func getString(prompt string, def string) string {
	reader := bufio.NewReader(os.Stdin)
	if def != "" {
		prompt = prompt + " [" + def + "]"
	}
	fmt.Print(prompt + ": ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", -1)
	if text == "" {
		text = def
	}
	return text
}

func getInt(prompt string, def int) int {
	valStr := getString(prompt, strconv.Itoa(def))
	valInt, err := strconv.Atoi(valStr)
	if err != nil {
		valInt = 0
	}
	return valInt
}
