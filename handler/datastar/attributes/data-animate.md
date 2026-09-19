# data-animate

Animates element attributes over time.

## Syntax

```html
data-animate:name="expression"
```

## Description

The `data-animate` attribute animates element attributes over time using CSS transitions or animations. The expression should evaluate to an object containing the animation configuration.

## Example

```html
<div data-animate:slide="{ opacity: 0, transform: 'translateX(-100%)' }"></div>
```

## Requirements

- **Key**: Must be provided (the animation name)
- **Value**: Must be provided (the animation expression)

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-animate)