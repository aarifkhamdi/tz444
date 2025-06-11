Auto:

1. make build
2. make docker-build
3. cd docker
4. docker-compose up -d

Interactive mode:

1. make run-server
2. make run-client

possible commands: https://github.com/aarifkhamdi/tz444/blob/ece6b2bd5555bc0bc61829111df70413d5483398/internal/client/cli/cli.go#L76

When choosing a protection mechanism, you should consider: the possibility of using acis, gpu and other specialized devices. It also should not be easy to parallelize.
I choosed argon2 since it's a good balance between security and performance. It uses memory-intensive approach to resist pow attacks.

In case you want to change the POW algorithm, you simply implement new https://github.com/aarifkhamdi/tz444/blob/e587e0d854af95a25a2febb65ba0e15dd5a8d137/internal/server/handler/pow.go#L9 and https://github.com/aarifkhamdi/tz444/blob/e587e0d854af95a25a2febb65ba0e15dd5a8d137/internal/server/handler/pow.go#L13
