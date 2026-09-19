# data-on:blur

Runs an expression whenever an element loses focus.

## Syntax

```html
data-on:blur="expression"
data-on:blur__modifier="expression"
```

## Description

The `data-on:blur` attribute listens for blur events, which fire when an element loses focus. Useful for validation, saving data, or hiding UI.

## Examples

### Basic blur

```html
<input data-on:blur="$touched = true" data-bind="email" />
```

### Validate on blur

```html
<input 
  data-bind="email" 
  data-on:blur="$errors.email = validateEmail($event.target.value)"
  data-custom-validity="$errors.email"
/>
```

### Save on blur

```html
<input data-bind="name" data-on:blur="@post('/api/save', { name: $event.target.value })" />
```

## Event Properties

- `$event.target` - The element that lost focus
- `$event.relatedTarget` - The element that gained focus (if any)

## Common Modifiers

- `once` - Only trigger once

## Related Events

- `focusout` - Bubbles (use with `document` modifier)
- `focus` - When element gains focus
- `focusin` - Bubbles version of focus

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-on)
- [MDN Blur Event](https://developer.mozilla.org/en-US/docs/Web/API/Element/blur_event)