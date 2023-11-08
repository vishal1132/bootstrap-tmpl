## Tenders Backend

Backend service format. Starts with `main.go` and then does->
* load config // panic if error in loading the config.
* setup global logger // panic if error in setting up logging.
* utils hosts the common utils functions, generic sorting, generic error handling etc.