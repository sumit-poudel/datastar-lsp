# data-nonce

Enables Content Security Policy mode using a nonempty nonce. This attribute is only allowed on the `html` element.

## Syntax

```html
data-nonce="nonce-value"
```

## Description

The `data-nonce` attribute enables CSP (Content Security Policy) mode by providing a nonce value. This attribute must be placed on the `<html>` element and the nonce must match the one in your CSP header.

## Example

```html
<html data-nonce="abc123">
  <head>
    <script nonce="abc123" src="https://cdn.example.com/datastar.js"></script>
  </head>
  <body>...</body>
</html>
```

## Requirements

- **Key**: Denied (not allowed)
- **Value**: Must be provided (the nonce string)

## Value Kind

- `string` - The value must be a string

## Element Restriction

- Only allowed on the `html` element

## Highlight

- `false` - Not highlighted in syntax

## References

- [DataStar Reference](https://data-star.dev/reference/security#content-security-policy)