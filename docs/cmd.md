# Rootchain commands

Rootchain commands.

## `rcx chain init` - initializes a new chain

### Usage

```sh
$ rcx chain init -h
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

### Example

Initialization of chain with power delegated to two of three signers.

```sh
$ rcx chain init -k default/x/first:1:1 -k default/x/second:1:1 -a zFNScYMH2PZmRrdEF3aP7HrHVM2HegvKwMa2yFMKeZ2wwQR35EXy:1
Wallet "default" password: *********
{
  "header": {
    "timestamp": "2018-07-05T10:27:01.5903054+02:00",
    "head_hash": "zHeadAt956yHU2WfzQrh626bh2Fy6UVf5SrxL9Px1x6gCaLDweWJ",
    "exec_hash": "zFuncm69BbwDVxbj9yUR3wPJXz8hW4W1k1pAj6tsY6eNB2vDGCyW",
    "state_hash": "z45oqTS7zRzv38NNimH3f7f6aVSMwUYX5KGz56rcpof5Dbpqgn4",
    "signed_hash": "zFSec2XVBVTrSFkwnqHBEEYQr31yuHGrugqq419XfK4smHGZq4p6"
  },
  "exec_ops": [
    "OP_GENESIS",
    "OP_ASSIGN_POWER [ OP_UINT64 1 OP_PUBKEY 0x03c4f082b4c4706f08da72fa507891ad05e03f24747065ec2dd52da9912c092282 ]",
    "OP_ASSIGN_POWER [ OP_UINT64 1 OP_PUBKEY 0x022a67842b81daa122e6fdb39ce9ed1b6be3d2ff43061e94acb4b423e50732e0f8 ]",
    "OP_ASSIGN_POWER [ OP_UINT64 1 OP_CID zFNScYMH2PZmRrdEF3aP7HrHVM2HegvKwMa2yFMKeZ2wwQR35EXy ]",
    "OP_SIGNED [ OP_ASSIGN_POWER [ OP_UINT64 1 ] OP_SIGNATURE 0x1b0288e8ec73236b734c26d314f35c906e4f8934f157f6cdde0151486c3195346344de689ad7de5345e427010fa54c4350eaae1aa90545f13d83e9404229a21049 ]",
    "OP_SIGNED [ OP_ASSIGN_POWER [ OP_UINT64 1 ] OP_SIGNATURE 0x1b126a9964e8cc1ba77968a4b0b4ca29f602e3a2634db9cccee7a68a0158996a98238c3c011bd329c76d8809defd563938b07f0160d5c7cf94a8bdd707e8ed1017 ]"
  ],
  "signatures": [
    "HMRJFr8Bz4gMuspY/jHxNy5+o8IA94NlP/ducIXjr0nmCCodwAL7/tXyl/Oq7uCFS/tasBkeoJ2yyEnfi7bQGj8=",
    "GzC6XBpGsAv2Oc0mejJPYOapdIbnLaQSzy3aFF15wchuaafGLBXc9bD1TpI22gHed/qzzzW9SSEz6gIsFRBBzas="
  ]
}
```