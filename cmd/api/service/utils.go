package service

type Data struct {
	UserId int
	Id     int
	Title  string
	Body   string
}

type Payload struct {
	Data []Data
}

func GetAll() ([]Data, error) {

	return
}

func GetById(id int32) (Data, error) {
	return
}
