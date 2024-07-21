package Component

import (
		"errors"
	)

type Search struct {
	Map       *[][]*[]IComponent
	ChunkSize int
}

func Create(xSize, ySize, chunkSize int) *Search {
	xChunk := xSize / chunkSize
	yChunk := ySize / chunkSize
	matrix := make([][]*[]IComponent, xChunk)
	for i := range matrix {
		matrix[i] = make([]*[]IComponent, yChunk)
	}
	return &Search{
		Map: &matrix,
	}
}

func (s *Search) Add(comp IComponent) error {
	return comp.insertingToMap(s)
}

func (s *Search) Search(x, y int) (IComponent, error) {
	xiChunk := x / s.ChunkSize
	yiChunk := y / s.ChunkSize
	eles := (*s.Map)[xiChunk][yiChunk]
	if eles == nil {
		return nil, errors.New("no node founded")
	}
	for _, ele := range *eles {
		if ele.isOn(x, y) {
			return ele, nil
		}
	}
	return nil, errors.New("no node founded")
}
