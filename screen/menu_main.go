package screen

import (
	"eklase/state"

	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// MainMenu defines a main menu screen layout.
func MainMenu(th *material.Theme, state *state.State) Screen {
	var (
		add   widget.Clickable
		add2  widget.Clickable
		list  widget.Clickable
		list2 widget.Clickable
		quit  widget.Clickable
	)
	return func(gtx layout.Context) (Screen, layout.Dimensions) {
		matAddBut := material.Button(th, &add, "Add student")
		matAddBut.Font = text.Font{Variant: "Mono", Style: text.Regular}
		matAdd2But := material.Button(th, &add2, "Add classes")
		matAdd2But.Font = text.Font{Variant: "Mono", Style: text.Regular}
		matListBut := material.Button(th, &list, "List students")
		matListBut.Font = text.Font{Variant: "Mono", Style: text.Regular}
		matList2But := material.Button(th, &list2, "List classes")
		matList2But.Font = text.Font{Variant: "Mono", Style: text.Regular}
		matQuitBut := material.Button(th, &quit, "Quit")
		matQuitBut.Font = text.Font{Variant: "Mono", Style: text.Regular}

		d := layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(rowInset(matAddBut.Layout)),
			layout.Rigid(rowInset(matAdd2But.Layout)),
			layout.Rigid(rowInset(matListBut.Layout)),
			layout.Rigid(rowInset(matList2But.Layout)),
			layout.Rigid(rowInset(matQuitBut.Layout)),
		)
		if add.Clicked() {
			return AddStudent(th, state), d
		}
		if add2.Clicked() {
			return AddClass(th, state), d
		}
		if list.Clicked() {
			return ListStudent(th, state), d
		}
		if list2.Clicked() {
			return ListClass(th, state), d
		}
		if quit.Clicked() {
			state.Quit()
		}
		return nil, d
	}
}
