# Rootchain commands

Rootchain commands.

## `chain init`

```sh
$ rcx chain init -h
Initializes a new chain.

See wallet usage for more information on key derivation path.

Usage:
  rcx chain init [config] [flags]

Examples:
  $ rcx chain init -n mychain -k wallet:1e6:1e6 -k default/x/test:1e6:0
  $ rcx chain init -a zFNScYMGz4wQocWbvHVqS1HcbzNzJB5JK3eAkzF9krbSLZiV8cNr:1

Flags:
  -h, --help           help for init
  -a, --addr strings   address and power in addr:power format
  -k, --key strings    key path and power in key:power:delegated format
```

## `wallet create`

```sh
$ rcx wallet create -h
Creates new wallet seed.

Default wallet name used is "default".

Usage:
  rcx wallet create [name] [flags]

Examples:
  $ rcx wallet create exampple

Flags:
  -h, --help    help for create
  -p, --print   Print mnemonic
```

## `wallet derive`

```sh
$ rcx wallet create -h
Derives key from seed and path or mnemonic name.

Path is defined as: "/<purpose>'/<coin_type>'/<account>'/<change>/<address_index>".

Mnemonic can be used for path by using --hash or -x flag.

Default wallet name used is "default".

Usage:
  rcx wallet derive [flags]

Examples:
  $ rcx wallet derive -xd memo
  $ rcx wallet derive -k wallet/x/memo
  $ rcx wallet derive -w wallet -xd memo
  $ rcx wallet derive -d m/44'/138'/0'/0/0

Flags:
  -h, --help                 help for derive
  -d, --derive-path string   derive BIP32 hierarchical path
  -x, --hash                 derive hash path
  -k, --key-path string      wallet key path (<wallet>/<x|m>/<path>)
  -p, --print                prints derivation info (default true)
  -w, --wallet string        wallet name (default "default")
```

## `wallet export`

```sh
$ rcx wallet export -h
Exports wallet private key and mnmemonic seed.

Default wallet name used is "default".

Usage:
  rcx wallet export [name] [flags]

Flags:
  -h, --help   help for export
```

## `wallet import`

```sh
$ rcx wallet import -h
Imports wallet private key or mnemonic seed.

Default wallet name used is "default".

Usage:
  rcx wallet import [name] [flags]

Flags:
  -h, --help   help for import
  -k, --key    import key instead of seed (default true)
```

## `wallet list`

```sh
$ rcx wallet list -h
Prints names of all available wallets.

Usage:
  rcx wallet list [flags]

Flags:
  -h, --help   help for list
```

## `sign`

```sh
$ rcx sign -h
Signs content with key derived from wallet.

Reads console stdin on empty arguments and -f file path flag.

Default wallet name used is "default".

Usage:
  rcx sign [content] [flags]

Examples:
  $ rcx wallet sign -w example -xd mnemonic '{"value": "0xd"}'
  $ rcx wallet sign -d m/44'/138'/0'/0/0 '{"value": "0xd"}'
  $ rcx wallet sign -w example -d m/44'/138'/0'/0/0 '{"value": "0xd"}'

Flags:
  -h, --help              help for sign
  -f, --file string       path of file to sign
  -k, --key-path string   wallet key path (<wallet>/<x|m>/<path>) (default "default")
```
