package stringutil

import "fmt"

const name = "Nipun"

var name2 = "Avinash"

var Fav []string
var fav2 [10]string
func init() {
	fmt.Println("init is called")
	Fav := make([]string, 4)
	Fav[0] = "First Entry"
	Fav[1] = "2nd Entry"
	Fav[2] = "Third Entry now"

	fav2[0] = "Ist entry"
	fav2[1] = "2nd entry"
	fav2[2] = "3rd Entry"
}

func AddFavourite(s string) {
	Fav = append(Fav, s)
}

func GetAllFavourites() []string {
	return Fav
}

func PrintFavourites() {
	for _, v := range Fav{
		fmt.Println(v)
	}
}

func AddFavourite2(s string) {
	fav2 = append(fav2, s)
}

func GetAllFavourites2() []string {
	return fav2
}

func PrintFavourites2() {
	for _, v := range fav2{
		fmt.Println(v)
	}
}