# iq

INI file query tool.


## Installation

    $ go get github.com/djui/iq


## Usage

    $ echo "[section]\nkey = value\n" | jq section.key -
    value

    $ jq section.key file.ini
    value
