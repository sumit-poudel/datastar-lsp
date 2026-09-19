# data-on:resize

Runs an expression whenever the window or element is resized.

## Syntax

```html
data-on:resize="expression"
data-on:resize__modifier="expression"
```

## Description

The `data-on:resize` attribute listens for resize events. By default, this listens on the window. Use the `window` or `document` modifiers explicitly, or use `data-on-resize` attribute for element-specific resize observation.

## Examples

### Window resize

```html
<div data-on:resize__window__throttle.100ms="$windowWidth = $event.target.innerWidth"></div>
```

### Responsive layout

```html
<div data-on:resize__window="$isMobile = $event.target.innerWidth < 768"></div>
```

## Event Properties

- `$event.target` - Window or element
- `$event.target.innerWidth` / `$event.target.innerHeight` - Viewport dimensions

## Common Modifiers

- `window` - Attach to window (default for this event)
- `document` - Attach to document
- `throttle` - **Highly recommended** - Throttle to avoid performance issues
- `debounce` - Debounce the handler

## Note

For element resize observation, prefer the `data-on-resize` attribute which uses ResizeObserver:

```html
<div data-on-resize__debounce.100ms="$width = $event.target.clientWidth">...</div>
```

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-on)
- [MDN Resize Event](https://developer.mozilla.org/en-US/docs/Web/API/Window/resize_event)