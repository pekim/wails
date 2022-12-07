package application

import "github.com/wailsapp/wails/exp/pkg/options"

type windowImpl interface {
	setTitle(title string)
	setSize(width, height int)
	setAlwaysOnTop(alwaysOnTop bool)
	Run() error
}

type Window struct {
	options *options.Window
	impl    windowImpl
}

func NewWindow(options *options.Window) *Window {
	return &Window{
		options: options,
	}
}

func (w *Window) SetTitle(title string) {
	if w.impl == nil {
		w.options.Title = title
		return
	}
	w.impl.setTitle(title)
}

func (w *Window) SetSize(width, height int) {
	if w.impl == nil {
		w.options.Width = width
		w.options.Height = height
		return
	}
	w.impl.setSize(width, height)
}

func (w *Window) Run() error {
	w.impl = newWindowImpl(w.options)
	return w.impl.Run()
}

func (w *Window) SetAlwaysOnTop(b bool) {
	if w.impl == nil {
		w.options.AlwaysOnTop = b
		return
	}
	w.impl.setAlwaysOnTop(b)
}