## bootstrap_go

Use [tera templates](https://docs.rs/tera/latest/tera/) to bootstrap a service in a language(first class support for `go` so far) because of the commands ran after running through the templates.
* `go fmt ./...` for formatting the code later.
* `go mod tidy` for tidying the dependencies.

You can plugin your own templates to start bootstrapping with the `module name` etc. 

A default structure is shipped with this repository as well (in [go](./go/README.md)).

The help for this looks like->
```sh
(base) bootstrap  🍣 main 📝 ×4🗃️  ×41🦀 v1.70.0 🐏 12GiB/16GiB | 9GiB/10GiB on ☁️  (us-east-1) 
🔋98% 🕙 22:45:06 ❯ ./target/release/bootstrap -h                
Usage: bootstrap --args <ARGS> --source <SOURCE> --dest <DEST>

Options:
  -a, --args <ARGS>      
  -s, --source <SOURCE>  
  -d, --dest <DEST>      
  -h, --help             Print help
  -V, --version          Print version
```

A sample command to bootstrap using templates in `go/` to `abcd/efgh` looks like->
`./target/release/bootstrap -s go -d abcd/efgh -a 'postgres_enabled=true redis_enabled=true service_name=abcd module_name=github.com/vishal1132/bootstrap'`