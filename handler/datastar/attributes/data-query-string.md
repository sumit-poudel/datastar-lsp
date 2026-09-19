# data-query-string

Syncs query string params to signal values on page load, and syncs signal values to query string params on change.

## Syntax

```html
data-query-string=""
data-query-string__modifier="expression"
```

## Description

The `data-query-string` attribute synchronizes URL query parameters with signals. On page load, query params are read into signals. When signals change, the URL is updated (without page reload).

## Examples

### Basic sync (all signals)

```html
<div data-query-string="">...</div>
```

### With history support

```html
<div data-query-string__history="">...</div>
```

### Filter empty values

```html
<div data-query-string__filter="">...</div>
```

### Combined modifiers

```html
<div data-query-string__filter__history="">...</div>
```

## Requirements

- **Key**: Denied (not allowed)
- **Value**: Allowed (optional expression)

## Value Kind

- `expression` - The value must be a valid expression

## Modifiers

- `filter` - Filters out empty values when syncing signal values to query string params
- `history` - Enables history support – each time a matching signal changes, a new entry is added to the browser's history stack. Signal values are restored from the query string params on popstate events.

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-query-string)