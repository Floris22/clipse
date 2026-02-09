package shell

const (
	listenCmd      = "-listen"
	listenShellCmd = "--listen-shell"
	pgrepCmd       = "pgrep 'clipse'"
	psCmd          = "ps -o command"
	wlCopyHandler  = "wl-copy"
	wlPasteHandler = "wl-paste"
	wlPasteWatcher = "--watch"
	wlCopyImgCmd   = "wl-copy -t image/png < %s"
	wlPasteImgCmd  = "wl-paste -t image/png > %s"
	wlStoreCmd     = "--wl-store"
	wlTypeSpec     = "--type"
)
