# data-match-media

Sets a signal to whether a media query matches and keeps it in sync whenever the query changes.

## Syntax

```html
data-match-media:name="expression"
```

## Description

The `data-match-media` attribute creates a signal that reflects the match state of a CSS media query. The signal updates automatically when the media query match state changes (e.g., window resize).

## Example

```html
<div data-match-media:isMobile="(max-width: 768px)">
  <span data-show="$isMobile">Mobile view</span>
  <span data-show="!$isMobile">Desktop view</span>
</div>
```

## Requirements

- **Key**: Must be provided (the signal name)
- **Value**: Must be provided (the media query expression)

## Value Kind

- `expression` - The value must be a valid media query expression

## Modifiers

- `case` - Converts the casing of the signal name
  - `case.camel` - Camel case: `mySignal` (default)
  - `case.kebab` - Kebab case: `my-signal`
  - `case.snake` - Snake case: `my_signal`
  - `case.pascal` - Pascal case: `MySignal`

## Signals

- `key` - Used as key only

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-match-media)