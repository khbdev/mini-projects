package domain

import "redis-clone/internal/model"

type Storage interface {
	Set(key string, data model.Data) error
	Get(key string) (model.Data) 
	Del(key string) error
}