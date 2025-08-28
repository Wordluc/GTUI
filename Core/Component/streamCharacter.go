package Component

type StreamCharacter struct {
	Get       func() chan string
	IChannel  int
	Delete    func()
	IsDeleted bool
}
