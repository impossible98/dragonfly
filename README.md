# Dragonfly

## Docker

[![docker](https://github.com/impossible98/dragonfly/actions/workflows/docker.yml/badge.svg?branch=master)](https://github.com/impossible98/dragonfly/actions/workflows/docker.yml)

```bash
docker run --detach \
    --name dragonfly \
    --publish 8080:8080 \
    --restart always \
    impossible98/dragonfly
```
