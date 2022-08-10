package cmd

import (
	"fmt"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"github.com/spf13/cobra"
)

// Keys:
// + password
// + key (name, path, binary, or text)
// + key with passphrase

type KeyGenerate struct {
	cfg   Config
	name  string
	email string
	ktype string
	bits  int
}

func (g KeyGenerate) Command() *cobra.Command {
	c := &cobra.Command{
		Use:     "generate",
		Aliases: []string{"create", "g"},
		Short:   "Generate a new key",
		RunE: func(cmd *cobra.Command, args []string) error {
			return g.run()
		},
	}
	c.Flags().StringVar(&g.name, "name", "", "password to use")
	c.Flags().StringVar(&g.email, "email", "", "password to use")
	c.Flags().StringVar(&g.ktype, "type", "rsa", "password to use")
	c.Flags().IntVar(&g.bits, "bits", 4096, "password to use")
	return c
}

func (g KeyGenerate) run() error {
	key, err := crypto.GenerateKey(g.name, g.email, g.ktype, g.bits)
	if err != nil {
		return fmt.Errorf("cannot generate key: %v", err)
	}
	b, err := key.Serialize()
	if err != nil {
		return fmt.Errorf("cannot serialize key: %v", err)
	}
	_, err = g.cfg.Write(b)
	return err
}
