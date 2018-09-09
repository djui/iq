# iq

INI file query tool.


## Installation

    $ brew install https://raw.githubusercontent.com/djui/iq/master/iq.rb

or

    $ go get github.com/djui/iq


## Usage

    $ echo "[section]\nkey = value\n" | iq section.key -
    value

    $ iq section.key file.ini
    value

    $ export AWS_ACCESS_KEY_ID=$(iq default.aws_access_key_id ~/.aws/credentials)
    $ export AWS_SECRET_ACCESS_KEY=$(iq default.aws_secret_access_key ~/.aws/credentials)
