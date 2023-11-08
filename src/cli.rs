use clap::Parser;

#[derive(Debug, Parser)]
#[command(name = "bootstrap", author, version, about, long_about)]
pub struct GoGen {
    #[clap(short, long)]
    pub(crate) args: String,

    #[clap(long, short)]
    pub(crate) source: String,

    #[clap(long, short)]
    pub(crate) dest: String,
}
