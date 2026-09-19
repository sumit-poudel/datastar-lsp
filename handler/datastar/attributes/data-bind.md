# data-bind

Creates a new signal and enables two-way binding between it and the element.

## Syntax

```html
data-bind="signalName"
data-bind:name
```

## Description

The `data-bind` attribute creates a signal (if one doesn't already exist) and sets up two-way data binding between it and an element's current bound state. When the signal changes, Datastar writes that value to the element. When one of the bind events fires, Datastar reads the element's current bound property/value and writes that back to the signal.

## Examples

### Basic binding

```html
<input data-bind="username" />
```

### Keyed syntax

```html
<input data-bind:username />
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
- `prop` - Binds to a specific property instead of the default binding. Must not be a read-only property.
- `event` - Defines which events sync the element property back to the signal.

## Signals

- `key-or-value` - Can be used as either key or value

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-bind)