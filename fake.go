package fake

import (
	"fmt"
)

func fake() {

	fmt.Println("Hello", random.Intn(10))
	for i := 0; i < 250; i++ {
		fmt.Println(Name.FullName(Any))
		fmt.Println(Address.Australian())
		fmt.Println(Pharmacy.FullDrug())
	}
}
