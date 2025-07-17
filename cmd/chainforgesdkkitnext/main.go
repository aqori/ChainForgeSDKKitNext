// cmd/chainforgesdkkitnext/main.go
package main

import (
"flag"
"log"
"os"

"chainforgesdkkitnext/internal/chainforgesdkkitnext"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := chainforgesdkkitnext.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
