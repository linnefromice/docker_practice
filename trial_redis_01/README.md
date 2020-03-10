# Command

- use redis-cli
  - docker-compose up
  - docker container ls
  - docker exec -it (CONTAINER ID) /bin/bash
  - redis-cli
  - keys * -> (empty list or set)
  - set "sample_key" "bell" -> OK
  - get "sample_key" -> "bell"

# Reference

- [docker-composeでredis環境をつくる](https://qiita.com/uggds/items/5e4f8fee180d77c06ee1)
- [【Redis】Go言語で高速呼び出しKVS【Redigo】](https://qiita.com/chan-p/items/5c5e7cc1e966f8a90422)