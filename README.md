# Go Build Wrapper

This tool helps you move your Go build artifacts to your PATH for easier execution.

## How to use
```shell
git clone https://github.com/u-kai/go-build-wrapper.git
cd go-build-wrapper
# gbw allows you to replace a string of your choice (optional).Default is project root directory name.
go run main.go gbw
cd path/to/your-go-project
# your-go-project-name allows you to replace a string of your choice (optional).Default is project root directory name.
gbw your-go-project-name
# Now you can run your Go project directly
your-go-project-name
```
