# Rootchain tutorial

## Initialize chain

Initialization of chain with power delegated to three signers.

```sh
$ rcx chain init -k default/x/first:1:1 -k default/x/second:1:1 -k default/x/third:1:1
Wallet "default" password: ***
{
  "header": {
    "timestamp": "2018-07-12T00:32:08.1844726+02:00",
    "head_hash": "zHeadAt93LVDzUVUcAM4dtTV19EE3XDGZpeb3Nq6ut5DhUfmZvL2",
    "exec_hash": "zFuncm695mVfdP9x43zEtMpvoc3F5vo1weWBnQUPUxUqBwdBjrRP",
    "state_hash": "z45oqTS7TS5VYK58xbDZToVyg4z6L7etNv7u2QPSWAv2ZBn7dtv",
    "signed_hash": "zFSec2XV9h8irDxUpwiMUEjLrMx8haVXfCq7pnF33eLnAVqPbBBQ"
  },
  "exec_ops": [
    "OP_ASSIGN_POWER [ OP_UINT64 1000000 OP_CID zFNScYMHAiBJPJcT9ri2cNwKTJfitbRcE2PSfMKQj272C6A5Fs25 ]",
    "OP_ASSIGN_POWER [ OP_UINT64 1000000 OP_CID zFNScYMH8j9JuHxLR5KLsNP528LuLBi7ToCrh9tmdLb85pWno5Bg ]",
    "OP_ASSIGN_POWER [ OP_UINT64 1000000 OP_CID zFNScYMH9HiRhyWx8nhwLmo4UfxegoSUTE7nQ17Lji2hzDVJnszf ]",
    "OP_SIGNED [ OP_DELEGATE_POWER [ OP_UINT64 1000000 ] OP_SIGNATURE 0x1b3f3b976e8798abdb51b4291605bf35b15e97404afe7872e0232138842c17ff1d443e6d9234931f7d56cd86bf8112d961280cfeb19bda20fb5ea56c743f3f5ae1 ]",
    "OP_SIGNED [ OP_DELEGATE_POWER [ OP_UINT64 1000000 OP_NONCE 1 ] OP_SIGNATURE 0x1c89b63a590c283a2630a531a9451676d6b3435da1a49be5750c346d7d3e2e31665b46a770e6c99253059afc26eb2ecde013a6e8f1a1976dcfac52abdcacb4cacb ]"
  ],
  "signatures": [
    "GwlJq70m0Ae4aRGFtOXi+kUEHAyT5Zxo4L7C414d/e28DR7KejqzEFhsQY0UwRxrpXBQa3MXRNAJMpy535JfoXE=",
    "G3aXDA+3ucg7DUvsTRcy32dMAni+jqJZRIFHSyI5wx0cZIeqwT9Fq5NIYjjb93Wg9VDL5UhieEWD5JnjN2aLxAI="
  ]
}
```
