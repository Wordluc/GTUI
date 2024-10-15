package Core

import "testing"
type mockElementTree struct {
	x int
	y int
	xSize int
	ySize int
	name string
}
func createMockElementTree(x,y,xSize,ySize int,name string) mockElementTree {
	return mockElementTree{
		x: x,
		y: y,
		xSize: xSize,
		ySize: ySize,
		name: name,
	}
}
func (e mockElementTree) GetPos() (int, int) {
	return e.x, e.y
}
func (e mockElementTree) GetSize() (int, int) {
	return e.xSize, e.ySize
}
func testElement(x,y int,t *testing.T,tree *TreeManager[mockElementTree] ,nametest string,expected ...string) {
	result,err := tree.Search(x,y)
	if err != nil {
		t.Error(err)
		return
	}
	if len(result) != len(expected) {
		t.Error(nametest,":Expected 1, got ",len(result))
		return
	}
	for i := 0; i < len(expected); i++ {
		if result[i].name != expected[i] {
			t.Error(nametest,":Expected ",expected[i]," got ",result[i].name)
		}
	}
}
func TestTree1(t *testing.T) {
	tree := CreateTreeManager[mockElementTree]()
	tree.AddElement(createMockElementTree(5,5,10,10,"test1"))
	tree.AddElement(createMockElementTree(7,7,10,10,"test2"))
	testElement(6,6,t,tree,"prima verifica 0","test1")
	testElement(8,8,t,tree,"prima verifica 1","test1","test2")
	testElement(16,16,t,tree,"prima verifica 2","test2")
}
func TestTree2(t *testing.T) {
	tree := CreateTreeManager[mockElementTree]()
	tree.AddElement(createMockElementTree(7,7,10,10,"test2"))
	tree.AddElement(createMockElementTree(5,5,10,10,"test1"))
//	testElement(6,6,t,tree,"seconda verifica 0","test1")
//	testElement(8,8,t,tree,"seconda verifica 1","test2","test1")
	testElement(16,16,t,tree,"seconda verifica 2","test2")
}
