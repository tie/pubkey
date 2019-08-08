package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/crypto/ssh"

	"bitbucket.org/dchapes/mode"
)

var (
	input  = ""
	output = ""
	chmod  = "u=rw,g=r"
)

func main() {
	log.SetFlags(0)

	home, err := os.UserHomeDir()
	if err == nil {
		// Make default input path relative to the working directory.
		cwd, err := os.Getwd()
		if err == nil {
			relHome, err := filepath.Rel(cwd, home)
			if err == nil {
				home = relHome
			}
		}
	}
	if err != nil {
		// Fallback to sshd host key.
		input = filepath.FromSlash("/etc/ssh/ssh_host_rsa_key")
	} else {
		input = filepath.Join(home, ".ssh", "id_rsa")
	}

	flag.StringVar(&input, "f", input, "private key file path")
	flag.StringVar(&output, "o", output, "output file (default stdout)")
	flag.StringVar(&chmod, "c", chmod, "chmod output file")
	flag.Parse()

	fmode, err := mode.Parse(chmod)
	if err != nil {
		log.Fatalf("parse chmod: %v", err)
		return
	}
	perm := fmode.Apply(0)

	var pem []byte
	if input != "" {
		pem, err = ioutil.ReadFile(input)
	} else {
		pem, err = ioutil.ReadAll(os.Stdin)
		input = "stdin"
	}
	if err != nil {
		log.Fatalf("read %q file: %v", input, err)
		return
	}

	signer, err := ssh.ParsePrivateKey(pem)
	if err != nil {
		log.Fatalf("parse %q file: %v", input, err)
		return
	}

	pubkey := ssh.MarshalAuthorizedKey(
		signer.PublicKey(),
	)
	if output != "" {
		err = ioutil.WriteFile(output, pubkey, perm)
	} else {
		_, err = os.Stdout.Write(pubkey)
		output = "stdout"
	}
	if err != nil {
		log.Fatalf("write %q file: %v", output, err)
		return
	}
}
