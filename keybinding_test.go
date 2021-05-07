package keybinding

import (
	"testing"

	"github.com/awesome-gocui/gocui"
)

func TestParse(t *testing.T) {
	var table = []struct {
		input    string
		key      interface{}
		modifier gocui.Modifier
		errStr   string
	}{
		{"ctrl + A", gocui.KeyCtrlA, gocui.ModNone, ""},
		{"Ctrl + a", gocui.KeyCtrlA, gocui.ModNone, ""},
		{"Ctl + a", gocui.KeyCtrlA, gocui.ModNone, ""},
		{"ctl + A", gocui.KeyCtrlA, gocui.ModNone, ""},
		{"f2", gocui.KeyF2, gocui.ModNone, ""},
		{"ctrl +    [", gocui.KeyCtrlLsqBracket, gocui.ModNone, ""},
		{"    ctrl +    ]   ", gocui.KeyCtrlRsqBracket, gocui.ModNone, ""},
		{"ctrl + /", gocui.KeyCtrlSlash, gocui.ModNone, ""},
		{"ctrl + \\", gocui.KeyCtrlBackslash, gocui.ModNone, ""},
		// {"left", gocui.KeyArrowLeft, gocui.ModNone, ""},
		{"PageUp", gocui.KeyPgup, gocui.ModNone, ""},
		{"PgUp", gocui.KeyPgup, gocui.ModNone, ""},
		{"pageup", gocui.KeyPgup, gocui.ModNone, ""},
		{"pgup", gocui.KeyPgup, gocui.ModNone, ""},
		{"tab", gocui.KeyTab, gocui.ModNone, ""},
		{"escape", gocui.KeyEsc, gocui.ModNone, ""},
		{"enter", gocui.KeyEnter, gocui.ModNone, ""},
		{"space", gocui.KeySpace, gocui.ModNone, ""},
		{"ctrl + alt + z", gocui.KeyCtrlZ, gocui.ModAlt, ""},
		{"f22", 0, gocui.ModNone, "unsupported keybinding: KeyF22"},
		{"ctrl + alt + !", 0, gocui.ModAlt, "unsupported keybinding: KeyCtrl! (+1)"},
		{"q", rune('q'), gocui.ModNone, ""},
		{" q", rune('q'), gocui.ModNone, ""},
	}

	for idx, trial := range table {
		actualKey, actualErr := Parse(trial.input)

		if actualKey.Value != trial.key {
			t.Errorf("Expected key '%+v' but got '%+v' (trial %d) %+v", trial.key, actualKey, idx, actualErr)
		}

		if actualKey.Modifier != trial.modifier {
			t.Errorf("Expected modifier '%+v' but got '%+v' (trial %d)", trial.modifier, actualKey, idx)
		}

		if actualErr == nil && trial.errStr != "" {
			t.Errorf("Expected error message of '%s' but got no message (trial %d)", trial.errStr, idx)
		} else if actualErr != nil && actualErr.Error() != trial.errStr {
			t.Errorf("Expected error message '%s' but got '%s' (trial %d)", trial.errStr, actualErr.Error(), idx)
		}
	}
}

//
// func TestParseAll(t *testing.T) {
// 	var table = []struct {
// 		input    string
// 		key      Key
// 		errStr   string
// 	}{
// 		{"ctrl + A, ctrl + B", Key{Value: gocui.KeyCtrlA, Modifier:gocui.ModNone}, ""},
// 		// {"Ctrl + a", gocui.KeyCtrlA, gocui.ModNone, ""},
// 		// {"Ctl + a", gocui.KeyCtrlA, gocui.ModNone, ""},
// 		// {"ctl + A", gocui.KeyCtrlA, gocui.ModNone, ""},
// 	}
//
// 	for idx, trial := range table {
// 		actualKey, actualErr := Parse(trial.input)
//
// 		if actualKey.Value != trial.key {
// 			t.Errorf("Expected key '%+v' but got '%+v' (trial %d)", trial.key, actualKey, idx)
// 		}
//
// 		if actualKey.Modifier != trial.modifier {
// 			t.Errorf("Expected modifier '%+v' but got '%+v' (trial %d)", trial.modifier, actualKey, idx)
// 		}
//
// 		if actualErr == nil && trial.errStr != "" {
// 			t.Errorf("Expected error message of '%s' but got no message (trial %d)", trial.errStr, idx)
// 		} else if actualErr != nil && actualErr.Error() != trial.errStr {
// 			t.Errorf("Expected error message '%s' but got '%s' (trial %d)", trial.errStr, actualErr.Error(), idx)
// 		}
// 	}
// }
//
