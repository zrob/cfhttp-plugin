# cfhttp-plugin
A CF cli plugin to provide CF context to HTTPie

## Installation
1. Install httpie, see [instructions here](https://github.com/jakubroztocil/httpie#2installation)
2. git clone the repo to your desktop
3. In the repo, run `go build` to compile a binary
4. run `cf install-plugin <path-to-binary>`

## Usage
You can now use httpie like syntax for your data payloads

e.g. 

Instead of 
`cf curl -X POST /v2/organizations -d '{"name": "test-org"}`

You can now run
`cf http POST /v2/organizations name=test-org`


