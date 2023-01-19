## A simple wordbook written in go

### Install

```sh
go build -o "$GOPATH/bin/wb"
```

### Configuration

You can get an auth from eduic here: http://my.eudic.net/OpenAPI/Authorization

Then put it in a config file which is located in *$HOME/.wbcfg.toml*, here is the format:

```toml
auth = ""
```

### Usage

#### Add new word(s) to wordbook

```sh
wb add word
```

![gBmbc.jpeg](https://i.328888.xyz/2023/02/21/gBmbc.jpeg)

#### Query words with page

```sh
wb lst 0 1
```

![gBz3q.jpeg](https://i.328888.xyz/2023/02/21/gBz3q.jpeg)

#### Delete word(s) from wordbook

```sh
wb del word
```

![gBh6x.jpeg](https://i.328888.xyz/2023/02/21/gBh6x.jpeg)
