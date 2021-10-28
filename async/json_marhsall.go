package async

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Origin struct {
	Name string `json:"name"`
}

type Location struct {
	Name string `json:"name"`
}

type Character struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Gender   string   `json:"gender"`
	Origin   Origin   `json:"origin"`
	Location Location `json:"location"`
}

func FetchCharacterById(client HttpClient, id int) Character {
	res, err := client.Get(fmt.Sprintf("https://rickandmortyapi.com/api/character/%d", id), nil)
	if err != nil {
		panic(err)
	}

	body, _ := ioutil.ReadAll(res.Body)

	var character Character
	json.Unmarshal(body, &character)
	fmt.Printf("character fetched as %v", character)
	return character
}
