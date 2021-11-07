See it in action: https://www.youtube.com/watch?v=_aU_xkLx5eg

![](img/usage_example.gif)

The gist of it is that the left stick moves your character, and the right stick moves your mouse. I've set it up in a way that feels pretty good to use by allowing you to move in one direction, aim with the right stick, and cast. The program flicks your mouse to the aimed position and returns it to the original movement position after, letting you move in a pretty similar way as you can with a typical mouse + kb. 

Configuration options can be found in the path-of-gamepad.yaml file created when the app is launched (the file will be created in the same directory as the executable, or created manually). By default, right click is bound to right bumper, flasks are bound to the D-Pad and left bumper, escape is back, and opening the inventory is start. Left click is performed automatically with the left stick but can also be 'clicked' as needed (menus, items) with "A". 

```yaml
settings:
  screen_width_px: 1920 # Make sure to adjust your screen resolution to match your monitor!
  screen_height_px: 1080
  character_y_offset_px: 100 # Character is slightly higher than the screen center
  walk_circle_radius_px: 250 # Radius of walk circle
  attack_circle_radius_px: 250 # While attacking with a "holdable" key, radius can changes
  free_mouse_sensitivity_px: 8 # Max speed in pixels while free mousing
  dead_zone_percentage: 0.17 # Radius of controller joystick deadzone 
buttons:
  # a,b,x,y need to be quoted
  'a': LeftClick # LeftClick and RightClick are special and are only supported on a/b/x/y/bumper_right/bumper_left
  'b': ""
  'x': ""
  'y': "MiddleClick"
  back: esc # Extra keys can be found here: https://github.com/go-vgo/robotgo/blob/master/docs/keys.md
  bumper_left: "1" # Multiple keybinds can be added to a single button with comma separation, i.e "1,2,3", but it's against the rules
  bumper_right: RightClick
  stick_button_left: ""
  stick_button_right: ""
  dpad_down: "3"
  dpad_left: "4"
  dpad_right: "5"
  dpad_up: "2"
  trigger_right: "shift"
  trigger_left: "control"
  start: i

# Designating a key as "held" is the difference between a flask (activates, 
# doesn't impact movement), and a main skill (which you want to aim, stop and hold, etc)
held: 
- bumper_right
- trigger_right
- trigger_left

```

Feel free to post any issues, PRs, feature requests - no promises I'll be able to get to them all (I also have a day job!) but I'll try my best to fix anything glaring asap. 
