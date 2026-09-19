# data-on:mouseenter

Runs an expression whenever the mouse enters an element.

## Syntax

```html
data-on:mouseenter="expression"
data-on:mouseenter__modifier="expression"
```

## Description

The `data-on:mouseenter` attribute listens for mouseenter events. Unlike `mouseover`, this does not bubble and does not fire when moving over child elements.

## Examples

### Basic hover

```html
<div data-on:mouseenter="$hovered = true" data-on:mouseleave="$hovered = false">Hover me</div>
```

### Show tooltip

```html
<button data-on:mouseenter="$showTooltip = true" data-on:mouseleave="$showTooltip = false">?</button>
<div data-show="$showTooltip" class="tooltip">Help text</div>
```

## Event Properties

- `$event.target` - The element the mouse entered
- `$event.relatedTarget` - The element the mouse came from
- `$event.clientX` / `$event.clientY` - Mouse position

## Common Modifiers

- `once` - Only trigger once

## Difference from `mouseover`

| Event | Bubbles | Fires for children |
|-------|---------|-------------------|
| `mouseenter` | No | No |
| `mouseover` | Yes | Yes |

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-on)
- [MDN MouseEnter Event](https://developer.mozilla.org/en-US/docs/Web/API/Element/mouseenter_event)