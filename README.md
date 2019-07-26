RAL (CLI Tool)
==============

Barebones CLI to read / post on textboards driven by the RAL software

Executive Summary
-----------------

RAL is software which powers textboards; there is one canonical instance
running at [RalEE.org](https://ralee.org/) which (as part of the standard
install) sports an API which this package (and my [RAL Textboard API for
Golang](https://github.com/wesleycoakley/ral) leverage).

This is not a fancy reader; it is aimed at
[techies and hackers](http://catb.org/jargon/html/H/hacker.html) who
do not shy away from UNIX-style utilities and who use pipes,
filters, and file redirections regularly. This makes the software very
extensible and scriptable.

Installation
------------

Ensure your `$GOPATH` is configured, then run:

```
go get github.com/wesleycoakley/ral
```

This will pull all dependencies and install the `raleexplorer` binary in
`$GOPATH/bin`; there are only a handful of deps. (which you don't
need to worry about unless you're interested) but I'll list them here
for easy reference:

- Naoina's [TOML parser / encoder](https://github.com/naoina/toml)
- My [RAL API for Golang](https://github.com/wesleycoakley/ral-api)

Configuration
-------------

The RAL CLI tool reads from a TOML-formatted configuration file; it searches
the following locations (in order) for a proper file:

1. ./config.toml
2. ~/.ral/config.toml
3. ~/.ralrc

Please copy the included `config.template.toml` file to one of the above
places and edit it to match your use case.

Contributing
------------
Contributions / PRs are welcome; e-mail me at
[w@wesleycoakley.com](mailto:w@wesleycoakley.com) if you
have a suggestion or want to work together.

License
-------

X11 License (available in the source tree as `/LICENSE`)
