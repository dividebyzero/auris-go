# Auris Go Quick Start

Auris Go is a native Go/Fyne implementation of the Auris interface system: warm amber, near-black surfaces, asymmetric chamfers, compact HUD typography, and restrained glow.

## Install

```bash
go get github.com/dividebyzero/auris-go
```

Auris currently targets Fyne v2.6.x.

## Start an application

```go
package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"

    "github.com/dividebyzero/auris-go/auris"
)

func main() {
    a := app.New()
    a.Settings().SetTheme(auris.NewTheme())

    w := a.NewWindow("Auris")
    w.Resize(fyne.NewSize(800, 600))

    status := auris.NewBadge("online", auris.BadgeSuccess)
    progress := auris.NewProgress(.68, 20, "Shield integrity", "68 / 100", auris.ProgressPrimary)
    panel := auris.NewPanel(
        "System",
        "SYS-01",
        container.NewVBox(status, progress),
        fyne.NewSize(620, 180),
        true,
    )

    w.SetContent(container.NewPadded(panel))
    w.ShowAndRun()
}
```

## Theme

Install `auris.NewTheme()` once at application startup. Auris components use semantic colors rather than raw palette values.

Explicit scheme control is available when needed:

```go
auris.UseDarkScheme()
auris.UseLightScheme()
```

Prefer semantic scheme fields such as `CurrentScheme().PrimaryActive` when building additional Auris-compatible UI.

## Core components

The kit includes Container, Panel, Badge, Notification, StatCard, DataRow, ProgressBar, Switch, Radio/RadioGroup, Select, StepIndicator, Terminal, HexOrnament, and ScanBracket.

The executable showcase in `cmd/showcase` is the visual component catalogue.

## Motion and reduced motion

Auris centralizes timing:

```go
auris.AnimateValue(auris.Motion{}, auris.DurationNormal, 0, 1, func(v float32) {
    // update rendered state
})
```

Use `Motion{Reduced: true}` when motion should be suppressed. Reduced motion resolves immediately to the final state.

The intended separation is:

**application state -> motion/interpolation -> component rendering**

Do not put business state into Auris widgets.

## Design rules

- Top-left and bottom-right corners are chamfered; the other two stay square.
- Use glow as tight depth/highlight, never as a giant blur.
- Rajdhani is display typography, Exo 2 is body copy, Share Tech Mono is data/status text.
- Prefer Auris semantic colors and primitives over one-off styling.
- Keep controls keyboard accessible and preserve disabled/focus states.
- Respect reduced motion.
- Auris is presentational. Routing, persistence, networking, and domain logic belong outside the kit.

## Validate

```bash
gofmt -w .
go vet ./...
go test -tags migrated_fynedo ./...
go build -tags migrated_fynedo ./cmd/showcase
```

See `AGENTS.md` when an AI coding agent will design or implement Auris UI.
