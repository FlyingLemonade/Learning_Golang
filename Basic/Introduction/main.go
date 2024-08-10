package main

import "fmt"


func main(){
	
	// Print
	fmt.Println("Hello World");
	fmt.Println("Type Hello World");

	// Fast Variable Declare
	sandal := "Fast Variable Declared"
	fmt.Println(sandal);

	// Function Call
	sandalJepit();
	fmt.Println(buySandal(2));
	

	// Declaring array
	var arrSandal [3]string;
	arrSandal[0] = "Swallow"
	arrSandal[1] = "Croc"
	arrSandal[2] = "Nike"
	fmt.Println(arrSandal)
	
	var arrHarga [3]int32 = [3]int32{5000,3000,2000}
	fmt.Println(arrHarga)

	// Fast Array Declaration
	arrHargaDiskon := [...]int32{4000,2000,1000}
	fmt.Printf("Discount 1000 setiap Kamis : %v\n",arrHargaDiskon)

	// Copy Array
	arrHargaCopy := arrHarga[:]
	fmt.Println(arrHargaCopy)

	// Slicing Array
	sliceExample := []int32{4,5,6}
	sliceExample = append(sliceExample, 7)
	fmt.Println(sliceExample)

	sliceExample_2 := []int32{8,9}
	sliceExample = append(sliceExample,sliceExample_2...)
	fmt.Println(sliceExample)

	// Map in Golang
	map_example := map[string]uint8{"Adam" : 70, "Sandal":50}
	fmt.Println(map_example["Adam"])

	// Check if the key value exists
	checkValue,ok := map_example["Sukri"]
	if ok {
		fmt.Printf("The age is %v \n",checkValue)
	} else {
		fmt.Println("Value not found")

	}

	// Loop in Golang
	for name:= range map_example{
		fmt.Printf("Name : %v\n", name)
	}

	for name, age:= range map_example{
		fmt.Printf("Name : %v, Age : %v \n", name, age)
	}

	// For loop
	for i := 0; i < 10; i++{
		fmt.Printf("For Loop Index : %v \n",i)
	}

	i := 0;
	// While Loop
	for i < 10{
		fmt.Printf("While Loop Index : %v \n", i)
		i++;
	}


	// String Rune

	var myString = "sandal"
	var indexed = myString[0]
	fmt.Printf("%v, %T", indexed, indexed)
	for i, v := range myString{
		fmt.Println(i,v)
	}


	// Pointer

	var p *int32 = new (int32);
	




}


func sandalJepit(){

	fmt.Println("\"Ini Sandal Baik Hati");
	fmt.Println("Sangatlah Baik Hait");
	fmt.Println("Sandal Pulanglah Kamu Keluargamu Merindukanmu\"");
	
}


func buySandal(jumlah int) int{
	hargaSandal := 3000
	return jumlah*hargaSandal
	
}