package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/modanisatech/marketplace/service-template/internal/config"
)

func TestGivenTestConfigFileWhenICallNewThenItShouldReturnConfig(t *testing.T) {
	configPath := "../../test/testdata"
	configName := "test-config"

	actualConfig, _ := config.New(configPath, configName)

	expectedConfig := config.Config{
		Appname: "service-template-test",
		Server:  config.Server{Port: ":1111"},
	}

	assert.Equal(t, expectedConfig, actualConfig)
}

func TestGivenNotExistingConfigFileWhenICallNewThenItShouldReturnError(t *testing.T) {
	configPath := "../../testdata/fake"
	notExistingConfigName := "nothing"

	_, err := config.New(configPath, notExistingConfigName)

	assert.NotNil(t, err)
}
