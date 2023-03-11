package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {

	// Bewegungsraketen: 365, -1
	// Turnzwerge 2 Donnerstag: 297, selector id 2
	url := ""
	selector := -1
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 2 {
		url = fmt.Sprintf("https://kursverwaltung.mtv-bs.de/de/kurse/alle-kurse/course-details/?displayCourse=%s", argsWithoutProg[0])
		s, err := strconv.Atoi(argsWithoutProg[1])
		if err != nil {
			log.Fatal(fmt.Sprintf("could not parse selector: %v", err))
		}
		selector = s
	} else {
		log.Fatal("no url provided")
	}
	_ = selector
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 40*time.Second)
	defer cancel()

	// navigate to main page of the course
	//var buf []byte
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Select the correct hour if needed
	if selector != -1 {
		err = chromedp.Run(ctx,
			chromedp.SetAttributeValue(fmt.Sprintf(`(//select[@name="registration-date"]//option)[%d]`, selector), `selected`, `selected`),
		)
		if err != nil {
			log.Fatal(err)
		}
	}

	// click on 'Jetzt anmelden' to register
	err = chromedp.Run(ctx,
		chromedp.Click("//a[normalize-space() = 'Jetzt anmelden']", chromedp.BySearch),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Fill login data and submit the for to login.
	err = chromedp.Run(ctx,
		chromedp.SendKeys(`//input[@name="login[loginFieldValue]"]`, os.Getenv("MTV_MAIL")),
		chromedp.SendKeys(`//input[@name="login[loginPassword]"]`, os.Getenv("MTV_PASS")),
		chromedp.Click(`//input[@name="submit"]`),
	)
	if err != nil {
		log.Fatal(err)
	}

	// The first date will be pre-selected, no you just accept the terms and can submit
	err = chromedp.Run(ctx,
		chromedp.Click(`//input[@id="generalTerms-generalTermsConfirm"]`),
	)
	if err != nil {
		log.Fatal(err)
	}

	// for testing make a screenshot of what we see, otherwise we could just
	// click submit to register.
	err = chromedp.Run(ctx, //chromedp.FullScreenshot(&buf, 90),
		chromedp.Click(`//input[@value="Verbindlich anmelden"]`, chromedp.BySearch),
	)
	if err != nil {
		log.Fatal(err)
	}
	// err = chromedp.Run(ctx,
	// 	chromedp.FullScreenshot(&buf, 90),
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if err := os.WriteFile("fullScreenshot.png", buf, 0o644); err != nil {
	// 	log.Fatal(err)
	// }

}
