package main

import(
	"fmt"
)

func Sqrt(x float64) float64{
	var z float64 = 1
	var y float64 = x

	for i:=0;i<40;i++{
		//fmt.Printf("i=%d\n",i)
		//fmt.Printf("y=%f\n",y)
		fmt.Println("i=")
		fmt.Println(i)
		fmt.Println("y=")
		fmt.Println(y)

		z -= (z*z - x)/(2*z)

		//fmt.Printf("z=%f\n",z)
		fmt.Println("z=")
		fmt.Println(z)

		if z == y{
			fmt.Println("beark\n")
			fmt.Println(i)
			break
		}
		y = z
	}
	return z
}

func main(){
	fmt.Println(Sqrt(2))
}
