package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Response1 struct {
	Page   int
	Fruits []string
	Qtd    map[string]int
}

type Response2 struct {
	Page   int            `json:"page"`
	Fruits []string       `json:"fruits"`
	Qtd    map[string]int `json:"quantity"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(1.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("golang")
	fmt.Println(string(strB))

	slcD := []string{"go", "gopher", "goroutines"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apples": 7, "pears": 3}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &Response1{
		Page:   1,
		Fruits: slcD,
		Qtd:    mapD}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &Response2{
		Page:   1,
		Fruits: slcD,
		Qtd:    mapD}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"version": 1.6, "advantages": ["concurrency", "easy to learn", "fast"]}`)
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	version := dat["version"].(float64)
	fmt.Printf("%T: ", version)
	fmt.Println(version)

	advantages := dat["advantages"].([]interface{})
	advantages1 := advantages[0].(string)
	fmt.Println(advantages1)

	str := `{"page": 1, "fruits": ["apple", "peach"], "quantity": {"apples": 7, "pears": 3}}`
	res := Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])
	fmt.Println(res.Qtd["apples"])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
