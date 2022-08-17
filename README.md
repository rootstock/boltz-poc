# boltz-poc


docker exec -it --user lnd polar-n1-backend1 /bin/bash

bitcoin-cli -chain=regtest -rpcuser=polaruser -rpcpassword=polarpass -generate 6

docker exec -it polar-n1-backend1 bitcoin-cli -chain=regtest -rpcuser=polaruser -rpcpassword=polarpass -generate 6

https://github.com/lightningnetwork/lnd/issues/1177


docker exec -it --user lnd polar-n1-alice /bin/bash

docker exec -it --user lnd polar-n1-alice lncli --network=regtest connect 02045b28b45f0b8efdac9381287b728d7c6897aa8c4d26fa5e9570078dc949e11b@host.docker.internal:9739
docker exec -it --user lnd polar-n1-alice lncli --network=regtest connect 03b45c2032206051af4e130a8901575107bc1441287700ec5a8b6421d70a863a24@host.docker.internal:9740
docker exec -it --user lnd polar-n1-bob lncli --network=regtest connect 03b45c2032206051af4e130a8901575107bc1441287700ec5a8b6421d70a863a24@host.docker.internal:9740

lnd@alice:/$ lncli --network=regtest connect 03b45c2032206051af4e130a8901575107bc1441287700ec5a8b6421d70a863a24@host.docker.internal:9740

alice:
02e7dd429d9148b6fde3e026d39ac63d2d76a3d3a5469ad8a64e0310fc6f62f302@host.docker.internal:9735

bob: 
02045b28b45f0b8efdac9381287b728d7c6897aa8c4d26fa5e9570078dc949e11b@host.docker.internal:9739

carol:
03b45c2032206051af4e130a8901575107bc1441287700ec5a8b6421d70a863a24@host.docker.internal:9740
