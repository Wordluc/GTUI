package Core

import "math"

type ElementTree interface {
	GetPos() (int, int)
	GetSize() (int, int)
	GetLayer() Layer
}
type MatrixElement[T ElementTree] struct {
	object T
}
type matrix[T any] [][][]MatrixElement[ElementTree] //rows,columns,elements
type MatrixElements struct {
	matrix                    matrix[MatrixElement[ElementTree]]
	nColumns, nRows, sizeCell int
}

func CreateMatrixElements(nColumns, nRows, size int) *MatrixElements {
	var matrix matrix[MatrixElement[ElementTree]] = make(matrix[MatrixElement[ElementTree]], nRows+1)
	for y := range nRows {
		matrix[y] = make([][]MatrixElement[ElementTree], nColumns+1)
	}
	return &MatrixElements{
		nColumns: nColumns,
		nRows:    nRows,
		matrix:   matrix,
		sizeCell: size,
	}
}

// start thinking how to link stuff together, the objects have to know that on the right there are some objects ecc....
// so maybe use a wrapper for ElementTree, where i can store information about my neighbour
func (m *MatrixElements) addElement(element ElementTree) {
	divide := func(a, b int) int {
		at := float64(a)
		bt := float64(b)
		return int(math.Ceil(at / bt))
	}

	x, y := element.GetPos()
	width, height := element.GetSize()

	startingColumn := x / m.sizeCell
	startingRow := y / m.sizeCell

	column := divide(width, m.sizeCell)
	row := divide(height, m.sizeCell)

	for ir := range row {
		for iw := range column {
			if ir+startingRow >= m.nRows {
				continue
			}
			if iw+startingColumn >= m.nColumns {
				continue
			}
			m.matrix[ir+startingRow][iw+startingColumn] =
				append(m.matrix[ir+startingRow][iw+startingColumn], MatrixElement[ElementTree]{object: element})
		}
	}
}

func isCollidingWith(xToSearch, yToSearch int, xElement, yElement, wElement, hElement int) bool {
	if xToSearch >= xElement && xToSearch < xElement+wElement && yToSearch >= yElement && yToSearch < yElement+hElement {
		return true
	}
	return false
}

func (m *MatrixElements) search(x, y int) (res []ElementTree) {
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
	elements []ElementTree
	layers   []*MatrixElements
	w, h     int
}

func CreateMatrixHandler(w, h int) *MatrixHandler {
	return &MatrixHandler{
		w: w,
		h: h,
	}
}

func (m *MatrixHandler) SearchInAllLayers(x, y int) (res [][]ElementTree) {
	for i := range m.layers {
		if m.layers[i] != nil {
			res = append(res, m.layers[i].search(x, y))
		}
	}
	return res
}

func (m *MatrixHandler) AddElement(eles ...ElementTree) {
	for _, ele := range eles {
		layer := ele.GetLayer()
		for int(layer) >= len(m.layers) {
			m.layers = append(m.layers, nil)
		}
		if m.layers[layer] == nil {
			m.layers[layer] = CreateMatrixElements(m.w, m.h, 10)
		}
		m.layers[layer].addElement(ele)
		m.elements = append(m.elements, ele)
	}
}

func (m *MatrixHandler) GetLayerN() int {
	return len(m.layers)
}

func (m *MatrixHandler) Refresh() {
	handler := CreateMatrixHandler(m.w, m.h)
	handler.AddElement(m.elements...)
	*m = *handler
}

func (m *MatrixHandler) ExecuteOnLayer(layer int, cond func(ele ElementTree) bool) {

	if m.layers[layer] == nil {
		return
	}
	for iy := range m.h {
		for ix := range m.w {
			for _, ele := range m.layers[layer].matrix[iy][ix] {
				if !cond(ele.object) {
					return
				}
			}
		}
	}
}
