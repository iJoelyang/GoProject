package main

import "fmt"

// map 使用哈希表，必须可以比较相等
// 除了slice， map，function 的内置类型都可以作为key
// struct 类型不包含上述字段，也可以作为key

func main() {
	m := map[string]string{
		"name":    "paddi",
		"course":  "golang",
		"site":    "github",
		"quality": "good",
	}

	m2 := make(map[string]int) //m2 == empty map
	var m3 map[string]int
	fmt.Println(m, m2, m3)

	fmt.Println("traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("getting values")
	fmt.Println("name: ", m["name"])
	fmt.Println("course: ", m["course"])

	if site, ok := m["sites"]; ok{
		fmt.Println("site: ", site)
	}else {
		fmt.Println("key does not exist.")
	}

	fmt.Println("deleting values")
	name, ok := m["name"]
	fmt.Println("name: ", name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println("name: ", name, ok)
}
