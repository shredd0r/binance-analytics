package cache

import (
	"log/slog"
	"os"

	"github.com/shredd0r/binance-analytics/utils"
	"gopkg.in/yaml.v3"
)

const pathToCache = "./output/cache.yml"

type Provider interface {
	Get() (Cache, error)
	Save(cache Cache) error
	Delete() error
}

type ProviderImpl struct {
	logger *slog.Logger
}

func (p *ProviderImpl) Get() (Cache, error) {
	c := &Cache{}
	p.logger.Debug("start read cache file")
	err := utils.ReadYamlFile(pathToCache, c)

	if err != nil {
		p.logger.Error("failed read cache file", slog.Any("err", err.Error()))
		return Cache{}, err
	}

	p.logger.Debug("get cache successful")

	return *c, nil
}

func (p *ProviderImpl) Save(cache Cache) error {
	yamlBytes, err := yaml.Marshal(cache)
	if err != nil {
		p.logger.Error("failed marshal cache for save", slog.Any("err", err.Error()))
		return err
	}

	cacheFile, err := os.Create(pathToCache)
	if err != nil {
		p.logger.Error("failed create/open cache file", slog.Any("err", err.Error()))
		return err
	}

	_, err = cacheFile.Write(yamlBytes)
	if err != nil {
		p.logger.Error("failed write new cache in file", slog.Any("err", err.Error()))
	}

	p.logger.Debug("save new cache successful")
	return nil
}

func (p *ProviderImpl) Delete() error {
	err := os.Remove(pathToCache)
	if err != nil {
		p.logger.Error("failed delete cache", slog.Any("err", err.Error()))
		return err
	}

	p.logger.Debug("delete cache successful")
	return nil
}
