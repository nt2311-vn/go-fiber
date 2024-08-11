document.addEventListener("DOMContentLoaded", () => {
	const html = document.documentElement;
	const themeToggleButton = document.getElementById("theme-toggle-button");
	const sunIcon = document.getElementById("sun-icon");
	const moonIcon = document.getElementById("moon-icon");
	const prefersDarkScheme = window.matchMedia(
		"(prefers-color-scheme: dark)",
	).matches;

	const userMenuButton = document.getElementById("user-menu-button");
	const userMenu = document.getElementById("user-menu");

	function applyTheme(theme) {
		html.setAttribute("data-theme", theme);
		localStorage.setItem("theme", theme);

		if (theme === "corporate") {
			sunIcon.classList.add("hidden");
			moonIcon.classList.remove("hidden");
		} else {
			moonIcon.classList.add("hidden");
			sunIcon.classList.remove("hidden");
		}
	}

	const savedTheme = localStorage.getItem("theme");
	if (savedTheme) {
		applyTheme(savedTheme);
	} else {
		applyTheme(prefersDarkScheme ? "business" : "corporate");
	}

	themeToggleButton.addEventListener("click", () => {
		const currentTheme = html.getAttribute("data-theme");
		const newTheme = currentTheme === "corporate" ? "business" : "corporate";
		applyTheme(newTheme);
	});

	if (userMenuButton) {
		userMenuButton.addEventListener("click", () => {
			userMenu.classList.toggle("hidden");
		});
	}
});
