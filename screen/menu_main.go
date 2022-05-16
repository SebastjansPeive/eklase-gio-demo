package screen

import (
	"eklase/state"

	"image/color"

	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// MainMenu defines a main menu screen layout.
func MainMenu(th *material.Theme, state *state.State) Screen {
	var (
		addStudent   widget.Clickable
		addClass     widget.Clickable
		listStudents widget.Clickable
		listClasses  widget.Clickable
		listGroups   widget.Clickable
		quit         widget.Clickable
	)
	return func(gtx layout.Context) (Screen, layout.Dimensions) {
		matAddStudentButton := material.Button(th, &addStudent, "Add student")
		matAddStudentButton.Background = color.NRGBA{A: 0xff, R: 0x5e, G: 0x9c, B: 0x94}
		matAddStudentButton.Font = text.Font{Variant: "Smallcaps", Weight: text.Bold, Style: text.Italic}
		matAddClassButton := material.Button(th, &addClass, "Add class")
		matAddClassButton.Background = color.NRGBA{A: 0xff, R: 0x5e, G: 0x9c, B: 0x84}
		matAddClassButton.Font = text.Font{Variant: "Smallcaps", Weight: text.Bold, Style: text.Italic}
		matListStudentsButton := material.Button(th, &listStudents, "List students")
		matListStudentsButton.Background = color.NRGBA{A: 0xff, R: 0x5e, G: 0x9c, B: 0x74}
		matListStudentsButton.Font = text.Font{Variant: "Smallcaps", Weight: text.Bold, Style: text.Italic}
		matListClassesButton := material.Button(th, &listClasses, "List classes")
		matListClassesButton.Background = color.NRGBA{A: 0xff, R: 0x5e, G: 0x9c, B: 0x64}
		matListClassesButton.Font = text.Font{Variant: "Smallcaps", Weight: text.Bold, Style: text.Italic}
		matListGroupsButton := material.Button(th, &listGroups, "List groups")
		matListGroupsButton.Background = color.NRGBA{A: 0xff, R: 0x5e, G: 0x9c, B: 0x54}
		matListGroupsButton.Font = text.Font{Variant: "Smallcaps", Weight: text.Bold, Style: text.Italic}
		matQuitBut := material.Button(th, &quit, "Quit")
		matQuitBut.Background = color.NRGBA{A: 0xff, R: 0x5e, G: 0x9c, B: 0x54}
		matQuitBut.Font = text.Font{Variant: "Smallcaps", Weight: text.Bold, Style: text.Italic}

		d := layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(rowInset(matAddStudentButton.Layout)),
			layout.Rigid(rowInset(matAddClassButton.Layout)),
			layout.Rigid(rowInset(matListStudentsButton.Layout)),
			layout.Rigid(rowInset(matListClassesButton.Layout)),
			layout.Rigid(rowInset(matListGroupsButton.Layout)),
			layout.Rigid(rowInset(matQuitBut.Layout)),
		)
		if addStudent.Clicked() {
			return AddStudent(th, state), d
		}
		if addClass.Clicked() {
			return AddClass(th, state), d
		}
		if listStudents.Clicked() {
			return ListStudent(th, state), d
		}
		if listClasses.Clicked() {
			return ListClass(th, state), d
		}
		if listGroups.Clicked() {
			return ListGroup(th, state), d
		}
		if quit.Clicked() {
			state.Quit()
		}
		return nil, d
	}
}
