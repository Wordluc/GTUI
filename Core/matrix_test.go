package Core

import (
	"log"
	"testing"
)

func createMockElementMatrix(x, y, xSize, ySize int, name string) mockElementMatrix {
	return mockElementMatrix{
		x:     x,
		y:     y,
		xSize: xSize,
		ySize: ySize,
		name:  name,
	}
}

func testElement(x, y int, t *testing.T, matrix *MatrixHandler, nametest string, expected ...string) {
	result := matrix.SearchInAllLayers(x, y)
	for l := range result {
		for _, c := range result[l] {
			log.Println(c.GetLayer(), c.(mockElementMatrix).name)
		}
		log.Println("--")
	}
	if len(result[0]) != len(expected) {
		t.Error(nametest, "Expected ", len(expected), "elements ,got ", len(result), ";pos", x, y)
		return
	}
	for i := range expected {
		if result[0][i].(mockElementMatrix).name != expected[i] {
			t.Error(nametest, ":Expected ", expected[i], " got ", result[0][i].(mockElementMatrix).name)
		}
	}
}
func TestMatrix1(t *testing.T) {
	matrix := CreateMatrixHandler(100, 100, 5)
	matrix.AddElement(createMockElementMatrix(5, 5, 10, 10, "test1"))
	matrix.AddElement(createMockElementMatrix(7, 7, 10, 10, "test2"))
	testElement(6, 6, t, matrix, "prima verifica 0", "test1")
	testElement(8, 8, t, matrix, "prima verifica 1", "test1", "test2")
	testElement(16, 16, t, matrix, "prima verifica 2", "test2")
}
func TestMatrixSeparatedComponents1(t *testing.T) {
	matrix := CreateMatrixHandler(100, 100, 5)
	matrix.AddElement(createMockElementMatrix(5, 5, 10, 10, "test1"))
	matrix.AddElement(createMockElementMatrix(17, 17, 10, 10, "test2"))
	testElement(6, 6, t, matrix, "terza verifica 0", "test1")
	testElement(8, 8, t, matrix, "terza verifica 1", "test1")
	testElement(18, 18, t, matrix, "terza verifica 2", "test2")
}
func TestMatrixSeparatedComponents2(t *testing.T) {
	matrix := CreateMatrixHandler(100, 100, 5)
	matrix.AddElement(createMockElementMatrix(17, 17, 10, 10, "test2"))
	matrix.AddElement(createMockElementMatrix(0, 0, 10, 10, "test1"))
	testElement(6, 6, t, matrix, "quarta verifica 0", "test1")
	testElement(8, 8, t, matrix, "quarta verifica 1", "test1")
	testElement(18, 18, t, matrix, "quarta verifica 2", "test2")
}
func TestMatrixMixedCase1(t *testing.T) {
	matrix := CreateMatrixHandler(100, 100, 5)
	matrix.AddElement(createMockElementMatrix(17, 17, 10, 10, "test2"))
	matrix.AddElement(createMockElementMatrix(20, 20, 10, 10, "test3"))
	matrix.AddElement(createMockElementMatrix(0, 0, 10, 10, "test1"))
	testElement(2, 2, t, matrix, "quinta verifica 0", "test1")
	testElement(8, 8, t, matrix, "quinta verifica 1", "test1")
	testElement(18, 18, t, matrix, "quinta verifica 2", "test2")
	testElement(21, 21, t, matrix, "quinta verifica 3", "test2", "test3")
}
func TestMatrixMixedCase2(t *testing.T) {
	matrix := CreateMatrixHandler(100, 100, 5)
	matrix.AddElement(createMockElementMatrix(70, 0, 20, 30, "button"))
	matrix.AddElement(createMockElementMatrix(0, 0, 50, 60, "textBox"))
	matrix.AddElement(createMockElementMatrix(85, 0, 10, 10, "button1"))
	testElement(1, 1, t, matrix, "sesta verifica 0", "textBox")
}
func TestMatrixRefresh(t *testing.T) {
	matrix := CreateMatrixHandler(100, 100, 5)
	matrix.AddElement(createMockElementMatrix(70, 0, 20, 30, "button"))
	matrix.AddElement(createMockElementMatrix(0, 0, 50, 60, "textBox"))
	matrix.AddElement(createMockElementMatrix(85, 0, 10, 10, "button1"))
	matrix = matrix.Refresh(100, 100, 5)
	testElement(1, 1, t, matrix, "sesta verifica 0", "textBox")
	testElement(1, 1, t, matrix, "settima verifica 0", "textBox")
}
func TestMatrixElementAbove(t *testing.T) {
	matrix := CreateMatrixHandler(100, 100, 5)
	matrix.AddElement(createMockElementMatrix(40, 40, 40, 40, "button1"))
	matrix.AddElement(createMockElementMatrix(40, 0, 10, 10, "textBox"))
	testElement(41, 1, t, matrix, "ottava verifica 0", "textBox")
	testElement(41, 41, t, matrix, "ottava verifica 1", "button1")
}
func TestMatrixElementNextToAnotherElement(t *testing.T) {
	matrix := CreateMatrixHandler(100, 100, 5)
	matrix.AddElement(createMockElementMatrix(10, 5, 10, 5, "button1"))
	matrix.AddElement(createMockElementMatrix(25, 0, 50, 50, "textBox"))
	testElement(26, 11, t, matrix, "nona verifica 1", "textBox")
	for i := 1; i < 50; i = i + 10 {
		testElement(26, i, t, matrix, "nona verifica 1", "textBox")
		print("---\n")
	}
}
func TestMatrixElementNextToAnotherElementSameH(t *testing.T) {
	matrix := CreateMatrixHandler(100, 100, 5)
	matrix.AddElement(createMockElementMatrix(10, 5, 10, 5, "button1"))
	matrix.AddElement(createMockElementMatrix(30, 5, 50, 5, "textBox"))
	testElement(31, 6, t, matrix, "nona verifica 1", "textBox")
	for i := 6; i < 5; i++ {
		testElement(31, i, t, matrix, "nona verifica 1", "textBox")
		print("---\n")
	}
}
func TestAdjacentComponentX(t *testing.T) {
	matrix := CreateMatrixHandler(100, 100, 20)
	matrix.AddElement(createMockElementMatrix(30, 10, 5, 5, "A"))
	matrix.AddElement(createMockElementMatrix(31, 10, 5, 5, "B"))
	matrix.AddElement(createMockElementMatrix(33, 10, 5, 5, "C"))
	//	matrix.AddElement(createMockElementMatrix(31, 10, 5, 5, "B"))
	//	matrix.AddElement(createMockElementMatrix(32, 10, 5, 5, "D"))
	res := matrix.searchInAllLayersRaw(34, 11)

	getEle := func(name string, elements []WrapperElement[ElementMatrix]) WrapperElement[ElementMatrix] {
		for i := range elements {

			if elements[i].object.(mockElementMatrix).name == name {
				return elements[i]
			}
		}
		return WrapperElement[ElementMatrix]{}
	}
	var ele = getEle("A", res[0])
	for ele.object != nil {
		println(ele.object.(mockElementMatrix).name)
		ele = *ele.right
	}
}
