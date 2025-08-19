package Core

type ElementTree interface {
	GetPos() (int, int)
	GetSize() (int, int)
	GetLayer() Layer
}

type WrapperElement[T ElementTree] struct {
	object T
}
type matrix[T any] [][][]T //rows,columns,elements
type MatrixLayer struct {
	matrix                    matrix[WrapperElement[ElementTree]]
	nColumns, nRows, sizeCell int
	elements                  []ElementTree
}

func CreateMatrixLayer(nColumns, nRows, size int) *MatrixLayer {
	nColumns++
	nRows++
	var matrix matrix[WrapperElement[ElementTree]] = make(matrix[WrapperElement[ElementTree]], nRows)
	for y := range nRows {
		matrix[y] = make([][]WrapperElement[ElementTree], nColumns)
	}
	return &MatrixLayer{
		nColumns: nColumns,
		nRows:    nRows,
		matrix:   matrix,
		sizeCell: size,
	}
}

// start thinking how to link stuff together, the objects have to know that on the right there are some objects ecc....
// so maybe use a wrapper for ElementTree, where i can store information about my neighbour
func (m *MatrixLayer) addElement(element ElementTree) {
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
				append(m.matrix[currentRow][currentColumn], WrapperElement[ElementTree]{object: element})
		}
	}
}

func isCollidingWith(xToSearch, yToSearch int, xElement, yElement, wElement, hElement int) bool {
	if xToSearch >= xElement && xToSearch < xElement+wElement && yToSearch >= yElement && yToSearch < yElement+hElement {
		return true
	}
	return false
}

func (m *MatrixLayer) search(x, y int) (res []ElementTree) {
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

// TODO: i have to pass the size of the screen and from that compute the w and h
func CreateMatrixHandler(w, h, sizeCell int) *MatrixHandler {
	return &MatrixHandler{
		nColumns: w / sizeCell,
		nRows:    h / sizeCell,
		size:     sizeCell,
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

func (m *MatrixHandler) AddElement(elements ...ElementTree) {
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

func (m *MatrixHandler) ExecuteOnLayer(layer int, cond func(ele ElementTree) bool) {
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
func (m *MatrixHandler) GetCollidingElement(layer int, elementWhichCollides ElementTree) (res []ElementTree) {
	x, y := elementWhichCollides.GetPos()
	width, height := elementWhichCollides.GetSize()
	var tElement WrapperElement[ElementTree]
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
