package Component

import (
	"errors"
	"fmt"
)

type ComponentM struct {
	Map       *[][]*[]IComponent
	ChunkSize int
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
	}
}

func (s *ComponentM) Add(comp IComponent) error {
	return comp.insertingToMap(s)
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
	eles := (*s.Map)[xiChunk][yiChunk]
	if eles == nil {
		return nil, fmt.Errorf("no node found at %d, %d", x, y)
	}
	res := []IComponent{}
	for _, ele := range *eles {
		if !ele.isOn(x, y) {
		   continue
		}
		res = append(res, ele)
	}
	return res, nil
}
