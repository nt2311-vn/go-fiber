package handlers

import "github.com/gofiber/fiber/v2"

type SidebarItem struct {
	Path string
	Name string
	Icon string
}

var sidebarItems []SidebarItem = []SidebarItem{
	{Path: "/app/dashboard", Name: "Dashboard", Icon: "/static/icon/list.svg"},
	{Path: "/app/stock", Name: "Stock", Icon: "/static/icon/stock.svg"},
	{Path: "/app/sales", Name: "Sales", Icon: "/static/icon/dollar-sign.svg"},
	{Path: "/app/pnp", Name: "Purchase and Payment", Icon: "/static/icon/shopping-cart.svg"},
	{Path: "/app/budget", Name: "Budget", Icon: "/static/icon/piggy-bank.svg"},
}

func HomePage(c *fiber.Ctx) error {
	return c.Render("home", nil)
}

func DashboardPage(c *fiber.Ctx) error {
	return c.Render("dashboard", fiber.Map{"CurrentPath": c.Path(), "SidebarItems": sidebarItems})
}

func StockPage(c *fiber.Ctx) error {
	return c.Render("stock", fiber.Map{"CurrentPath": c.Path(), "SidebarItems": sidebarItems})
}

func SalesPage(c *fiber.Ctx) error {
	return c.Render("sales", fiber.Map{"CurrentPath": c.Path(), "SidebarItems": sidebarItems})
}

func PnPPage(c *fiber.Ctx) error {
	return c.Render("pnp", fiber.Map{"CurrentPath": c.Path(), "SidebarItems": sidebarItems})
}

func BudgetPage(c *fiber.Ctx) error {
	return c.Render("budget", fiber.Map{"CurrentPath": c.Path(), "SidebarItems": sidebarItems})
}
