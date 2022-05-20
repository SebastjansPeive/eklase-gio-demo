package screen

import (
	"eklase/state"
	"fmt"
	"image"
	"image/color"
	"log"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// ListStudent defines a screen layout for listing existing students.
func ListStudent(th *material.Theme, state *state.State) Screen {
	var close widget.Clickable
	list := widget.List{List: layout.List{Axis: layout.Vertical}}

	th.ContrastBg = color.NRGBA{A: 0xff, R: 0x5e, G: 0x9c, B: 0x64}
	lightContrast := th.ContrastBg
	lightContrast.A = 0x33
	darkContrast := th.ContrastBg
	darkContrast.A = 0x55

	students, err := state.Students()
	if err != nil {
		// TODO: Show user an error toast.
		log.Printf("failed to fetch students: %v", err)
		return nil
	}

	studentsLayout := func(gtx layout.Context) layout.Dimensions {
		return material.List(th, &list).Layout(gtx, len(students), func(gtx layout.Context, index int) layout.Dimensions {
			student := students[index]
			return layout.Stack{}.Layout(gtx,
				layout.Expanded(func(gtx layout.Context) layout.Dimensions {
					color := lightContrast
					if index%2 == 0 {
						color = darkContrast
					}
					max := image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Min.Y)
					paint.FillShape(gtx.Ops, color, clip.Rect{Max: max}.Op())
					return layout.Dimensions{Size: gtx.Constraints.Min}
				}),
				layout.Stacked(rowInset(material.Body1(th, fmt.Sprintf("%v %s %s", student.ID, student.Surname, student.Name)).Layout)),
			)
		})
	}

	return func(gtx layout.Context) (Screen, layout.Dimensions) {
		matCloseBut := material.Button(th, &close, "Close")
		matCloseBut.Background = color.NRGBA{A: 0xff, R: 0x5e, G: 0x9c, B: 0x64}
		matCloseBut.Font = text.Font{Variant: "Smallcaps", Weight: text.Bold, Style: text.Italic}
		d := layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Flexed(0.05, rowInset(material.Body1(th, fmt.Sprintf("%s %s %s", "ID", "Surname", "Name")).Layout)),
			layout.Flexed(1, rowInset(studentsLayout)),
			layout.Rigid(rowInset(matCloseBut.Layout)),
		)
		if close.Clicked() {
			return MainMenu(th, state), d
		}
		return nil, d
	}
}

func ListClass(th *material.Theme, state *state.State) Screen {
	var close widget.Clickable
	list := widget.List{List: layout.List{Axis: layout.Vertical}}

	th.ContrastBg = color.NRGBA{A: 0xff, R: 0x5e, G: 0x9c, B: 0x64}
	lightContrast := th.ContrastBg
	lightContrast.A = 0x33
	darkContrast := th.ContrastBg
	darkContrast.A = 0x55

	classes, err := state.Classes()
	if err != nil {
		log.Printf("failed to fetch classes: %v", err)
		return nil
	}

	classesLayout := func(gtx layout.Context) layout.Dimensions {
		return material.List(th, &list).Layout(gtx, len(classes), func(gtx layout.Context, index int) layout.Dimensions {
			class := classes[index]
			return layout.Stack{}.Layout(gtx,
				layout.Expanded(func(gtx layout.Context) layout.Dimensions {
					color := lightContrast
					if index%2 == 0 {
						color = darkContrast
					}
					max := image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Min.Y)
					paint.FillShape(gtx.Ops, color, clip.Rect{Max: max}.Op())
					return layout.Dimensions{Size: gtx.Constraints.Min}
				}),
				layout.Stacked(rowInset(material.Body1(th, fmt.Sprintf("%v %s %s", class.ID, class.Year, class.Modifier)).Layout)),
			)
		})
	}

	return func(gtx layout.Context) (Screen, layout.Dimensions) {
		matCloseBut := material.Button(th, &close, "Close")
		matCloseBut.Background = color.NRGBA{A: 0xff, R: 0x5e, G: 0x9c, B: 0x64}
		matCloseBut.Font = text.Font{Variant: "Smallcaps", Weight: text.Bold, Style: text.Italic}
		d := layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Flexed(0.05, rowInset(material.Body1(th, fmt.Sprintf("%s %s %s", "ID", "Year", "Modifier")).Layout)),
			layout.Flexed(1, rowInset(classesLayout)),
			layout.Rigid(rowInset(matCloseBut.Layout)),
		)
		if close.Clicked() {
			return MainMenu(th, state), d
		}
		return nil, d
	}
}

func ListGroup(th *material.Theme, state *state.State) Screen {
	var close widget.Clickable
	list := widget.List{List: layout.List{Axis: layout.Vertical}}

	th.ContrastBg = color.NRGBA{A: 0xff, R: 0x5e, G: 0x9c, B: 0x64}
	lightContrast := th.ContrastBg
	lightContrast.A = 0x33
	darkContrast := th.ContrastBg
	darkContrast.A = 0x55

	groups, err := state.Groups()
	if err != nil {
		log.Printf("failed to fetch groups: %v", err)
		return nil
	}

	assign := make([]widget.Clickable, len(groups))

	groupsLayout := func(gtx layout.Context) layout.Dimensions {
		return material.List(th, &list).Layout(gtx, len(groups), func(gtx layout.Context, index int) layout.Dimensions {
			group := groups[index]
			return layout.Stack{}.Layout(gtx,
				layout.Expanded(func(gtx layout.Context) layout.Dimensions {
					color := lightContrast
					if index%2 == 0 {
						color = darkContrast
					}
					max := image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Min.Y)
					paint.FillShape(gtx.Ops, color, clip.Rect{Max: max}.Op())
					return layout.Dimensions{Size: gtx.Constraints.Min}
				}),
				layout.Stacked(rowInset(func(gtx layout.Context) layout.Dimensions {
					matAssignBut := material.Button(th, &assign[index], "Assign class to student")
					matAssignBut.Background = color.NRGBA{A: 0xff, R: 0x5e, G: 0x9c, B: 0x64}
					matAssignBut.Font = text.Font{Variant: "Smallcaps", Weight: text.Bold, Style: text.Italic}
					return layout.Flex{}.Layout(gtx,
						layout.Rigid(rowInset(material.Body1(th, fmt.Sprintf("%s %s %s %s ", group.Name.String, group.Surname.String, group.Year.String, group.Modifier.String)).Layout)),
						layout.Rigid(rowInset(matAssignBut.Layout)),
					)
				})),
			)
		})
	}

	return func(gtx layout.Context) (Screen, layout.Dimensions) {
		matCloseBut := material.Button(th, &close, "Close")
		matCloseBut.Background = color.NRGBA{A: 0xff, R: 0x5e, G: 0x9c, B: 0x64}
		matCloseBut.Font = text.Font{Variant: "Smallcaps", Weight: text.Bold, Style: text.Italic}
		d := layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Flexed(0.05, rowInset(material.Body1(th, fmt.Sprintf("%s %s %s %s", "Name", "Surname", "Year", "Modifier")).Layout)),
			layout.Flexed(1, rowInset(groupsLayout)),
			layout.Rigid(rowInset(matCloseBut.Layout)),
		)
		for i := range assign {
			if assign[i].Clicked() {
				return AssignClassToStudent(th, state, i+1), d
			}
		}
		if close.Clicked() {
			return MainMenu(th, state), d
		}
		return nil, d
	}
}
