package main

import (
    "os"
    "fmt"
    "github.com/urfave/cli"
)

func AddNetwork(c *cli.Context) error {
    netString := c.Args().Get(0)
    net, err := CreateNet(netString)
    if err != nil {
        fmt.Println(err)
        return nil
    }
    fmt.Fprintf(os.Stderr, "Net %s created.\n", net.Id)

    return nil
}

func RemoveNetwork(c *cli.Context) error {
    fmt.Println("removed task network: ", c.Args().First())
    return nil
}

func ListNetwork(c *cli.Context) error {
    fmt.Println("list network: ")
    GetAllNetworks()
    return nil
}

func AssignIP(c *cli.Context) error {
    fmt.Println("Assign IP address ")
    ret, err := GetIP()
    if err != nil {
        fmt.Fprintf(os.Stderr, "assign IP error: %s\n", err)
    } else {
        fmt.Fprintf(os.Stderr, "%s\n", ret)
    }
    return nil
}

func ReleaseIP(c *cli.Context) error {
    fmt.Fprintf(os.Stderr, "Release IP address: %s ", c.Args().Get(0))
    ipaddr := c.Args().Get(0)
    ret, err := DeleteIP(ipaddr)
    if err != nil {
        fmt.Fprintf(os.Stderr, "delete IP error: %s\n", err)
    } else {
        fmt.Fprintf(os.Stderr, "%s\n", ret)
    }

    return nil
}

func main() {
  app := cli.NewApp()

  app.Commands = []cli.Command{
    {
      Name:        "network",
      Aliases:     []string{"t"},
      Usage:       "options for task networks",
      Subcommands: []cli.Command{
        {
          Name:  "add",
          Usage: "add a new network. For example : add 192.168.1.1/24",
          Action: AddNetwork,
        },
        {
          Name:  "remove",
          Usage: "remove an existing network",
          Action: RemoveNetwork,
        },
        {
          Name:  "list",
          Usage: "list all network.",
          Action: ListNetwork,
        },
      },
    },
    {
      Name:        "ipaddr",
      Aliases:     []string{"t"},
      Usage:       "options for task networks",
      Subcommands: []cli.Command{
        {
          Name:  "get",
          Usage: "get a free IP address.",
          Action: AssignIP,
        },
        {
          Name:  "release",
          Usage: "release a IP address",
          Action: ReleaseIP,
        },
      },
    },
  }

  app.Run(os.Args)
}
