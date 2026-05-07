package configloader

import (
	"log"
	"os"
	"path/filepath"

	"github.com/zeromicro/go-zero/core/conf"
	"gopkg.in/yaml.v3"
)

const commonConfigFile = "common.yaml"

// MustLoad loads configFile and merges a sibling common.yaml first when present.
// Values in configFile override common.yaml recursively.
func MustLoad(configFile string, v any) {
	if err := Load(configFile, v); err != nil {
		log.Fatalf("error: config file %s, %s", configFile, err.Error())
	}
}

func Load(configFile string, v any) error {
	commonFile := filepath.Join(filepath.Dir(configFile), commonConfigFile)
	if _, err := os.Stat(commonFile); err != nil {
		if os.IsNotExist(err) {
			return conf.Load(configFile, v)
		}

		return err
	}

	base, err := loadYamlMap(commonFile)
	if err != nil {
		return err
	}

	override, err := loadYamlMap(configFile)
	if err != nil {
		return err
	}

	merged := mergeMaps(base, override)
	content, err := yaml.Marshal(merged)
	if err != nil {
		return err
	}

	return conf.LoadFromYamlBytes(content, v)
}

func loadYamlMap(file string) (map[string]any, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var m map[string]any
	if err := yaml.Unmarshal(content, &m); err != nil {
		return nil, err
	}

	if m == nil {
		m = make(map[string]any)
	}

	return m, nil
}

func mergeMaps(base, override map[string]any) map[string]any {
	for key, overrideValue := range override {
		baseValue, ok := base[key]
		if !ok {
			base[key] = overrideValue
			continue
		}

		baseMap, baseIsMap := baseValue.(map[string]any)
		overrideMap, overrideIsMap := overrideValue.(map[string]any)
		if baseIsMap && overrideIsMap {
			base[key] = mergeMaps(baseMap, overrideMap)
			continue
		}

		base[key] = overrideValue
	}

	return base
}
