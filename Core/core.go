package Core

import (
	"errors"
	"sort"
)

type typeTreeNode bool

const (
	ByX typeTreeNode = true
	ByY typeTreeNode = false
)

type ElementTree interface {
	GetPos() (int, int)
	GetSize() (int, int)
}
type TreeNode[T ElementTree] struct {
	nodeType typeTreeNode
	xPos     int
	yPos     int
	width    int
	height   int
	element  T
	bigger   *TreeNode[T]
	smaller  *TreeNode[T]
}

func CreateNode[T ElementTree](element T, typeNode typeTreeNode) *TreeNode[T] {
	xPos, yPos := element.GetPos()
	xSize, ySize := element.GetSize()
	return &TreeNode[T]{
		nodeType: typeNode,
		bigger:   nil,
		smaller:  nil,
		element:  element,
		xPos:     xPos,
		yPos:     yPos,
		width:    xSize,
		height:   ySize,
	}
}

func (t *TreeNode[T]) isCollidingWithGroup(x, y int, width, height int) bool {
	if x > t.xPos && x < t.xPos+t.width && y > t.yPos && y < t.yPos+t.height {
		return true
	}
	if x+width > t.xPos && x < t.xPos+t.width && y > t.yPos && y < t.yPos+t.height {
		return true
	}
	if x > t.xPos && x < t.xPos+t.width && y+height > t.yPos && y+height < t.yPos+t.height {
		return true
	}
	if x+width > t.xPos && x < t.xPos+t.width && y+height > t.yPos && y+height < t.yPos+t.height {
		return true
	}
	return false
}
func (t *TreeNode[T]) isElementColliding(x, y int) bool {
	width, height := t.element.GetSize()
	xPos, yPos := t.element.GetPos()
	if x > xPos && x < xPos+width && y > yPos && y < yPos+height {
		return true
	}
	return false
}

func addElement[T ElementTree](destination **TreeNode[T], element T, _type typeTreeNode) *TreeNode[T] {
	if *destination == nil {
		*destination = CreateNode(element, _type)
		return *destination
	}
	(*destination).addNode(element)
	return nil
}

func (t *TreeNode[T]) addNode(element T) {
	xPos, yPos := element.GetPos()
	width, height := element.GetSize()
	xPosElNode, yPosElNode := t.element.GetPos()
	if (t.nodeType == ByX && yPos >= yPosElNode) || (t.nodeType == ByY && xPos >= xPosElNode) {
		addElement(&t.bigger, element, !t.nodeType)
	} else {
		addElement(&t.smaller, element, !t.nodeType)
	}
	if t.isCollidingWithGroup(xPos, yPos, width, height) {
		xSize, ySize := element.GetSize()
		if (xPos + xSize) > t.xPos+t.width {
			t.width = xPos + xSize - t.xPos
		} else {
			t.width = t.xPos + t.width - xPos
		}
		if (yPos + ySize) > t.yPos+t.height {
			t.height = yPos + ySize - t.yPos
		} else {
			t.height = t.yPos + t.height - yPos
		}

		if xPos < t.xPos {
			t.xPos = xPos
		}
		if yPos < t.yPos {
			t.yPos = yPos
		}
	}
}
func (d *TreeNode[T]) addNodes(elements []T) {
	for _, e := range elements {
		d.addNode(e)
	}
}
func (d *TreeNode[T]) search(x, y int) []T {
	var result []T
	d.execute(x, y, func(node *TreeNode[T]) {
		if isColliding := node.isElementColliding(x, y); isColliding {
			if isColliding {
				result = append(result, node.element)
			}
		}
	})
	return result
}
func (d *TreeNode[T]) execute(x, y int, do func(*TreeNode[T])) {
	do(d)
	xPosElNode, yPosElNode := d.element.GetPos()
	if d.smaller != nil && ((d.nodeType == ByY && x < xPosElNode) || (d.nodeType == ByX && y < yPosElNode) || d.smaller.isCollidingWithGroup(x, y, 0, 0)) {
		d.smaller.execute(x, y, do)
	}
	if d.bigger != nil && ((d.nodeType == ByY && x >= xPosElNode) || (d.nodeType == ByX && y >= yPosElNode) || d.bigger.isCollidingWithGroup(x, y, 0, 0)) {
		d.bigger.execute(x, y, do)
	}
}

// Iterate over all over the tree,if the function returns false the iteration stops
func (d *TreeNode[T]) executeForAll(do func(node *TreeNode[T]) bool) *TreeNode[T] { //TODO:auto adjust tree structure
	if !do(d) {
		return nil
	}
	if d.bigger != nil {
		d.bigger.executeForAll(do)
	}
	if d.smaller != nil {
		d.smaller.executeForAll(do)
	}
	return nil
}

func (d *TreeNode[T]) GetElement() T {
	return d.element
}

func (d *TreeNode[T]) GetElements() []T {
	result := make([]T, 0)
	d.executeForAll(func(node *TreeNode[T]) bool {
		result = append(result, node.element)
		return true
	})
	return result
}

func Sort[T ElementTree](elements []T, _type typeTreeNode) {
	sort.Slice(elements, func(i, j int) bool {
		x1, y1 := elements[i].GetPos()
		x2, y2 := elements[j].GetPos()
		if _type == ByX {
			return x1 < x2
		}
		return y1 < y2
	})
}

func (d *TreeNode[T]) GetPos() (int, int) {
	return d.xPos, d.yPos
}

func (d *TreeNode[T]) GetSize() (int, int) {
	return d.width, d.height
}

type collisionChache[T ElementTree] struct {
	collisionElement []T
	node             *TreeNode[T]
}

type TreeManager[T ElementTree] struct {
	root *TreeNode[T]
	chachedCollision [5]collisionChache[T]
	nextIndexToCache        int
}

func CreateTreeManager[T ElementTree]() *TreeManager[T] {
	
	return &TreeManager[T]{chachedCollision: [5]collisionChache[T]{}, nextIndexToCache: 0}
}

func (d *TreeManager[T]) AddElement(element T) {
	if d.root == nil {
		d.root = CreateNode(element, ByX)
		return
	}
	d.root.addNode(element)
}

// Iterate over all nodes, if return false, it will stop iteration
func (d *TreeManager[T]) Execute(cond func(node *TreeNode[T]) bool) {
	d.root.executeForAll(cond)
}

func (d *TreeManager[T]) Search(x, y int) ([]T, error) {
	var result []T
	if d.root == nil {
		return nil, errors.New("Tree is empty")
	}
	result = d.root.search(x, y)
	return result, nil
}

func (d *TreeManager[T]) GetCollidingElement(elementWhichCollides *TreeNode[T]) []T {
	var result []T
	for i := range d.chachedCollision {
		if d.chachedCollision[i].node == elementWhichCollides {
			return d.chachedCollision[i].collisionElement
		}
	}
	d.root.executeForAll(func(node *TreeNode[T]) bool {
		x,y := node.element.GetPos()
		xSize, ySize := node.element.GetSize()
		if elementWhichCollides.isCollidingWithGroup(x, y, xSize, ySize) {
			result = append(result, node.element)
		}
		return true
	})
			print("\a")
	d.nextIndexToCache = (d.nextIndexToCache + 1) % len(d.chachedCollision)
	d.chachedCollision[d.nextIndexToCache] = collisionChache[T]{node: elementWhichCollides, collisionElement: result}
	return result
}
func (d *TreeManager[T]) Refresh() {
	var newRoot *TreeNode[T]
	d.root.executeForAll(func(node *TreeNode[T]) bool {
		if newRoot == nil {
			newRoot = CreateNode(node.element, ByX)
		} else {
			newRoot.addNode(node.element)
		}
		return true
	})
	d.root = newRoot
}
