package initialize

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/eris-ltd/eris-cli/Godeps/_workspace/src/github.com/eris-ltd/common/go/common"
)

func cloneRepo(name, location string) error {
	if _, err := os.Stat(location); !os.IsNotExist(err) {
		logger.Debugf("The location exists. Attempting to pull instead.\n")
		if err := pullRepo(location); err != nil {
			return err
		} else {
			return nil
		}
	}
	src := "https://github.com/eris-ltd/" + name
	c := exec.Command("git", "clone", src, location)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		return err
	}
	return nil
}

func pullRepo(location string) error {
	var input string
	logger.Printf("Looks like the %s directory exists.\nWould you like the marmots to pull in any recent changes? (Y/n): ", location)
	fmt.Scanln(&input)

	if input == "Y" || input == "y" || input == "YES" || input == "Yes" || input == "yes" {
		prevDir, _ := os.Getwd()
		if err := os.Chdir(location); err != nil {
			return fmt.Errorf("Error:\tCould not move into the directory (%s)\n", location)
		}
		c := exec.Command("git", "pull", "origin", "master")
		// c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		if err := c.Run(); err != nil {
			return err
		}
		if err := os.Chdir(prevDir); err != nil {
			return fmt.Errorf("Error:\tCould not move into the directory (%s)\n", location)
		}
	}
	return nil
}

func dropDefaults() error {
	if err := writeDefaultFile(common.ServicesPath, "keys.toml", DefaultKeys); err != nil {
		return fmt.Errorf("Cannot add keys: %s.\n", err)
	}
	if err := writeDefaultFile(common.ServicesPath, "ipfs.toml", DefaultIpfs); err != nil {
		return fmt.Errorf("Cannot add ipfs: %s.\n", err)
	}
	if err := writeDefaultFile(common.ServicesPath, "do_not_use.toml", DefaultIpfs2); err != nil {
		return fmt.Errorf("Cannot add ipfs: %s.\n", err)
	}
	if err := writeDefaultFile(common.ActionsPath, "do_not_use.toml", defAct); err != nil {
		return fmt.Errorf("Cannot add default action: %s.\n", err)
	}
	return nil
}

func dropChainDefaults() error {
	defChainDir := filepath.Join(common.BlockchainsPath, "config", "default")
	if err := writeDefaultFile(common.BlockchainsPath, "default.toml", DefChainService); err != nil {
		return fmt.Errorf("Cannot add default chain definition: %s.\n", err)
	}
	if err := writeDefaultFile(defChainDir, "config.toml", DefChainConfig); err != nil {
		return fmt.Errorf("Cannot add default config.toml: %s.\n", err)
	}
	if err := writeDefaultFile(defChainDir, "genesis.json", DefChainGen); err != nil {
		return fmt.Errorf("Cannot add default genesis.json: %s.\n", err)
	}
	if err := writeDefaultFile(defChainDir, "priv_validator.json", DefChainKeys); err != nil {
		return fmt.Errorf("Cannot add default priv_validator.json: %s.\n", err)
	}
	if err := writeDefaultFile(defChainDir, "server_conf.toml", DefChainServConfig); err != nil {
		return fmt.Errorf("Cannot add default server_conf.toml: %s.\n", err)
	}
	if err := writeDefaultFile(defChainDir, "genesis.csv", DefChainCSV); err != nil {
		return fmt.Errorf("Cannot add default genesis.csv: %s.\n", err)
	}
	return nil
}

func writeDefaultFile(savePath, fileName string, toWrite func() string) error {
	if err := os.MkdirAll(savePath, 0777); err != nil {
		return err
	}
	writer, err := os.Create(filepath.Join(savePath, fileName))
	defer writer.Close()
	if err != nil {
		return err
	}
	writer.Write([]byte(toWrite()))
	return nil
}
