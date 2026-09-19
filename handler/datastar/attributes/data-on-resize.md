# data-on-resize

Runs an expression whenever an element's dimensions change.

## Syntax

```html
data-on-resize="expression"
data-on-resize__modifier="expression"
```

## Description

The `data-on-resize` attribute uses the ResizeObserver API to trigger an expression when an element's size changes. Useful for responsive layouts, adjusting charts, or dynamic sizing.

## Examples

### Basic resize

```html
<div data-on-resize="$width = $event.target.clientWidth"></div>
```

### Debounced resize

```html
<canvas data-on-resize__debounce.100ms="resizeCanvas($el)"></canvas>
```

### Throttled resize

```html
<div data-on-resize__throttle.50ms="$layout = calculateLayout($el)"></div>
```

## Requirements

- **Key**: Denied (not allowed)
- **Value**: Must be provided (the expression)

## Value Kind

- `expression` - The value must be a valid expression

## Modifiers

- `debounce` - Debounce the event listener
  - `debounce.500ms` - Debounce for 500 milliseconds (accepts any integer)
  - `debounce.1s` - Debounce for 1 second (accepts any integer)
  - `debounce.leading` - Debounce with leading edge (must come after timing)
  - `debounce.notrailing` - Debounce without trailing edge (must come after timing)
- `throttle` - Throttle the event listener
  - `throttle.500ms` - Throttle for 500 milliseconds (accepts any integer)
  - `throttle.1s` - Throttle for 1 second (accepts any integer)
  - `throttle.noleading` - Throttle without leading edge (must come after timing)
  - `throttle.trailing` - Throttle with trailing edge (must come after timing)

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-on-resize)