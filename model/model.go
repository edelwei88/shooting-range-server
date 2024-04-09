package model

const queries = [2]string{
	"v",
	"a",
}

type Auth struct {
	Login    string `json:"Login"`
	Password string `json:"Password"`
}

type Ammo struct {
	ID              uint32 `json:"ID"`
	Name            string `json:"Name"`
	Quantity        uint32 `json:"Quantity"`
	Id_manufacturer uint32 `json:"Id_manufacturer"`
}

type Client struct {
	ID      uint32 `json:"ID"`
	Name    string `json:"Name"`
	Surname string `json:"Surname"`
}

type Employee struct {
	ID      uint32 `json:"ID"`
	Name    string `json:"Name"`
	Surname string `json:"Surname"`
	Phone   string `json:"Phone"`
}

type Gun struct {
	ID             uint32 `json:"ID"`
	Id_ammo        uint32 `json:"Id_ammo"`
	Price_per_hour uint32 `json:"Price_per_hour"`
	Name           string `jsom:"Name"`
	Quantity       uint32 `json:"Quantity"`
}

type Lease struct {
	Id_reservation uint32 `json:"Id_reservation"`
	Price          uint32 `json:"Price"`
}

type Line struct {
	ID     uint32 `json:"ID"`
	Length uint   `json:"Length"`
}

type Manufacturer struct {
	ID   uint32 `json:"ID"`
	Name string `json:"Name"`
}

type Reservation struct {
	ID            uint32 `json:"ID"`
	Id_client     uint32 `json:"Id_client"`
	Id_employee   uint32 `json:"Id_employee"`
	Id_line       uint32 `json:"Id_line"`
	Id_gun        uint32 `json:"Id_gun"`
	Ammo_quantity uint32 `json:"Ammo_quantity"`
	Date          string `json:"Date"`
	Time_hours    uint32 `json:"Time_hours"`
}

var Auths = []Auth{
	{Login: "login1", Password: "password1"},
	{Login: "login2", Password: "password2"},
	{Login: "login3", Password: "password3"},
}

var Employees = []Employee{
	{ID: 1, Name: "name1", Surname: "surname1", Phone: "123456789"},
	{ID: 2, Name: "name2", Surname: "surname2", Phone: "123456789"},
	{ID: 3, Name: "name3", Surname: "surname3", Phone: "123456789"},
}
