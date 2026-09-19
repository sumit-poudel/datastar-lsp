# data-ignore-morph

Ignores an element when patching an element using the `morph` mode.

## Syntax

```html
data-ignore-morph
```

## Description

The `data-ignore-morph` attribute prevents an element from being morphed during DOM updates when using morph mode. The element will be left as-is during patches.

## Example

```html
<div data-ignore-morph>
  <canvas id="chart"></canvas>
</div>
```

## Requirements

- **Key**: Denied (not allowed)
- **Value**: Allowed (but not used)

## Value Kind

- `string` - The value must be a string

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-ignore-morph)