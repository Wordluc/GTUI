package Core

import (
	"errors"
	"sync"
)

type typeTreeNode bool

const (
	ByX typeTreeNode = true
	ByY typeTreeNode = false
)

type ElementTree interface {
	GetPos() (int, int)
	GetSize() (int, int)
	GetLayer() Layer
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

func (t *TreeNode[T]) isCollidingWithNode(x, y int, width, height int) bool {
	if x > t.xPos && x < t.xPos+t.width && y > t.yPos && y < t.yPos+t.height {
		return true
	}
	if x+width > t.xPos && x+width < t.xPos+t.width && y > t.yPos && y < t.yPos+t.height {
		return true
	}
	if x > t.xPos && x < t.xPos+t.width && y+height > t.yPos && y+height < t.yPos+t.height {
		return true
	}
	if x+width > t.xPos && x+width < t.xPos+t.width && y+height > t.yPos && y+height < t.yPos+t.height {
		return true
	}
	return false
}

func addElement[T ElementTree](destination **TreeNode[T], element T, _type typeTreeNode) {
	if *destination == nil {
		*destination = CreateNode(element, _type)
		return
	}
	(*destination).addNode(element)
}

func (t *TreeNode[T]) addNode(element T) {
	xPos, yPos := element.GetPos()
	xPosElNode, yPosElNode := t.element.GetPos()
	if (t.nodeType == ByX && yPos >= yPosElNode) || (t.nodeType == ByY && xPos >= xPosElNode) {
		addElement(&t.bigger, element, !t.nodeType)
	} else {
		addElement(&t.smaller, element, !t.nodeType)
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
		if isColliding := node.isCollidingWithNode(x, y, 0, 0); isColliding {
			result = append(result, node.element)
		}
	})
	return result
}

func (d *TreeNode[T]) execute(x, y int, do func(*TreeNode[T])) {
	xPosElNode, yPosElNode := d.element.GetPos()
	smallerDone := false
	biggerDone := false

	if d.smaller != nil && d.smaller.isCollidingWithNode(x, y, 0, 0) {
		d.smaller.execute(x, y, do)
		smallerDone = true
	}else
	if d.bigger != nil && d.bigger.isCollidingWithNode(x, y, 0, 0) {
		d.bigger.execute(x, y, do)
		biggerDone = true
	}

	if d.smaller != nil && !smallerDone && ((d.nodeType == ByY && x < xPosElNode) || (d.nodeType == ByX && y < yPosElNode)) {
		d.smaller.execute(x, y, do)
	} else {
		if d.bigger != nil && !biggerDone {
			d.bigger.execute(x, y, do)
		}
	}
	do(d)
}

// Iterate over all over the tree,if the function returns false the iteration will stop
func (d *TreeNode[T]) executeForAll(do func(node *TreeNode[T]) bool) *TreeNode[T] { //TODO:auto adjust tree structure
	if d == nil {
		return nil
	}
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

type collisionChache[T ElementTree] struct {
	collisionElement []T
	node             *TreeNode[T]
}

type TreeManager[T ElementTree] struct {
	root             []*TreeNode[T]
	chachedCollision [5]collisionChache[T]
	nextIndexToCache int
}
func (d *TreeManager[T]) GetLayerN() int {
	return len(d.root)+1
	
}
func CreateTreeManager[T ElementTree]() *TreeManager[T] {

	return &TreeManager[T]{
		chachedCollision: [5]collisionChache[T]{},
		nextIndexToCache: 0,
		root:             make([]*TreeNode[T], 1),
	}
}

func (d *TreeManager[T]) AddElement(element T) {
	layer := element.GetLayer()
	for int(layer) >= len(d.root) {
		d.root = append(d.root, nil)
	}
	if d.root[layer] == nil {
		d.root[layer] = CreateNode(element, ByX)
		return
	}
	d.root[layer].addNode(element)
}

// Iterate over all nodes, if return false, it will stop iteration
func (d *TreeManager[T]) Execute(layer Layer, cond func(node *TreeNode[T]) bool) {
	if int(layer) >= len(d.root) {
		return
	}
	d.root[layer].executeForAll(cond)
}

func (d *TreeManager[T]) Search(layer Layer, x, y int) ([]T, error) {
	var result []T
	if int(layer) >= len(d.root) {
		return result,nil
	}
	if d.root[layer] == nil {
		return nil, errors.New("Tree is empty")
	}
	result = d.root[layer].search(x, y)
	return result, nil
}

func (d *TreeManager[T]) SearchAll(x, y int) ([]T, error) {
	var results [][]T = make([][]T, len(d.root))
	var group sync.WaitGroup = sync.WaitGroup{}
	for layer := 0; layer < len(d.root); layer++ {
		if d.root[layer] == nil {
			continue
		}
		group.Add(1)
		go func() {
			results[layer] = d.root[layer].search(x, y)
			group.Done()
		}()
	}
	group.Wait()
	var result []T
	for _, res := range results {
		result = append(result, res...)
	}
	return result, nil
}

func (d *TreeManager[T]) GetCollidingElement(layer Layer, elementWhichCollides *TreeNode[T]) []T {
	var result []T
	defer func() {
		d.nextIndexToCache = (d.nextIndexToCache + 1) % len(d.chachedCollision)
		d.chachedCollision[d.nextIndexToCache] = collisionChache[T]{node: elementWhichCollides, collisionElement: result}
	}()

	for i := range d.chachedCollision {
		if d.chachedCollision[i].node == elementWhichCollides {
			result = d.chachedCollision[i].collisionElement
			return result
		}
	}

	d.root[layer].executeForAll(func(node *TreeNode[T]) bool {
		x, y := node.element.GetPos()
		xSize, ySize := node.element.GetSize()
		if elementWhichCollides.isCollidingWithNode(x, y, xSize, ySize) {
			result = append(result, node.element)
		}
		return true
	})
	return result
}

func (d *TreeManager[T]) Refresh() {
	newTree := CreateTreeManager[T]()
	for layer := 0; layer < len(d.root); layer++ {
		d._refresh(Layer(layer), newTree)
	}
	(*d) = (*newTree) //deep copy
}

func (d *TreeManager[T]) _refresh(layer Layer, tree *TreeManager[T]) {
	d.root[layer].executeForAll(func(node *TreeNode[T]) bool {
		tree.AddElement(node.element)
		return true
	})
}
