/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ["./views/**/*.html"],
	theme: {
		extend: {
			fontFamily: {
				sans: ["Nurito Sans", "sans-serif"],
			},
		},
	},
	plugins: [require("daisyui")],
	daisyui: {
		themes: ["corporate", "business"],
	},
};
