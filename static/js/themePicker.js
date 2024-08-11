document.addEventListener("DOMContentLoaded", () => {
	const html = document.documentElement;

	// Function to apply the theme based on user preference
	function applyTheme() {
		const prefersDarkScheme = window.matchMedia(
			"(prefers-color-scheme: dark)",
		).matches;
		html.setAttribute(
			"data-theme",
			prefersDarkScheme ? "business" : "corporate",
		);
	}

	// Apply theme on page load
	applyTheme();

	// Optional: Add a listener to apply theme changes if the user changes system theme
	window
		.matchMedia("(prefers-color-scheme: dark)")
		.addEventListener("change", applyTheme);
});
