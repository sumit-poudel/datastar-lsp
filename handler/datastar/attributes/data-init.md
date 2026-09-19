# data-init

Runs an expression whenever the attribute is initialised.

## Syntax

```html
data-init="expression"
```

## Description

The `data-init` attribute executes an expression when the element is initialized. This runs once when the element is first processed by Datastar. Useful for setting up initial state, fetching data, or registering event listeners.

## Examples

### Initialize signals

```html
<div data-init="$count = 0, $items = []"></div>
```

### Fetch initial data

```html
<div data-init="@get('/api/user').then(r => $user = r)"></div>
```

### With delay modifier

```html
<div data-init__delay.1s="$notification = 'Welcome!'"></div>
```

## Requirements

- **Key**: Denied (not allowed)
- **Value**: Must be provided (the expression)

## Value Kind

- `expression` - The value must be a valid expression

## Modifiers

- `delay` - Delay the event listener
  - `delay.500ms` - Delay for 500 milliseconds (accepts any integer)
  - `delay.1s` - Delay for 1 second (accepts any integer)
- `viewtransition` - Wraps the expression in `document.startViewTransition()` when the View Transition API is available.

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-init)