package display

var DisplayServer = GetDisplayServer()

type DS interface {
	Runtime() string
	ReadClipboard() string
	CopyText(string)
	CopyImage(string)
	Paste()
	RunListener()
	RunDetachedListener()
	SendPasteKey(string)
}

func GetDisplayServer() DS {
	return &WaylandDS{
		runtime: "wayland",
	}
}
