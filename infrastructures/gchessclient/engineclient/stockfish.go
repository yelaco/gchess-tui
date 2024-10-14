package engineclient

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

type StockfishClient struct {
	execPath string
	in       io.WriteCloser
	out      io.ReadCloser
}

func NewStockfishClient(execPath string) EngineClient {
	return &StockfishClient{
		execPath: execPath,
	}
}

func (c *StockfishClient) Start() error {
	var err error
	cmd := exec.Command(c.execPath)
	if c.in, err = cmd.StdinPipe(); err != nil {
		return err
	}
	if c.out, err = cmd.StdoutPipe(); err != nil {
		return err
	}

	return cmd.Start()
}

func (c *StockfishClient) IsReady() bool {
	_, err := fmt.Fprintln(c.in, "isready")
	if err != nil {
		return false
	}
	scanner := bufio.NewScanner(c.out)
	scanner.Scan()
	return strings.Contains(scanner.Text(), "readyok")
}

func (c *StockfishClient) SetupEngine() error {
	return nil
}

func (c *StockfishClient) GetBestMove(fen string) (string, error) {
	_, err := fmt.Fprintf(c.in, "posiiton fen %s\ngo movetime 200\n", fen)
	if err != nil {
		return "", err
	}
	scanner := bufio.NewScanner(c.out)
	scanner.Scan()

	move := scanner.Text()
	if !strings.HasPrefix(move, "bestmove") {
		return "", fmt.Errorf("unexpected response from engine: %s", move)
	}
	return strings.Split(move, " ")[1], nil
}

func (c *StockfishClient) Evaluate(fen string) (int, error) {
	_, err := fmt.Fprintf(c.in, "posiiton fen %s\neval\n", fen)
	if err != nil {
		return 0, err
	}
	scanner := bufio.NewScanner(c.out)
	scanner.Scan()

	return 0, nil
}
