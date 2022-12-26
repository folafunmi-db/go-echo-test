package service

import (
	"encoding/json"
	"io/ioutil"
)

type Data struct {
	UserId string
	Id     string
	Title  string
	Body   string
}

type Payload struct {
	Data []Data
}

func raw() ([]Data, error) {
	r, err := ioutil.ReadFile("data.json")

	if err != nil {
		return nil, err
	}

	var payload Payload
	err = json.Unmarshal(r, &payload.Data)

	if err != nil {
		return nil, err
	}

	return payload.Data, nil
}

func GetAll() ([]Data, error) {
	data, err := raw()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetById(id int) (any, error) {
	data, err := raw()

	if err != nil {
		return nil, err
	}

	if id > len(data) {
		res := make([]string, 0)
		return res, nil
	}

	return data[id], nil
}
