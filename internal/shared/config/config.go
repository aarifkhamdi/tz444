package config

import (
	"errors"
	"fmt"
	"os"
	"reflect"

	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

func bindTags(v *viper.Viper, iface any) error {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)

	if ifv.Kind() == reflect.Ptr {
		ift = ift.Elem()
	}

	for i := range ift.NumField() {
		field := ift.Field(i)
		if envName, ok := field.Tag.Lookup("env"); ok {
			if err := v.BindEnv(field.Name, envName); err != nil {
				return fmt.Errorf("failed to bind env %s: %w", envName, err)
			}
		}
	}
	return nil
}

func New[T any](cfg T) T {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	if err := bindTags(v, cfg); err != nil {
		panic(err)
	}
	if err := v.ReadInConfig(); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			panic(err)
		}
	}

	if err := v.UnmarshalExact(cfg); err != nil {
		panic(err)
	}

	if err := validator.New().Struct(cfg); err != nil {
		panic(err)
	}

	return cfg
}
