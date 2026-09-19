# data-on:click

Attaches an event listener to an element, executing an expression whenever the event is triggered.

## Syntax

```html
data-on:click="expression"
data-on:click__modifier="expression"
```

## Description

The `data-on:click` attribute listens for click events on an element. This is one of the most commonly used events for handling user interactions.

## Examples

### Basic click

```html
<button data-on:click="$count++">Increment</button>
```

### Toggle boolean

```html
<button data-on:click="$visible = !$visible">Toggle</button>
```

### Function call

```html
<button data-on:click="@post('/api/action', { id: $item.id })">Action</button>
```

### With prevent default

```html
<a href="/link" data-on:click__prevent="$navigate('/page')">Link</a>
```

## Event Properties

- `$event.target` - The clicked element
- `$event.clientX` / `$event.clientY` - Mouse position relative to viewport
- `$event.offsetX` / `$event.offsetY` - Mouse position relative to element
- `$event.ctrlKey`, `$event.shiftKey`, `$event.altKey`, `$event.metaKey` - Modifier keys

## Common Modifiers

- `prevent` - Calls `preventDefault()` (useful for links/forms)
- `stop` - Calls `stopPropagation()`
- `once` - Only trigger once
- `capture` - Use capture phase
- `outside` - Trigger when clicking outside the element
- `debounce` - Debounce rapid clicks
- `throttle` - Throttle rapid clicks

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-on)
- [MDN Click Event](https://developer.mozilla.org/en-US/docs/Web/API/Element/click_event)