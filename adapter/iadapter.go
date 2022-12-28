package adapter

// IAdapter интерфейс объединяющие методы двух объектов
type IAdapter interface {
	MakeRequestGet() ([]User, error)
	MakeRequestCreate(user User) (User, error)
	MakeRequestUpdate(user User) (User, error)
	MakeRequestDelete(idMax int) (User, error)
	GetRezultDocumentation() ([]Document, error)
	GetDirectoriesSlice() ([]Directory, error)
	Max(p []User) int
	Min(p []User) int
	Close()
}
