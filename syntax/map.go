package main

import "fmt"

// func main() {
// 	var countryCapitalMap map[string]string
// 	countryCapitalMap = make(map[string]string)

// 	// map插入键值对 
// 	countryCapitalMap["China"] = "Beijing"
// 	countryCapitalMap["France"] = "Paris"

// 	// 使用key输出map值
// 	for country := range countryCapitalMap {
// 		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
// 	}

// 	// 查看元素在集合中是否存在
// 	captial, ok := countryCapitalMap["China"]
	
// 	fmt.Println(captial);
// 	fmt.Println(ok);
// }

func main() {
	countryCapitalMap := map[string] string {"France": "Paris", "Italy": "Rome"}
	fmt.Println("原始map")
	
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	delete(countryCapitalMap, "France")

	fmt.Println("Entry for France is deleted")  
   
    fmt.Println("删除元素后 map")

	for country := range countryCapitalMap {
		fmt.Println("Capital of",country,"is",countryCapitalMap[country])
	}
}