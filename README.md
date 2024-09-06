<p align="center">
  <a href="https://github.com/federicotorres233">
    <picture>
      <source height="250" media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/FedericoTorres233/linuxjourney/main/public/logo.jpeg">
      <img height="250" alt="Linux" src="https://raw.githubusercontent.com/FedericoTorres233/linuxjourney/main/public/logo.jpeg">
    </picture>
  </a>
  <br>
<p align="center">
  <em><b>Linux Journey</b> is a terminal based tutorial that guides you through the amazing world of <b>Linux</b></em>
</p>

---

## ⭐️ To do

- Add progress
- Add continue
- Store progress
- Full vim compatibility
- Add more pages
- Command execution
- Tests
- Add colors

## ⚙️ Installation

The easiest way to start your journey is to get the latest binary from the github repository.

```bash
wget -qO- https://github.com/FedericoTorres233/linuxjourney/releases/download/v0.1.0/linuxjourney.tar.gz | tar xvfz -
```

Done! You can start now learning ^-^:

```bash
./bin/linuxjourney
```

### Installing from source

Linux Journey requires **Go version `1.20` or higher** to run. If you need to install or upgrade Go, visit the [official Go download page](https://go.dev/dl/). Clone the repository and cd into it:

```bash
git clone https://github.com/FedericoTorres233/linuxjourney && cd linuxjourney
```

Install project dependencies:

```bash
go mod tidy
```

Next, build the project from source. Be sure you have the `make` command installed:

```
make
```

Done! You can start using it:

```bash
./bin/linuxjourney
```

## 🎯 Chapters
Journey chapters are located in the `public/pages` folder:

```
linuxjourney/
│
└── public/
    │
    └── pages/
        │
        ├── chapter1/ # Start
        ├── chapter2/ # Intro to linux
        ├── chapter3/ # Basic commands
  
```

<!-- 
-   Something [Something](example.com)
--->
## 👍 Contribute

If you want to say **Thank You** and/or support the project:

1. Add a [GitHub Star](https://github.com/federicotorres233/linuxjourney/stargazers) to the project.
2. Submit a [pull request](https://github.com/FedericoTorres233/linuxjourney/pulls) to add a new feature
3. Create an [issue](https://github.com/FedericoTorres233/linuxjourney/issues) if you find any bug
