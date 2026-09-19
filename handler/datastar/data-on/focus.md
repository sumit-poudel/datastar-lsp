# data-on:focus

Runs an expression whenever an element receives focus.

## Syntax

```html
data-on:focus="expression"
data-on:focus__modifier="expression"
```

## Description

The `data-on:focus` attribute listens for focus events, which fire when an element receives focus (via click, tab, or programmatic focus).

## Examples

### Basic focus

```html
<input data-on:focus="$focused = true" data-on:blur="$focused = false" />
```

### Show help on focus

```html
<input data-bind="password" data-on:focus="$showHelp = true" data-on:blur="$showHelp = false" />
<div data-show="$showHelp">Password must be 8+ characters</div>
```

### Select text on focus

```html
<input data-on:focus="$event.target.select()" data-bind="code" />
```

## Event Properties

- `$event.target` - The element that received focus
- `$event.relatedTarget` - The element that lost focus (if any)

## Common Modifiers

- `capture` - Use capture phase (focus doesn't bubble by default)
- `once` - Only trigger once

## Related Events

- `focusin` - Bubbles (use with `capture` or `document` modifier)
- `blur` - When element loses focus
- `focusout` - Bubbles version of blur

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-on)
- [MDN Focus Event](https://developer.mozilla.org/en-US/docs/Web/API/Element/focus_event)