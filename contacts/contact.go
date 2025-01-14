package contacts

type Contact struct {
	Name string
}

func GetContact() Contact {
	return Contact{
		Name: "John Doe",
	}
}
