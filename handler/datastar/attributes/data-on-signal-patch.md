# data-on-signal-patch

Runs an expression whenever one or more signals are patched.

## Syntax

```html
data-on-signal-patch="expression"
data-on-signal-patch__modifier="expression"
```

## Description

The `data-on-signal-patch` attribute triggers an expression when signals are updated via patch operations. Can be filtered to watch specific signals using the `filter` key or `data-on-signal-patch-filter`.

## Examples

### Watch all signal patches

```html
<div data-on-signal-patch="console.log('Signals changed')"></div>
```

### With filter key

```html
<div data-on-signal-patch="filter: 'user.*'" data-on-signal-patch="console.log('User signals changed')"></div>
```

### Debounced

```html
<div data-on-signal-patch__debounce.100ms="@post('/sync', $signals)"></div>
```

## Requirements

- **Key**: Allowed (optional filter key)
- **Value**: Must be provided (the expression)

## Value Kind

- `expression` - The value must be a valid expression

## Keys

- `filter` - Filters which signals to watch

## Modifiers

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

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-on-signal-patch)