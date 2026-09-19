# data-on:dblclick

Runs an expression whenever a double-click event is triggered.

## Syntax

```html
data-on:dblclick="expression"
data-on:dblclick__modifier="expression"
```

## Description

The `data-on:dblclick` attribute listens for double-click events. Fires after two click events within a short time interval.

## Examples

### Basic double-click

```html
<div data-on:dblclick="$editing = true">Double-click to edit</div>
```

### Select word on double-click

```html
<p data-on:dblclick="window.getSelection().selectAllChildren($event.target)">...</p>
```

## Event Properties

- `$event.target` - The double-clicked element
- `$event.detail` - Number of clicks (2 for dblclick)
- `$event.clientX` / `$event.clientY` - Mouse position

## Common Modifiers

- `prevent` - Calls `preventDefault()`
- `stop` - Calls `stopPropagation()`
- `once` - Only trigger once

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-on)
- [MDN DblClick Event](https://developer.mozilla.org/en-US/docs/Web/API/Element/dblclick_event)