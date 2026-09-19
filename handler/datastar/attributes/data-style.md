# data-style

Adds or removes one or more inline CSS styles from an element using a set of key-value pairs that map to the style name and expression.

## Syntax

```html
data-style="{ styleName: expression }"
data-style:name="expression"
```

## Description

The `data-style` attribute sets inline CSS styles reactively. Styles are applied when the expression evaluates to a truthy value and removed when falsy (for boolean-like values) or updated with the new value.

## Examples

### Object syntax (multiple styles)

```html
<div data-style="{ color: $textColor, 'font-size': $fontSize + 'px' }"></div>
```

### Keyed syntax (single style)

```html
<div data-style:opacity="$isVisible ? 1 : 0.5"></div>
```

### Conditional styles

```html
<div data-style="{ display: $show ? 'block' : 'none', transform: 'rotate(' + $angle + 'deg)' }"></div>
```

## Requirements

- **Key**: Allowed (optional style name)
- **Value**: Must be provided (the expression)

## Value Kind

- `expression` - The value must be a valid expression

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-style)