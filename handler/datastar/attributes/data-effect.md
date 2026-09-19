# data-effect

Executes an expression on page load and whenever any signals in the expression change.

## Syntax

```html
data-effect="expression"
```

## Description

The `data-effect` attribute runs a side effect whenever any reactive signals referenced in the expression change. It also runs once on initial page load. Use for logging, API calls, or other side effects.

## Example

```html
<div data-effect="console.log('User changed:', $user)"></div>
<div data-effect="@post('/api/save', { data: $formData })"></div>
```

## Requirements

- **Key**: Denied (not allowed)
- **Value**: Must be provided (the expression)

## Value Kind

- `expression` - The value must be a valid expression

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-effect)