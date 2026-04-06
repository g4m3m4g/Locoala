package hosts

import (
	"fmt"
	"os"
	"strings"
)

const hostsPath = `C:\Windows\System32\drivers\etc\hosts`

func Add(domain string) error {
	entry := "127.0.0.1 " + domain

	content, err := os.ReadFile(hostsPath)
	if err != nil {
		return err
	}
	fmt.Println(string(content))

	if strings.Contains(string(content), entry) {
		fmt.Println("Hosts entry already exists")
		return nil
	}

	f, err := os.OpenFile(hostsPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("need admin privileges to modify hosts file, run cmd as administrator and try again \nOR manually add to hosts file")
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

	content, err := os.ReadFile(hostsPath)
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

	err = os.WriteFile(hostsPath, []byte(strings.Join(newLines, "\n")), 0644)
	if err != nil {
		return fmt.Errorf("need admin privileges to modify hosts file, run cmd as administrator and try again \nOR manually remove from hosts file")
	}

	fmt.Println("Removed from hosts:", entry)
	return nil
}