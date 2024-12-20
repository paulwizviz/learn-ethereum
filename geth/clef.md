# Creating an EOA using Clef

Please refer to [Clef official documentation](https://geth.ethereum.org/docs/tools/clef/introduction).

To create an EOA follow this steps:

1. `cd` to the root of this project.
1. Run the command `./scripts/gethclient.sh build` to build the image.
1. Run the command `./scripts/gethclient.sh shell` to opem a shell in the container.
1. In the shell run the following command:
```sh
clef newaccount --keystore ./data
```
1. You will find an account created in the directory `./tmp`

At the end of the step you will find a file created in the following format `UTC-<timestamp>-<account id in hash>`. For example: 
```
UTC--2024-03-24T17-54-26.876362756Z--0caf5082c30ffafdbd2941fa5ee64d2156d2a0ab
```

You can view the file in `./tmp`.

The file contains in a JSON format (representing the account state) looking like this:
```json
{
    "address":"3e056424aaf3e6cb819571eacfcdaad23e967fe1",
    "crypto": {
        "cipher":"aes-128-ctr",        "ciphertext":"04e097119722c127a931712878daedd8a41a2f9f311ecd8a7d5f43e808be07cc",
        "cipherparams":{
            "iv":"c85f8bfd8476b275a690c7042771a447"
        },
        "kdf":"scrypt",
        "kdfparams":{
            "dklen":32,
            "n":262144,
            "p":1,
            "r":8,
            "salt":"e4616683a81c751b0d2b577fb8beb7cd9960f7e6368888ad6624907a22b16278"
        },
        "mac":"3fecde41d0c4d04c1a23a7e29bdb2b7e8e433df8070f446677691f4a3570cdcd"
    },
    "id":"eb822fd0-0857-4043-ad08-fefd972fdb0a",
    "version":3
}
```