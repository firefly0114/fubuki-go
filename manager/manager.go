package manager

import (
	"errors"
	"fmt"
)

var Manager = newManager()

type manager struct {
}

func newManager() *manager {
	return new(manager)
}

func (m *manager) Run() error {
	// TODO: 实现manager逻辑，manager.Run 应返回错误
	return errors.New("error")
}

func Run() error {
	err := Manager.Run()
	return fmt.Errorf("manager exited, %w", err)
}
