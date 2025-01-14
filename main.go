package main

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Static("/", "./frontend")

	app.Post("/submit", func(c *fiber.Ctx) error {
		email := c.FormValue("email")
		if strings.Contains(email, ".com") && strings.Contains(email, "@") {
			html:=`
			<main class="md:w-1/3 p-5 bg-white text-black sm:h-screen md:h-auto" style="border-radius: 20px;">
  <div class="grid gap-5 p-5">
    <img src="assets/images/icon-list.svg" width="50px" height="50px" alt="tick">
    <h1 class="text-5xl font-bold">Thanks for subscribing!</h1>
  </div>
  
  <p class="p-5">

    A confirmation email has been sent to <b>`+email+`</b>
    Please open it and click the button inside to confirm your subscription.
  </p>
  <div class="w-full flex justify-center align-center p-5">

    <button  class="w-full my-4 p-5 text-white font-bold rounded-xl"
    style="background-color: hsl(234, 29%, 20%)">Dismiss message</button>
  </div>
  
</main>
			
			
			`
			return c.SendString(html)
		}
		html := `
		    <main
      class="bg-white text-black  md:rounded-xl md:p-5 sm:w-100%  md:flex justify-center align-center  "
      style="text-align: left; id="response""
    >
    <div
    class="md:hidden h-64 w-100% p-5"
    style="
      background-image: url(assets/images/illustration-sign-up-mobile.svg);
      background-size: cover;
      border-radius: 0 0 25px 25px;
    "
  ></div>
      <div class="p-5">
        <h1 class="font-bold px-5 py-2" style="font-size: 3rem">
          Stay updated!
        </h1>
        <p class="p-5">
          Join 60,000+ product managers receiving monthly updates on:
        </p>
        <ul class="px-5">
          <li>
            <img src="assets/images/icon-list.svg" alt="tick" />
            <p>Product discovery and building what matters</p>
          </li>
          <li>
            <img src="assets/images/icon-list.svg" alt="tick" />
            <p>Measuring to ensure updates are a success</p>
          </li>
          <li>
            <img src="assets/images/icon-list.svg" alt="tick" />
            <p>And much more!</p>
          </li>
        </ul>

        <form hx-post="/submit" hx-swap="innerHTML" hx-target="#response">
		<div class="px-5 py-2" >
            <label for="email" class="text-sm flex" style="justify-content: space-between;"><p>Email Address</p> <p style="color:orangered;">Valid email required</p></label>
            <br />
            <input
              type="text"
              name="email"
              id="email"
              placeholder="email@company.com"
              required
              class="p-4 rounded-lg w-full "
              style="border:2px solid orangered;"
            />
			<style>
			input:invalid{
			background-color:#FFD580;
			color:orangered;
			}
			</style>            
			<button
              class="w-full my-4 p-2 text-white font-bold rounded-xl"
              style="background-color: hsl(234, 29%, 20%)"
            >
            Subscribe to monthly newsletter
            </button>
          </div>
		      </form>
        
      </div>
      <div
        class=" rounded-xl sm:hidden md:block"
        style="
          background-image: url(assets/images/illustration-sign-up-desktop.svg);
          background-size: cover;
          min-width: 20vw;
        "
      ></div> 
		  `
		return c.SendString(html)
	})
	

	app.Listen(":9000")
}