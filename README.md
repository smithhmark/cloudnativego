# cloudnativego

This repo is me following through examples from the book "Cloud Native Go" by Matthew Titmus.

I treat the worked examples as suggestions and jumping off points to grow my understanding.

# Running examples
this project has a makefile, running `make help` should explain what targets are available. However, because I like keeping the various examples around, any given commit might have multiple commands available. therefore, please note:
- running `make build` will enumerage all of the ./cmd/<binary> that are available
- running `make run bin=<binary>` will execute the binary whose name (from `make build` above) is assiend to the envvar "bin"


