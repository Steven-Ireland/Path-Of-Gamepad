package config

import (
	"fmt"
	"strconv"

	"github.com/spf13/viper"
)

func Load() {
	viper.SetDefault("buttons", map[string]string{
		"a":             "LeftClick",
		"b":             "",
		"x":             "",
		"y":             "MiddleClick",
		"bumper_right":  "RightClick",
		"bumper_left":   "1",
		"back":          "esc",
		"start":         "i",
		"dpad_up":       "2",
		"dpad_down":     "3",
		"dpad_left":     "4",
		"dpad_right":    "5",
		"trigger_right": "shift",
		"trigger_left":  "control",
	})
	viper.SetDefault("held", []string{
		"bumper_right",
	})
	viper.SetDefault("settings", map[string]string{
		"screen_width_px":           "1920",
		"screen_height_px":          "1080",
		"character_y_offset_px":     "100",
		"character_x_offset_px":     "0",
		"walk_circle_radius_px":     "250",
		"attack_circle_radius_px":   "250",
		"free_mouse_sensitivity_px": "8",
		"dead_zone_percentage":      "0.17",
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

func IsKeyHoldable(button string) bool {
	for _, b := range Holdable() {
		if b == button {
			return true
		}
	}
	return false
}

func ScreenWidth() int {
	i, err := strconv.Atoi(viper.GetStringMapString("settings")["screen_width_px"])
	if err != nil {
		fmt.Printf("Error reading screen_width_px, value %s is not an integer\n", err)
		i = 1920
	}

	return i
}

func ScreenHeight() int {
	i, err := strconv.Atoi(viper.GetStringMapString("settings")["screen_height_px"])
	if err != nil {
		fmt.Printf("Error reading screen_height_px, value %s is not an integer\n", err)
		i = 1080
	}

	return i
}

func CharacterOffsetY() int {
	i, err := strconv.Atoi(viper.GetStringMapString("settings")["character_y_offset_px"])
	if err != nil {
		fmt.Printf("Error reading character_y_offset_px, value %s is not an integer\n", err)
		i = 100
	}

	return i
}

func CharacterOffsetX() int {
	i, err := strconv.Atoi(viper.GetStringMapString("settings")["character_x_offset_px"])
	if err != nil {
		fmt.Printf("Error reading character_x_offset_px, value %s is not an integer\n", err)
		i = 0
	}

	return i
}

func WalkCircleRadius() int {
	i, err := strconv.Atoi(viper.GetStringMapString("settings")["walk_circle_radius_px"])
	if err != nil {
		fmt.Printf("Error reading walk_circle_radius_px, value %s is not an integer\n", err)
		i = 250
	}

	return i
}

func AttackCircleRadius() int {
	i, err := strconv.Atoi(viper.GetStringMapString("settings")["attack_circle_radius_px"])
	if err != nil {
		fmt.Printf("Error reading attack_circle_radius_px, value %s is not an integer\n", err)
		i = 250
	}

	return i
}

func FreeMouseSensitivity() int {
	i, err := strconv.Atoi(viper.GetStringMapString("settings")["free_mouse_sensitivity_px"])
	if err != nil {
		fmt.Printf("Error reading free_mouse_sensitivity_px, value %s is not an integer\n", err)
		i = 8
	}

	return i
}

func DeadZonePercentage() float64 {
	f, err := strconv.ParseFloat(viper.GetStringMapString("settings")["dead_zone_percentage"], 64)
	if err != nil {
		fmt.Printf("Error reading dead_zone_percentage, value %s is not a float \n", err)
		f = 0.17
	}

	return f
}
