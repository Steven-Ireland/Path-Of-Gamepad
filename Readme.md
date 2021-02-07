See it in action: https://www.youtube.com/watch?v=_aU_xkLx5eg

![](https://media0.giphy.com/media/LSFB29dKFn7SRgdZ2q/giphy.gif)

Configuration options can be found in the path-of-gamepad.yaml file created when the app is launched (the file will be created in the same directory as the executable)
```yaml
settings:
  screen_width_px: 1920
  screen_height_px: 1080
  character_y_offset_px: 100 # Character is slightly higher than the screen center
  walk_circle_radius_px: 250 # Radius of walk circle
  attack_circle_radius_px: 250 # While attacking with a "holdable" key, radius can changes
  free_mouse_sensitivity_px: 8 # Max speed in pixels while free mousing
  dead_zone_percentage: 0.17 # Radius of controller joystick deadzone 
buttons:
  a: LeftClick # LeftClick and RightClick are special and are only supported on a/b/x/y/bumper_right/bumper_left
  b: ""
  x: ""
  y: ""
  back: esc # Extra keys can be found here: https://github.com/go-vgo/robotgo/blob/master/docs/keys.md
  bumper_left: "1" # Multiple keybinds can be added to a single button with comma separation, i.e "1,2,3", but it's against the rules
  bumper_right: RightClick
  dpad_down: "3"
  dpad_left: "4"
  dpad_right: "5"
  dpad_up: "2"
  start: i

# Designating a key as "held" is the difference between a flask (activates, 
# doesn't impact movement), and a main skill (which you want to aim, stop and hold, etc)
held: 
- bumper_right
```

Feel free to post any issues, PRs, feature requests - no promises I'll be able to get to them all (I also have a day job!) but I'll try my best to fix anything glaring asap. 