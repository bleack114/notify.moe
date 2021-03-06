import { Diff } from "./Diff"

class LoadOptions {
	addToHistory?: boolean
	forceReload?: boolean
}

export class Application {
	fadeOutClass: string
	activeLinkClass: string
	content: HTMLElement
	loading: HTMLElement
	currentPath: string
	originalPath: string
	lastRequest: XMLHttpRequest | null

	constructor() {
		this.currentPath = window.location.pathname
		this.originalPath = window.location.pathname
		this.activeLinkClass = "active"
		this.fadeOutClass = "fade-out"
	}

	init() {
		document.addEventListener("DOMContentLoaded", () => {
			let links = document.getElementsByTagName("a")

			this.markActiveLinks(links)
			this.ajaxify(links)
		})
	}

	find(id: string): HTMLElement | null {
		return document.getElementById(id)
	}

	get(url: string): Promise<string> {
		// return fetch(url, {
		// 	credentials: "same-origin"
		// }).then(response => response.text())

		if(this.lastRequest) {
			this.lastRequest.abort()
			this.lastRequest = null
		}

		return new Promise((resolve, reject) => {
			let request = new XMLHttpRequest()

			request.onerror = () => reject(new Error("You are either offline or the requested page doesn't exist."))
			request.ontimeout = () => reject(new Error("The page took too much time to respond."))
			request.onload = () => {
				if(request.status < 200 || request.status >= 400)
					reject(request.responseText)
				else
					resolve(request.responseText)
			}

			request.open("GET", url, true)
			request.send()

			this.lastRequest = request
		})
	}

	load(url: string, options?: LoadOptions) {
		// Start sending a network request
		let request = this.get("/_" + url).catch(error => error)

		// Parse options
		if(!options) {
			options = new LoadOptions()
		}

		if(options.addToHistory === undefined) {
			options.addToHistory = true
		}

		// Set current path
		this.currentPath = url

		// Add to browser history
		if(options.addToHistory) {
			history.pushState(url, "", url)
		}

		// Mark active links
		this.markActiveLinks()

		let onTransitionEnd = (e: Event) => {
			// Ignore transitions of child elements.
			// We only care about the transition event on the content element.
			if(e.target !== this.content) {
				return
			}

			// Outdated response.
			if(this.currentPath !== url) {
				return
			}

			// Remove listener after we finally got the correct event.
			this.content.removeEventListener("transitionend", onTransitionEnd)

			// Wait for the network request to end.
			request.then(html => {
				// Set content
				this.setContent(html)
				this.scrollToTop()

				// Fade animations
				this.content.classList.remove(this.fadeOutClass)
				this.loading.classList.add(this.fadeOutClass)

				// Send DOMContentLoaded Event
				this.emit("DOMContentLoaded")
			})
		}

		this.content.addEventListener("transitionend", onTransitionEnd)

		this.content.classList.add(this.fadeOutClass)
		this.loading.classList.remove(this.fadeOutClass)

		return request
	}

	setContent(html: string) {
		this.content.innerHTML = html
	}

	markActiveLinks(links?: NodeListOf<HTMLAnchorElement>) {
		if(!links) {
			links = document.getElementsByTagName("a")
		}

		for(let i = 0; i < links.length; i++) {
			let link = links[i]

			Diff.mutations.queue(() => {
				if(link.getAttribute("href") === this.currentPath) {
					link.classList.add(this.activeLinkClass)
				} else {
					link.classList.remove(this.activeLinkClass)
				}
			})
		}
	}

	ajaxify(links?: NodeListOf<HTMLAnchorElement>) {
		if(!links) {
			links = document.getElementsByTagName("a")
		}

		for(let i = 0; i < links.length; i++) {
			let link = links[i] as HTMLAnchorElement

			// Don't ajaxify links to a different host
			if(link.hostname !== window.location.hostname) {
				if(!link.target) {
					link.target = "_blank"
				}

				continue
			}

			// Don't ajaxify invalid links, links with a target or links that disable ajax specifically
			if(link.href === "" || link.href.includes("#") || link.target.length > 0 || link.dataset.ajax === "false") {
				continue
			}

			let self = this

			link.onclick = function(e) {
				// Middle mouse button should have standard behaviour
				if(e.which === 2) {
					return
				}

				let url = this.getAttribute("href")

				e.preventDefault()

				if(!url || url === self.currentPath) {
					return
				}

				// Load requested page
				self.load(url)
			}
		}
	}

	scrollToTop() {
		let parent : HTMLElement | null = this.content

		Diff.mutations.queue(() => {
			while(parent = parent.parentElement) {
				parent.scrollTop = 0
			}
		})
	}

	emit(eventName: string) {
		document.dispatchEvent(new Event(eventName, {
			"bubbles": true,
			"cancelable": true
		}))
	}
}