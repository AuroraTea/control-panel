# 规范
- 一个文件对应一个 handle 函数
- 文件和函数名同名(文件名本身是无影响的)
- HTTP Status 用 http 包内的常量
- `http.StatusInternalServerError`的直接用`panic`

---

# Style Guide
- One file for one handle function.
- File and function names of the same name.
- HTTP Status uses the constants in the http package.
- For `http.StatusInternalServerError` use `panic`.