package main

import (
	"io"
	"os/exec"
	"fmt"
	"bytes"
	"bufio"
)

func main() {
	runCmd()
	fmt.Println()
	runCmdWithPipe()
}

func runCmd() {
	useBufferedIO := true
	fmt.Println("Run command `echo -n \"My first command comes from golang\"`")
	cmd0 := exec.Command("echo", "-n", "My first command comes from golang")
	stdout0, err :=  cmd0.StdoutPipe()
	if err != nil {
		fmt.Printf("command error\n")
		return
	}
	if err := cmd0.Start(); err != nil {
		fmt.Printf("command error\n")
		return
	}

	if !useBufferedIO {
		var outputBuffer0 bytes.Buffer
		for {
			tempOutput := make([]byte, 5)
			n, err := stdout0.Read(tempOutput)
			if err != nil {
				if err == io.EOF {
					break
				} else {
					fmt.Printf("command error\n")
					return
				}
			}
			if n > 0 {
				outputBuffer0.Write(tempOutput[:n])
			}
		}
		fmt.Printf("%s\n", outputBuffer0.String())
	} else {
		outputBuffer0 := bufio.NewReader(stdout0)
		output0, _, err := outputBuffer0.ReadLine()
		if err != nil {
			fmt.Printf("command error\n")
			return
		}
		fmt.Printf("%s\n", string(output0))
	}
}

func runCmdWithPipe() {
	fmt.Println("Run command `ps aux | grep apipe`: ")
	cmd1 := exec.Command("ps", "aux")
	cmd2 := exec.Command("grep", "apipe")
	var outputBuf1 bytes.Buffer
	cmd1.Stdout = &outputBuf1
	if err := cmd1.Start(); err != nil {
		fmt.Printf("command error\n")
		return
	}
	if err := cmd1.Wait(); err != nil {
		fmt.Printf("command error\n")
		return
	}
	cmd2.Stdin = &outputBuf1
	var outputBuf2 bytes.Buffer
	cmd2.Stdout = &outputBuf2
	if err := cmd2.Start(); err != nil {
		fmt.Printf("command error\n")
		return
	}
	if err := cmd2.Wait(); err != nil {
		fmt.Printf("command error\n")
		return
	}
	fmt.Printf("%s\n", outputBuf2.Bytes())
}