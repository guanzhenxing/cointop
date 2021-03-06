package cointop

import (
	"fmt"

	"github.com/miguelmota/gocui"
)

// UpdateUI takes a callback which updates the view
func (ct *Cointop) UpdateUI(f func() error) {
	ct.debuglog(fmt.Sprintf("UpdateUI()"))

	if ct.g == nil {
		return
	}

	ct.g.Update(func(g *gocui.Gui) error {
		return f()
	})
}
