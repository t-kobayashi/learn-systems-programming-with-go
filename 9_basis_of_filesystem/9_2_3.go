package main
import "os"

func main() {
	// create directory
	os.Mkdir("setting", 0755)
	// create deep directory
	os.MkdirAll("setting/myapp/networksettings", 0750)
}
