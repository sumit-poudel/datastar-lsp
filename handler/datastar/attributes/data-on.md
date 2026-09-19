# data-on

Runs an expression whenever an event is triggered on an element.

## Syntax

```html
data-on:event="expression"
data-on:event__modifier="expression"
```

## Description

The `data-on` attribute attaches event listeners to elements. When the specified event occurs, the expression is evaluated. Supports all standard DOM events plus custom events.

## Examples

### Click event

```html
<button data-on:click="$count++">Increment</button>
```

### Multiple modifiers

```html
<input data-on:input__debounce.300ms="$search = $event.target.value" />
```

### Window event

```html
<div data-on:resize__window="console.log('Window resized')"></div>
```

### Prevent default

```html
<form data-on:submit__prevent="@post('/submit', $formData)">...</form>
```

## Requirements

- **Key**: Must be provided (the event name)
- **Value**: Must be provided (the expression)

## Value Kind

- `expression` - The value must be a valid expression

## Modifiers

- `once` - Only trigger the event listener once
- `passive` - Do not call `preventDefault` on the event listener
- `capture` - Use a capture event listener
- `case` - Converts the casing of the event
  - `case.camel` - Camel case: `myEvent`
  - `case.kebab` - Kebab case: `my-event` (default)
  - `case.snake` - Snake case: `my_event`
  - `case.pascal` - Pascal case: `MyEvent`
- `delay` - Delay the event listener
  - `delay.500ms` - Delay for 500 milliseconds (accepts any integer)
  - `delay.1s` - Delay for 1 second (accepts any integer)
- `debounce` - Debounce the event listener
  - `debounce.500ms` - Debounce for 500 milliseconds (accepts any integer)
  - `debounce.1s` - Debounce for 1 second (accepts any integer)
  - `debounce.leading` - Debounce with leading edge (must come after timing)
  - `debounce.notrailing` - Debounce without trailing edge (must come after timing)
- `throttle` - Throttle the event listener
  - `throttle.500ms` - Throttle for 500 milliseconds (accepts any integer)
  - `throttle.1s` - Throttle for 1 second (accepts any integer)
  - `throttle.noleading` - Throttle without leading edge (must come after timing)
  - `throttle.trailing` - Throttle with trailing edge (must come after timing)
- `viewtransition` - Wraps the expression in `document.startViewTransition()` when the View Transition API is available
- `window` - Attaches the event listener to the `window` element
- `document` - Attaches the event listener to the `document` element
- `outside` - Triggers when the event is outside the element
- `prevent` - Calls `preventDefault` on the event listener
- `stop` - Calls `stopPropagation` on the event listener

## Native Events

All standard DOM events are supported including: `click`, `keydown`, `keyup`, `input`, `change`, `submit`, `focus`, `blur`, `mouseenter`, `mouseleave`, `scroll`, `resize`, and many more.

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-on)