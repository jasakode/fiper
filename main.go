package fiper

import (
	"path/filepath"
	"plugin"
)

var fiperPlugin *plugin.Plugin

func New(soFileLock string) error {
	// START
	abs, err := filepath.Abs("./" + soFileLock)
	if err != nil {
		return err
	}
	p, err := plugin.Open(abs)
	if err != nil {
		return err
	}
	fiperPlugin = p

	v, err := fiperPlugin.Lookup("New")
	if err != nil {
		panic(err)
	}

	result := v.(func() error)()
	if result != nil {
		return result
	}

	return nil
}
