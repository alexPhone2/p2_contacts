package contacts

type Contact struct {
	Name        string
	PhoneNumber string
}

func GetContact() Contact {
	return Contact{
		Name:        "John Doe",
		PhoneNumber: "123-456-7890",
	}
}
