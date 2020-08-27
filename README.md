# Go Network Manager client

Network Manager DBus API client generated from DBus bindings using [dbus-codegen-go](https://github.com/amenzhinsky/dbus-codegen-go).

## Usage

See [_examples](./_examples)

## Generating code / Updating API

```sh
# install dependencies
make setup
# copy XML interfaces
make copy
# generate code
make generate

# one-shot
make
```

## Generating from enum documentation

Types documentation is available at https://developer.gnome.org/NetworkManager/stable/nm-dbus-types.html

To extract, open the console in your browser and extract content using the following script and save in enum.json.

Run `go run gen/enum.go` to generate go sources from enum.json

**JS Snippet**
```js

let list = []
let clean = (v) => v.trim("\n\t ")
$(".refsect2").each((i, el) => {

  const title = clean($(el).find("h3").text()).replace("enum ", "")
  const description = clean($(el).find("h3").next().text())
  const items = []
  
  $(el).find("table tr").each((i, row) => {

    const label = clean($(row).find("td:eq(0)").text())
    const value = clean($(row).find("td:eq(1)").text()).replace("= ", "")
    const desc = clean($(row).find("td:eq(2)").text())

    items.push({label, value, desc})
  })

  list.push({
    title, description, items
  })
})

console.log(JSON.stringify(list))

```

## License

Apache 2