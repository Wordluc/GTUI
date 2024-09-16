package Component

import (
	"GTUI/Core/Utils"
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
type ComponentManager struct {
	Map       *[][]*[]IComponent
	ChunkSize int
	nChunkX   int
	nChunkY   int
}

func Create(xSize, ySize, chunkSize int) (*ComponentManager, error) {
	totalError := Utils.NewError()
	xChunk := xSize/chunkSize + 1
	yChunk := ySize/chunkSize + 1
	if totalError.AddIf((xSize == 0 || ySize == 0), errors.New("invalid chunk size")) {
		return nil, totalError
	}

	matrix := make([][]*[]IComponent, xChunk)
	for i := range xChunk {
		matrix[i] = make([]*[]IComponent, yChunk)
	}
	return &ComponentManager{
		Map:       &matrix,
		ChunkSize: chunkSize,
		nChunkX:   xChunk,
		nChunkY:   yChunk,
	}, nil
}

func (s *ComponentManager) Add(comp IComponent) error {
	shape, e := comp.getShape()
	if e != nil {
		return e
	}
	finalX, finalY := shape.xPos+shape.Width, shape.yPos+shape.Height
	for i := shape.xPos; i < finalX; i++ {
		for j := shape.yPos; j < finalY; j++ {
			xC, yC := i/s.ChunkSize, j/s.ChunkSize
			if (xC >= s.nChunkX || yC >= s.nChunkY){
				return errors.New("component out of range")
			}
			comp.OnLeave()
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

func (s *ComponentManager) Search(x, y int) ([]IComponent, error) {
	var xiChunk int = x / s.ChunkSize
	var yiChunk int = y / s.ChunkSize
	totalError := Utils.NewError()
	totalError.AddIf(xiChunk >= len(*s.Map) || yiChunk >= len((*s.Map)[xiChunk]), errors.New("component out of range"))
	totalError.AddIf(xiChunk < 0 || yiChunk < 0, errors.New("component out of range"))
	totalError.AddIf(xiChunk >= s.nChunkX || yiChunk >= s.nChunkY, errors.New("component out of range"))
	eles := (*s.Map)[xiChunk][yiChunk]
	totalError.AddIf(eles == nil, errors.New("no node found at "+fmt.Sprintf("%d, %d", x, y)))

	if totalError.HasError() {
		return nil, totalError
	}

	res := []IComponent{}
	for _, ele := range *eles {
		shape, e := ele.getShape()
		if e != nil {
			return nil, e
		}
		if x >= shape.xPos && x < shape.xPos+shape.Width && y >= shape.yPos && y < shape.yPos+shape.Height {
			res = append(res, ele)
		}
	}
	return res, nil
}
