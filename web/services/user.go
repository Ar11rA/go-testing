package services

type User struct {
}

func CreateUser(user User) (bool, error) {
	return true, nil
}

func GetUsers() []User {
	return []User{}
}
