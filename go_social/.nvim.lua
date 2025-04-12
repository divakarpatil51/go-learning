
local auto_save_file_group = vim.api.nvim_create_augroup('AutoSavePyFileGroup', { clear = true })
vim.api.nvim_create_autocmd({ 'BufWritePost' }, {
    group = auto_save_file_group,
    callback = function()
        local file_name = vim.fn.expand('%:p')
        local file_extension = file_name:match("^.+(%..+)$")

        if file_extension == ".go" then
            vim.cmd("silent !go fmt %")
        else
            vim.api.nvim_command("lua vim.lsp.buf.format()")
        end
    end
})
