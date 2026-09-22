package auris

import (
	"testing"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func TestScanPulseAlphaBounds(t *testing.T) {
	child:=canvas.NewRectangle(Void)
	for _,p:=range []float32{-1,0,1,2} {
		obj:=NewScanBracketState(child,fyne.NewSize(100,50),14,2,p)
		if obj==nil { t.Fatal("scan bracket state must render") }
	}
}

func TestTerminalCursorStateRenders(t *testing.T) {
	lines:=[]TerminalLine{{Text:"ready",Type:TerminalOK}}
	if NewTerminalState("log","LIVE",lines,fyne.NewSize(300,120),true)==nil { t.Fatal("visible cursor terminal must render") }
	if NewTerminalState("log","LIVE",lines,fyne.NewSize(300,120),false)==nil { t.Fatal("hidden cursor terminal must render") }
}
