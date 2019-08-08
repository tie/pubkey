# Pubkey

Extract public key from SSH private key.

---

Let’s suppose that you’ve accidentally `rm`’ed your `~/.ssh/id_ed25519.pub` file.
Don‘t panic, ssh would still work because private key file contains public key.
And if you need to recover `.pub` file, you’ve found the right tool.

## Usage

All flags are optional. The input file `-f` flag defaults to `$HOME/.ssh/id_rsa`, or `/etc/ssh/ssh_host_rsa_key`
if `$HOME` is not set. Empty input file path implies standard input. The public key is written to standard output
if output file flag `-o` is not specified. Chmod-style permissions for output file may be specified using `-c` flag.

- Write the public key of `id_ed25519` identity to the `id_ed25519.pub` file.

  ```
  pubkey -f ~/.ssh/id_ed25519 -o ~/.ssh/id_ed25519.pub -c u=r,g=r
  ```

- You can also pipe your private key through `cat` without exposing file descriptors to `pubkey`.

  ```
  < ~/.ssh/id_ed25519 cat | pubkey -f= | cat > ~/.ssh/id_ed25519.pub
  ```
