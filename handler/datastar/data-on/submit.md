# data-on:submit

Runs an expression whenever a form is submitted.

## Syntax

```html
data-on:submit="expression"
data-on:submit__modifier="expression"
```

## Description

The `data-on:submit` attribute listens for form submit events. Typically used with the `prevent` modifier to prevent default form submission and handle it via AJAX.

## Examples

### Basic form submit

```html
<form data-on:submit__prevent="@post('/api/submit', $formData)">
  <input data-bind="email" />
  <button type="submit">Submit</button>
</form>
```

### With validation

```html
<form data-on:submit__prevent="$valid && @post('/api/submit', $formData)">
  <input data-bind="email" data-custom-validity="$valid ? '' : 'Invalid email'" />
  <button type="submit">Submit</button>
</form>
```

### Using FormData

```html
<form data-on:submit__prevent="const fd = new FormData($event.target); @post('/api/submit', fd)">
  <input data-bind="name" />
  <button type="submit">Submit</button>
</form>
```

## Event Properties

- `$event.target` - The form element
- `$event.target.elements` - Form controls collection
- `$formData` - Special variable containing FormData (when using data-bind)

## Common Modifiers

- `prevent` - **Essential** - Prevents default form submission (page reload)
- `stop` - Calls `stopPropagation()`
- `once` - Only trigger once

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-on)
- [MDN Submit Event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLFormElement/submit_event)