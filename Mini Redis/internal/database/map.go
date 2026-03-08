package database

import "redis-clone/internal/model"




var Memory = make(map[string]string)


type Data struct {
  memory map[string]string
}

func NewData() *Data{
	return &Data{
		memory: Memory,
	}
}



func (d *Data) Set(key string, data model.Data) error {
	d.memory[key] = data.Value
	return  nil
}

func (d *Data) Get(key string) (model.Data) {
	 data := d.memory[key]
	return  model.Data{Value: data}
}

func (d *Data) Del(key string) error {
	delete(d.memory, key)

	return  nil
}