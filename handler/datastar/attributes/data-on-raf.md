# data-on-raf

Runs an expression on every `requestAnimationFrame` event.

## Syntax

```html
data-on-raf="expression"
data-on-raf__modifier="expression"
```

## Description

The `data-on-raf` attribute executes an expression on every browser animation frame (typically 60fps). Useful for smooth animations, game loops, or real-time visual updates. Use throttling to reduce frequency.

## Examples

### Basic RAF

```html
<canvas data-on-raf="drawFrame($canvas)"></canvas>
```

### Throttled RAF

```html
<div data-on-raf__throttle.16ms="$updatePosition()"></div>
```

### Without leading edge

```html
<div data-on-raf__throttle.noleading="$render()"></div>
```

## Requirements

- **Key**: Denied (not allowed)
- **Value**: Must be provided (the expression)

## Value Kind

- `expression` - The value must be a valid expression

## Modifiers

- `throttle` - Throttle the event listener
  - `throttle.500ms` - Throttle for 500 milliseconds (accepts any integer)
  - `throttle.1s` - Throttle for 1 second (accepts any integer)
  - `throttle.noleading` - Throttle without leading edge (must come after timing)
  - `throttle.trailing` - Throttle with trailing edge (must come after timing)

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-on-raf)