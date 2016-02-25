package main
import "fmt"

type hotel struct  {
	Name, Address, City, Zip, Region string
}

type hotels struct {
	Hotels []hotel

}

func main()  {
	h := []hotels{
		hotel{
			Name: "Hotel California",
			Address: "4010 St Remy",
			City: "Merced",
			Zip: "95348",
			Region: "South",
		},
		hotel{
			Name: "Hotel Banana",
			Address: "4010 St Cognac",
			City: "Maderra",
			Zip: "95349",
			Region: "South",
		},
	}

	fmt.Println(h);
}