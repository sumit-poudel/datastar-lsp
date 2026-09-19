# data-custom-validity

Adds custom validity to an element using an expression.

## Syntax

```html
data-custom-validity="expression"
```

## Description

The `data-custom-validity` attribute sets a custom validation message on a form element. When the expression evaluates to a non-empty string, the element becomes invalid and displays the message.

## Example

```html
<input 
  data-bind="email" 
  data-custom-validity="$email.includes('@') ? '' : 'Email must contain @'" 
/>
```

## Requirements

- **Key**: Denied (not allowed)
- **Value**: Must be provided (the expression)

## Value Kind

- `expression` - The value must be a valid expression

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-custom-validity)