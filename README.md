# Gofer
Yet another Discord - Telegram Bridge, built with Go!

## Getting Started

* First, download the latest release from [GitHub Releases](/srevinsaju/gofer/releases)
```bash
wget https://github.com/srevinsaju/gofer/releases/download/continuous/gofer
chmod +x gofer
```

* Create a `gofer` configuration file

```bash
./gofer create
```
Enter the values when asked for. Type `EXIT` once done.
A `gofer.json` will be created in the current directory.

* Run `gofer` against the generated configuration

```bash
./gofer path/to/gofer.json
```

* Enjoy

## Why `gofer` ?
### Why yet another telegram-discord bridge?
Well, I was learning `golang`, and I wanted a project to get my hands dirty on 😂, so a discord-telegram bridge looked like a cool stuff I could make!

### Why the name?
`gofer`, because of its prefix `go-` and because `gofer` is literally known as a messenger / person who runs errands. As `gofer` gets messages from discord and posts them to telegram and vice-versa, it sounded like an appropriate name. Also, fun fact, take a look at [Gopher's page](https://blog.golang.org/gopher) too!

## License
[This project](https://github.com/srevinsaju/gofer) is licensed under the MIT License.
For more information, see [LICENSE](./LICENSE).
