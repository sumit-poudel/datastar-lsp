# data-on:keydown

Runs an expression whenever a key is pressed.

## Syntax

```html
data-on:keydown="expression"
data-on:keydown__modifier="expression"
```

## Description

The `data-on:keydown` attribute listens for keydown events on an element. The `$event` variable contains the KeyboardEvent with properties like `key`, `code`, `ctrlKey`, `shiftKey`, etc.

## Examples

### Basic keydown

```html
<input data-on:keydown="console.log('Key:', $event.key)" />
```

### Enter key

```html
<input data-on:keydown__prevent="$event.key === 'Enter' && $submit()" />
```

### Escape key

```html
<div data-on:keydown="$event.key === 'Escape' && ($modal = false)">...</div>
```

### With modifiers

```html
<input data-on:keydown__capture="$handleKey($event)" />
```

## Event Properties

- `$event.key` - The key value (e.g., "Enter", "a", "ArrowLeft")
- `$event.code` - The physical key code (e.g., "KeyA", "Enter")
- `$event.ctrlKey` - Whether Ctrl was pressed
- `$event.shiftKey` - Whether Shift was pressed
- `$event.altKey` - Whether Alt was pressed
- `$event.metaKey` - Whether Meta/Cmd was pressed

## Common Modifiers

- `prevent` - Calls `preventDefault()`
- `stop` - Calls `stopPropagation()`
- `capture` - Use capture phase
- `once` - Only trigger once
- `debounce` - Debounce the handler
- `throttle` - Throttle the handler

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-on)
- [MDN KeyboardEvent](https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent)