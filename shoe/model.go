package shoe

type Model struct {
	Brand string `json:"brand"` // en este tag se le dice que si manda a json que cambie la propiedad de mayuscula a minuscula
	Price int    `json:"price"`
	Color string `json:"color"`
}

/*simulando una db*/

// slides similar a los arrays pero dinamicos

// mapa donde la clave será un string  el valor sera un Model
type Storage map[string]*Model

func (storage Storage) Create(model *Model) {
	storage[model.Brand] = model
}

func (storage Storage) GetAll() Storage {
	return storage
}
func (storage Storage) GetByBrand(brand string) *Model {
	if value, ok := storage[brand]; ok {
		return value
	}
	return nil
}

func (storage Storage) Delete(brand string) {
	delete(storage, brand)
}

func (stotage Storage) Update(brand string, model *Model) {
	stotage[brand] = model
}

var storage Storage

func init() {
	storage = make(map[string]*Model)
}
