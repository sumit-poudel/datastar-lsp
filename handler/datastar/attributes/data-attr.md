# data-attr

Sets one or more attribute values using a set of key-value pairs that map to the attribute name and expression.

## Syntax

```html
data-attr="{ attributeName: expression }"
data-attr:name="expression"
```

## Description

The `data-attr` attribute sets HTML attribute values reactively. You can set multiple attributes at once using an object syntax, or a single attribute using the keyed syntax.

## Examples

### Object syntax (multiple attributes)

```html
<input data-attr="{ disabled: $loading, readonly: $readonly }" />
```

### Keyed syntax (single attribute)

```html
<button data-attr:aria-label="$buttonLabel">Submit</button>
```

## Requirements

- **Key**: Allowed (optional attribute name)
- **Value**: Must be provided (the expression)

## Modifiers

- `case` - Converts the casing of the attribute name
  - `case.camel` - Camel case (default)
  - `case.kebab` - Kebab case
  - `case.snake` - Snake case
  - `case.pascal` - Pascal case

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-attributes)