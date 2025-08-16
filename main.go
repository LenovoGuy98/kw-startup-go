package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/user"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Kindworks Startup")

	// Kindworks Image
	img := canvas.NewImageFromFile("kindworks.png")
	img.FillMode = canvas.ImageFillContain
	img.SetMinSize(fyne.NewSize(200, 200))

	// Link to Kindworks website
	link := widget.NewHyperlink("https://dokindworks.org", parseURL("https://dokindworks.org"))
	println("Welcome to your new Kindworks laptop")

	// Update Software
	//	updateOutputLabel := widget.NewLabel("")
	//	updateButton := widget.NewButton("Update Software", func() {
	//		cmd := exec.Command("sudo", "apt", "update")
	//		output, err := cmd.CombinedOutput()
	//		if err != nil {
	//			updateOutputLabel.SetText(fmt.Sprintf("Error updating: %s\n%s", err, string(output)))
	//			return
	//		}
	//		cmd = exec.Command("sudo", "apt", "upgrade", "-y")
	//		output, err = cmd.CombinedOutput()
	//		if err != nil {
	//			updateOutputLabel.SetText(fmt.Sprintf("Error upgrading: %s\n%s", err, string(output)))
	//			return
	//		}
	//		updateOutputLabel.SetText(string(output))
	//	})

	// Learn about host

	learnButton2 := widget.NewButton("Learn about the Host ", func() {
		learnWindow2 := myApp.NewWindow("Learn the Host")
		learnText2 := widget.NewLabel(getSystemInfo())
		learnWindow2.SetContent(container.NewVBox(learnText2))
		learnWindow2.Resize(fyne.NewSize(400, 100))
		learnWindow2.Show()
	})

	// Learn ZorinOS
	learnButton := widget.NewButton("Learn about your Laptop", func() {
		learnWindow := myApp.NewWindow("Learn the OS")
		learnText := widget.NewLabel("ZorinOS is Zorina user-friendly Linux distribution based on Ubuntu.\n\nHere are some basics:\n- The panel at the bottom is similar to Windows and macOS.\n- The 'Start' menu gives you access to all your applications.\n- You can install new software from the 'Software' application.")
		learnWindow.SetContent(container.NewVBox(learnText))
		learnWindow.Resize(fyne.NewSize(400, 200))
		learnWindow.Show()
	})

	// Layout
	content := container.NewVBox(
		img,
		link,
		widget.NewLabel("Welcome to your Kindworks Startup screen"),
		widget.NewSeparator(),
		learnButton,
		learnButton2,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(400, 400))
	myWindow.ShowAndRun()
}

func parseURL(urlStr string) *url.URL {
	link, err := url.Parse(urlStr)
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}
	return link
}

func getSystemInfo() string {
	currentUser, err := user.Current()
	if err != nil {
		return "Could not get user info"
	}
	hostname, err := os.Hostname()
	if err != nil {
		return "Could not get hostname"
	}

	// Get memory info
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal(err)
	}

	// Get disk info
	diskInfo, err := disk.Usage("/")
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("OS: %s \nArchitecture: %s \nUsername: %s\nHostname: %s\nMemory: %d MB \nDiskInfo: %d GB", runtime.GOOS, runtime.GOARCH, currentUser.Username, hostname, memInfo.Total/1024/1024, diskInfo.Total/1024/1024)
}
