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
type InteractiveShape struct {
	xPos   int
	yPos   int
	Width  int
	Height int
}
type ComponentM struct {
	Map       *[][]*[]IComponent
	ChunkSize int
	nChunkX   int
	nChunkY   int
}

func Create(xSize, ySize, chunkSize int) *ComponentM {
	xChunk := xSize / chunkSize
	yChunk := ySize / chunkSize
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

func (s *ComponentM) Add(comp IComponent) error {
	if s == nil {
		return errors.New("component manager is nil")
	}
	shape, e := comp.getShape()
	if e != nil {
		return e
	}
	finalX, finalY := shape.xPos+shape.Width, shape.yPos+shape.Height
	for i := shape.xPos; i <= finalX; i += 1 {
		for j := shape.yPos; j <= finalY; j += 1 {
			xC, yC := i/s.ChunkSize, j/s.ChunkSize
			ele := (*s.Map)[xC][yC]
			if ele == nil {
				(*s.Map)[xC][yC] = &[]IComponent{comp}
			} else {
				if (*ele)[len((*ele))-1] != comp {
					(*ele) = append(*ele, comp)
				}
			}
		}
	}
	return nil
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
			if x>= shape.xPos && x < shape.xPos+shape.Width && y >= shape.yPos && y < shape.yPos+shape.Height {
			res = append(res, ele)
		}
	}
	return res, nil
}
