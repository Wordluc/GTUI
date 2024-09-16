package GTUI

import (
	"GTUI/Core"
	"GTUI/Core/Component"
	"GTUI/Core/Utils"
	"GTUI/Core/Utils/Color"
	"GTUI/Keyboard"
	"GTUI/Terminal"
	"errors"
	"strings"
	"time"
)

type Gtui struct {
	globalColor      Color.Color
	term             Terminal.ITerminal
	keyb             Keyboard.IKeyBoard
	buff             []Core.IEntity
	componentManager *Component.ComponentManager
	xCursor          int
	yCursor          int
	xSize            int
	ySize            int
	cursorVisibility bool
	loop             Keyboard.Loop
}

func NewGtui(loop Keyboard.Loop, keyb Keyboard.IKeyBoard, term Terminal.ITerminal) (*Gtui, error) {
	xSize, ySize := term.Size()
	var totalError Utils.CError
	totalError.AddIf(loop == nil || keyb == nil || term == nil , errors.New("invalid size"))
	componentManager,e:= Component.Create(xSize, ySize, 5)
	totalError.Add(e)
	if totalError.HasError() {
		return nil, totalError
	}
	return &Gtui{
		globalColor:      Color.GetDefaultColor(),
		cursorVisibility: true,
		loop:             loop,
		term:             term,
		keyb:             keyb,
		buff: make([]Core.IEntity,0),
		componentManager: componentManager,
		xCursor:          0,
		yCursor:          0,
		xSize:            xSize,
		ySize:            ySize,
	}, nil
}

func (c *Gtui) SetCur(x, y int)error {
	if x < 0 || y < 0 || x >= c.xSize || y >= c.ySize {
		return errors.New("cursor out of range")
	}
	c.SetVisibilityCursor(true)
	compPreSet, _ := c.componentManager.Search(c.xCursor, c.yCursor)
	compsPostSet, _ := c.componentManager.Search(x, y)
	inPreButNotInPost := Utils.GetDiff(compsPostSet, compPreSet)
	inPostButNotInPre := Utils.GetDiff(compPreSet, compsPostSet)
	for _, e := range inPostButNotInPre {
		e.OnHover()
	}
	for _, comp := range inPreButNotInPost {
		if ci, ok := comp.(Component.IWritableComponent); ok {
			if ci.IsTyping() {
				ci.SetCurrentPosCursor(x, y)
				return nil
			}else{
		      ci.(Component.IComponent).OnLeave()
			}
		}
	}

	c.yCursor = y
	c.xCursor = x
	for _, comp := range compsPostSet {
		if ci, ok := comp.(Component.IWritableComponent); ok {
			if ci.IsTyping() {
				c.xCursor, c.yCursor = ci.SetCurrentPosCursor(x, y)
				break
			}
		}
	}
	c.term.SetCursor(c.xCursor+1, c.yCursor+1)
	c.SetVisibilityCursor(false)
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
	c.term.Clear()
	c.term.Stop()
}

func (c *Gtui) InsertEntity(entity Core.IEntity) {
	c.buff = append(c.buff, entity)
}

func (c *Gtui) InsertComponent(component Component.IComponent) error {
	c.buff = append(c.buff, component.GetGraphics())
	return c.componentManager.Add(component)
}
func (c *Gtui) EventOn(x, y int, event func(Component.IComponent)) error {
	resultArray, e := c.componentManager.Search(x, y)
	if e != nil {
		return e
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
		resultArray[i].OnClick()
		time.AfterFunc(time.Millisecond*1000, func() {
			resultArray[i].OnRelease()
			c.IRefreshAll()
		})
	}
	return nil
}

func (c *Gtui) AllineCursor() {
	c.SetVisibilityCursor(false)
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
	c.SetVisibilityCursor(true)
}

func (c *Gtui) IRefreshAll() {
	c.SetVisibilityCursor(false)
	var str strings.Builder
	for _, b := range c.buff {
		str.WriteString(b.GetAnsiCode(c.globalColor))
		str.WriteString(c.globalColor.GetAnsiColor())
	}
	c.term.PrintStr(str.String())
	c.term.SetCursor(c.xCursor+1, c.yCursor+1)
	c.SetVisibilityCursor(true)
	return
}

func (c *Gtui) innerLoop(keyb Keyboard.IKeyBoard) bool {
	//Keyboard
	c.IRefreshAll()
	c.AllineCursor()
	if !c.loop(c.keyb) {
		return false
	}
	c.IClear()
	c.IRefreshAll()
	return true
}
