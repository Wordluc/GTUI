package Component
 
type IComponent interface {
	OnClick()
	OnRelease()
	OnHover()
	OnLeave()
	insertingToMap(*Search)error
	isOn(x,y int)bool
}
