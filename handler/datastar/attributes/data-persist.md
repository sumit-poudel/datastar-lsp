# data-persist

Persists signals in local storage.

## Syntax

```html
data-persist=""
data-persist:name="expression"
```

## Description

The `data-persist` attribute saves signal values to localStorage (or sessionStorage) and restores them on page load. This enables state persistence across page reloads.

## Examples

### Persist all signals

```html
<body data-persist="">...</body>
```

### Persist specific signals with custom key

```html
<div data-persist:myApp="{ user: $user, preferences: $prefs }"></div>
```

### Use session storage

```html
<div data-persist__session="$tempData"></div>
```

## Requirements

- **Key**: Allowed (optional custom storage key)
- **Value**: Allowed (optional expression for specific signals)

## Value Kind

- `expression` - The value must be a valid expression

## Modifiers

- `session` - Persists signals in session storage instead of local storage

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-persist)