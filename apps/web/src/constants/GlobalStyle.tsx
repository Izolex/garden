import { createGlobalStyle } from 'styled-components'
// @ts-expect-error: Cannot find module
import Comic from 'url:./fonts/comic.ttf'
// @ts-expect-error: Cannot find module
import ComicBold from 'url:./fonts/comic-bold.ttf'

export const GlobalStyles = createGlobalStyle`
	:root {
		--pink-old: rgb(243,107,126);
		--pink: ${(props) => props.theme.color.pink};
		--green: rgba(145,232,66,1);
		--blue: ${(props) => props.theme.color.blue};
		--orange: ${(props) => props.theme.color.orange};
	}

	@font-face {
		font-family: 'Comic';
		src: local('Comic'), url(${Comic}) format('woff');
	}

	@font-face {
		font-family: 'Comic';
		font-weight: bold;
		src: local('Comic'), url(${ComicBold}) format('woff');
	}

	* {
		padding: 0;
		border: 0;
		margin: 0;
		outline: 0;
	}

	 body {
		background: rgb(255,255,255);
		background: linear-gradient(180deg, rgba(255,255,255,1) 0%, var(--green) 100%);
		color: black;
		font-family: "Comic", "Comic Sans MS", "Comic Sans", cursive;
		text-align: center;
	}

	a {
		color: black;
	}

	h1, h2, h3, h4, h5, h6 {

  font-weight: bold;
  }
`
