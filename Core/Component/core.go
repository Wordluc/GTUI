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

type ComponentM struct {
	Map       *[][]*[]IComponent
	components []IComponent
	ChunkSize int
	nChunkX   int
	nChunkY   int
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

func (s *ComponentM) Add(comp IComponent) error {
	if s == nil {
		return errors.New("component manager is nil")
	}
	shape, e := comp.getShape()
	if e != nil {
		return e
	}
	for _, baseShape := range shape.getShapes() {
		finalX, finalY := baseShape.xPos+baseShape.Width, baseShape.yPos+baseShape.Height
		for i := baseShape.xPos; i < finalX; i++ {
			for j := baseShape.yPos; j < finalY; j++ {
				xC, yC := i/s.ChunkSize, j/s.ChunkSize
				if xC >= s.nChunkX || yC >= s.nChunkY {
					return errors.New("component out of range")
				}
				comp.OnOut(i, j)
				ele := (*s.Map)[xC][yC]
				if ele == nil {
					(*s.Map)[xC][yC] = &[]IComponent{comp}
				} else {
					if (*ele)[len((*ele))-1] != comp {
						(*ele) = append(*ele, comp)
					}
				}
				s.components = append(s.components, comp)
			}
		}
	}
	return nil
}
func (s *ComponentM) Refresh()  {
	s.Clean()
	for _,comp:= range s.components{
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
