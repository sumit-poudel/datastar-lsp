# data-on:change

Runs an expression whenever the change event is triggered.

## Syntax

```html
data-on:change="expression"
data-on:change__modifier="expression"
```

## Description

The `data-on:change` attribute listens for change events, which fire when an element's value is committed (e.g., on blur for inputs, on selection for selects). Unlike `input`, this fires less frequently.

## Examples

### Select change

```html
<select data-on:change="$selected = $event.target.value">
  <option value="a">Option A</option>
  <option value="b">Option B</option>
</select>
```

### Checkbox change

```html
<input type="checkbox" data-on:change="$agreed = $event.target.checked" />
```

### File input change

```html
<input type="file" data-on:change="$file = $event.target.files[0]" />
```

## Event Properties

- `$event.target.value` - The new value
- `$event.target.checked` - For checkboxes/radios
- `$event.target.files` - For file inputs

## Common Modifiers

- `prevent` - Calls `preventDefault()`
- `stop` - Calls `stopPropagation()`
- `debounce` - Debounce the handler
- `throttle` - Throttle the handler

## Difference from `input`

| Event | Triggers | Use Case |
|-------|----------|----------|
| `input` | Every keystroke/change | Real-time feedback, filtering |
| `change` | On commit (blur/select) | Final value, form submission |

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-on)
- [MDN Change Event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/change_event)