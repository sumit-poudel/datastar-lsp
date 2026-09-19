# data-computed

Creates one or more signals that are computed based on a set of key-value pairs that map to the signal name and expression.

## Syntax

```html
data-computed="{ signalName: expression }"
data-computed:name="expression"
```

## Description

The `data-computed` attribute creates derived signals that automatically update when their dependencies change. Computed signals are read-only and recalculated whenever any signal in their expression changes.

## Examples

### Object syntax (multiple computed signals)

```html
<div data-computed="{ fullName: '$firstName + \" \" + $lastName', itemCount: '$items.length' }"></div>
```

### Keyed syntax (single computed signal)

```html
<div data-computed:total="$price * $quantity"></div>
```

## Requirements

- **Key**: Allowed (optional signal name)
- **Value**: Must be provided (the expression)

## Value Kind

- `expression` - The value must be a valid expression

## Modifiers

- `case` - Converts the casing of the signal name
  - `case.camel` - Camel case: `mySignal` (default)
  - `case.kebab` - Kebab case: `my-signal`
  - `case.snake` - Snake case: `my_signal`
  - `case.pascal` - Pascal case: `MySignal`

## Signals

- `key-or-object` - Can be used as key with object value

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-computed)