# data-on:mouseleave

Runs an expression whenever the mouse leaves an element.

## Syntax

```html
data-on:mouseleave="expression"
data-on:mouseleave__modifier="expression"
```

## Description

The `data-on:mouseleave` attribute listens for mouseleave events. Unlike `mouseout`, this does not bubble and does not fire when moving to child elements.

## Examples

### Basic hover

```html
<div data-on:mouseenter="$hovered = true" data-on:mouseleave="$hovered = false">Hover me</div>
```

### Hide tooltip

```html
<button data-on:mouseenter="$showTooltip = true" data-on:mouseleave="$showTooltip = false">?</button>
<div data-show="$showTooltip" class="tooltip">Help text</div>
```

## Event Properties

- `$event.target` - The element the mouse left
- `$event.relatedTarget` - The element the mouse entered
- `$event.clientX` / `$event.clientY` - Mouse position

## Common Modifiers

- `once` - Only trigger once

## Difference from `mouseout`

| Event | Bubbles | Fires for children |
|-------|---------|-------------------|
| `mouseleave` | No | No |
| `mouseout` | Yes | Yes |

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-on)
- [MDN MouseLeave Event](https://developer.mozilla.org/en-US/docs/Web/API/Element/mouseleave_event)