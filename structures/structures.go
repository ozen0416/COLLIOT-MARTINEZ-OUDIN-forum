package structures

type Topic struct {
	Id        int
	CatId     int
	Content   string
	PubliDate string
}

type Auth struct {
	Email string
	Pass  string
}
