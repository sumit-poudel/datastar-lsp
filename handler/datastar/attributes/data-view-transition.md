# data-view-transition

Sets the value of `view-transition-name` for use with the View Transition API.

## Syntax

```html
data-view-transition="expression"
```

## Description

The `data-view-transition` attribute sets the CSS `view-transition-name` property on an element, enabling smooth transitions between page states using the View Transition API. The expression should evaluate to a unique name for the transition.

## Example

```html
<div data-view-transition="'user-card-' + $userId">
  <h2 data-text="$user.name"></h2>
  <img data-attr:src="$user.avatar" />
</div>
```

```css
@view-transition {
  navigation: auto;
}
```

## Requirements

- **Key**: Denied (not allowed)
- **Value**: Must be provided (the transition name expression)

## Value Kind

- `expression` - The value must be a valid expression

## References

- [DataStar Reference](https://data-star.dev/reference/attributes#data-view-transition)