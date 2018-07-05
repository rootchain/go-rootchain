# Rootchain tutorial

## Initialize chain

Initialization of chain with power delegated to three signers.

```sh
$ rcx chain init -k default/x/first:1:1 -k default/x/second:1:1 -k default/x/third:1:1
Wallet "default" password: ***
{
  "header": {
    "timestamp": "2018-07-05T11:45:20.9946756+02:00",
    "head_hash": "zHeadAt9D2JoqSgRz5ahvvjw9uFQpPyEkPgx7g1VnLQkYXPuTfEq",
    "exec_hash": "zFuncm69AfZUCm4rSekZEgDKBu75Tzi1ZNGjoozcPRW4DoA2eAKt",
    "state_hash": "z45oqTRvr4SrB4zZ2V3aUiteMShMJ6JGD5TzDLCC2UpjWLwNaBD",
    "signed_hash": "zFSec2XV1x4ccDgnErohJh2QG8mKF5FbMzKXEarW9gcZkpbnZjWg"
  },
  "exec_ops": [
    "OP_GENESIS",
    "OP_ASSIGN_POWER [ OP_UINT64 1 OP_PUBKEY 0x03c4f082b4c4706f08da72fa507891ad05e03f24747065ec2dd52da9912c092282 ]",
    "OP_ASSIGN_POWER [ OP_UINT64 1 OP_PUBKEY 0x022a67842b81daa122e6fdb39ce9ed1b6be3d2ff43061e94acb4b423e50732e0f8 ]",
    "OP_ASSIGN_POWER [ OP_UINT64 1 OP_PUBKEY 0x0344b07e90213f018429856b9fb08f6b80f15630b60c9b6e5dcdd380a11aaef527 ]",
    "OP_SIGNED [ OP_DELEGATE_POWER [ OP_UINT64 1 ] OP_SIGNATURE 0x1b0288e8ec73236b734c26d314f35c906e4f8934f157f6cdde0151486c3195346344de689ad7de5345e427010fa54c4350eaae1aa90545f13d83e9404229a21049 ]",
    "OP_SIGNED [ OP_DELEGATE_POWER [ OP_UINT64 1 ] OP_SIGNATURE 0x1b126a9964e8cc1ba77968a4b0b4ca29f602e3a2634db9cccee7a68a0158996a98238c3c011bd329c76d8809defd563938b07f0160d5c7cf94a8bdd707e8ed1017 ]",
    "OP_SIGNED [ OP_DELEGATE_POWER [ OP_UINT64 1 ] OP_SIGNATURE 0x1c308dd202beb5b622c48e49f80f00ba171e197a7d22dfcb73e6551040b0e46b515930545291188f5c4018400460fb49a297d19de6e7322ac65ceb6df9880eaf4d ]"
  ],
  "signatures": [
    "GxYW46UgDfCZIBJV63jY7Zv2K4GK7mh41HXPpDNA42qhIwap9zIcSJVmFDaqKOc/mccsyfSlUo1R4NSp+1ec3rA=",
    "HAbwI1cJUO0alivKbu03sH/oGmG5iGXiR3wzGEnbGao4OZDiO4QE9j9g/RTP98c5CNnj0dGRmZqWVZXZplpPEx0=",
    "HPxse4KQQr9uqN7ULjRmgwNZAs9ligYH2VkJLxwNmfXfG0+qULhFQ1fQzzpGE7t5ov7xKOtLo10rhKPGHSeVOXk="
  ]
}
```