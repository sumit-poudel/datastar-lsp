# data-scroll-into-view

Scrolls the element into view.

## Syntax

```html
data-scroll-into-view
data-scroll-into-view="modifier"
```

## Description

The `data-scroll-into-view` attribute scrolls the element into the viewport. Can be triggered by signal changes or used with modifiers to control scroll behavior.

## Examples

### Basic scroll

```html
<div id="target" data-scroll-into-view>Scroll to me</div>
<button data-on:click="$showTarget = true">Show Target</button>
```

### Smooth scroll

```html
<div data-scroll-into-view__smooth>Scroll smoothly</div>
```

### Scroll to center

```html
<div data-scroll-into-view__hcenter__vcenter>Scroll to center</div>
```

### Focus after scroll

```html
<input data-scroll-into-view__focus data-bind="search" />
```

## Requirements

- **Key**: Denied (not allowed)
- **Value**: Allowed (optional modifier string)

## Value Kind

- `string` - The value must be a string (modifier)

## Modifiers

- `smooth` - Scrolling is animated smoothly
- `instant` - Scrolling is instant
- `auto` - Scrolling is determined by the computed `scroll-behavior` CSS property
- `hstart` - Scrolls to the left of the element
- `hcenter` - Scrolls to the horizontal center of the element
- `hend` - Scrolls to the right of the element
- `hnearest` - Scrolls to the nearest horizontal edge of the element
- `vstart` - Scrolls to the top of the element
- `vcenter` - Scrolls to the vertical center of the element
- `vend` - Scrolls to the bottom of the element
- `vnearest` - Scrolls to the nearest vertical edge of the element
- `focus` - Focuses the element after scrolling

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-scroll-into-view)