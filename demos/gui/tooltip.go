package gui

import (
	"math/rand"
	"time"

	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/window"
	"github.com/g3n/g3nd/app"
)

func init() {
	app.DemoMap["gui.tooltip"] = &GuiTooltip{}
}

type GuiTooltip struct {
	labelOne   *gui.Label
	labelTwo   *gui.Label
	labelThree *gui.Label
	labelFour  *gui.Label
}

var yOffset float32

// Start is called once at the start of the demo.
func (t *GuiTooltip) Start(a *app.App) {

	yOffset = 10

	// Show and enable demo panel
	a.DemoPanel().SetRenderable(true)
	a.DemoPanel().SetEnabled(true)

	// Example one
	t.labelOne = tooltipLabel()

	tooltipLabelOne := tooltipPanel("(*Tooltip).SetFollow(true)\n" + "(*Tooltip).SetDelay(1000)\n" + "(*Tooltip).SetPositionOffset(0,20)")

	tooltipOne := gui.NewTooltip(tooltipLabelOne)
	tooltipOne.SetPositionOffset(0, 20)
	tooltipOne.SetFollow(true)
	tooltipOne.SetDelay(1000)

	t.labelOne.SetTooltip(tooltipOne)
	a.DemoPanel().Add(t.labelOne)

	// Example two
	t.labelTwo = tooltipLabel()
	tooltipLabelTwo := tooltipPanel("(*Tooltip).SetFollow(false)\n" + "(*Tooltip).SetDelay(250)")

	tooltipTwo := gui.NewTooltip(tooltipLabelTwo)
	tooltipTwo.SetFollow(false)
	tooltipTwo.SetDelay(250)

	t.labelTwo.SetTooltip(tooltipTwo)

	a.DemoPanel().Add(t.labelTwo)

	// Example three
	t.labelThree = tooltipLabel()
	tooltipLabelThree := tooltipPanel("(*Tooltip).SetDelay(0)\n" + "(*Tooltip).SetPositionFixed(100,100)")

	tooltipThree := gui.NewTooltip(tooltipLabelThree)
	tooltipThree.SetPositionFixed(100, 100)
	tooltipThree.SetDelay(0)

	t.labelThree.SetTooltip(tooltipThree)

	a.DemoPanel().Add(t.labelThree)

	// Example four
	t.labelFour = tooltipLabel()
	tooltipLabelFour := tooltipPanel("(*Tooltip).SetDelay(0)\n" + "(*Tooltip).SetPositionCustom( func ) => (randomX, randomY)")

	tooltipFour := gui.NewTooltip(tooltipLabelFour)
	tooltipFour.SetDelay(0)
	tooltipFour.SetPositionCustom(func(event *window.CursorEvent) (newX float32, newY float32) {
		// Random position generator
		wW, wH := window.Get().GetSize()
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)

		newX = float32(r.Intn(wW - int(t.labelFour.Width())))
		newY = float32(r.Intn(wH - int(t.labelFour.Height())))
		return
	})

	t.labelFour.SetTooltip(tooltipFour)

	a.DemoPanel().Add(t.labelFour)
}

// Update is called every frame.
func (t *GuiTooltip) Update(a *app.App, deltaTime time.Duration) {}

// Cleanup is called once at the end of the demo.
func (t *GuiTooltip) Cleanup(a *app.App) {

	t.labelOne.Dispatch(gui.OnCursorLeave, nil)
	t.labelTwo.Dispatch(gui.OnCursorLeave, nil)
	t.labelThree.Dispatch(gui.OnCursorLeave, nil)
	t.labelFour.Dispatch(gui.OnCursorLeave, nil)
}

// Create "Hover me" labels
func tooltipLabel() *gui.Label {

	l := gui.NewLabel("Hover me!")
	l.SetPosition(60, yOffset)
	l.SetBgColor(math32.NewColor("lightgreen"))
	l.SetColor(math32.NewColor("black"))
	l.SetPaddings(24, 64, 24, 64)
	yOffset += 90
	return l
}

// Create the actual tooltips
func tooltipPanel(t string) gui.IPanel {

	r := gui.NewLabel(t)
	r.SetBorders(1, 1, 1, 1)
	r.SetBordersColor(math32.NewColor("black"))
	r.SetPaddings(6, 6, 6, 6)
	r.SetColor(math32.NewColor("black")) // text
	r.SetBgColor(math32.NewColor("lightblue"))
	return r
}
