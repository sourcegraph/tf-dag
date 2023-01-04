# tf-dag

Fork of [github.com/terraform/hashicorp/internal/dag](https://github.com/hashicorp/terraform/tree/main/internal/dag). Directed acyclic graph (DAG) implementation in Go.

## Why the fork?

The `dag` package from [terraform/hashicorp] is one of the most well maintained and tested DAG implementations in Go, and [many projects depend on it](https://sourcegraph.com/search?q=context:global+lang:Go+%22github.com/hashicorp/terraform/dag%22&patternType=standard&sm=1&groupBy=repo). However, it was [made internal](https://github.com/hashicorp/terraform/commit/70eebe3521d2f1ffcb5c12b75e90f1b82db94551). Additionally, upstream package contains a lot of extra dependencies where most projects do not need.

By forking it, we can import it as a standlone package and introduce custom changes to make it more generic.

### Changes

Below is a list of changes we made to the upstream package:

- [x] replace the use of [`tfdiags`](https://github.com/hashicorp/terraform/tree/main/internal/tfdiags) with `error`
- [ ] remove logging

## Usage

```sh
go get github.com/sourcegraph/tf-dag
```

[terraform/hashicorp]: https://github.com/hashicorp/terraform
