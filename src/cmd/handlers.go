package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
    "tutor/cmd/views"
	views "tutor/cmd/views/pages"
	"tutor/cmd/models"

	"github.com/labstack/echo/v4"
)

var cachedReviews []models.Review

func RegisterHandlers(e *echo.Echo, db DB) {
	fs, _ := embed.GetStaticDirFS()

	e.StaticFS("/static", fs)
	HomePage(e)
	ServicesPage(e)
	ReviewsPage(e, db)
	CreateReview(e, db)
	ContactUs(e)
}

func HomePage(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return views.Base("Home Page", views.Home()).Render(c.Request().Context(), c.Response().Writer)
	})
}

func ServicesPage(e *echo.Echo) {
	e.GET("/services", func(c echo.Context) error {
		return views.Base("Services", views.Services()).Render(c.Request().Context(), c.Response().Writer)
	})
}

func ContactUs(e *echo.Echo) {
	e.GET("/contact-us", func(c echo.Context) error {
		return views.Base("Contact Us", views.ContactUs()).Render(c.Request().Context(), c.Response().Writer)
	})

	e.POST("/contact-us", func(c echo.Context) error {

		name := c.FormValue("name")
		number := c.FormValue("number")
		body := c.FormValue("body")

		var nameError string
		var bodyError string
		var numberError string
		var isError bool

		if len(name) == 0 {

			nameError = `

        <div class="error" id="name-error" :style="{ 'width': nameinput.offsetWidth + 'px'}">Name cannot be empty</div>

        `
			isError = true
		}

		if len(body) == 0 {

			bodyError = `


        <div class="error" id="body-error" :style="{ 'width': nameinput.offsetWidth + 'px'}">Message cannot be empty</div>
        `
			isError = true
		}
		if len(number) == 0 {

			numberError = `

        <div class="error" id="body-error" :style="{ 'width': nameinput.offsetWidth + 'px'}">Number cannot be empty</div>
        `
			isError = true
		}

		re := regexp.MustCompile("^[0-9]+$")

		if !re.MatchString(number) {
			numberError = `

        <div class="error" id="body-error" :style="{ 'width': nameinput.offsetWidth + 'px'}">Number must contain only digits</div>
        `
			isError = true
		}

		messageSent := ""

		if !isError {
			messageSent = "<p>Message sent, you will be contacted shortly</p>"
		}

		formHTML := fmt.Sprintf(`
    <form id="form" x-data="{ nameinput : $refs.nameinput}" class="flex column align-center full-width form-container"
        hx-post="/contact-us" hx-target="form" hx-swap="outerHTML">
        <h2>Request a callback</h2>
        <label :style="{ 'width': nameinput.offsetWidth + 'px'}" for="name">Name</label>
            %s
        <input x-ref="nameinput" type="text" name="name" value="%s" >
        <label :style="{ 'width': nameinput.offsetWidth + 'px'}" for="number">Number</label>
            %s
        <input x-ref="nameinput" type="text" name="number" value="%s" >
        <label :style="{ 'width': nameinput.offsetWidth + 'px'}" for="body" >Message</label>
            %s
        <textarea :style="{ 'width': nameinput.offsetWidth + 'px', 'min-height': '4rem', 'margin-bottom': '1rem' }" type="text" name="body"
            value="%s">%s</textarea>
        <button class="submit-button" type="submit">Submit</button>
            %s
    </form>

        `, nameError, name, numberError, number, bodyError, body, body, messageSent)

		if !isError {
			EmailService("Callback Request", fmt.Sprintf("From: %s \n Message: %s \n Callback on: %s", name, body, number))
		}

		c.Request().Header.Set("Content-Type", "text/html")

		return c.HTML(http.StatusOK, formHTML)
	})
}

func ReviewsPage(e *echo.Echo, db DB) {
	e.GET("/reviews", func(c echo.Context) error {

		if len(cachedReviews) == 0 {

			var isError bool
			reviews, err := db.GetReviews()
			if err != nil {
				e.Logger.Print(err)
				isError = true
			}
			if isError {
				return views.Base("Home Page", views.Home()).Render(c.Request().Context(), c.Response().Writer)
			}
			cachedReviews = reviews
		}

		return views.Base("Reviews", views.Reviews(cachedReviews)).Render(c.Request().Context(), c.Response().Writer)
	})
}

func CreateReview(e *echo.Echo, db DB) {
	e.POST("/reviews/create", func(c echo.Context) error {
		name := c.FormValue("name")
		body := c.FormValue("body")
		stars := c.FormValue("stars")

		var nameError string
		var bodyError string
		var isError bool

		if len(name) == 0 {

			nameError = `

        <div class="error" id="name-error" :style="{ 'width': nameinput.offsetWidth + 'px'}">Name cannot be empty</div>

        `
			isError = true
		}

		if len(body) == 0 {

			bodyError = `


        <div class="error" id="body-error" :style="{ 'width': nameinput.offsetWidth + 'px'}">Review content cannot be empty</div>
        `
			isError = true
		}

		starsF, err := strconv.ParseFloat(stars, 32)
		if err != nil {
			isError = true
		}

		sentForApproval := ""

		if !isError {
			sentForApproval = "<p>Review submitted for approval</p>"
		}

		formHTML := fmt.Sprintf(`
    <form id="form" x-data="{ nameinput : $refs.nameinput}" class="flex column align-center full-width form-container" hx-post="/reviews/create" hx-target="form" hx-swap="outerHTML">
        <h2>Send a Review!</h2>
        <label :style="{ 'width': nameinput.offsetWidth + 'px'}" for="Name">Name</label>
            %s
        <input x-ref="nameinput" type="text" name="name" value="%s" x-on:keydown="document.getElementById('name-error').innerHTML = ''">
        <label :style="{ 'width': nameinput.offsetWidth + 'px'}" for="body">Review</label>
            %s
        <textarea :style="{ 'width': nameinput.offsetWidth + 'px', 'min-height': '4rem' }" type="text" name="body" x-on:keydown="document.getElementById('body-error').innerHTML = ''"
            value="%s">%s</textarea>
        <div class="star-rating" x-init="document.getElementById('5-stars').checked = true;" >
            <input name="stars" type="radio" id="5-stars" name="rating" value="5">
            <label for="5-stars" class="star">★</label>

            <input name="stars" type="radio" id="4-stars" name="rating" value="4">
            <label for="4-stars" class="star">★</label>

            <input name="stars" type="radio" id="3-stars" name="rating" value="3">
            <label for="3-stars" class="star">★</label>

            <input name="stars" type="radio" id="2-stars" name="rating" value="2">
            <label for="2-stars" class="star">★</label>

            <input name="stars" type="radio" id="1-stars" name="rating" value="1">
            <label for="1-stars" class="star">★</label>
        </div>
        <button class="submit-button" type="submit">Submit</button>
            %s
    </form>

        `, nameError, name, bodyError, body, body, sentForApproval)

		if !isError {
			db.CreatePendingReview(models.Review{Name: name, Body: body, Stars: starsF})
		}

		c.Request().Header.Set("Content-Type", "text/html")

		return c.HTML(http.StatusOK, formHTML)
	})
}
