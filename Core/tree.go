package Core

import (
	"errors"
	"math"
	"sync"
)

type typeTreeNode bool

const (
	ByX typeTreeNode = true
	ByY typeTreeNode = false
)

type elementToCompare interface {
	GetPos() (int, int)
	GetSize() (int, int)
}

type ElementTree interface {
	GetPos() (int, int)
	GetSize() (int, int)
	GetLayer() Layer
}

func isCollidingWithElement[T elementToCompare](x, y int, width, height int, t T) bool {
	xEl, yEl := t.GetPos()
	widthEl, heightEl := t.GetSize()
	if x > xEl && x < xEl+widthEl && y >= yEl && y <yEl+heightEl {
		return true
	}
	if x+width > xEl && x+width < xEl+widthEl && y >= yEl && y < yEl+heightEl {
		return true
	}
	if x > xEl && x < xEl+widthEl && y+height >= yEl && y+height < yEl+heightEl {
		return true
	}
	if x+width > xEl && x+width < xEl+widthEl && y+height >= yEl && y+height < yEl+heightEl {
		return true
	}
	return false
}

// A	B
//
// C	D
type TreeNode[T ElementTree] struct {
	x      int
	y      int
	width  int
	height int

	deepth   int
	elements []T
	A        *TreeNode[T]
	B        *TreeNode[T]
	C        *TreeNode[T]
	D        *TreeNode[T]
}

func CreateNode[T ElementTree](x, y int, width, height int, deepth int) *TreeNode[T] {
	if deepth == 0 {
		return &TreeNode[T]{
			x:      x,
			y:      y,
			width:  width,
			height: height,
			deepth: deepth,
		}
	}
	return &TreeNode[T]{
		x:      x,
		y:      y,
		width:  width,
		height: height,
		deepth: deepth,
		A:      CreateNode[T](x, y, width/2, height/2, deepth-1),
		B:      CreateNode[T](x+width/2, y, width/2, height/2, deepth-1),
		C:      CreateNode[T](x, y+height/2, width/2, height/2, deepth-1),
		D:      CreateNode[T](x+width/2, y+height/2, width/2, height/2, deepth-1),
	}
}

func (t *TreeNode[T]) GetPos() (int, int) {
	return t.x, t.y
}

func (t *TreeNode[T]) GetSize() (int, int) {
	return t.width, t.height
}

func (t *TreeNode[T]) isCollidingWithNode(x, y, width, height int) bool {
	return isCollidingWithElement(x, y, width, height, t)
}

func (t *TreeNode[T]) addNode(element T, x int, y int) {
	if t.deepth == 0 {
		t.elements = append(t.elements, element)
		return
	}
	if x >= t.x+t.width/2 {
		if y >= t.y+t.height/2 {
			t.D.addNode(element, x, y)
		} else {
			t.B.addNode(element, x, y)
		}
	} else {
		if y >= t.y+t.height/2 {
			t.C.addNode(element, x, y)
		} else {
			t.A.addNode(element, x, y)
		}
	}
}

func (d *TreeNode[T]) search(x, y int) []T {
	var result []T
	d.execute(x, y, func(node *TreeNode[T]) {
		for i := range node.elements {
			if isCollidingWithElement(x, y, 0, 0, node.elements[i]) {
				result = append(result, node.elements[i])
			}
		}
	})

	return result
}

func (d *TreeNode[T]) execute(x, y int, do func(*TreeNode[T])) {
	call := func(d *TreeNode[T], x, y int, do func(*TreeNode[T])) {
		if d == nil {
			return
		}
		d.execute(x, y, do)
	}
	if d.deepth == 0 {
		do(d)
	}
	if x >= d.x+d.width/2 {
		if y > d.y+d.height/2 {
			call(d.D, x, y, do)
		} else {
			call(d.B, x, y, do)
		}
	} else {
		if y >= d.y+d.height/2 {
			call(d.C, x, y, do)
		} else {
			call(d.A, x, y, do)
		}
	}
}

// Iterate over all over the tree,if the function returns false the iteration will stop
func (d *TreeNode[T]) executeForAll(do func(node *TreeNode[T]) bool) { //TODO:auto adjust tree structure
	call := func(d *TreeNode[T]) {
		if d == nil {
			return
		}
		if d.deepth == 0 { //TODO: da verificare
			do(d) 
		} else {
			d.executeForAll(do)
		}
	}
	if d == nil {
		return
	}
	call(d.A)
	call(d.B)
	call(d.C)
	call(d.D)
}

func (d *TreeNode[T]) GetElements() []T {
	return d.elements
}

type collisionChache[T ElementTree] struct {
	collisionElement []T
	element          *T
}

type TreeManager[T ElementTree] struct {
	root             []*TreeNode[T]
	chachedCollision [5]collisionChache[T]
	nextIndexToCache int
	nLayer           Layer
	width            int
	height           int
	deepth           int
	elements         []T
}

func CreateTreeManager[T ElementTree](nLayer Layer, width, height int, deepth int) *TreeManager[T] {

	return &TreeManager[T]{
		chachedCollision: [5]collisionChache[T]{},
		nextIndexToCache: 0,
		root:             make([]*TreeNode[T], nLayer),
		nLayer:           nLayer,
		width:            width,
		height:           height,
		deepth:           deepth,
	}
}

func (d *TreeManager[T]) AddElement(element T) {
	layer := element.GetLayer()
	if d.root[layer] == nil {
		d.root[layer] = CreateNode[T](0, 0, d.width, d.height, d.deepth)
	}
	x, y := element.GetPos()
	width, height := element.GetSize()
	paceX := int(d.width/int(math.Pow(float64(2), float64(d.deepth))))
	paceY := int(d.height/int(math.Pow(float64(2), float64(d.deepth))))
	for yToInsert := 0; yToInsert*paceY < y+height; yToInsert++ {
		for xToInsert := 0; xToInsert*paceX < x+width; xToInsert++ {
			d.root[layer].addNode(element, x+xToInsert*(paceX), y+yToInsert*(paceY))
		}
	}
	d.elements = append(d.elements, element)
}

// Iterate over all nodes, if return false, it will stop iteration
func (d *TreeManager[T]) Execute(layer Layer, cond func(node *TreeNode[T]) bool) {
	d.root[layer].executeForAll(cond)
}

func (d *TreeManager[T]) Search(layer Layer, x, y int) ([]T, error) {
	var result []T
	if d.root[layer] == nil {
		return nil, errors.New("Tree is empty")
	}
	result = d.root[layer].search(x, y)
	return result, nil
}

func (d *TreeManager[T]) SearchAll(x, y int) ([]T, error) {
	var results [][]T = make([][]T, d.nLayer)
	var group sync.WaitGroup = sync.WaitGroup{}
	for layer := 0; layer < int(d.nLayer); layer++ {
		if d.root[layer] != nil {
			group.Add(1)
			go func() {
				d.root[layer].execute(x, y, func(node *TreeNode[T]) {
						for i := range node.elements {
							if isCollidingWithElement( x, y,0,0,node.elements[i]) {
								results[layer] = append(results[layer], node.elements[i])
							}
							
						}
				})
				group.Done()
			}()
		}
	}
	group.Wait()
	var result []T
	for _, res := range results {
		result = append(result, res...)
	}
	return result, nil
}

func (d *TreeManager[T]) GetCollidingElement(layer Layer, elementWhichCollides T) []T {//da sistemare
	var result []T

	defer func() {
		d.nextIndexToCache = (d.nextIndexToCache + 1) % len(d.chachedCollision)
		d.chachedCollision[d.nextIndexToCache] = collisionChache[T]{element: &elementWhichCollides, collisionElement: result}
	}()

	for i := range d.chachedCollision {
		if d.chachedCollision[i].element == &elementWhichCollides {
			result = d.chachedCollision[i].collisionElement
			return result
		}
	}
	x, y := elementWhichCollides.GetPos()
	width, height := elementWhichCollides.GetSize()
	for _, el := range d.elements {
		if el.GetLayer()==layer && isCollidingWithElement(x, y, width, height, el) {
			result = append(result, el)
		}
	}

	return result
}

func (d *TreeManager[T]) Refresh() {
	newTree := CreateTreeManager[T](d.nLayer, d.width, d.height, d.deepth)
	for i := range d.elements {
		newTree.AddElement(d.elements[i])
	}
	*d = *newTree
}

func (d *TreeManager[T]) _refresh(layer Layer, tree *TreeManager[T]) {
}
