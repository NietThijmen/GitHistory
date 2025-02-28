# Git history over multiple directories
Have you ever just wanted to see the amount of commits you've done over a period of time.

With this you can!

## Usage
```bash
./GitHistory -name <name> -remote-url <github.com/bitbucket.org> -this-month=true
```

## Flags
- name: The name of the user you are searching for (this is a contains so it can also be a part of the name)
- remote-url: The remote url of the repository (let's say you only want to check for repo's that are actually yours E.G. github.com/nietthijmen)
- this-month: If you want to see the commits of this month (default is false) (This was just a thing I wanted to have in git so I added it)


## Build
```bash
go build
```

## Add to path
```bash
mv GitHistory /usr/local/bin
```

## Author
Thijmen Rierink <thijmen@rierink.dev><br/>
[GitHub](https://github.com/nietthijmen) |
[Buy me a coffee](https://www.buymeacoffee.com/nietthijmen) |
[My website](https://rierink.dev)