package main

import (
	"os/exec"
	"strings"
)

const hledger = "hledger"

func listAccounts() ([]string, error) {
	out, err := exec.Command(hledger, "accounts", "--tree").Output()
	if err != nil {
		return nil, err
	}

	accounts := strings.Split(string(out), "\n")

	return accounts, nil
}

func register() ([][]string, error) {
	return [][]string{
		{"hello", "world", "this", "is", "fyne"},
		{"talking", "to", "world", "now", "hi"},
		{"goodbye", "to", "the", "world", "hey"},
		{"goodbye", "to", "the", "world", "hey"},
		{"goodbye", "to", "the", "world", "hey"},
	}, nil
}

func register1() ([][]string, error) {
	out, err := exec.Command(hledger, "register").Output()
	if err != nil {
		return nil, err
	}

	lineTransactions := strings.Split(string(out), "\n")
	lineTransactions = lineTransactions[:10] // TODO temporary truncation for testing

	transactions := make([][]string, 0, 10)
	for _, line := range lineTransactions {
		columns := strings.Split(line, " ")
		transactions = append(transactions, columns)
	}

	return transactions, nil
}

func registerLines() ([]string, error) {
	out, err := exec.Command(hledger, "register").Output()
	if err != nil {
		return nil, err
	}

	lineTransactions := strings.Split(string(out), "\n")
	lineTransactions = lineTransactions[:10] // TODO temporary truncation for testing

	return lineTransactions, nil
}
