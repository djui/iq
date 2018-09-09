# iq

INI file query tool.


## Installation

    $ go get github.com/djui/iq


## Usage

    $ echo "[section]\nkey = value\n" | iq section.key -
    value

    $ iq section.key file.ini
    value
