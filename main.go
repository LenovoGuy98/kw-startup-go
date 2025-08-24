package main

import (
	"log"
	"os/exec"
	"runtime"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	// Initialize GTK without command-line arguments.
	gtk.Init(nil)

	settings, err := gtk.SettingsGetDefault()
	if err != nil {
		log.Fatal("Unable to get settings:", err)
	}
	settings.SetProperty("gtk-theme-name", "Adwaita-dark")

	// Create a new top-level window.
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetProperty("gtk-theme-name", "Adwaita-dark")
	win.SetTitle("Kindworks Startup")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	win.SetDefaultSize(300, 200)
	win.SetPosition(gtk.WIN_POS_CENTER)

	// Create a vertical box container.
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 6)
	if err != nil {
		log.Fatal("Unable to create box:", err)
	}
	win.Add(box)

	// Add Kindworks image.
	img, err := gtk.ImageNewFromFile("kindworks.png")
	if err != nil {
		log.Println("Could not load image:", err)
	} else {
		box.PackStart(img, false, false, 0)
	}

	// Add a link button to dokindworks.org.
	linkButton, err := gtk.LinkButtonNewWithLabel("https://dokindworks.org", "Visit Kindworks")
	if err != nil {
		log.Fatal("Unable to create link button:", err)
	}
	box.PackStart(linkButton, false, false, 0)

	// Create a new label and handle the error.
	myLabel, err := gtk.LabelNew("Welcome to your new computer! ")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	myLabel.SetMarkup("<b>Welcome to your new computer </b>")

	box.PackStart(myLabel, false, false, 0)

	//NewLabel, err

	// ewLabel("Welcome to your Kindworks Startup screen"),

	// Add a button to show system information.
	sysInfoButton, err := gtk.ButtonNewWithLabel("Show System Information")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	sysInfoButton.Connect("clicked", showSystemInfo)
	box.PackStart(sysInfoButton, false, false, 0)

	// Add a button to open the PDF file.

	pdfButton, err := gtk.ButtonNewWithLabel("Open PDF")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}

	pdfButton.Connect("clicked", openPDF)
	box.PackStart(pdfButton, false, false, 0)

	// Create a horizontal box for application buttons.
	appBox, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 6)
	if err != nil {
		log.Fatal("Unable to create box:", err)
	}
	box.PackStart(appBox, false, false, 0)

	// Add a button for Firefox.

	firefoxButton, err := gtk.ButtonNewWithLabel("Firefox")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	firefoxButton.Connect("clicked", launchFirefox)
	appBox.PackStart(firefoxButton, true, true, 0)

	// Add a button for LibreOffice.

	libreOfficeButton, err := gtk.ButtonNewWithLabel("LibreOffice")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	libreOfficeButton.Connect("clicked", launchLibreOffice)
	appBox.PackStart(libreOfficeButton, true, true, 0)

	// Add a button for Zoom.

	zoomButton, err := gtk.ButtonNewWithLabel("Zoom")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	zoomButton.Connect("clicked", launchZoom)
	appBox.PackStart(zoomButton, true, true, 0)

	// Show the window and all its widgets.
	win.ShowAll()

	// Start the GTK main loop.
	gtk.Main()
}

func showSystemInfo() {
	// Create a new dialog.
	dialog, err := gtk.DialogNew()
	if err != nil {
		log.Println("Unable to create dialog:", err)
		return
	}
	dialog.SetTitle("System Information")
	dialog.SetDefaultSize(400, 300)

	// Get the content area of the dialog.
	contentArea, err := dialog.GetContentArea()
	if err != nil {
		log.Println("Unable to get content area:", err)
		return
	}

	// Create a new text view.
	textView, err := gtk.TextViewNew()
	if err != nil {
		log.Println("Unable to create text view:", err)
		return
	}
	textView.SetEditable(false)
	textView.SetWrapMode(gtk.WRAP_WORD)

	// Get the text buffer.
	buffer, err := textView.GetBuffer()
	if err != nil {
		log.Println("Unable to get buffer:", err)
		return
	}

	// Get system information.
	info := "OS: " + runtime.GOOS + "\n"
	info += "Arch: " + runtime.GOARCH + "\n"
	cmd := exec.Command("hostname")
	out, err := cmd.Output()
	if err == nil {
		info += "Hostname: " + string(out)
	}
	cmd = exec.Command("uname", "-a")
	out, err = cmd.Output()
	if err == nil {
		info += "Uname: " + string(out)
	}
	cmd = exec.Command("sudo", "dmidecode", "-s", "system-serial-number")
	out, err = cmd.Output()
	if err == nil {
		info += "Serial Number: " + string(out)
	} else {
		info += "Serial Number: Not available (requires root privileges)"
	}

	buffer.SetText(info)

	// Add the text view to the dialog.
	contentArea.Add(textView)

	// Add a close button.
	dialog.AddButton("Close", gtk.RESPONSE_CLOSE)
	dialog.Connect("response", func() {
		dialog.Destroy()
	})

	// Show the dialog.
	dialog.ShowAll()
}

func openPDF() {
	err := exec.Command("xdg-open", "Your-Linux-system.pdf").Start()
	if err != nil {
		log.Println("Could not open PDF:", err)
	}
}

func launchFirefox() {
	err := exec.Command("firefox").Start()
	if err != nil {
		log.Println("Could not launch application:", err)
	}
}

func launchLibreOffice() {
	err := exec.Command("libreoffice").Start()
	if err != nil {
		log.Println("Could not launch application:", err)
	}
}

func launchZoom() {
	err := exec.Command("zoom-client").Start()
	if err != nil {
		log.Println("Could not launch application:", err)
	}
}
