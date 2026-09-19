# data-ref

Creates a signal whose value references an element.

## Syntax

```html
data-ref="signalName"
data-ref:name
```

## Description

The `data-ref` attribute creates a signal that holds a reference to the DOM element. This allows you to access the element directly in expressions for methods like `focus()`, `scrollIntoView()`, or measuring dimensions.

## Examples

### Basic ref

```html
<input data-ref="myInput" />
<button data-on:click="$myInput.focus()">Focus Input</button>
```

### Keyed syntax

```html
<input data-ref:myInput />
<button data-on:click="$myInput.scrollIntoView()">Scroll to Input</button>
```

### Measuring element

```html
<div data-ref:box></div>
<div data-effect="console.log('Box size:', $box.getBoundingClientRect())"></div>
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

- [DataStar Reference](https://data-star.dev/reference/attributes#data-ref)