# Azure ip updater

Este es un script para actualizar las ip en el firewall de los servidores de SQL en azure. Utiliza el azure-cli, por lo cual lo debes de tener ya [instalado](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli) y debes estar [logueado](https://docs.microsoft.com/en-us/cli/azure/authenticate-azure-cli) para usar este script.

## Objetivo

La idea de este script es reducir la longitud del comando de original de Azure por que me da pereza escribir todo:

```
// Comando de azure:
$ az sql server firewall-rule update -g miGrupo -s miServidor -n nombreDeLaRegla  --start-ip-address 127.0.0.1 --start-ip-address 127.0.0.1

// Comando del script:
azure_ip_updater -ip 127.0.0.1 -rule miRegla -g miGurpo -s Servidor 
```

## Requerimientos

* Tener instalado Go 1.16 en adelante.
* Tener [instalado azure-cli](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli)
* [Estar logueado en azure-cli](https://docs.microsoft.com/en-us/cli/azure/authenticate-azure-cli)
* Tener la vareable de entorno _AZURE_CLI_PATH_ en tu .zshrc .profile, .bashrc o el que uses, con la dirección del bin de la instalación de tu azure-cli:

```
$ echo "export AZURE_CLI_PATH=/opt/homebrew/Cellar/azure-cli/2.23.0/bin" >> .zshrc
```

## Instalación

Dentro del repo, darle el comando:

```
$ go install
```

Para revisar que se instaló bien, puedes ejecutar el comando:

```
$ ls $GOBIN | grep azure
```

Listo! te debe de responder con un _azure_ip_updater_

## Ayuda

Puedes ejecutar el comando ```$ azure_ip_updater -h```:
```
Usage of azure_ip_updater:
  -g string
    	Grupo de recursos en el que está el servidor SQL
  -ip string
    	IP que se actualizará
  -rule string
    	Nombre de la regla que recibirá la ip
  -s string
    	Servidor de SQL
```
