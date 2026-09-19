# data-class

Adds or removes one or more classes from an element using a set of key-value pairs that map to the class name and expression.

## Syntax

```html
data-class="{ className: expression }"
data-class:name="expression"
```

## Description

The `data-class` attribute adds or removes CSS classes reactively based on expression evaluation. Classes are added when the expression evaluates to truthy and removed when falsy.

## Examples

### Object syntax (multiple classes)

```html
<div data-class="{ active: $isActive, disabled: $isDisabled }"></div>
```

### Keyed syntax (single class)

```html
<div data-class:hidden="$showHidden"></div>
```

## Requirements

- **Key**: Allowed (optional class name)
- **Value**: Must be provided (the expression)

## Value Kind

- `expression` - The value must be a valid expression

## Modifiers

- `case` - Converts the casing of the class name
  - `case.camel` - Camel case: `myClass`
  - `case.kebab` - Kebab case: `my-class` (default)
  - `case.snake` - Snake case: `my_class`
  - `case.pascal` - Pascal case: `MyClass`

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-class)