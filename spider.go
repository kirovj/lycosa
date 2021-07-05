package main

import "github.com/go-rod/rod"

const ShotPath = "shot"

// Spider the config of rod
type Spider struct {
	Url      string
	Name     string
	Timeout  uint8
	Children []*Spider
}

// Prey the result of rod
type Prey struct {
	Html string
}

// Netting rod build grub mission
func Netting(s *Spider) *Prey {
	page := rod.New().MustConnect().MustPage(s.Url)
	page.MustWaitLoad().MustScreenshot(ShotPath + s.Name + ".png")
	return &Prey{Html: "ok"}
}

//func Grub() *Prey {
//
//}
