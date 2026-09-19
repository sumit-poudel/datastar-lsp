# data-on:input

Runs an expression whenever the input event is triggered.

## Syntax

```html
data-on:input="expression"
data-on:input__modifier="expression"
```

## Description

The `data-on:input` attribute listens for input events, which fire whenever the value of an `<input>`, `<select>`, or `<textarea>` element changes. This is the preferred event for real-time input handling.

## Examples

### Basic input

```html
<input data-on:input="$value = $event.target.value" />
```

### Debounced search

```html
<input data-on:input__debounce.300ms="$search = $event.target.value" />
```

### Form binding

```html
<input data-bind="name" data-on:input__debounce.100ms="@validate('name', $event.target.value)" />
```

## Event Properties

- `$event.target.value` - The current input value
- `$event.target.name` - The input's name attribute
- `$event.target.type` - The input type
- `$event.inputType` - The type of input (e.g., "insertText", "deleteContentBackward")

## Common Modifiers

- `debounce` - Debounce the handler (recommended for input)
- `throttle` - Throttle the handler
- `prevent` - Calls `preventDefault()`
- `stop` - Calls `stopPropagation()`

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-on)
- [MDN Input Event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/input_event)