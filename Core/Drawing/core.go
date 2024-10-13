package Drawing

import (
	"github.com/Wordluc/GTUI/Core"
)
type typeTreeNode bool
const (
	ByX typeTreeNode = true
	ByY typeTreeNode = false
)
type TreeNode struct {
	nodeType typeTreeNode
	xPos int
	yPos int
	width int
	height int
	element []Core.IEntity
	nodeOrderedByX *TreeNode
	nodeOrderedByY *TreeNode
}
func CreateNode(element Core.IEntity,typeNode typeTreeNode) *TreeNode {
	xPos,yPos:=element.GetPos()
	xSize,ySize:=element.GetSize()
	return &TreeNode{
		nodeType: typeNode,
		element:  []Core.IEntity{element},
		xPos:     xPos,
		yPos:     yPos,
		width:    xSize,
		height:   ySize,
	}
}
func (t *TreeNode) isCollide(x,y int) bool {
	return x >= t.xPos && x < t.xPos+t.width && y >= t.yPos && y < t.yPos+t.height
}

func (t *TreeNode) addNode(element Core.IEntity) {
	if !t.isCollide(element.GetPos()){
		if t.nodeType==ByX{
			if t.nodeOrderedByX==nil{
				t.nodeOrderedByX=CreateNode(element,ByX)
			}else{
				t.nodeOrderedByX.addNode(element)
			}
		}else
		if t.nodeType==ByY{
			if t.nodeOrderedByY==nil{
				t.nodeOrderedByY=CreateNode(element,ByY)
			}else{
				t.nodeOrderedByY.addNode(element)
			}
		}
	}
	t.element = append(t.element, element)
	xPos,yPos:=element.GetPos()
	xSize,ySize:=element.GetSize()
	if (xPos+xSize)>t.xPos+t.width{
		t.width=xPos+xSize-t.xPos
	}
	if (yPos+ySize)>t.yPos+t.width{
		t.height=yPos+ySize-t.yPos
	}
	if xPos<t.xPos{
		t.xPos=xPos
	}
	if yPos<t.yPos{
		t.yPos=yPos
	}
}
func (d *TreeNode) addNodes(element []Core.IEntity) {
	for _,e:=range element{
		d.addNode(e)
	}
}
func (d *TreeNode) execute(do func(node *TreeNode)) *TreeNode {//TODO:auto adjust tree structure
	do(d)
	if d.nodeOrderedByX!=nil{
		d.nodeOrderedByX.execute(do)
	}
	if d.nodeOrderedByY!=nil{
		d.nodeOrderedByY.execute(do)
	}
	return nil
}
func (d *TreeNode) GetElements() ([]Core.IEntity){
	return d.element
}
func (d *TreeNode) GetPos() (int,int){
	return d.xPos,d.yPos
}
func (d *TreeNode) GetSize() (int,int){
	return d.width,d.height
}
type DrawingManager struct {
	root *TreeNode
}

func CreateDrawingManager() *DrawingManager {
	return &DrawingManager{

	}
}
func (d *DrawingManager) AddElement(element Core.IEntity) {
	if d.root==nil{
		d.root=CreateNode(element,ByX)
	}
	d.root.addNode(element)
}
func (d *DrawingManager) Execute(cond func(node *TreeNode)) {
	d.root.execute(cond)
}

func (d *DrawingManager) Refresh() {
	var newRoot *TreeNode
	d.root.execute(func(node *TreeNode){
		if newRoot==nil{
			newRoot=node
		}
		newRoot.addNodes(node.GetElements())
	})
	d.root=newRoot
}
