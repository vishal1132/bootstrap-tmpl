use clap::Parser;
use cli::GoGen;
use std::fs;
use tera::{Context, Tera};
use xshell::{cmd, Shell};

mod cli;

fn main() -> color_eyre::Result<()> {
    let cli = GoGen::parse();

    let templates: Tera = Tera::new(format!("{}/**/*", cli.source).as_str())?;

    let mut context = Context::new();
    cli.args.split(' ').for_each(|k| {
        let mut kv = k.split('=');
        let (key, val) = (kv.next().unwrap_or("test"), kv.next().unwrap_or("test"));
        context.insert(key, val);
    });

    let cwd = std::env::current_dir()?;
    let cwd = cwd.to_str().unwrap();
    templates.templates.iter().for_each(|(key, val)| {
        let full_path = val.path.clone().unwrap();
        let file_name = full_path.split('/').last().unwrap();

        let relative_path = full_path
            .as_str()
            .strip_suffix(file_name) // path without filename
            .unwrap()
            .strip_prefix(format!("{}/{}", cwd, cli.source).as_str())
            .unwrap() // path without cwd and source.
            .strip_prefix('/')
            .unwrap_or("")
            .strip_suffix('/')
            .unwrap_or("");

        let dir_path = format!("{}/{}/{}", cwd, cli.dest.clone(), relative_path);
        fs::create_dir_all(dir_path.clone()).expect("can't create directory");

        let full_path_file = format!("{}/{}", dir_path, file_name);
        fs::File::create(full_path_file.clone()).expect("can't create file");

        let content = templates.render(key, &context).unwrap();
        if content.trim().is_empty() {
            fs::remove_file(full_path_file).expect("can't delete file");
        } else {
            fs::write(full_path_file, content).expect("can't write content");
        }
    });

    std::env::set_current_dir(format!("{}/{}", cwd, cli.dest))?;

    let sh = Shell::new()?;
    cmd!(sh, "go fmt ./...").run()?;
    cmd!(sh, "go mod tidy").run()?;
    Ok(())
}
