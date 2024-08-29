package GTUI

import (
	"GTUI/Core"
	"GTUI/Core/Component"
	"GTUI/Core/Utils"
	"GTUI/Core/Utils/Color"
	"GTUI/Keyboard"
	"GTUI/Terminal"
	"strings"
	"time"
)

type Gtui struct {
	globalColor      Color.Color
	term             Terminal.ITerminal
	keyb             Keyboard.IKeyBoard
	buff             []Core.IEntity
	componentManager *Component.ComponentM
	xCursor          int
	yCursor          int
	xSize            int
	ySize            int
	cursorVisibility bool
	loop             Keyboard.Loop
}

func (c *Gtui) SetCur(x, y int) {
	if x < 0 || y < 0 || x >= c.xSize || y >= c.ySize {
		return
	}
	c.SetVisibilityCursor(true)
	compPreSet, _ := c.componentManager.Search(c.xCursor, c.yCursor)
	comps, _ := c.componentManager.Search(x, y)
	inPreButNotInPost := Utils.Diff(comps, compPreSet)
	inPostButNotInPre := Utils.Diff(compPreSet, comps)
	for _, e := range inPreButNotInPost {
		e.OnLeave()
	}
	for _, e := range inPostButNotInPre {
		e.OnHover()
	}

	c.xCursor = x
	c.yCursor = y

	for _, comp := range comps {
		if ci, ok := comp.(Component.ICursorInteragibleComponent); ok {
			if ci.IsWritable() {
				deltax, deltay := ci.DiffTotalToXY(x, y)
				if deltax > 0 {
					deltax = 0
				}
				if deltay > 0 {
					deltay = 0
				}
				c.yCursor = y + deltay
				c.xCursor = x + deltax
				ci.SetCurrentPos(c.xCursor, c.yCursor)
				break
			}
		}
	}
	c.term.SetCursor(c.xCursor+1, c.yCursor+1)
	c.SetVisibilityCursor(false)
}
func (c *Gtui) GetCur() (int, int) {
	return c.xCursor, c.yCursor
}
func NewGtui(loop Keyboard.Loop,keyb Keyboard.IKeyBoard,term Terminal.ITerminal) (*Gtui, error) {
	xSize, ySize := term.Size()
	componentManager := Component.Create(xSize, ySize, 5)
	return &Gtui{globalColor: Color.GetDefaultColor(), cursorVisibility: true, loop: loop, term: term,keyb: keyb, buff: make([]Core.IEntity, 0), componentManager: componentManager, xCursor: 0, yCursor: 0, xSize: xSize, ySize: ySize}, nil
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

func (c *Gtui) Start(){
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

func (c *Gtui) InsertComponent(component Component.IComponent) {
	c.buff = append(c.buff, component.GetGraphics())
	c.componentManager.Add(component)
}

func (c *Gtui) Click(x, y int) error {
	resultArray, e := c.componentManager.Search(x, y)
	if e != nil {
		return e
	}
	for i := range resultArray {
		resultArray[i].OnClick() //dare la possibilita' di scegliere il timer
		time.AfterFunc(time.Millisecond*1000, func() {
			resultArray[i].OnRelease()
			c.IRefreshAll()
		})
	}
	return nil
}
func (c *Gtui) AllineCursor() {
	x, y := c.GetCur()
	comps, _ := c.componentManager.Search(x, y)
	for _, comp := range comps {
		if ci, ok := comp.(Component.ICursorInteragibleComponent); ok {
			if ci.IsWritable() {
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
	c.SetVisibilityCursor(false)
	var str strings.Builder
	for _, b := range c.buff {
		str.WriteString(b.GetAnsiCode(c.globalColor))
		str.WriteString(c.globalColor.GetAnsiColor())
	}
	c.term.PrintStr(str.String())
	c.SetVisibilityCursor(true)

	c.term.SetCursor(c.xCursor+1, c.yCursor+1)
	return
}

func (c *Gtui) innerLoop(keyb Keyboard.IKeyBoard) bool {
	c.IRefreshAll()
	c.AllineCursor()
   if !c.loop(c.keyb){
	  return false
	}
	c.IClear()
	c.IRefreshAll()
	return true
}
