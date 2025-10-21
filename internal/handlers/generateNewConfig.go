package handlers

import (
	"github.com/KrystofJan/tempus/internal/config"
)

func GenerateConfig() error {
	if err := config.Delete(); err != nil {
		return err
	}
	_, err := config.New()
	if err != nil {
		return err
	}
	return nil
}
