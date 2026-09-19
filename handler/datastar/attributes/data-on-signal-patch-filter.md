# data-on-signal-patch-filter

Filters which signals to watch when using the `data-on-signal-patch` attribute.

## Syntax

```html
data-on-signal-patch-filter="expression"
```

## Description

The `data-on-signal-patch-filter` attribute provides a filter expression to limit which signal changes trigger the `data-on-signal-patch` handler. The expression should evaluate to a pattern or function that matches signal names.

## Example

```html
<div 
  data-on-signal-patch-filter="$signal.startsWith('user.')" 
  data-on-signal-patch="console.log('User signal changed:', $signal)"
></div>
```

## Requirements

- **Key**: Denied (not allowed)
- **Value**: Must be provided (the filter expression)

## Value Kind

- `expression` - The value must be a valid expression

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-on-signal-change)