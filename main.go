package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Navapon/play-with-go/member"
)

func main() {

	resp, err := http.Get("https://raw.githubusercontent.com/whs/bnk48json/master/bnk48.json")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	bnkMembers := make([]*member.Bnk, 0)
	err = json.Unmarshal([]byte(body), &bnkMembers)
	if err != nil {
		fmt.Println("Error Unmarshal", err.Error)
		return
	}

	fmt.Println("get:\n", bnkMembers)
	fmt.Println("status code ", resp.StatusCode)

	for _, member := range bnkMembers {
		fmt.Println(member)
	}
}
