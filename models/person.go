package models

type Person struct {
	Name string
	Age  int `json:"age"`
}

var People []Person

func AddPerson(p Person) {
	People = append(People, p)
}

func DeletePersonByName(name string) bool {
	for i, p := range People {
		if p.Name == name {
			People = append(People[:i], People[i+1:]...)
			return true
		}
	}
	return false
}
