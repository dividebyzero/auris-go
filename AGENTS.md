# Building With Auris Go — Agent Guide

This file is guidance for coding agents working in repositories that use Auris Go, or modifying Auris itself.

## Mission

Preserve the Auris visual language while producing native, maintainable Fyne UI. Auris is a presentation system, not an application architecture.

Do not move domain logic, persistence, transport, routing, or long-lived application state into Auris components.

## Visual contract

Treat these as invariants unless the task explicitly changes the design system:

1. **Geometry** — signature asymmetric chamfer: top-left and bottom-right cut at 45 degrees; top-right and bottom-left remain square.
2. **Depth** — glow replaces conventional elevation. Keep it tight to the surface or glyph.
3. **Palette** — consume semantic `Scheme` roles. Do not scatter raw hex colors through widgets.
4. **Typography** — display/heading: Rajdhani; body/labels: Exo 2; machine/data/status: Share Tech Mono.
5. **Motion** — fast 120 ms, normal 200 ms, slow 350 ms. Use the shared motion helpers and honor reduced motion.
6. **Interaction** — keyboard, focus, disabled, and selection states are part of the component, not optional polish.
7. **Progression** — segmented progress has a subdued filled trail and a brighter leading cell.
8. **Text glow** — glow glyphs, not their rectangular bounds.

## Before creating a new widget

First check whether the UI can be composed from existing Auris primitives/components.

Prefer, in order:

- semantic text helpers: `DisplayText`, `BodyText`, `DataText`
- `CurrentScheme()`
- chamfer/slant primitives
- depth/glow primitives
- existing components
- a small new Auris primitive
- only then a new standalone widget

Do not wrap a stock Fyne widget if doing so visibly breaks the Auris language. Native Fyne behavior is useful; native Fyne appearance is not automatically Auris appearance.

## State model

Keep semantic state separate from rendered state.

For animated UI:

```text
application owns target state
        ↓
Auris Motion interpolates
        ↓
component renders current frame
```

This makes reduced motion deterministic and keeps animation timing out of domain logic.

Avoid unmanaged goroutines inside widgets. UI updates from background timing must return to Fyne's UI thread. This repository opts into Fyne's `migrated_fynedo` model.

## Theme behavior

Never hard-code `DarkScheme()` inside a reusable component.

Use:

```go
s := CurrentScheme()
```

A component must remain meaningful in both dark and light schemes. Use semantic roles such as `SurfacePanel`, `TextBright`, `PrimaryActive`, `Danger`, and `Success`.

## Layout

Avoid placeholder widths and arbitrary fixed dimensions when content can establish a useful minimum size.

A caller may request a concrete size for HUD composition, but reusable components should:

- provide sensible minimum/intrinsic dimensions;
- resize without clipping essential text or controls;
- derive separators/borders from actual available width;
- avoid assuming the showcase window size.

## Accessibility

Interactive components must account for:

- keyboard activation/navigation;
- visible focus;
- disabled input;
- readable semantic contrast;
- reduced motion.

Do not use animation as the only way to communicate state.

## Repository discipline

Keep implementation files small and focused. Split primitives, behavior, and tests rather than growing giant component files.

When modifying Auris:

1. create a branch;
2. make focused commits;
3. open a pull request;
4. wait for PR CI;
5. inspect failed job/step logs and fix on the branch;
6. merge only after CI is green.

Validation commands:

```bash
gofmt -w .
go vet ./...
go test -tags migrated_fynedo ./...
go build -tags migrated_fynedo ./cmd/showcase
```

Add or update tests for behavioral changes. For visual changes, update the showcase so the affected state is visible.

## Definition of done

A change is not done merely because it compiles. It should:

- look recognizably Auris;
- use semantic scheme roles;
- behave correctly with keyboard/focus/disabled states where applicable;
- honor reduced motion where applicable;
- resize reasonably;
- keep application logic outside the UI kit;
- have focused tests for behavior;
- pass repository CI.

When uncertain, use the Flutter Auris repository as the visual and behavioral reference, but implement the result idiomatically in Go/Fyne rather than translating Flutter line by line.
