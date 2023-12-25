# repbash

Execute bash by replace with [replace variables](https://github.com/puutaro/CommandClick/blob/master/md/developer/set_replace_variables.md)


```sh.sh
Usage: repbash [--src-tsv-path SRC-TSV-PATH] [--args-con ARGS-CON] [--save-repbash-line] [--import IMPORT] LAUNCHSHELLPATH

Positional arguments:
  LAUNCHSHELLPATH        launch shell path

Options:
  --src-tsv-path SRC-TSV-PATH, -t SRC-TSV-PATH
                         val name to val value tsv paths
  --args-con ARGS-CON, -a ARGS-CON
                         args contetns key to value map list (Replace ${REPBASH_ARGS_CON} with this in script)
  --save-repbash-line, -s
                         save flag for repbash cmd line (Ordinary, replace 'exec repbash ~' with '### REPBASH_CON')
  --import IMPORT, -i IMPORT
                         import shell file (include http url9
  --help, -h             display this help and exit
```
