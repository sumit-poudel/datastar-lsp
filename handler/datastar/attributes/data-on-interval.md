# data-on-interval

Runs an expression at a regular interval.

## Syntax

```html
data-on-interval="expression"
data-on-interval__modifier="expression"
```

## Description

The `data-on-interval` attribute executes an expression repeatedly at a specified interval. Useful for polling, timers, or periodic updates.

## Examples

### Basic interval (1 second default)

```html
<div data-on-interval="$time = new Date()"></div>
```

### Custom interval duration

```html
<div data-on-interval__duration.500ms="$counter++"></div>
```

### Execute immediately on start

```html
<div data-on-interval__duration.leading="$fetchData()"></div>
```

### With view transition

```html
<div data-on-interval__viewtransition="$updateChart()"></div>
```

## Requirements

- **Key**: Denied (not allowed)
- **Value**: Must be provided (the expression)

## Value Kind

- `expression` - The value must be a valid expression

## Modifiers

- `duration` - Sets the interval duration
  - `duration.500ms` - Interval duration of 500 milliseconds (accepts any integer)
  - `duration.1s` - Interval duration of 1 second (default)
  - `duration.leading` - Execute the first interval immediately
- `viewtransition` - Wraps the expression in `document.startViewTransition()` when the View Transition API is available

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-on-interval)