# data-signals

Merges one or more signals into the existing signals using a set of key-value pairs that map to the signal name and expression.

## Syntax

```html
data-signals="{ signalName: expression }"
data-signals:name="expression"
```

## Description

The `data-signals` attribute initializes or updates multiple signals at once. It merges the provided signals into the existing signal store. Useful for setting initial state or batch updating signals.

## Examples

### Object syntax (multiple signals)

```html
<div data-signals="{ count: 0, name: '', items: [] }"></div>
```

### Keyed syntax (single signal)

```html
<div data-signals:theme="'dark'"></div>
```

### With ifmissing modifier

```html
<div data-signals__ifmissing="{ config: { theme: 'light', lang: 'en' } }"></div>
```

## Requirements

- **Key**: Allowed (optional signal name)
- **Value**: Allowed (optional expression)

## Value Kind

- `expression` - The value must be a valid expression

## Modifiers

- `case` - Converts the casing of the signal name
  - `case.camel` - Camel case: `mySignal` (default)
  - `case.kebab` - Kebab case: `my-signal`
  - `case.snake` - Snake case: `my_signal`
  - `case.pascal` - Pascal case: `MySignal`
- `ifmissing` - Only patches signals if their keys do not already exist. This is useful for setting defaults without overwriting existing values.

## Signals

- `key-or-object` - Can be used as key with object value

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-signals)