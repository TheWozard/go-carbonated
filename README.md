# Go-Carbonated

A Wrapper around [Bubble Tea](https://github.com/charmbracelet/bubbletea) that converts it into a component stack renderer.

## Features
- Blur and Focus events when switching components.
  - Focus events include Sizing information for ensuring sizing on focus change.
- Ability to consume a message.
  - Enables nesting of components and ergonomically handling multiple components with the same message triggers.

## Components
- Text: Renders text that can be closed with "esc", "backspace", "enter", "space".
