# sentinel-api

<p>
  <a
    href="https://github.com/dlbarduzzi/sentinel-api/actions/workflows/ci.yaml"
    target="_blank"
    rel="noopener">
    <img
      src="https://github.com/dlbarduzzi/sentinel-api/actions/workflows/ci.yaml/badge.svg"
      alt="ci"
    />
  </a>
</p>

A centralized control plane for managing and synchronizing Prometheus alerts across multiple clusters.

## Docker

1. Use the containerization tool of your choice (i.e. docker, podman)

2. Build the docker image:

```sh
docker build -t ghcr.io/dlbarduzzi/sentinel-api:__VERSION__ . -f Dockerfile
```

3. Running container in your local machine:

```sh
docker run --rm \
  --name sentinel-api \
  -p 8090:8090 \
  ghcr.io/dlbarduzzi/sentinel-api:__VERSION__
```

## Acknowledgements

This project is heavily inspired by the open-source project
[PocketBase](https://github.com/pocketbase/pocketbase).

Several design patterns and core features are adapted from that project,
with modifications and extensions to meet the goals of this application.

The primary purpose of this project is for personal exploration and experimentation.

## License

[MIT](./LICENSE)
