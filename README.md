# Go-Carbonated

A Wrapper around [Bubble Tea](https://github.com/charmbracelet/bubbletea) that converts it into a component stack renderer.

## Features
- Blur and Focus events when switching components.
  - Focus events include Sizing information for ensuring sizing on focus change.
- Ability to consume a message.
  - Enables nesting of components and ergonomically handling multiple components with the same message triggers.

## Example Components
- Button: Adds the configured components to the stack on Enter/Space
- Dynamic Text: Updates the text with each call
- Styled: Wraps output View() in a lipgloss style
- Text: Renders text
- Watcher: Allows action off of messages in the stack
- Wrapper: Renders multiple components at the same time, only sending keyboard inputs to one.
