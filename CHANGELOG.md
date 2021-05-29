# CHANGELOG

## Unreleased

- use prettier default config
- rm `READEME.md` from templates
- add `.prettierignore`

## 0.5.4

- add install.sh
- add `arm64` build in goreleaser

## 0.5.3

- update goreleaser filename

## 0.5.2

- switch to `github.com/swiftcarrot/gqlgen`

## 0.5.1

- fix `defer` in database template
- rm `cmd/pack`
- replace `packr` with go embed
- add server generator
- no more ignoring `DS_Store`
- switch to `github.com/swiftcarrot/gqlgen` internally

## 0.5.0

- remove Pagination & items in scaffold
- js `export default function()` one liner
- use go embed for project migration files, add database package

## 0.4.2

- fix gqlgen model plugin missing

## 0.4.1

- remove `g webpacker:install`
- move webpacker generator to packages generator
- add `packages/app` by default

## 0.4.0

- remove `g storybook:install`
- add `g packages`
- default storybook integration
- storybook v6
- add `storybook-preset` package
- fix html minify options in webpacker

## 0.3.0

- add `g webpacker:install` for custom webpack & babel configs

## 0.2.1

- improve error handling in handler
- generating a `GET /_health` endpoint
- add `dashi/server` package

## 0.2.0

- add `g storybook:install`
- add config package
- upgrade to apollo client v3
- improve pluralization in [flect](https://github.com/swiftcarrot/flect)
