package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// ex
	testTree := [][]int{
		{59},
		{73, 41},
		{52, 40, 53},
		{26, 53, 6, 34},
	}

	url := "https://raw.githubusercontent.com/7-solutions/backend-challenge/main/files/hard.json"
	data, err := getData(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}

	fmt.Println("Test Output:", maxPathSum(testTree))  // 237
	fmt.Println("Hard JSON Output:", maxPathSum(data)) // 7273
}

func maxPathSum(nodes [][]int) int {
	if len(nodes) == 0 {
		return 0
	}

	for i := len(nodes) - 2; i >= 0; i-- {
		for j := 0; j < len(nodes[i]); j++ {
			nodes[i][j] += maxVal(nodes[i+1][j], nodes[i+1][j+1])
		}
	}

	return nodes[0][0]
}

func maxVal(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func getData(url string) ([][]int, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var data [][]int
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
