package interfaces

type ILevel interface {
	SetObjectOnBoard(uint, uint, rune)
	AddActorToRenderList(Actor)
	RenderOnConsole()
	GetLevelSize() (uint, uint)
}
