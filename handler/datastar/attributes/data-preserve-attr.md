# data-preserve-attr

Preserves the value of an attribute when patching elements using the `morph` mode.

## Syntax

```html
data-preserve-attr="attributeName"
```

## Description

The `data-preserve-attr` attribute prevents a specific attribute from being overwritten during DOM morphing. This is useful when you have attributes managed by external libraries that shouldn't be touched by Datastar's morphing algorithm.

## Example

```html
<canvas data-preserve-attr="width" data-preserve-attr="height"></canvas>
```

## Requirements

- **Key**: Denied (not allowed)
- **Value**: Must be provided (the attribute name to preserve)

## Value Kind

- `string` - The value must be a string (attribute name)

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-preserve-attr)