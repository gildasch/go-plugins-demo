#HSLIDE

# Go plugins

Gildas Chabot, leboncoin

#HSLIDE

## Demo!

#HSLIDE

- Released in 1.8
- Add Go code at runtime
- Works only on linux for now
- Based on C's `dlfcn.h` (dynamic linking) with `cgo`
  - See [`src/plugin/plugin_dlopen.go`](https://tip.golang.org/src/plugin/plugin_dlopen.go)

#HSLIDE

## Before

- Add plugins at compile time (Caddy)
- HTTP/RPC calls

#HSLIDE

## Now: Plugin side

- Nothing special on the code of the plugin
  - `-buildmode=plugin` build option
- Creates a `.so`

#HSLIDE

## Now: Caller side

- Only two functions:
  - `Open(path string) (*Plugin, error)`
  - `(p  *Plugin) Lookup(symName string) (Symbol, error)`
- Two types:
  - `Plugin`
  - `Symbol`

#HSLIDE

## Example

```
p, err := plugin.Open("plugin.so")

fs, err := p.Lookup("Transform")
f, ok := fs.(func(image.Image) image.Image)

vs, err := p.Lookup("Priority")
v, ok = *vs.(*int)
```

#HSLIDE

## Need?

- Add ---

#HSLIDE

## Safety?

- Plugin == safe?
- Same rights as caller <!-- .element: class="fragment" -->
- Check the sources of the plugins! <!-- .element: class="fragment" -->

#HSLIDE

## Stable?

- What happens on `panic`?
- --- <!-- .element: class="fragment" -->

#HSLIDE

## Go + C?

- Go plugin used in C
  - Already done with ---  <!-- .element: class="fragment" -->
- C `.so` loaded in Go
  - For now, error --- <!-- .element: class="fragment" -->

#HSLIDE

## Thank you

Questions?

Find the code and slides on Github: https://github.com/GildasCh/go-plugins-demo
