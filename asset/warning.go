package asset

import (
	"fmt"
)

func Warn() {

	// tool title
	title :=
		`
	▓█████   ▄████   ▄████  ██░ ██  █    ██  ███▄    █ ▄▄▄█████▓
	▓█   ▀  ██▒ ▀█▒ ██▒ ▀█▒▓██░ ██▒ ██  ▓██▒ ██ ▀█   █ ▓  ██▒ ▓▒
	▒███   ▒██░▄▄▄░▒██░▄▄▄░▒██▀▀██░▓██  ▒██░▓██  ▀█ ██▒▒ ▓██░ ▒░
	▒▓█  ▄ ░▓█  ██▓░▓█  ██▓░▓█ ░██ ▓▓█  ░██░▓██▒  ▐▌██▒░ ▓██▓ ░ 
	░▒████▒░▒▓███▀▒░▒▓███▀▒░▓█▒░██▓▒▒█████▓ ▒██░   ▓██░  ▒██▒ ░ 
	░░ ▒░ ░ ░▒   ▒  ░▒   ▒  ▒ ░░▒░▒░▒▓▒ ▒ ▒ ░ ▒░   ▒ ▒   ▒ ░░   
	░ ░  ░  ░   ░   ░   ░  ▒ ░▒░ ░░░▒░ ░ ░ ░ ░░   ░ ▒░    ░    
	░   ░ ░   ░ ░ ░   ░  ░  ░░ ░ ░░░ ░ ░    ░   ░ ░   ░      
	░  ░      ░       ░  ░  ░  ░   ░              ░          
    `
	// User's Warning
	// ASCII WARNING taken from
	// https://patorjk.com/software/taag/#p=display&f=Big%20Money-sw&t=WARNING
	warning1 := `
 __       __   ______   _______   __    __  ______  __    __   ______  
/  |  _  /  | /      \ /       \ /  \  /  |/      |/  \  /  | /      \ 
$$ | / \ $$ |/$$$$$$  |$$$$$$$  |$$  \ $$ |$$$$$$/ $$  \ $$ |/$$$$$$  |
$$ |/$  \$$ |$$ |__$$ |$$ |__$$ |$$$  \$$ |  $$ |  $$$  \$$ |$$ | _$$/ 
$$ /$$$  $$ |$$    $$ |$$    $$ |$$$$  $$ |  $$ |  $$$$  $$ |$$ |/   |
$$ $$/$$ $$ |$$$$$$$$ |$$$$$$$  |$$ $$ $$ |  $$ |  $$ $$ $$ |$$ |$$$$ |
$$$$/  $$$$ |$$ |  $$ |$$ |  $$ |$$ |$$$$ | _$$ |_ $$ |$$$$ |$$ \__$$ |
$$$/    $$$ |$$ |  $$ |$$ |  $$ |$$ | $$$ |/ $$   |$$ | $$$ |$$    $$/ 
$$/      $$/ $$/   $$/ $$/   $$/ $$/   $$/ $$$$$$/ $$/   $$/  $$$$$$/  

`
	warning2 :=
		`
		(Unless you have permission)
		Use this tool only on your own systems, and no one else's.
		I created this tool as a fun way to harden my own network.
		This tool is not meant to be used to break laws.
		You have been warned! 
		`

	// tool starts
	fmt.Println(Red(warning1))
	fmt.Println(Yellow(warning2))
	fmt.Println(title)
}
