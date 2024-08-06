package main

import "fmt"


func main(){
	
	fmt.Println("Hello World");
	fmt.Println("Type Hello World");

	sandal := "Fast Variable Declared"

	fmt.Println(sandal);
	sandalJepit();
	fmt.Println(buySandal(2));
	
	var arrSandal [3]string;
	arrSandal[0] = "Swallow"
	arrSandal[1] = "Croc"
	arrSandal[2] = "Nike"
	fmt.Println(arrSandal)
	
	var arrHarga [3]int32 = [3]int32{5000,3000,2000}
	fmt.Println(arrHarga)

	arrHargaDiskon := [...]int32{4000,2000,1000}
	fmt.Printf("Discount 1000 setiap Kamis : %v\n",arrHargaDiskon)


	arrHargaCopy := arrHarga[:]
	fmt.Println(arrHargaCopy)

	
	sliceExample := []int32{4,5,6}
	sliceExample = append(sliceExample, 7)
	fmt.Println(sliceExample)

	sliceExample_2 := []int32{8,9}
	sliceExample = append(sliceExample,sliceExample_2...)
	fmt.Println(sliceExample)

	map_example := map[string]uint8{"Adam" : 70, "Sandal":50}
	fmt.Println(map_example["Adam"])

	checkValue,ok := map_example["Sukri"]
	if ok {
		fmt.Printf("The age is %v \n",checkValue)
	} else {
		fmt.Println("Value not found")

	}

	for name:= range map_example{
		fmt.Printf("Name : %v\n", name)
	}


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