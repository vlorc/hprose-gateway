package session

type SessionFactory interface {
	Instance(token string) (Session, error)
}

type Session interface {
	Appid() string
}
