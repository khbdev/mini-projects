package database

import (
	"redis-clone/internal/model"
	"sync"
)




var Memory = make(map[string]string)


type Data struct {
  memory map[string]string
  mu sync.Mutex
}

func NewData() *Data{
	return &Data{
		memory: Memory,
	}
}



func (d *Data) Set(key string, data model.Data) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.memory[key] = data.Value
	return  nil
}

func (d *Data) Get(key string) (model.Data) {
		d.mu.Lock()
	defer d.mu.Unlock()
	 data := d.memory[key]
	return  model.Data{Value: data}
}

func (d *Data) Del(key string) error {
		d.mu.Lock()
	defer d.mu.Unlock()
	delete(d.memory, key)
	return  nil
}