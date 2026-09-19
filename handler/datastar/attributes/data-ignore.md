# data-ignore

Ignores an element and its descendants from being processed.

## Syntax

```html
data-ignore
data-ignore="self"
```

## Description

The `data-ignore` attribute tells Datastar to skip processing an element and all its descendants. This is useful for preventing Datastar from interfering with third-party components or complex DOM structures.

## Examples

### Ignore element and descendants

```html
<div data-ignore>
  <third-party-component></third-party-component>
</div>
```

### Ignore only the element itself (not descendants)

```html
<div data-ignore="self">
  <span>This will still be processed</span>
</div>
```

## Requirements

- **Key**: Denied (not allowed)
- **Value**: Allowed (optional modifier)

## Value Kind

- `string` - The value must be a string (only "self" is valid)

## Modifiers

- `self` - Only ignore the element itself, not its descendants.

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-ignore)