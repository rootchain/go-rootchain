# Rootchain tutorial

## Initialize chain

Initialization of chain with power delegated to two signers.

```sh
$ rcx chain init -a default/x/first:1:1 -a default/x/second:1:1 -a default/x/third:1
Wallet "default" password: ***
{
  "header": {
    "timestamp": "2018-07-13T11:48:01.7661194+02:00",
    "head_hash": "zFuncm68waMq245BAyUgZxrfBmYBCsPcQmjkqSMsChpFXZe1h5QR",
    "exec_hash": "zFuncm69EAW3NE2Mf8cD5EwT1X4cSUdY9Dv794X5ZMo2M3Jyn76x",
    "state_hash": "z45oqTRvbf6EgnwyEsAh4a5Mk43pVYTiG8x3uBEjhSvMLE923m7",
    "signed_hash": "zFSec2XV2KL8rxWHi9YroTNvFatrxuxtgsZMgVgyfhmRsuCiGkFC"
  },
  "exec_ops": [
    "OP_ASSIGN_POWER [ OP_UINT64 1 OP_CID zFNScYMHAiBJPJcT9ri2cNwKTJfitbRcE2PSfMKQj272C6A5Fs25 ]",
    "OP_ASSIGN_POWER [ OP_UINT64 1 OP_CID zFNScYMH8j9JuHxLR5KLsNP528LuLBi7ToCrh9tmdLb85pWno5Bg ]",
    "OP_ASSIGN_POWER [ OP_UINT64 1 OP_CID zFNScYMH9HiRhyWx8nhwLmo4UfxegoSUTE7nQ17Lji2hzDVJnszf ]",
    "OP_SIGNED [ OP_DELEGATE_POWER [ OP_UINT64 1 ] OP_SIGNATURE 0x1bac70b921c4f6ccbcc50e901372d97dd3bc62fe9ebcc6fea68dd97bc72d4fccb00112f0879fc452c275457f91be3cdcb4eeb1f79d16f355e654f316144b33e7de ]",
    "OP_SIGNED [ OP_DELEGATE_POWER [ OP_UINT64 1 ] OP_SIGNATURE 0x1b01f2877e1998e25566c2e3f2bcf7a130f1636aa7cff0298d7e71615ec9eb5ddd2ba6b47fd1b1c85056b722ffdba869c9123d77a7fda2b32c5bc6818288678feb ]"
  ],
  "signatures": [
    "HJ2h0LwFpwocQ+NQ1oMo8bJZNNZrv4upBJq57hozWrCSAMmBjEjE9HzeZphL2p2PeG2i4X65cTusSw2GJkyicjA=",
    "HJzn3xOELS8q+FPRVTVZqeL+LlL/F5MbX/QJJu9dRwkDQ29u96kn+THroxnHkQna4zj2i83b0Tjuq8wrNc1ceBA="
  ]
}
```

It still requires two signers to seal the block because ceil(signers * 51%) is still two.
