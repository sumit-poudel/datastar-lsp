# data-show

Shows or hides an element based on whether an expression evaluates to `true` or `false`.

## Syntax

```html
data-show="expression"
```

## Description

The `data-show` attribute toggles the visibility of an element by setting `display: none` when the expression is falsy. Unlike `data-class` or CSS, this directly manipulates the element's style display property.

## Examples

### Basic show/hide

```html
<div data-show="$isVisible">Content</div>
<button data-on:click="$isVisible = !$isVisible">Toggle</button>
```

### Conditional display

```html
<div data-show="$user && $user.isAdmin">Admin Panel</div>
<div data-show="$items.length > 0">Items: {{$items.length}}</div>
```

## Requirements

- **Key**: Denied (not allowed)
- **Value**: Must be provided (the expression)

## Value Kind

- `expression` - The value must be a valid expression

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-show)