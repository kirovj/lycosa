package goal

import (
	"github.com/go-rod/rod"
	"time"
)

func Try() {
	page := rod.New().MustConnect().MustPage("https://www.wikipedia.org/").MustWindowFullscreen()

	page.MustElement("#searchInput").MustInput("earth")
	page.MustElement("#search-form > fieldset > button").MustClick()

	page.MustWaitLoad().MustScreenshot("a.png")
	time.Sleep(time.Minute)
}
