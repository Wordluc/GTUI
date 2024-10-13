package Component

import (
	"errors"
	"fmt"
)
type StreamCharacter struct {
	Get      func() chan string
	IChannel int
	Delete   func()
}
type IInteractiveShape interface {
	isOn(x, y int) bool
	getShapes() []BaseInteractiveShape
}
type BaseInteractiveShape struct {
	xPos   int
	yPos   int
	Width  int
	Height int
}
func (b BaseInteractiveShape) GetPos() (int, int) {
	return b.xPos, b.yPos
}
func (b BaseInteractiveShape) GetSize() (int, int) {
	return b.Width, b.Height
}

func (b *BaseInteractiveShape) isOn(x, y int) bool {
	return x >= b.xPos && x < b.xPos+b.Width && y >= b.yPos && y < b.yPos+b.Height
}

func (b *BaseInteractiveShape) getShapes() []BaseInteractiveShape {
	return []BaseInteractiveShape{*b}
}

type ComplexInteractiveShape struct {
	shapes []IInteractiveShape
}

func (c *ComplexInteractiveShape) isOn(x, y int) bool {
	for _, shape := range c.shapes {
		if shape.isOn(x, y) {
			return true
		}
	}
	return false
}

func (c *ComplexInteractiveShape) getShapes() []BaseInteractiveShape {
	var shapes []BaseInteractiveShape
	for _, shape := range c.shapes {
		shapes = append(shapes, shape.getShapes()...)
	}
	return shapes
}

func (c *ComplexInteractiveShape) addShape(shape IInteractiveShape) {
	c.shapes = append(c.shapes, shape)
}
//TODO: da ottimizzare
type ComponentM struct {
	Map        *[][]*[]IComponent
	components []IComponent
	ChunkSize  int
	nChunkX    int
	nChunkY    int
}

func Create(xSize, ySize, chunkSize int) *ComponentM {
	xChunk := xSize/chunkSize + 1
	yChunk := ySize/chunkSize + 1
	matrix := make([][]*[]IComponent, xChunk)
	for i := range xChunk {
		matrix[i] = make([]*[]IComponent, yChunk)
	}
	return &ComponentM{
		Map:       &matrix,
		ChunkSize: chunkSize,
		nChunkX:   xChunk,
		nChunkY:   yChunk,
	}
}

func (s *ComponentM) Clean() {
	for i := range s.nChunkX {
		(*s.Map)[i] = make([]*[]IComponent, s.nChunkY)
	}
}

func (s *ComponentM) Add(comp IComponent) error {//TODO: da ottimizzare
	if s == nil {
		return errors.New("component manager is nil")
	}
	shape, e := comp.getShape()
	if e != nil {
		return e
	}
	for _, baseShape := range shape.getShapes() {
		nChunkX := baseShape.Width / s.ChunkSize+1
		nChunkY := baseShape.Height / s.ChunkSize+1
		for i := 0; i < nChunkX; i++ {
			for j := 0; j < nChunkY; j++ {
				comp.OnOut(i, j)
				ele := (*s.Map)[i+baseShape.xPos/s.ChunkSize][j+baseShape.yPos/s.ChunkSize]
				if ele == nil {
					s.components = append(s.components, comp)
					(*s.Map)[i+baseShape.xPos/s.ChunkSize][j+baseShape.yPos/s.ChunkSize] = &[]IComponent{comp}
				} else {
					if (*ele)[len((*ele))-1] != comp {
						s.components = append(s.components, comp)
						(*ele) = append(*ele, comp)
					}
				}
			}
		}
	}
	return nil
}

func (s *ComponentM) Refresh() {
	s.Clean()
	for _, comp := range s.components {
		s.Add(comp)
	}
}

func (s *ComponentM) Search(x, y int) ([]IComponent, error) {
	if s == nil {
		return nil, errors.New("component manager is nil")
	}
	var xiChunk int = x / s.ChunkSize
	var yiChunk int = y / s.ChunkSize
	if xiChunk >= len(*s.Map) || yiChunk >= len((*s.Map)[xiChunk]) {
		return nil, fmt.Errorf("no node found at %d, %d", x, y)
	}
	//Check range
	if xiChunk < 0 || yiChunk < 0 {
		return nil, errors.New("no node found at " + fmt.Sprintf("%d, %d", x, y))
	}
	if xiChunk >= s.nChunkX || yiChunk >= s.nChunkY {
		return nil, errors.New("no node found at " + fmt.Sprintf("%d, %d", x, y))
	}
	eles := (*s.Map)[xiChunk][yiChunk]
	if eles == nil {
		return nil, errors.New("no node found at " + fmt.Sprintf("%d, %d", x, y))
	}
	res := []IComponent{}
	for _, ele := range *eles {
		shape, e := ele.getShape()
		if e != nil {
			return nil, e
		}
		if shape.isOn(x, y) {
			res = append(res, ele)
		}
	}
	return res, nil
}
