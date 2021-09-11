package util

import (
	"fmt"
	"go.uber.org/config"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
)

func FindFirstExistingFile(filePaths []string) string {
	for _, file := range filePaths {
		if _, err := os.Stat(file); err == nil {
			return file
		}
	}
	return ""
}

func GetConfigProviderFromFiles(filenames []string) config.Provider {
	opts := []config.YAMLOption {
		config.Permissive(),
		config.Expand(os.LookupEnv),
	}

	for _, name := range filenames {
		opts = append(opts, config.File(name))
	}
	provider, err := config.NewYAML(opts...)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return provider
}

func getFilesInDir(dir string) []fs.FileInfo {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	return files
}

func GetConfigProviderFromDir(configDir string) config.Provider {
	files := getFilesInDir(configDir)

	var filenames []string
	for _, file := range files {
		filenames = append(filenames, path.Join(configDir, file.Name()))
	}
	return GetConfigProviderFromFiles(filenames)
}

func GetStaticConfigProvider(val interface{}) (provider config.Provider, err error) {
	opt := config.Static(val)
	return config.NewYAML(opt)
}