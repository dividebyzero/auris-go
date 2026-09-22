# auris-go

Native Go/Fyne port of the Auris UI design system.

The Flutter implementation at `point-source/auris` is the visual and behavioral reference. This port preserves the Auris design language while using Fyne-native mechanisms.

## Status

**Auris Go v1 implementation complete.** The public kit includes semantic dark/light schemes, embedded Auris typography, chamfer/depth primitives, motion with reduced-motion support, and the core Auris widget set.

## Components

Container, Panel, Badge, Notification, StatCard, DataRow, ProgressBar, Switch, Radio/RadioGroup, Select, StepIndicator, Terminal, HexOrnament, and ScanBracket.

## Theme

Install `auris.NewTheme()` as the Fyne application theme. Components resolve their presentation from the active Auris semantic scheme. `UseDarkScheme()` and `UseLightScheme()` are also available for explicit control.

## Motion

`Animate` and `AnimateValue` provide the shared Auris timing/easing layer. Reduced motion resolves immediately to the final state. The project opts into Fyne's v2.6 UI-thread migration.

## Reference and license

This port derives from PointSource Auris and retains its BSD 3-Clause license. The bundled Rajdhani, Exo 2, and Share Tech Mono font assets are the same assets shipped by the reference implementation.

## Validation

CI checks formatting, `go vet`, tests, and a build of `cmd/showcase`. The showcase is the visual acceptance surface for the port.
