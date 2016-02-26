package main

import (
    "os"
    "io"
    "fmt"
    "html/template"
    "strings"
    "github.com/codegangsta/cli"
)

func main() {

    var parsedStringSlice []string
    var output_fp io.Writer


    app := cli.NewApp()
    app.Name = "stencil"
    app.Version = "0.0.1"
    app.Usage = "Take a go template and variables from the commandline, create output"

    app.Flags = []cli.Flag{
        cli.StringFlag{
            Name: "template",
            Value: "",
            Usage: "Template file",
        },
        cli.StringFlag{
            Name: "output",
            Value: "",
            Usage: "Specify output file",
        },
        cli.StringSliceFlag{
            Name: "var",
            Value: &cli.StringSlice{},
            Usage: "Variables to be used in template",
        },
    }

    app.Action = func(c *cli.Context) {
        parsedStringSlice = c.StringSlice("var")
        input := c.String("template")
        output := c.String("output")


        if output != "" {
            var err error
            output_fp, err = os.OpenFile(output, os.O_CREATE|os.O_WRONLY, 0744)
            if err != nil {
                panic(err)
            }
        } else {
            output_fp = os.Stdout
        }

        tmpl, err := template.ParseFiles(input)
        if err != nil {
            fmt.Printf("Error reading template file: %s\n", err)
            os.Exit(1)
        }

        context := make(map[string]string)

        for _, v := range parsedStringSlice {
            arr := strings.Split(v, "=")
            context[arr[0]] = arr[1]
        }

        err = tmpl.Execute(output_fp, context)

        if err != nil {
            fmt.Println("Error!")
            fmt.Println(err)
        }
    }

    app.Run(os.Args)
}
