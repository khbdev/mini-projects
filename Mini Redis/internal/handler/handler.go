package handler

import (
	"redis-clone/internal/database"
	"redis-clone/internal/model"
	"strings"
)




type Handler struct {
	database database.Data
}



func NewHandler(database database.Data) *Handler{
	return  &Handler{database: database}
}



func (h *Handler) Handle(input string) string {
		parts := strings.Fields(strings.TrimSpace(input))
	if len(parts) == 0 {
		return "ERR empty command"
	}
	cmd := strings.ToUpper(parts[0])

	switch cmd {
	case "SET":
			if len(parts) < 3 {
			return "ERR usage: SET key value"
		}
		key := parts[1]
		value := strings.Join(parts[2:], " ")

		
		mod := model.Data{
			Value: value,
		}

		h.database.Set(key, mod)
		return "OK"
		
			case "GET":
		if len(parts) < 2 {
			return "ERR usage: GET key"
		}

		key := parts[1]
		value := h.database.Get(key)
		return value.Value


	case "DEL":
		if len(parts) < 2 {
			return "ERR usage: DEL key"
		}

		key := parts[1]
		h.database.Del(key)
		return "OK"

		
	default:
		return "ERR unknown command"

	}
}