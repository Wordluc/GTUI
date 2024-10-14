package Core

import (
	"errors"
	"fmt"
	"reflect"
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
	element  []T
	bigger   *TreeNode[T]
	smaller  *TreeNode[T]
}

func CreateNode[T ElementTree](element T, typeNode typeTreeNode) *TreeNode[T] {
	xPos, yPos := element.GetPos()
	xSize, ySize := element.GetSize()
	return &TreeNode[T]{
		nodeType: typeNode,
		element:  []T{element},
		xPos:     xPos,
		yPos:     yPos,
		width:    xSize,
		height:   ySize,
	}
}

func (t *TreeNode[T]) isCollidingWithGroup(x, y int) bool {
	//fmt.Printf("x: %d, y: %d, pos: (%d,%d) ; size: (%d,%d)\n", x, y,t.xPos, t.yPos, t.width, t.height)
	return x > t.xPos && x < t.xPos+t.width && y > t.yPos && y < t.yPos+t.height
}

func (t *TreeNode[T]) returnCollidingElements(x, y int) []T {
	var result []T
	for _, element := range t.element {
		xSize, ySize := element.GetSize()
		xPos, yPos := element.GetPos()
		if x > xPos && x < xPos+xSize && y > yPos && y < yPos+ySize {
			result = append(result, element)
		}
	}
	return result
}

func addElement[T ElementTree](destination **TreeNode[T], element T, _type typeTreeNode) {
	if *destination == nil {
		*destination = CreateNode(element, _type)
	}
	print("\a")
	(*destination).addNode(element)
}

func (t *TreeNode[T]) addNode(element T) {
	xPos, yPos := element.GetPos()
	if !t.isCollidingWithGroup(element.GetPos()) {
		if t.nodeType == ByX {
			if yPos < t.yPos {
				addElement(&t.smaller, element, ByY)
			} else {
				addElement(&t.bigger, element, ByY)
			}
		} else //if t.nodeType == ByY
		if xPos < t.xPos {
			addElement(&t.smaller, element, ByX)
		} else {
			addElement(&t.bigger, element, ByX)
		}
		return
	}
	t.element = append(t.element, element)
	xSize, ySize := element.GetSize()
	if (xPos + xSize) > t.xPos+t.width {
		t.width = xPos + xSize - t.xPos
	}
	if (yPos + ySize) > t.yPos+t.width {
		t.height = yPos + ySize - t.yPos
	}
	if xPos < t.xPos {
		t.xPos = xPos
	}
	if yPos < t.yPos {
		t.yPos = yPos
	}
}
func (d *TreeNode[T]) addNodes(elements []T) {
	for _, e := range elements {
		d.addNode(e)
	}
}
func (d *TreeNode[T]) search(x, y int) []T {
	if d.isCollidingWithGroup(x, y) {
		if r := d.returnCollidingElements(x, y); len(r) > 0 {
			return r
		}
	}
	var result []T
	if d.nodeType == ByX {
		if x > d.xPos {
			if d.bigger != nil {
				if r := d.bigger.search(x, y); r != nil {
					result = append(result, r...)
				}
			}
		} else {
			if d.smaller != nil {
				if r := d.smaller.search(x, y); r != nil {
					result = append(result, r...)
				}
			}
		}
	} else {
		if y > d.yPos {
			if d.bigger != nil {
				if r := d.bigger.search(x, y); r != nil {
					result = append(result, r...)
				}
			}
		} else {
			if d.smaller != nil {
				if r := d.smaller.search(x, y); r != nil {
					result = append(result, r...)
				}
			}
		}
	}
	return result

}
func (d *TreeNode[T]) executeForAll(do func(node *TreeNode[T])) *TreeNode[T] { //TODO:auto adjust tree structure
	do(d)
	if d.bigger != nil {
		d.bigger.executeForAll(do)
	}
	if d.smaller != nil {
		d.smaller.executeForAll(do)
	}
	return nil
}

func (d *TreeNode[T]) GetElements() []T {
	return d.element
}

func (d *TreeNode[T]) GetPos() (int, int) {
	return d.xPos, d.yPos
}

func (d *TreeNode[T]) GetSize() (int, int) {
	return d.width, d.height
}

type TreeManager[T ElementTree] struct {
	root *TreeNode[T]
}

func CreateTreeManager[T ElementTree]() *TreeManager[T] {
	return &TreeManager[T]{}
}

func (d *TreeManager[T]) AddElement(element T) {
	fmt.Printf("adding element %v\n", reflect.TypeOf(element).String())
	if d.root == nil {
		d.root = CreateNode(element, ByX)
	}
	d.root.addNode(element)
}

func (d *TreeManager[T]) Execute(cond func(node *TreeNode[T])) {
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
func (d *TreeManager[T]) Refresh() {
	var newRoot *TreeNode[T]
	d.root.executeForAll(func(node *TreeNode[T]) {
		if newRoot == nil {
			newRoot = node
		}
		newRoot.addNodes(node.GetElements())
	})
	d.root = newRoot
}
