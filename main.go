package main

import (
	"fmt"
	"net/url"
	"os"
	"os/user"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
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
		// learnText2 := widget.NewLabel("ZorinOS is Zorina user-friendly Linux distribution based on Ubuntu.\n\nHere are some basics:\n- The panel at the bottom is similar to Windows and macOS.\n- The 'Start' menu gives you access to all your applications.\n- You can install new software from the 'Software' application.")
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
		//	widget.NewLabel("Create Account"),
		//	nameEntry,
		//		registerButton,
		//		accountLabel,
		//		widget.NewSeparator(),
		// widget.NewLabel("System Information"),
		//sysInfoLabel,
		widget.NewSeparator(),
		//	updateButton,
		//		updateOutputLabel,
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

	return fmt.Sprintf("OS: %s\nArchitecture: %s\nUsername: %s\nHostname: %s", runtime.GOOS, runtime.GOARCH, currentUser.Username, hostname)
}
