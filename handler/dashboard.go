package handler

/*func Root(c *fiber.Ctx) error {
	if c.Get("Authorization") != "" {
		c.Redirect("/dashboard")
	} else {
		c.Redirect("/authentication")
	}
	return c.SendString("")
}*/

/*func Dashboard(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["email"].(string)
	return c.SendString("Welcome " + claims["firstName"].(string) + name)
}*/
