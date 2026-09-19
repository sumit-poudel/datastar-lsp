# data-json-signals

Sets the text content of an element to a reactive JSON stringified version of signals.

## Syntax

```html
data-json-signals="expression"
```

## Description

The `data-json-signals` attribute serializes signals to JSON and sets the element's text content. The output updates reactively when signals change. Useful for debugging or displaying signal state.

## Examples

### Display all signals

```html
<pre data-json-signals="$signals()"></pre>
```

### Display specific signals with terse modifier

```html
<pre data-json-signals__terse="{ user: $user, settings: $settings }"></pre>
```

## Requirements

- **Key**: Denied (not allowed)
- **Value**: Allowed (optional expression)

## Value Kind

- `expression` - The value must be a valid expression

## Modifiers

- `terse` - Outputs a more compact JSON format without extra whitespace. Useful for displaying filtered data inline.

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-json-signals)