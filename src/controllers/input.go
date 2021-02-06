package controllers

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Input struct {
	DPad                                       DPad
	Start, Back                                bool
	Start_PRESS, Back_PRESS                    bool
	Start_UNPRESS, Back_UNPRESS                bool
	A, B, X, Y                                 bool
	A_PRESS, B_PRESS, X_PRESS, Y_PRESS         bool
	A_UNPRESS, B_UNPRESS, X_UNPRESS, Y_UNPRESS bool

	Left  Analog
	Right Analog
}

type Analog struct {
	Direction      Vector2
	Hold           bool
	Hold_PRESS     bool
	Bumper         bool
	Bumper_PRESS   bool
	Bumper_UNPRESS bool
	Trigger        float64
}

type DPad struct {
	Up    bool
	Down  bool
	Left  bool
	Right bool

	Up_PRESS    bool
	Down_PRESS  bool
	Left_PRESS  bool
	Right_PRESS bool
}

func (dpad DPad) Direction() (r Vector2) {
	if dpad.Down {
		r.Y -= 1
	}
	if dpad.Up {
		r.Y += 1
	}
	if dpad.Left {
		r.X -= 1
	}
	if dpad.Right {
		r.X += 1
	}
	return
}

type Gamepad struct {
	Id       glfw.Joystick
	DeadZone float64
}

func ApplyDeadZone(v float64, deadZone float64) float64 {
	if v < -deadZone {
		return (v + deadZone) / (1 - deadZone)
	}
	if v > deadZone {
		return (v - deadZone) / (1 - deadZone)
	}
	return 0.0
}

func IsDeadZone(axes Vector2) bool {
	return axes.Y == 0 && axes.X == 0
}

func Read(gamepad Gamepad, lastInput Input) Input {
	joy := glfw.Joystick(gamepad.Id)

	axes := joy.GetAxes()
	buttons := joy.GetButtons()

	input := Input{}

	button := func(i int) bool {
		if i >= len(buttons) {
			return false
		}
		return buttons[i] == 1
	}

	axis := func(ix, iy int) Vector2 {
		if ix >= len(axes) || iy >= len(axes) {
			return Vector2{}
		}
		return Vector2{
			X: ApplyDeadZone(float64(axes[ix]), gamepad.DeadZone),
			Y: -ApplyDeadZone(float64(axes[iy]), gamepad.DeadZone),
		}
	}

	// fmt.Printf("Axes: %v\n", axes)

	input.DPad.Up = button(10)
	input.DPad.Up_PRESS = input.DPad.Up && !lastInput.DPad.Up
	input.DPad.Right = button(11)
	input.DPad.Right_PRESS = input.DPad.Right && !lastInput.DPad.Right
	input.DPad.Down = button(12)
	input.DPad.Down_PRESS = input.DPad.Down && !lastInput.DPad.Down
	input.DPad.Left = button(13)
	input.DPad.Left_PRESS = input.DPad.Left && !lastInput.DPad.Left

	input.A = button(0)
	input.A_PRESS = input.A && !lastInput.A
	input.A_UNPRESS = !input.A && lastInput.A

	input.B = button(1)
	input.B_PRESS = input.B && !lastInput.B
	input.B_UNPRESS = !input.B && lastInput.B

	input.X = button(2)
	input.X_PRESS = input.X && !lastInput.X
	input.X_UNPRESS = !input.X && lastInput.X

	input.Y = button(3)
	input.Y_PRESS = input.Y && !lastInput.Y
	input.Y_UNPRESS = !input.Y && lastInput.Y

	input.Back = button(6)
	input.Back_PRESS = input.Back && !lastInput.Back
	input.Back_UNPRESS = !input.Back && lastInput.Back

	input.Start = button(7)
	input.Start_PRESS = input.Start && !lastInput.Start
	input.Start_UNPRESS = !input.Start && lastInput.Start

	input.Left.Direction = axis(0, 1) // left thumb
	input.Left.Hold = button(8)       // left thumb pressed
	input.Left.Bumper = button(4)     // left bumper
	input.Left.Bumper_PRESS = input.Left.Bumper && !lastInput.Left.Bumper
	input.Left.Bumper_UNPRESS = !input.Left.Bumper && lastInput.Left.Bumper
	input.Left.Trigger = float64(axes[4]) // right trigger

	input.Right.Direction = axis(2, 3) // right thumb
	input.Right.Hold = button(9)       // right thumb pressed
	input.Right.Bumper = button(5)     // right bumper
	input.Right.Bumper_PRESS = input.Right.Bumper && !lastInput.Right.Bumper
	input.Right.Bumper_UNPRESS = !input.Right.Bumper && lastInput.Right.Bumper
	input.Right.Trigger = float64(axes[5]) // right trigger

	return input
}
