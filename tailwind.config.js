import daisyui from "daisyui";
/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ["./views/*.templ"],
	theme: {
		extend: {},
	},
	plugins: [daisyui],
};
