# DataStar lsp in go

```
-- in nvim
vim.api.nvim_create_autocmd("FileType", {
	pattern = { "html" },
	callback = function(args)
		vim.lsp.start({
			name = "datastar lsp",
			cmd = { "<path/to/your/binary>" },
			root_dir = vim.fs.root(args.buf, { ".git", "go.mod" }) or vim.fn.getcwd(),
		})
	end,
})
```

[Data Star](https://data-star.dev/)
