//go:build windows

package hosts

import (
	"fmt"
	"os"
	"strings"
)

func Add(domain string) error {
	entry := "127.0.0.1 " + domain
	path := GetHostsPath()

	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if strings.Contains(string(content), entry) {
		fmt.Println("Hosts entry already exists")
		return nil
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("need administrator privileges")
	}
	defer f.Close()

	_, err = f.WriteString("\n" + entry + "\n")
	if err != nil {
		return err
	}

	fmt.Println("Added to hosts:", entry)
	return nil
}

func Remove(domain string) error {
	entry := "127.0.0.1 " + domain
	path := GetHostsPath()

	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")
	var newLines []string

	for _, line := range lines {
		if strings.TrimSpace(line) != entry {
			newLines = append(newLines, line)
		}
	}

	err = os.WriteFile(path, []byte(strings.Join(newLines, "\n")), 0644)
	if err != nil {
		return fmt.Errorf("need administrator privileges")
	}

	fmt.Println("Removed from hosts:", entry)
	return nil
}