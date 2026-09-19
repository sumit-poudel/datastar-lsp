# data-on-intersect

Runs an expression on intersection with the viewport.

## Syntax

```html
data-on-intersect="expression"
data-on-intersect__modifier="expression"
```

## Description

The `data-on-intersect` attribute uses the Intersection Observer API to trigger an expression when an element enters or exits the viewport. Useful for lazy loading, infinite scrolling, or triggering animations.

## Examples

### Basic intersection

```html
<div data-on-intersect="@get('/api/load-more')">...</div>
```

### Trigger once only

```html
<img data-on-intersect__once="$loaded = true" data-src="$imageUrl" />
```

### Trigger on exit

```html
<div data-on-intersect__exit="console.log('Element left viewport')"></div>
```

### Threshold-based

```html
<div data-on-intersect__threshold.75="$visible = true">...</div>
```

## Requirements

- **Key**: Denied (not allowed)
- **Value**: Must be provided (the expression)

## Value Kind

- `expression` - The value must be a valid expression

## Modifiers

- `once` - Only triggers the event once
- `exit` - Only triggers the event when the element exits the viewport
- `half` - Triggers when half of the element is visible
- `full` - Triggers when the full element is visible
- `threshold` - Triggers when the element is visible by a certain percentage
  - `threshold.25` - Triggers when 25% of the element is visible
  - `threshold.75` - Triggers when 75% of the element is visible
- `delay` - Delay the event listener
  - `delay.500ms` - Delay for 500 milliseconds (accepts any integer)
  - `delay.1s` - Delay for 1 second (accepts any integer)
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
- `viewtransition` - Wraps the expression in `document.startViewTransition()` when the View Transition API is available

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-on-intersect)