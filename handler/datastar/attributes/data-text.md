# data-text

Sets the text content of an element to the evaluated expression.

## Syntax

```html
data-text="expression"
```

## Description

The `data-text` attribute sets the element's `textContent` to the result of the expression. The content updates reactively whenever signals in the expression change. HTML is escaped (not rendered).

## Examples

### Basic text binding

```html
<span data-text="$user.name"></span>
```

### Computed text

```html
<span data-text="'Hello, ' + $name + '!'"></span>
```

### Conditional text

```html
<span data-text="$count === 1 ? '1 item' : $count + ' items'"></span>
```

## Requirements

- **Key**: Denied (not allowed)
- **Value**: Must be provided (the expression)

## Value Kind

- `expression` - The value must be a valid expression

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-text)