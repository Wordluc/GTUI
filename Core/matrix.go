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

func (m *MatrixLayer) addElement(element *WrapperElement[ElementMatrix]) {
	x, y := element.object.GetPos()
	width, height := element.object.GetSize()

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
				append(m.matrix[currentRow][currentColumn], element)
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
	elements              []*WrapperElement[ElementMatrix]
}

func CreateMatrixHandler(w, h, sizeCell int) *MatrixHandler {
	return &MatrixHandler{
		nColumns: w / sizeCell,
		nRows:    h / sizeCell,
		size:     sizeCell,
	}
}

// TODO: refactor!!!!
func (m *MatrixHandler) createElement(element ElementMatrix) *WrapperElement[ElementMatrix] {
	res := &WrapperElement[ElementMatrix]{object: element}
	x, y := element.GetPos()
	var left, right, up, down *WrapperElement[ElementMatrix]
	var leftX, rightX, upY, downY int = int(math.Inf(1)), int(math.Inf(1)), int(math.Inf(1)), int(math.Inf(1))
	var disX, disY int
	isAbsoluteValueLessOrEqual := func(a, b int) bool {
		return math.Abs(float64(a)) <= math.Abs(float64(b))
	}
	for i := range m.elements {
		xM, yM := m.elements[i].object.GetPos()
		disX, disY = xM-x, yM-y
		if disX < 0 {
			//LEFT OF ELEMENT
			if isAbsoluteValueLessOrEqual(disX, leftX) {
				leftX = disX
				left = m.elements[i]
			}
		} else {
			//RIGHT OF ELEMENT
			if isAbsoluteValueLessOrEqual(disX, rightX) {
				rightX = disX
				right = m.elements[i]
			}
		}
		if disY < 0 {
			//UP OF ELEMENT
			if isAbsoluteValueLessOrEqual(disY, upY) {
				upY = disY
				up = m.elements[i]
			}
		} else {
			//DOWN OF ELEMENT
			if isAbsoluteValueLessOrEqual(disY, downY) {
				downY = disY
				down = m.elements[i]
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

func (m *MatrixHandler) SearchInAllLayers(x, y int) (res [][]ElementMatrix) {
	for i := range m.layers {
		if m.layers[i] != nil {
			res = append(res, m.layers[i].search(x, y))
		}
	}
	return res
}
func (m *MatrixHandler) AddElement(elements ...ElementMatrix) {
	for i := range elements {

		layer := elements[i].GetLayer()
		for int(layer) >= len(m.layers) {
			m.layers = append(m.layers, nil)
		}
		if m.layers[layer] == nil {
			m.layers[layer] = CreateMatrixLayer(m.nColumns, m.nRows, m.size)

		}
		ele := m.createElement(elements[i])
		m.layers[layer].addElement(ele)
		m.elements = append(m.elements, ele)
	}
}

func (m *MatrixHandler) GetLayerN() int {
	return len(m.layers)
}

func (m *MatrixHandler) Refresh(w, h, sizeCell int) *MatrixHandler {
	handler := CreateMatrixHandler(w, h, sizeCell)
	for i := range m.elements {
		handler.AddElement(m.elements[i].object)
	}
	return handler
}

func (m *MatrixHandler) ExecuteOnLayer(layer int, cond func(ele ElementMatrix) bool) {
	if m.layers[layer] == nil {
		return
	}
	for _, ele := range m.elements {
		if ele.object.GetLayer() != Layer(layer) {
			continue
		}
		if !cond(ele.object) {
			return
		}
	}
}

type Direction int8

const (
	Up    Direction = iota
	Down  Direction = iota
	Right Direction = iota
	Left  Direction = iota
)

func (m *MatrixHandler) GetNextElement(elementWhereTOStart ElementMatrix, whereToGo Direction) ElementMatrix {
	var ele *WrapperElement[ElementMatrix]
	for i := range m.elements {
		if m.elements[i].object == elementWhereTOStart {
			ele = m.elements[i]
			break
		}
	}
	if ele == nil {
		return nil
	}
	getObject := func(ele *WrapperElement[ElementMatrix]) ElementMatrix {
		if ele == nil {
			return nil
		}
		return ele.object
	}
	switch whereToGo {
	case Up:
		return getObject(ele.up)
	case Down:
		return getObject(ele.down)
	case Right:
		return getObject(ele.right)
	case Left:
		return getObject(ele.left)
	}
	return nil
}
