package fiper

import (
	"plugin"
)

var fiperPlugin *plugin.Plugin

func New() error {
	// START
	p, err := plugin.Open("fiper/fiper.so")
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
