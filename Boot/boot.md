# Fresh instalation of [Debian](https://www.debian.org) computer.

# My `.bashrc`

```bash
export EDITOR='nvim'
export VISUAL='nvim'

# Git aliases
alias gtpl='git pull'
alias gtps='git push'
alias gadd='git add .'
alias gcom='git commit'
alias gdif='git diff'

# Meson aliases
alias mess='meson setup build/'
alias mesc='meson compile -C build/'
alias mesi='meson install -C build/'

# Programming languages aliases
alias p='python3'
alias c='g++ -o main *.cpp'
alias h='ghci'
alias gor='go run *.go'
alias gof='go fmt *.go'

# Update and install Rust language.
alias rustinstall="curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh"

# Update and install Scala lang.
alias scalainstall="curl -fL https://github.com/coursier/coursier/releases/latest/download/cs-x86_64-pc-linux.gz | gzip -d > cs && chmod +x cs && ./cs setup"

# Update and install bun.
alias buninstall="curl -fsSL https://bun.sh/install | bash"
```

# Instalation of all apps

```bash
apt update
apt upgrade

general="firefox-esr thunderbird vlc freefilesync neovim"
media="gimp inkscape handbrake kid3-qt asunder obs-studio"
chat="telegram-desktop"
kde="kdiff3 kdenlive k3b kate karbon krita krename"

office="texstudio tikzit libreoffice texlive-full pandoc ipe xournalpp ghostwriter calligra"
tools="screenfetch htop tree curl exiftool bleachbit"

dev="git python3 ruby"
cpp="cpp cppcheck meson cmake make valgrind doxygen"
java="default-jre default-jdk"
math="sagemath octave jupyter polymake"
go="golang gopls"
perl="perl"
haskell="ghc ghc-prof ghc-doc"
prolog="swi-prolog-full"

apps="$general $media $chat $kde $office $tools $dev $cpp $java $math $go $perl $haskell $prolog"
apt install $apps

rustinstall
scalainstall

cargo install juliaup
cargo install mdbook
rustup install rust-analyzer
```
