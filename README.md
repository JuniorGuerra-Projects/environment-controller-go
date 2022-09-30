# environment-controller-go

this is a controller of environment

the use actual of environment is some hard in the call
but if we can use the tags for this call we can use this

example of use

import "app/Junior/environmentgo"

type environment struct {
  Url string `env:"URL"`
  AccountNumber string `env:"ACCNO"`
}

var Environment environment

func init() {
  environmentgo.New(&Environment)
}

func main() {
  fmt.Printf("Environment: %+v", Environment)
  // salida -> Url: `URL en environment` || AccountNumber: `ACCNO en environment`
}

esta pequenia logica nos ahorra el clasico:
os.getEnv para cada variable que queramos tener.
