package user

type Model struct {
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Storage map[string]*Model

func (storage Storage) Create(model *Model) *Model {
	storage[model.Email] = model
	return model
}

func (storage Storage) GetAll() Storage {
	return storage
}
func (storage Storage) GetByEmail(email string) *Model {
	if value, ok := storage[email]; ok {
		return value
	}
	return nil
}

func (storage Storage) Delete(email string) {
	delete(storage, email)
}

func (stotage Storage) Update(email string, model *Model) {
	stotage[email] = model
}
func (storage Storage) Login(email, password string) *Model {
	for _, v := range storage {
		if v.Email == email && v.Password == password {
			return v
		}
	}
	return nil

}

var storage Storage

func init() {
	storage = make(map[string]*Model)
	storage.Create(&Model{
		FirstName: "Thom",
		Email:     "thomtwd@gmail.com",
		Password:  "thom",
	})
}
