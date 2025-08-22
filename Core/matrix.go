package Core

import (
	"math"
)

type ElementMatrix interface {
	GetPos() (int, int)
	GetSize() (int, int)
	GetLayer() Layer
}

type WrapperElement[T ElementMatrix] struct {
	object T
	left   *WrapperElement[T]
	right  *WrapperElement[T]
	up     *WrapperElement[T]
	down   *WrapperElement[T]
}
type matrix[T any] [][][]T //rows,columns,elements
type MatrixLayer struct {
	matrix                    matrix[*WrapperElement[ElementMatrix]]
	nColumns, nRows, sizeCell int
	elements                  []ElementMatrix
}

func CreateMatrixLayer(nColumns, nRows, size int) *MatrixLayer {
	nColumns++
	nRows++
	var matrix matrix[*WrapperElement[ElementMatrix]] = make(matrix[*WrapperElement[ElementMatrix]], nRows)
	for y := range nRows {
		matrix[y] = make([][]*WrapperElement[ElementMatrix], nColumns)
	}
	return &MatrixLayer{
		nColumns: nColumns,
		nRows:    nRows,
		matrix:   matrix,
		sizeCell: size,
	}
}

// TODO: refactor!!!!
func (m *MatrixLayer) createElement(row, column int, element ElementMatrix) *WrapperElement[ElementMatrix] {
	res := &WrapperElement[ElementMatrix]{object: element}
	currentCell := m.matrix[row][column]
	x, y := element.GetPos()
	var left, right, up, down *WrapperElement[ElementMatrix]
	var leftX, rightX, upY, downY int = int(math.Inf(1)), int(math.Inf(1)), int(math.Inf(1)), int(math.Inf(1))
	var disX, disY int
	isAbsoluteValueLessOrEqual := func(a, b int) bool {
		return math.Abs(float64(a)) <= math.Abs(float64(b))
	}
	for i := range currentCell {
		xM, yM := currentCell[i].object.GetPos()
		disX, disY = xM-x, yM-y
		if disX < 0 {
			//LEFT OF ELEMENT
			if isAbsoluteValueLessOrEqual(disX, leftX) {
				leftX = disX
				left = currentCell[i]
			}
		} else {
			//RIGHT OF ELEMENT
			if isAbsoluteValueLessOrEqual(disX, rightX) {
				rightX = disX
				right = currentCell[i]
			}
		}
		if disY < 0 {
			//UP OF ELEMENT
			if isAbsoluteValueLessOrEqual(disY, upY) {
				upY = disY
				up = currentCell[i]
			}
		} else {
			//DOWN OF ELEMENT
			if isAbsoluteValueLessOrEqual(disY, downY) {
				downY = disY
				down = currentCell[i]
			}
		}
	}

	if isAbsoluteValueLessOrEqual(leftX, rightX) {
		if left != nil {
			if left.right == nil {
				left.right = res
				res.left = left
			} else {
				t := left.right
				left.right = res
				res.left = left

				t.left = res
				res.right = t
			}
		}
	} else {
		if right != nil {
			if right.left == nil {
				right.left = res
				res.right = right

			} else {
				t := right.left
				right.left = res
				res.right = right

				t.right = res
				res.left = t
			}
		}
	}

	if isAbsoluteValueLessOrEqual(upY, downY) {
		if up != nil {
			if up.down == nil {
				up.down = res
				res.up = up

			} else {
				t := up.down
				up.down = res
				res.up = up

				t.up = res
				res.down = t
			}
		}
	} else {
		if down != nil {
			if down.up == nil {
				down.up = res
				res.down = down

			} else {
				t := down.up
				down.up = res
				res.down = down

				t.down = res
				res.up = t
			}
		}
	}

	return res
}

func (m *MatrixLayer) addElement(element ElementMatrix) {
	m.elements = append(m.elements, element)
	x, y := element.GetPos()
	width, height := element.GetSize()

	startingColumn := x / m.sizeCell
	startingRow := y / m.sizeCell

	column := width / m.sizeCell
	row := height / m.sizeCell

	for ir := range row + 1 {
		for iw := range column + 1 {
			currentRow, currentColumn := ir+startingRow, iw+startingColumn
			if ir+startingRow >= m.nRows {
				continue
			}
			if iw+startingColumn >= m.nColumns {
				continue
			}
			m.matrix[currentRow][currentColumn] =
				append(m.matrix[currentRow][currentColumn], m.createElement(currentRow, currentColumn, element))
		}
	}
}

func isCollidingWith(xToSearch, yToSearch int, xElement, yElement, wElement, hElement int) bool {
	if xToSearch >= xElement && xToSearch < xElement+wElement && yToSearch >= yElement && yToSearch < yElement+hElement {
		return true
	}
	return false
}

func (m *MatrixLayer) searchRaw(x, y int) (res []WrapperElement[ElementMatrix]) {
	row := int(y / m.sizeCell)
	column := int(x / m.sizeCell)
	if row >= m.nRows {
		return res
	}
	if column >= m.nColumns {
		return res
	}
	var xEle, yEle, wEle, hEle int
	for i := range m.matrix[row][column] {
		xEle, yEle = m.matrix[row][column][i].object.GetPos()
		wEle, hEle = m.matrix[row][column][i].object.GetSize()
		if isCollidingWith(x, y, xEle, yEle, wEle, hEle) {
			res = append(res, *m.matrix[row][column][i])
		}
	}
	return res
}
func (m *MatrixLayer) search(x, y int) (res []ElementMatrix) {
	row := int(y / m.sizeCell)
	column := int(x / m.sizeCell)
	if row >= m.nRows {
		return res
	}
	if column >= m.nColumns {
		return res
	}
	var xEle, yEle, wEle, hEle int
	for i := range m.matrix[row][column] {
		xEle, yEle = m.matrix[row][column][i].object.GetPos()
		wEle, hEle = m.matrix[row][column][i].object.GetSize()
		if isCollidingWith(x, y, xEle, yEle, wEle, hEle) {
			res = append(res, m.matrix[row][column][i].object)
		}
	}
	return res
}

type MatrixHandler struct {
	layers                []*MatrixLayer
	nColumns, nRows, size int
}

func CreateMatrixHandler(w, h, sizeCell int) *MatrixHandler {
	return &MatrixHandler{
		nColumns: w / sizeCell,
		nRows:    h / sizeCell,
		size:     sizeCell,
	}
}

func (m *MatrixHandler) searchInAllLayersRaw(x, y int) (res [][]WrapperElement[ElementMatrix]) {
	for i := range m.layers {
		if m.layers[i] != nil {
			res = append(res, m.layers[i].searchRaw(x, y))
		}
	}
	return res
}

func (m *MatrixHandler) SearchInAllLayers(x, y int) (res [][]ElementMatrix) {
	for i := range m.layers {
		if m.layers[i] != nil {
			res = append(res, m.layers[i].search(x, y))
		}
	}
	return res
}
func (m *MatrixHandler) AddElement(elements ...ElementMatrix) {
	for _, ele := range elements {
		layer := ele.GetLayer()
		for int(layer) >= len(m.layers) {
			m.layers = append(m.layers, nil)
		}
		if m.layers[layer] == nil {
			m.layers[layer] = CreateMatrixLayer(m.nColumns, m.nRows, m.size)

		}
		m.layers[layer].addElement(ele)
	}
}

func (m *MatrixHandler) GetLayerN() int {
	return len(m.layers)
}

func (m *MatrixHandler) Refresh(w, h, sizeCell int) *MatrixHandler {
	handler := CreateMatrixHandler(w, h, sizeCell)
	for i := range m.layers {
		if m.layers[i] == nil {
			continue
		}
		handler.AddElement(m.layers[i].elements...)
	}
	return handler
}

func (m *MatrixHandler) ExecuteOnLayer(layer int, cond func(ele ElementMatrix) bool) {
	if m.layers[layer] == nil {
		return
	}
	for _, ele := range m.layers[layer].elements {
		if !cond(ele) {
			return
		}
	}
}

func (t *WrapperElement[ElementTree]) isCollidingWithNode(x, y int, width, height int) bool {
	elementX, elementY := t.object.GetPos()
	elementWidth, elementHeight := t.object.GetSize()
	if x >= elementX && x < elementX+elementWidth && y >= elementY && y < elementY+elementHeight {
		return true
	}
	if x+width >= elementX && x+width < elementX+elementWidth && y >= elementY && y < elementY+elementHeight {
		return true
	}
	if x >= elementX && x < elementX+elementWidth && y+height >= elementY && y+height < elementY+elementHeight {
		return true
	}
	if x+width >= elementX && x+width < elementX+elementWidth && y+height >= elementY && y+height < elementY+elementHeight {
		return true
	}
	return false
}

// to optimize ? maybe with some chacing
func (m *MatrixHandler) GetCollidingElements(layer int, elementWhichCollides ElementMatrix) (res []ElementMatrix) {
	x, y := elementWhichCollides.GetPos()
	width, height := elementWhichCollides.GetSize()
	var tElement WrapperElement[ElementMatrix]
	for i := range m.layers[layer].elements {
		if m.layers[layer].elements[i] == elementWhichCollides {
			continue
		}
		tElement.object = m.layers[layer].elements[i]
		if tElement.isCollidingWithNode(x, y, width, height) {
			res = append(res, m.layers[layer].elements[i])
		}
	}
	return res
}
