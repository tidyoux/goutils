package viper

import (
	"bufio"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tidyoux/goutils"
)

const (
	extConfigsKey = "extConfigs"
)

// MergeExtIfNecessary merges external configs if necessary.
func MergeExtIfNecessary() error {
	exts := GetStringSlice(extConfigsKey, nil)
	for _, ext := range exts {
		if !goutils.FileExist(ext) {
			log.Errorf("merge config %s failed, file not exist", ext)
			continue
		}

		err := goutils.WithReadFile(ext, func(reader *bufio.Reader) error {
			return viper.MergeConfig(reader)
		})
		if err != nil {
			log.Errorf("merge config %s failed, %v", ext, err)
		}
	}
	return nil
}
