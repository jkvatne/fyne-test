// Package main provides various examples of Fyne API capabilities
package main

import (
	"fmt"
	"net/url"

	"github.com/jkvatne/fyne"
	"github.com/jkvatne/fyne/app"
	"github.com/jkvatne/fyne/canvas"
	"github.com/jkvatne/fyne/cmd/fyne_demo/data"
	"github.com/jkvatne/fyne/cmd/fyne_demo/screens"
	"github.com/jkvatne/fyne/layout"
	"github.com/jkvatne/fyne/theme"
	"github.com/jkvatne/fyne/widget"
)

const preferenceCurrentTab = "currentTab"

func welcomeScreen(a fyne.App) fyne.CanvasObject {
	logo := canvas.NewImageFromResource(data.FyneScene)
	logo.SetMinSize(fyne.NewSize(228, 167))

	link, err := url.Parse("https://github.com/jkvatne/")
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	return widget.NewVBox(
		widget.NewLabelWithStyle("Welcome to the Fyne toolkit demo app", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		layout.NewSpacer(),
		widget.NewHBox(layout.NewSpacer(), logo, layout.NewSpacer()),
		widget.NewHyperlinkWithStyle("github.com/jkvatne", link, fyne.TextAlignCenter, fyne.TextStyle{}),
		layout.NewSpacer(),

		widget.NewGroup("Theme",
			fyne.NewContainerWithLayout(layout.NewGridLayout(2),
				widget.NewButton("Dark", func() {
					a.Settings().SetTheme(theme.DarkTheme())
				}),
				widget.NewButton("Light", func() {
					a.Settings().SetTheme(theme.LightTheme())
				}),
			),
		),
	)
}

func main() {
	a := app.NewWithID("io.fyne.demo")
	a.SetIcon(theme.FyneLogo())

	w := a.NewWindow("Fyne Demo")
	w.SetMainMenu(fyne.NewMainMenu(fyne.NewMenu("File",
		fyne.NewMenuItem("New", func() { fmt.Println("Menu New") }),
		// a quit item will be appended to our first menu
	), fyne.NewMenu("Edit",
		fyne.NewMenuItem("Cut", func() { fmt.Println("Menu Cut") }),
		fyne.NewMenuItem("Copy", func() { fmt.Println("Menu Copy") }),
		fyne.NewMenuItem("Paste", func() { fmt.Println("Menu Paste") }),
	)))
	w.SetMaster()

	tabs := widget.NewTabContainer(
		widget.NewTabItemWithIcon("Welcome", theme.HomeIcon(), welcomeScreen(a)),
		widget.NewTabItemWithIcon("Widgets", theme.ContentCopyIcon(), screens.WidgetScreen()),
		widget.NewTabItemWithIcon("Graphics", theme.DocumentCreateIcon(), screens.GraphicsScreen()),
		widget.NewTabItemWithIcon("Windows", theme.ViewFullScreenIcon(), screens.DialogScreen(w)),
		widget.NewTabItemWithIcon("Advanced", theme.SettingsIcon(), screens.AdvancedScreen(w)))
	tabs.SetTabLocation(widget.TabLocationLeading)
	tabs.SelectTabIndex(a.Preferences().Int(preferenceCurrentTab))
	w.SetContent(tabs)

	w.ShowAndRun()
	a.Preferences().SetInt(preferenceCurrentTab, tabs.CurrentTabIndex())
}
