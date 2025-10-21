package handlers

import (
	"github.com/KrystofJan/tempus/internal/config"
	"github.com/KrystofJan/tempus/internal/errors"
	"github.com/KrystofJan/tempus/internal/service"
)

func EnsureDefaultTaskExists() error {
	cfg, cfgErr := config.Get()
	if cfgErr != nil {
		if cfgErr.ErrorCode != errors.ConfigFileNoExists {
			return cfgErr
		} else {
			if err := GenerateConfig(); err != nil {
				return err
			}
		}
	}

	taskProvider, err := service.NewTaskProvider()
	if err != nil {
		return err
	}

	_, err = taskProvider.FindTaskByName(cfg.DefaultTask)
	if err != nil {
		if _, err := taskProvider.AddTask(cfg.DefaultTask); err != nil {
			return err
		}
	}

	return nil
}
