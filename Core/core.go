package Core

type typeTreeNode bool
const (
	ByX typeTreeNode = true
	ByY typeTreeNode = false
)
type ElementTree interface {
	GetPos() (int,int)
	GetSize() (int,int)
}
type TreeNode[T ElementTree] struct {
	nodeType typeTreeNode
	xPos int
	yPos int
	width int
	height int
	element []T
	nodeOrderedByX *TreeNode[T]
	nodeOrderedByY *TreeNode[T]
}
func CreateNode[T ElementTree](element T,typeNode typeTreeNode) *TreeNode[T] {
	xPos,yPos:=element.GetPos()
	xSize,ySize:=element.GetSize()
	return &TreeNode[T]{
		nodeType: typeNode,
		element:  []T{element},
		xPos:     xPos,
		yPos:     yPos,
		width:    xSize,
		height:   ySize,
	}
}

func (t *TreeNode[T]) isCollidingWithGroup(x,y int) bool {
	return x > t.xPos && x < t.xPos+t.width && y > t.yPos && y < t.yPos+t.height
}

func (t *TreeNode[T]) returnCollidingElements(x,y int)[]T{
	var result []T
	for _,element:=range t.element{
		xSize,ySize:=element.GetSize()
		xPos,yPos:=element.GetPos()
		if x > xPos && x < xPos+xSize && y > yPos && y < yPos+ySize{
			result=append(result,element)
		}
	}
	return result
}

func (t *TreeNode[T]) addNode(element T) {
	if !t.isCollidingWithGroup(element.GetPos()){
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
		return
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

func (d *TreeNode[T]) addNodes(element []T) {
	for _,e:=range element{
		d.addNode(e)
	}
}

func (d *TreeNode[T]) execute(do func(node *TreeNode[T])) *TreeNode[T] {//TODO:auto adjust tree structure
	do(d)
	if d.nodeOrderedByX!=nil{
		d.nodeOrderedByX.execute(do)
	}
	if d.nodeOrderedByY!=nil{
		d.nodeOrderedByY.execute(do)
	}
	return nil
}

func (d *TreeNode[T]) GetElements() ([]T){
	return d.element
}

func (d *TreeNode[T]) GetPos() (int,int){
	return d.xPos,d.yPos
}

func (d *TreeNode[T]) GetSize() (int,int){
	return d.width,d.height
}

type TreeManager[T ElementTree] struct {
	root *TreeNode[T]
}

func CreateTreeManager[T ElementTree]() *TreeManager[T] {
	return &TreeManager[T]{
	}
}

func (d *TreeManager[T]) AddElement(element T) {
	if d.root==nil{
		d.root=CreateNode(element,ByX)
	}
	d.root.addNode(element)
}

func (d *TreeManager[T]) Execute(cond func(node *TreeNode[T])) {
	d.root.execute(cond)
}

func (d *TreeManager[T]) Search(x,y int) ([]T,error) {
	var result []T
	d.root.execute(func(node *TreeNode[T]){
		if node.isCollidingWithGroup(x,y){
			if r:=node.returnCollidingElements(x,y);len(r)>0{
				result=append(result,r...)
			}
		}
	})
	return result,nil
}
func (d *TreeManager[T]) Refresh() {
	var newRoot *TreeNode[T]
	d.root.execute(func(node *TreeNode[T]){
		if newRoot==nil{
			newRoot=node
		}
		newRoot.addNodes(node.GetElements())
	})
	d.root=newRoot
}
