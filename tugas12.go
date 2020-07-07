package main

import (
	"fmt"
	"regexp"
)

func main() {
	var text = "jeruk mangga APEL"
	var regex, err = regexp.Compile(`[a-z]+`)
	var regexx, errx = regexp.Compile(`[A-Z]+`)

	if err != nil {
		fmt.Println(err.Error())
	}

	if errx != nil {
		fmt.Println(errx.Error())
	}

	var hasil = regex.FindAllString(text, 2)
	fmt.Println(hasil)

	var hasilx = regexx.FindAllString(text, 2)
	fmt.Println(hasilx)
}
