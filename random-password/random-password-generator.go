package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	// "crypto/rand"
	// "unsafe"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	i := 0
	
	var ListsList [][]int
	
	var SpecialCharsList [15]int
	var NumbersList [10]int
	var UpperCaseList [26]int
	var LowerCaseList [26]int
	var HexaDecimalList [16]int

	// Change Me:
	Length := 40         // Ascii:
	SpecialChars := false //  33-47, 15
	Numbers := false       //  48-57, 10
	UpperCase := false     //  65-90, 26
	LowerCase := false     //  97-122, 26
	
	HexaDecimal := true
		
	if SpecialChars {
		// for i = 33; i <= 47; i++ {
		// 	SpecialCharsList[i-33] += i
		// }
		SpecialCharsList = [15]int{33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47}
		ListsList = append(ListsList, SpecialCharsList[:])
	}

	if Numbers {
		// for i = 48; i <= 57; i++ {
		// 	NumbersList[i-48] += i
		// }
		NumbersList = [10]int{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}
		ListsList = append(ListsList, NumbersList[:])
	}
		
	if UpperCase {
		// for i = 65; i <= 90; i++ {
		// 	UpperCaseList[i-65] += i
		// }
		UpperCaseList = [26]int{65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90}
		ListsList = append(ListsList, UpperCaseList[:])
	}
			
	if LowerCase {
		// for i = 97; i <= 122; i++ {
		// 	LowerCaseList[i-97] += i
		// }
		LowerCaseList = [26]int{97, 98, 99, 100, 110, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122}
		ListsList = append(ListsList, LowerCaseList[:])
	}

	if HexaDecimal {
		HexaDecimalList = [16]int{48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 97, 98, 99, 100, 110, 102 }
		ListsList = append(ListsList, HexaDecimalList[:])
	}

	SumChars := 0
	for i := range ListsList {
		SumChars += len(ListsList[i])
	}

	// fmt.Println(ListsList)
	// fmt.Println(*rand_adress + ( int(unsafe.Sizeof(rand_adress)) ))
	// fmt.Println(unsafe.Sizeof(rand_adress))
	// fmt.Println(reflect.TypeOf((reflect.TypeOf(rand_adress).Size())))

	password := ""
	RandList := 0
	RandChar := 0
	for i = 1; i <= Length; i++ {
		RandList = rand.Intn(len(ListsList))
		RandChar = rand.Intn(len(ListsList[RandList]))
		password += string(ListsList[RandList][RandChar])
	}

	SumSub := 0.0
	for i = 1; i < Length; i++ {
		SumSub = math.Pow(float64(SumChars), float64(i))
	}

	combinations := math.Pow(float64(SumChars), float64(Length)) - float64(SumSub)
	fmt.Println("Password Strength:", combinations, "differant combinations")
	fmt.Println("Your Password is:")
	fmt.Println(password)
}
