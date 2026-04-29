package admin

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"time"

	"github.com/privacyfish/pfish-signup/internal/store"
)

func Run(args []string, dataDir string, usedStoreSize int64, out io.Writer) error {
	if len(args) == 0 {
		return errors.New("usage: pfish-signup admin <list-pending|export|consume|expire-old>")
	}
	pending := store.NewPending(filepath.Join(dataDir, "pending"), 0, nil)
	used := store.NewUsed(filepath.Join(dataDir, "used.bin"), usedStoreSize)
	switch args[0] {
	case "list-pending":
		items, err := pending.List()
		if err != nil {
			return err
		}
		now := time.Now()
		for _, item := range items {
			age := now.Sub(item.Record.CreatedAt).Round(time.Second)
			fmt.Fprintf(out, "%s\t%s\texpires=%s\tusername_hmac=%s\n",
				item.ID,
				age,
				item.Record.ExpiresAt.UTC().Format(time.RFC3339),
				item.Record.UsernameHMAC,
			)
		}
		return nil
	case "export":
		if len(args) != 2 {
			return errors.New("usage: pfish-signup admin export <id>")
		}
		record, err := pending.Read(args[1])
		if err != nil {
			return err
		}
		fmt.Fprintln(out, record.CiphertextAgeB64)
		return nil
	case "consume":
		if len(args) != 3 {
			return errors.New("usage: pfish-signup admin consume <id> <username-hmac-hex>")
		}
		if _, err := pending.Read(args[1]); err != nil {
			return err
		}
		if h, err := hex.DecodeString(args[2]); err != nil || len(h) != store.HMACSize {
			return errors.New("username-hmac-hex must be 64 hex chars")
		}
		slot, err := used.Consume(args[2])
		if err != nil {
			return err
		}
		if err := pending.Delete(args[1]); err != nil {
			return err
		}
		fmt.Fprintf(out, "consumed id=%s username_hmac=%s slot=%d\n", args[1], args[2], slot)
		return nil
	case "expire-old":
		n, err := pending.Sweep(time.Now())
		if err != nil {
			return err
		}
		fmt.Fprintf(out, "expired=%d\n", n)
		return nil
	case "sweep-credentials":
		fmt.Fprintln(out, "credential store is in-memory inside the running service")
		return nil
	default:
		return fmt.Errorf("unknown admin command %q", args[0])
	}
}
