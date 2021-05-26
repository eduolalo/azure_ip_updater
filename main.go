package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {

	ip := flag.String("ip", "", "IP que se actualizará")
	name := flag.String("rule", "", "Nombre de la regla que recibirá la ip")
	group := flag.String("g", "", "Grupo de recursos en el que está el servidor SQL")
	server := flag.String("s", "", "Servidor de SQL")

	flag.Parse()

	validator := regexp.MustCompile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	isValid := validator.MatchString(*ip)
	if !isValid {
		log.Println("La ip no es válida")
		return
	}

	var command strings.Builder
	command.WriteString("az sql server firewall-rule update -g ")
	command.WriteString(*group)
	command.WriteString(" -s ")
	command.WriteString(*server)
	command.WriteString(" -n ")
	command.WriteString(*name)
	command.WriteString(" --start-ip-address ")
	command.WriteString(*ip)
	command.WriteString(" --end-ip-address ")
	command.WriteString(*ip)

	log.Printf("Actualizando la regla \"%s\" del firewall de \"%s\"", *name, *server)

	cmd := exec.Command("zsh", "-c", command.String())
	cmd.Dir = os.Getenv("AZURE_CLI_PATH")

	out, err := cmd.Output()
	if err != nil {
		log.Printf("Error: %s \n", err)
	}

	log.Printf("Out: %s", out)

}
