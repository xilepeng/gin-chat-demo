# im



安装 Mongo DB

docker在启动时会自产生三种网络机制：bridge、host、none。

``` s
docker network create -d bridge xn  

➜  im git:(main) ✗ docker network create -d bridge xn
b65eff06eee9b33356434aa9290757b8cae68424fd3890848d52d86c7793426d

➜  im git:(main) ✗ docker network ls
NETWORK ID     NAME      DRIVER    SCOPE
d6869059c462   bridge    bridge    local
b75882e5c4b8   host      host      local
f34cb6eeca00   none      null      local
b65eff06eee9   xn        bridge    local




docker run -d --network xn --name mongo \
-e MONGO_INITDB_ROOT_USERNAME=admin \
-e MONGO_INITDB_ROOT_PASSWORD=admin \
-p 27017:27017 \
mongo
```