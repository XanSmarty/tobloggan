package html

const CSS = `
		<style>
			:root {
				color-scheme: light dark;
				--text-color: light-dark(#222, #ddd);
				--link-color: light-dark(#ddd, #222);
				--hint-color: light-dark(#555, #aaa);
				font-size: calc(1rem + 0.25vw); /* https://jameshfisher.com/2024/03/12/a-formula-for-responsive-font-size/ */
			}
			html {
				-webkit-text-size-adjust: none;
				text-size-adjust: none;
				-moz-text-size-adjust: none;
			}
			body {
				font-family: Tahoma, sans-serif;
				margin-left: auto;
				margin-right: auto;
				max-width: calc(30rem + 30vw);
				line-height: 1.5em;
				color: var(--text-color);
			}
			body p.article-entry {
				margin: 0em;
				text-overflow: ellipsis;
				white-space: nowrap;
				overflow: hidden; 
			}
			body p.article-entry .date {
				font-family: monospace;
				font-size: 0.8em;
				color: var(--hint-color);
			}
			header {
				font-size: 1.2em;
				margin-top: 1.2em;
				margin-bottom: 1.2em;
			}
			h1, h2, h3, h4, h5, h6 {
				font-family: Verdana;
				line-height: 2em;
			}
			header, footer .outro {
				margin: 2em 0.1em;
				border: 0.2em solid var(--text-color);
				border-radius: 0.5em;
				text-align: center;
				padding: 1em;
			}
			header .date {
				font-size: 0.8em;
			}
			header .tldr {
				font-size: 0.8em;
				font-style: italic;
				font-weight: normal;
				margin: 0 0 2em 0;
			}
			header a.topic {
				text-decoration: none;
				font-size: 0.8em;
				margin: 0.5em;
				line-height: 2.5em;
				padding: 0.2em 0.3em;
				border: 0.1em solid var(--text-color);
				border-radius: 0.3em;
			}
			footer .outro {
				margin-top: 8em;
			}
			a {
				padding: 0.2em;
			}
			a:link {
				color: var(--text-color);
			}
			a:visited {
				color: var(--text-color);
			}
			a:not(:has(img)):hover, header a.topic:hover {
				color: var(--link-color);
				background-color: var(--text-color);
				border-radius: 0.3em;
			}
			dt {
				margin-top: 2em;
			}
			pre {
				overflow-x: auto;
				margin: 0.5em;
				border: 0.1em solid var(--text-color);
				border-radius: 0.3em;
				display: inline-block;
			}
			pre code {
				border: none;
			}
			code {
				padding: 0.1em;
				line-height: 1.2em;
				font-family: monospace;
				border: 0.05em dotted;
				border-radius: 0.2em;
			}
			blockquote {
				padding-left:  0.5em; 
				margin-inline: 0.5em;
				border-left:   0.2em solid var(--text-color);
			}
			img {
				max-width: 100%;
				box-shadow: 0 0 1em var(--text-color);
				border-radius: 0.3em;
				margin: 2em auto;
			}
        </style>
`
