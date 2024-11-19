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
func (e mockElementTree) GetLayer() Layer {
	return L1
}
func (e mockElementTree) SetLayer(layer Layer){
	return
}

func testElement(x,y int,t *testing.T,tree *TreeManager[mockElementTree] ,nametest string,expected ...string) {
	result,err := tree.Search(L1,x,y)
	if err != nil {
		t.Error(err)
		return
	}
	if len(result) != len(expected) {
		t.Error(nametest,"Expected ",len(expected), ":got ",len(result),";pos",x,y)
		return
	}
	for i := 0; i < len(expected); i++ {
		if result[i].name != expected[i] {
			t.Error(nametest,":Expected ",expected[i]," got ",result[i].name)
		}
	}
}
func TestTree1(t *testing.T) {
	tree := CreateTreeManager[mockElementTree](6)
	tree.AddElement(createMockElementTree(5,5,10,10,"test1"))
	tree.AddElement(createMockElementTree(7,7,10,10,"test2"))
	testElement(6,6,t,tree,"prima verifica 0","test1")
	testElement(8,8,t,tree,"prima verifica 1","test1","test2")
	testElement(16,16,t,tree,"prima verifica 2","test2")
}
func TestTree2(t *testing.T) {
	tree := CreateTreeManager[mockElementTree](6)
	tree.AddElement(createMockElementTree(7,7,10,10,"test2"))
	tree.AddElement(createMockElementTree(5,5,10,10,"test1"))
	if tree.root[0].smaller.element.name!="test1"{
		t.Error("Expected test1 got ",tree.root[0].smaller.element.name)
	}
	testElement(6,6,t,tree,"seconda verifica 0","test1")
	testElement(8,8,t,tree,"seconda verifica 1","test2","test1")
	testElement(16,16,t,tree,"seconda verifica 2","test2")
}
func TestTreeSeparatedComponents1(t *testing.T) {
	tree := CreateTreeManager[mockElementTree](6)
	tree.AddElement(createMockElementTree(5,5,10,10,"test1"))
	tree.AddElement(createMockElementTree(17,17,10,10,"test2"))
	testElement(6,6,t,tree,"terza verifica 0","test1")
	testElement(8,8,t,tree,"terza verifica 1","test1",)
	testElement(18,18,t,tree,"terza verifica 2","test2")
}
func TestTreeSeparatedComponents2(t *testing.T) {
	tree := CreateTreeManager[mockElementTree](6)
	tree.AddElement(createMockElementTree(17,17,10,10,"test2"))
	tree.AddElement(createMockElementTree(0,0,10,10,"test1"))
	testElement(6,6,t,tree,"quarta verifica 0","test1")
	testElement(8,8,t,tree,"quarta verifica 1","test1",)
	testElement(18,18,t,tree,"quarta verifica 2","test2")
}
func TestTreeMixedCase1(t *testing.T) {
	tree := CreateTreeManager[mockElementTree](6)
	tree.AddElement(createMockElementTree(17,17,10,10,"test2"))
	tree.AddElement(createMockElementTree(20,20,10,10,"test3"))
	tree.AddElement(createMockElementTree(0,0,10,10,"test1"))
	testElement(2,2,t,tree,"quinta verifica 0","test1")
	testElement(8,8,t,tree,"quinta verifica 1","test1",)
	testElement(18,18,t,tree,"quinta verifica 2","test2")
	testElement(21,21,t,tree,"quinta verifica 3","test2","test3")
}
func TestTreeMixedCase2(t *testing.T) {
	tree := CreateTreeManager[mockElementTree](6)
	tree.AddElement(createMockElementTree(70,0,20,30,"button"))
	tree.AddElement(createMockElementTree(0,0,50,60,"textBox"))
	tree.AddElement(createMockElementTree(85,0,10,10,"button1"))
	testElement(1,1,t,tree,"sesta verifica 0","textBox")
}
func TestTreeRefresh(t *testing.T) {
	tree := CreateTreeManager[mockElementTree](6)
	tree.AddElement(createMockElementTree(70,0,20,30,"button"))
	tree.AddElement(createMockElementTree(0,0,50,60,"textBox"))
	tree.AddElement(createMockElementTree(85,0,10,10,"button1"))
	tree.Refresh()
	testElement(1,1,t,tree,"sesta verifica 0","textBox")
	testElement(1,1,t,tree,"settima verifica 0","textBox")
}
func TestTreeElementAbove(t *testing.T) {
	tree := CreateTreeManager[mockElementTree](6)
	tree.AddElement(createMockElementTree(40,40,40,40,"button1"))
	tree.AddElement(createMockElementTree(40,0,10,10,"textBox"))
	testElement(41,1,t,tree,"ottava verifica 0","textBox")
	testElement(41,41,t,tree,"ottava verifica 1","button1")
}
func TestTreeElementNextToAnotherElement(t *testing.T) {
	tree := CreateTreeManager[mockElementTree](6)
	tree.AddElement(createMockElementTree(10,5,10,5,"button1"))
	tree.AddElement(createMockElementTree(25,0,50,50,"textBox"))
	if tree.root[0].element.name!="button1"{
		t.Error("Expected button1 got ",tree.root[0].element.name)
	}
	if tree.root[0].smaller.element.name!="textBox"{
		t.Error("Expected textBox got ",tree.root[0].smaller.element.name)
	}
		testElement(26,11,t,tree,"nona verifica 1","textBox")
	for i := 1; i < 50; i=i+10 {
		testElement(26,i,t,tree,"nona verifica 1","textBox")
		print("---\n")
	}
}
