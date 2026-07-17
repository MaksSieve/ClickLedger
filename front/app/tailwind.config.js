/** @type {import('tailwindcss').Config} */
/*
 ** TailwindCSS Configuration File
 **
 ** Docs: https://tailwindcss.com/docs/configuration
 ** Default: https://github.com/tailwindcss/tailwindcss/blob/master/stubs/defaultConfig.stub.js
 */
 const colors = require('tailwindcss/colors');
 module.exports = {
	 purge: {
		 enabled: process.env.NODE_ENV === 'production',
		 content: [
			 'components/**/*.vue',
			 'layouts/**/*.vue',
			 'pages/**/*.vue',
			 'plugins/**/*.js',
			 'static/**/*.vue',
			 'store/**/*.vue',
		 ],
	 },
	 theme: {
		 extend: {
			 colors: {
				 teal: colors.teal,
				 'blue-gray': colors.blueGray,
				 orange: colors.orange,
			 },
			 spacing: {
				 13: '3.25rem',
			 },
			 screens: {
				 'dashboard-xl': '1900px',
				 '2xl': '1335px',
				 '3xl': '2200px',
			 },
			 width: {
				 '1/8': '12.5%',
				 '2/8': '25%',
				 '3/8': '37.5%',
				 '4/8': '50%',
				 '5/8': '62.5%',
				 '6/8': '75%',
				 '7/8': '87.5%',
			 },
			 height: {
				 'vp-max': 'calc(100vh - 128px)',
			 },
		 },
	 },
	 variants: {},
	//  plugins: [require('@tailwindcss/forms')],
 };
 