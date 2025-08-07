# go-fractal

This repo generates a fractal image, based on code written for the Golang
online tutorial.

## Equation

The formula basically boils down to:

```
z := (x^2 + y^2) % 256
```

## Result

<img src="https://raw.githubusercontent.com/leaf-node/go-fractal/refs/heads/main/fractal.png">
