# boltz-poc
# WARNING: The applications themselves are not production ready. The apps and tools in this repo are solely for demo and experimenting.

# 0. Requirements
Go (golang) is required [installation instructions](https://go.dev/doc/install)

Recent version of docker is also needed.
## Boltz/RSK Docker images
 - Clone Boltz repo: https://github.com/rootstock/boltz-backend
 - checkout the `rsk_integration` branch
 - install the dependencies with `npm install` (e.g this will also pull and build the dependency `boltz-core` from **our git repository**)
 - Copy contracts artifacts and hardhat config, which are different from the original boltz-core. We have a mintable (dummy) DOC contract for testing and the `EtherSwap` contract has been modified to add the minting option when claiming locking RBTC. 

```
mkdir -p contracts && cp -R ./hardhat.config.ts node_modules/boltz-core/scripts node_modules/boltz-core/artifacts contracts
```
 - Build Boltz image -> Run: `npm run docker:build`
 - Build RSK node image -> Run: `npm run docker:rskj:build`

Note: On M1 (Apple silicon) mac computers, there may be some npm installation errors (e.g. for grpc-tools). Since Rosetta is now included in Mac OS, it helps to run docker in "x86 mode". This can be done via setting environment variable to x86 linux (e.g. add to ~/.zprofile)

```
# optional, for M1 macs
export DOCKER_DEFAULT_PLATFORM=linux/amd64     
```
# 1. Setup environment

Return to directory of this repo i.e. `boltz-poc`

- Run: `npm run docker:volumes:reset` to clear any data from previous runs
- Run: `npm run docker:start`

If there are any errors of the form `[lncli] rpc error: code = Unknown desc = server is still in the process of starting`, then wait for a bit and run `npm run docker:connect` (or just `npm run docker:start` again)

We should see something like

```
Mine a BTC block.
{
  "address": "bcrt1qt5hu3sukzvggmnx0ap6r5mjvsghqryp8w63kqg",
  "blocks": [
    "32b362fc4d1b76a4112e0b919ab23ffcc3f7f8e165506c71509c92d57a896de1"
  ]
}
Connect LN-BTC nodes.
Connect LN-BTC alice->bob.
{

}
Connect LN-BTC alice->carol.
{

}
Connect LN-BTC bob->carol.
{

}
```

Repeat `npm run docker:connect` until all 3 connections are live.

# 2. Start server
This is the backend server. This handles payment requests from the user interfact (called `client` below). The server interacts with the Boltz service and an RSKj ndoe. It is implemented in go.

- Run: `npm run server`

# 3. Start client
This is a `next.js` app used to run the browswer front end for the payments service. There is a Readme.md in the `client` directory as well.

- Run `npm install` in the client directory for dependendencies

- Run: `npm run client`

- Go to: `http://localhost:3001` in the browser (check the port, the server should indicate the port it is running on).

Note: Since payments are made using the lightning network, this application should be accessed from a browser that has the alby lightning wallet extension installed.


# 4. Wallet setup

Install alby, a browser based lightning wallet. https://getalby.com

We run 3 lightning nodes (LND) in docker containers (each connected to the same bitcoin node). Use the alby extension to add three accounts: Alice, Bob, Carol. 

Bob is an intermediary lighting node in between Alice and Carol, which are the two main personas. Alice serves as the wallet for the payments service itself. We interact with the UI using Carol's wallet as the end user. The merchant receives paymens on the RSK side, so we do not need a lightning wallet for them.

- Go to http://localhost:8082|8083|8084
- Advance -> proceed to localhost (unsafe)
- Click on  'Not Secure' -> Site Settings
- Change 'Insecure content' to 'Allow'
- Taken from: [here](https://github.com/getAlby/lightning-browser-extension/issues/252#issuecomment-934632283)

NOTE: Once testing is complete, these security settings should be reset to the default ones. 

## Alice
- Url: https://localhost:8082
- Macaroon: 

```
0201036c6e6402f801030a10cb58f4d8fd84f5d5732d8bfec93a49421201301a160a0761646472657373120472656164120577726974651a130a04696e666f120472656164120577726974651a170a08696e766f69636573120472656164120577726974651a210a086d616361726f6f6e120867656e6572617465120472656164120577726974651a160a076d657373616765120472656164120577726974651a170a086f6666636861696e120472656164120577726974651a160a076f6e636861696e120472656164120577726974651a140a057065657273120472656164120577726974651a180a067369676e6572120867656e657261746512047265616400000620382c08951c5e0a768064d6c4c80fa4074b3826cd119a90743e3229836c687409
```

## Bob
- Url: https://localhost:8083
- Macaroon: 

```
0201036c6e6402f801030a10c4c8f39813c6eec7b4695b001840cc871201301a160a0761646472657373120472656164120577726974651a130a04696e666f120472656164120577726974651a170a08696e766f69636573120472656164120577726974651a210a086d616361726f6f6e120867656e6572617465120472656164120577726974651a160a076d657373616765120472656164120577726974651a170a086f6666636861696e120472656164120577726974651a160a076f6e636861696e120472656164120577726974651a140a057065657273120472656164120577726974651a180a067369676e6572120867656e657261746512047265616400000620927effb7c6689429c46999ffde40bc7d3640471f30d6de74525e5fbe00ffccfa
```

## Carol
- Url: https://localhost:8084
- Macaroon: 

```
0201036c6e6402f801030a1082cdcd0bc73577f45607a93d6a4d794c1201301a160a0761646472657373120472656164120577726974651a130a04696e666f120472656164120577726974651a170a08696e766f69636573120472656164120577726974651a210a086d616361726f6f6e120867656e6572617465120472656164120577726974651a160a076d657373616765120472656164120577726974651a170a086f6666636861696e120472656164120577726974651a160a076f6e636861696e120472656164120577726974651a140a057065657273120472656164120577726974651a180a067369676e6572120867656e657261746512047265616400000620382286a8f35a8aa21c1e5edf6a4cfa0f5abbd9a41e9d3c5f21817b8eb6f32b7b
```

# 5. Running the app

Simply open the PoC app in the browser using the URL and port from the client section above, and then make payments using alby wallet.

Golang backend server output (and also Boltz) will display the logs, the steps and status for each payment.
# 6. Stop environment
- Run: `npm run docker:stop`

It is also good to clear things up before the next run: `npm run docker:volumes:reset`

# X. Miscellaneous

### Generate contract.go 

If the smart contracts (in `boltz-core` repository) are modified, then their abis will change. To interact with the revised smart contracts using go, we need install `protoc`, a solidity compiler, and a go-ethereum tool `abigen` which will generate the go file. This is only needed if a contract's abi has changed, in which case we need to provide the modified `json` file with the new abi.

- install protoc ([here](https://grpc.io/docs/protoc-installation/))
- install solc ([here](https://docs.soliditylang.org/en/v0.8.9/installing-solidity.html) and [here](https://www.educative.io/answers/how-to-install-solidity-in-mac))
- Clone Geth / make / make devtools ([doc](https://goethereumbook.org/smart-contract-compile/))

If we make changes to the `EtherSwap.sol` contract in boltz-core, then generate the new golang file from the new json file.

```
abigen --abi=server/abi/EtherSwap.json --out=server/abi/EtherSwap.go --pkg=abi
```

and similarly, for the `ERC20Swal.sol` or any other smart contract that the server interacts with directly.  

### Fix common error: https://github.com/lightningnetwork/lnd/issues/1177

* Generate btc blocks

`docker exec -it polar-n1-backend1 bitcoin-cli -chain=regtest -rpcuser=polaruser -rpcpassword=polarpass -generate 1`

### Connect to docker images
* Connect to bitcoin node

`docker exec -it --user lnd polar-n1-backend1 /bin/bash`

* Connect to a LN node

`docker exec -it --user lnd polar-n1-alice /bin/bash`

### Connect btw LN nodes
* `docker exec -it --user lnd polar-n1-alice lncli --network=regtest connect 02045b28b45f0b8efdac9381287b728d7c6897aa8c4d26fa5e9570078dc949e11b@host.docker.internal:9739`
* `docker exec -it --user lnd polar-n1-alice lncli --network=regtest connect 03b45c2032206051af4e130a8901575107bc1441287700ec5a8b6421d70a863a24@host.docker.internal:9740`
* `docker exec -it --user lnd polar-n1-bob lncli --network=regtest connect 03b45c2032206051af4e130a8901575107bc1441287700ec5a8b6421d70a863a24@host.docker.internal:9740`

### LN nodes URLs:

* alice: `02e7dd429d9148b6fde3e026d39ac63d2d76a3d3a5469ad8a64e0310fc6f62f302@host.docker.internal:9735`
* bob: `02045b28b45f0b8efdac9381287b728d7c6897aa8c4d26fa5e9570078dc949e11b@host.docker.internal:9739`
* carol: `03b45c2032206051af4e130a8901575107bc1441287700ec5a8b6421d70a863a24@host.docker.internal:9740`
