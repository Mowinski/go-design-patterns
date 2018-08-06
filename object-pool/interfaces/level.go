package interfaces

type ILevel interface {
	SetObjectOnBoard(x, y uint, object rune)
	AddActorToRenderList(actor Actor)
	RenderOnConsole()
}
