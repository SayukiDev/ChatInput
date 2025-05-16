package entry

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type ExtEntry struct {
	OnTypedKey    func(*fyne.KeyEvent) bool
	OnFocusGained func() bool
	OnFocusLost   func() bool
	*widget.Entry
}

func NewExtEntry() *ExtEntry {
	e := &ExtEntry{
		Entry: &widget.Entry{
			MultiLine: true, Wrapping: fyne.TextWrap(fyne.TextTruncateClip),
		},
	}
	e.ExtendBaseWidget(e)
	return e
}

func (e *ExtEntry) TypedKey(key *fyne.KeyEvent) {
	if e.OnTypedKey != nil {
		if e.OnTypedKey(key) {
			return
		}
	}
	e.Entry.TypedKey(key)
}

func (e *ExtEntry) FocusGained() {
	if e.OnFocusGained != nil {
		if e.OnFocusGained() {
			return
		}
	}
	e.Entry.FocusGained()
}

func (e *ExtEntry) FocusLost() {
	if e.OnFocusLost != nil {
		if e.OnFocusLost() {
			return
		}
	}
	e.Entry.FocusLost()
}
