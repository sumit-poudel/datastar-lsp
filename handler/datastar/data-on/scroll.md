# data-on:scroll

Runs an expression whenever an element is scrolled.

## Syntax

```html
data-on:scroll="expression"
data-on:scroll__modifier="expression"
```

## Description

The `data-on:scroll` attribute listens for scroll events on scrollable elements. Fires continuously during scrolling.

## Examples

### Basic scroll

```html
<div data-on:scroll="$scrollTop = $event.target.scrollTop" class="scrollable">...</div>
```

### Infinite scroll

```html
<div data-on:scroll__throttle.100ms="$nearBottom && @loadMore()" class="list">...</div>
```

### Parallax effect

```html
<div data-on:scroll="$parallaxY = $event.target.scrollTop * 0.5">...</div>
```

## Event Properties

- `$event.target` - The scrolling element
- `$event.target.scrollTop` - Vertical scroll position
- `$event.target.scrollLeft` - Horizontal scroll position
- `$event.target.scrollHeight` - Total scrollable height
- `$event.target.clientHeight` - Visible height

## Common Modifiers

- `throttle` - **Highly recommended** - Throttle to avoid performance issues
- `debounce` - Debounce the handler
- `passive` - Mark as passive (browser optimization)
- `capture` - Use capture phase

## Performance Note

Scroll events fire rapidly. Always use `throttle` or `debounce` modifiers:

```html
<div data-on:scroll__throttle.16ms="$updatePosition()">...</div>
```

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-on)
- [MDN Scroll Event](https://developer.mozilla.org/en-US/docs/Web/API/Element/scroll_event)