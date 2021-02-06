package config

import (
	"github.com/spf13/viper"
)

func Load() {
	viper.SetDefault("buttons", map[string]string{
		"a":            "LeftClick",
		"b":            "",
		"x":            "",
		"y":            "",
		"bumper_right": "RightClick",
		"bumper_left":  "1",
		"back":         "esc",
		"start":        "i",
		"dpad_up":      "2",
		"dpad_down":    "3",
		"dpad_left":    "4",
		"dpad_right":   "5",
	})
	viper.SetDefault("held", []string{
		"bumper_right",
	})
	viper.SetConfigName("path-of-gamepad")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			viper.SafeWriteConfigAs("./path-of-gamepad.yaml")
		} else {

		}
	}

}

func Buttons() map[string]string {
	return viper.GetStringMapString("buttons")
}

func Holdable() []string {
	return viper.GetStringSlice("held")
}
