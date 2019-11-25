package healt

import (
	"github.com/dogukanayd/go-rest/app/utils"
	"os"
)

// Health..
type Health struct {
	MySQL utils.MySQLInterface
}

// New Health instance
func New() *Health {
	return &Health{MySQL: utils.NewMySQL().Connect(os.Getenv("DB_NAME"))}
}

// Check services
func (h *Health) Check() error  {
	err := h.checkMysql()

	if err != nil {
		return err
	}

	return nil
}

func (h *Health) checkMysql() error  {
	res := h.MySQL.Ping()
	_ = h.MySQL.Close()

	return res
}