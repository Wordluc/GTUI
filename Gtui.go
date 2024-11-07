package GTUI

import (
	"errors"
	"strings"
	"time"

	"github.com/Wordluc/GTUI/Core"
	"github.com/Wordluc/GTUI/Core/Component"
	"github.com/Wordluc/GTUI/Core/Drawing"
	"github.com/Wordluc/GTUI/Core/Utils"
	"github.com/Wordluc/GTUI/Core/Utils/Color"
	"github.com/Wordluc/GTUI/Keyboard"
	"github.com/Wordluc/GTUI/Terminal"
)

type Gtui struct {
	globalColor      Color.Color
	term             Terminal.ITerminal
	keyb             Keyboard.IKeyBoard
	drawingManager   *Core.TreeManager[Core.IEntity]
	componentManager *Core.TreeManager[Component.IComponent]
	xCursor          int
	yCursor          int
	xSize            int
	ySize            int
	cursorVisibility bool
	loop             Keyboard.Loop
}

func NewGtui(loop Keyboard.Loop, keyb Keyboard.IKeyBoard, term Terminal.ITerminal) (*Gtui, error) {
	xSize, ySize := term.Size()
	return &Gtui{
		globalColor:      Color.GetDefaultColor(),
		cursorVisibility: true,
		loop:             loop,
		term:             term,
		keyb:             keyb,
		drawingManager:   Core.CreateTreeManager[Core.IEntity](),
		componentManager: Core.CreateTreeManager[Component.IComponent](),
		xCursor:          0,
		yCursor:          0,
		xSize:            xSize,
		ySize:            ySize,
	}, nil
}

func (c *Gtui) SetCur(x, y int) error {
	if x < 0 || y < 0 || x >= c.xSize || y >= c.ySize {
		return errors.New("cursor out of range")
	}
	compPreSet, _ := c.componentManager.Search(c.xCursor, c.yCursor)
	compsPostSet, _ := c.componentManager.Search(x, y)
	inPreButNotInPost := Utils.GetDiff(compsPostSet, compPreSet)
	inPostButNotInPre := Utils.GetDiff(compPreSet, compsPostSet)

	for i, e := range inPostButNotInPre {
		if ci, ok := inPostButNotInPre[i].(Component.IWritableComponent); ok {
			if ci.IsTyping() {
				continue
			}
		}
		e.OnHover(x, y)
	}
	lenInPreButNotInPost := len(inPreButNotInPost)
	for i := 0; i < lenInPreButNotInPost; i++ {
		if ci, ok := inPreButNotInPost[i].(Component.IWritableComponent); ok {
			if ci.IsTyping() {
				ci.SetCurrentPosCursor(x, y)
				return nil
			} else {
				inPreButNotInPost[i].OnOut(x, y)
			}
		} else {
			inPreButNotInPost[i].OnOut(x, y)
		}
	}

	c.yCursor = y
	c.xCursor = x

	for _, comp := range compsPostSet {
		if ci, ok := comp.(Component.IWritableComponent); ok {
			if ci.IsTyping() {
				c.xCursor, c.yCursor = ci.SetCurrentPosCursor(x, y)
			}
			break
		}
	}

	c.term.SetCursor(c.xCursor+1, c.yCursor+1)
	return nil
}
func (c *Gtui) GetCur() (int, int) {
	return c.xCursor, c.yCursor
}

func (c *Gtui) CreateStreamingCharacter() Component.StreamCharacter {
	stream := Component.StreamCharacter{}
	stream.Get = func() chan string {
		i := c.keyb.NewChannel()
		stream.IChannel = i
		return c.keyb.GetChannels()[i]
	}
	stream.Delete = func() {
		c.keyb.DeleteChannel(stream.IChannel)
	}
	return stream
}

func (c *Gtui) Start() {
	c.term.Start()
	c.term.Clear()
	c.IRefreshAll()
	c.keyb.Start(c.innerLoop)
	c.term.Stop()
}

func (c *Gtui) InsertEntity(entityToAdd Core.IEntity) {
	if container, ok := entityToAdd.(*Drawing.Container); ok {
		for _, entity := range container.GetChildren() {
			c.InsertEntity(entity)
		}
		return
	}
	c.drawingManager.AddElement(entityToAdd)
}

func (c *Gtui) InsertComponent(componentToAdd Component.IComponent) error {
	if container, ok := componentToAdd.(*Component.Container); ok {
		for _, component := range container.GetComponents() {
			c.componentManager.AddElement(component)
			component.OnOut(0, 0)
		}
		c.InsertEntity(componentToAdd.GetGraphics())
	} else {
		componentToAdd.OnOut(0, 0)
		c.componentManager.AddElement(componentToAdd)
		c.InsertEntity(componentToAdd.GetGraphics())
	}
	return nil
}

func (c *Gtui) RefreshComponents() {
	c.drawingManager.Refresh()
	c.componentManager.Refresh()
}

func (c *Gtui) EventOn(x, y int, event func(Component.IComponent)) error {
	resultArray, e := c.componentManager.Search(x, y)
	if e != nil {
		return e
	}
	for i := 0; i < len(resultArray); i++ {
		if ci, ok := resultArray[i].(*Component.Container); ok {
			resultArray = append(resultArray, ci.GetComponents()...)
		}
	}
	for i := range resultArray {
		event(resultArray[i])
	}
	return nil
}

func (c *Gtui) Click(x, y int) error {
	resultArray, e := c.componentManager.Search(x, y)
	if e != nil {
		return e
	}
	for i := range resultArray {
		resultArray[i].OnClick(x, y)
		time.AfterFunc(time.Millisecond*1000, func() {
			resultArray[i].OnRelease(x, y)
			c.IRefreshAll()
		})
	}
	return nil
}

func (c *Gtui) AllineCursor() {
	x, y := c.GetCur()
	comps, _ := c.componentManager.Search(x, y)
	for _, comp := range comps {
		if ci, ok := comp.(Component.IWritableComponent); ok {
			if ci.IsTyping() {
				deltax, deltay := ci.DiffCurrentToXY(x, y)
				c.yCursor = y + deltay
				c.xCursor = x + deltax
				break
			}
		}
	}
	c.term.SetCursor(c.xCursor+1, c.yCursor+1)
}

func (c *Gtui) IRefreshAll() {
	var str strings.Builder
	cond := func(node *Core.TreeNode[Core.IEntity]) bool {
		var el Core.IEntity
		if el = node.GetElement(); !el.IsTouched() {
			return true
		}
		x, y := el.GetPos()
		width, height := el.GetSize()
		str.WriteString(c.ClearZone(x, y, width, height))
		str.WriteString(el.GetAnsiCode(c.globalColor))
		str.WriteString(c.globalColor.GetAnsiColor())
		for _, child := range c.drawingManager.GetCollidingElement(node) {
			if child.IsTouched() {
				str.WriteString(child.GetAnsiCode(c.globalColor))
				str.WriteString(c.globalColor.GetAnsiColor())
				child.Touch()
				continue
			}
			str.WriteString(child.GetAnsiCode(c.globalColor))
			str.WriteString(c.globalColor.GetAnsiColor())
		}
		return true
	}
	c.drawingManager.Execute(cond)
	c.term.PrintStr(str.String())
	c.term.SetCursor(c.xCursor+1, c.yCursor+1)
	return
}

func (c *Gtui) innerLoop(keyb Keyboard.IKeyBoard) bool {
	//Keyboard
	//	c.IRefreshAll()
	c.AllineCursor()
	if !c.loop(c.keyb) {
		return false
	}
	c.IRefreshAll()
	return true
}
