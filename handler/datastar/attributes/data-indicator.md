# data-indicator

Creates a signal to track in-flight backend requests.

## Syntax

```html
data-indicator="signalName"
data-indicator:name
```

## Description

The `data-indicator` attribute creates a signal that tracks the number of in-flight backend requests. The signal increments when a request starts and decrements when it completes. Useful for showing loading states.

## Examples

### Basic indicator

```html
<button data-indicator="saving" data-on:click="@post('/save', $form)">Save</button>
<div data-show="$saving">Saving...</div>
```

### Keyed syntax

```html
<button data-indicator:saving data-on:click="@post('/save', $form)">Save</button>
```

## Requirements

- **Key**: Exclusive (cannot be combined with other exclusive attributes)
- **Value**: Exclusive (signal name)

## Value Kind

- `signal-name` - The value must be a valid signal name

## Modifiers

- `case` - Converts the casing of the signal name
  - `case.camel` - Camel case: `mySignal` (default)
  - `case.kebab` - Kebab case: `my-signal`
  - `case.snake` - Snake case: `my_signal`
  - `case.pascal` - Pascal case: `MySignal`

## Signals

- `key-or-value` - Can be used as either key or value

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-indicator)