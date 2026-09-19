# data-replace-url

Replaces the URL in the browser with an evaluated expression.

## Syntax

```html
data-replace-url="expression"
```

## Description

The `data-replace-url` attribute updates the browser's URL without adding a history entry (uses `history.replaceState`). The expression should evaluate to a URL string. Useful for updating the URL to reflect current state without cluttering history.

## Example

```html
<div data-replace-url="'/user/' + $userId"></div>
<button data-on:click="$userId = 123; $replaceUrl = '/user/' + $userId">View User 123</button>
```

## Requirements

- **Key**: Denied (not allowed)
- **Value**: Must be provided (the URL expression)

## Value Kind

- `expression` - The value must be a valid expression

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-replace-url)