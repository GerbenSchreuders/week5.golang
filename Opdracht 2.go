package main

import (
	"fmt"
	//"github.com/davecgh/go-spew/spew"
)




type ShopItem struct {
		Naam 			string
		Verkooprijs 	float32
		Inkoopprijs 	float64
		VoorraadAantal 	int

	}



	type Sales struct {
		Naam			string
		ItemNaam		string
		GekochteItems 	int

	}

func main() {

//	x := ShopItem{
//		Kroket:         "Broodje Kroket",
//		Appel:          "Appel",
//		Gezond:         "Broodje gezond",
//		Verkooprijs:    0,
//		Inkoopprijs:    0,
//		VoorraadAantal: 0,



	BroodjeKroket:= ShopItem{
		Naam:				"Broodje Kroket",
		Verkooprijs:		1.5,
		Inkoopprijs:		0,
		VoorraadAantal:		25,
	}

	BroodjeGezond:= ShopItem{
		Naam:				"Broodje Gezond",
		Verkooprijs:		1.5,
		Inkoopprijs:		0,
		VoorraadAantal:		25,
	}

	Appel:= ShopItem{
		Naam:				"Broodje Appel",
		Verkooprijs:		1.5,
		Inkoopprijs:		0,
		VoorraadAantal:		25,
	}

	K1:= Sales{
		Naam:          "Bob",
		ItemNaam:      "Broodje Kroket",
		GekochteItems: 4,
	}


	/*K2:= Sales{
		Naam:          "Gerrit",
		ItemNaam:      "Broodje Gezond",
		GekochteItems: 3,
	}

	K3:= Sales{
		Naam:          "Bjorn",
		ItemNaam:      "Appel",
		GekochteItems: 1,
	}

	K4:= Sales{
		Naam:          "Student 1",
		ItemNaam:      "Broodje Kroket",
		GekochteItems: 5,
	}

	K5:= Sales{
		Naam:          "Student 2",
		ItemNaam:      "Broodje Kroket",
		GekochteItems: 7,
	}

	K6:= Sales{
		Naam:          "Student 3",
		ItemNaam:      "Broodje Kroket",
		GekochteItems: 4,
	}*/



fmt.Printf("%+v\n", BroodjeKroket, BroodjeGezond, Appel)
fmt.Println(K1)

//spew.Dump(K1)



	//s:= []ShopItem{BroodjeKroket,Appel}
	//fmt.Println(s)

	//fmt.Println(BroodjeKroket, BroodjeGezond, Appel)

}
