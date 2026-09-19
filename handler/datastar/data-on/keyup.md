# data-on:keyup

Runs an expression whenever a key is released.

## Syntax

```html
data-on:keyup="expression"
data-on:keyup__modifier="expression"
```

## Description

The `data-on:keyup` attribute listens for keyup events, which fire when a key is released. Useful for detecting when a key combination is complete.

## Examples

### Basic keyup

```html
<input data-on:keyup="console.log('Released:', $event.key)" />
```

### Enter to submit

```html
<input data-on:keyup="$event.key === 'Enter' && $submit()" />
```

### Escape to cancel

```html
<input data-on:keyup="$event.key === 'Escape' && ($editing = false)" />
```

## Event Properties

- `$event.key` - The key value
- `$event.code` - The physical key code
- `$event.ctrlKey`, `$event.shiftKey`, `$event.altKey`, `$event.metaKey` - Modifier keys

## Common Modifiers

- `prevent` - Calls `preventDefault()`
- `stop` - Calls `stopPropagation()`
- `once` - Only trigger once
- `debounce` - Debounce the handler
- `throttle` - Throttle the handler

## Difference from `keydown`

| Event | Fires | Use Case |
|-------|-------|----------|
| `keydown` | Key pressed (repeats on hold) | Games, continuous action |
| `keyup` | Key released (once) | Shortcuts, form submission |

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-on)
- [MDN KeyUp Event](https://developer.mozilla.org/en-US/docs/Web/API/Element/keyup_event)