let shortid = require('shortid')

const maxPostLength = 100000

exports.post = function*(request, response) {
	let user = request.user

	if(!user) {
		response.writeHead(409)
		response.end('Not logged in')
		return
	}
	
	let threadId = request.body.threadId
	
	if(!threadId) {
		response.writeHead(409)
		response.end('No thread specified')
		return
	}
	
	let thread = yield arn.get('Threads', threadId)
	
	if(!thread) {
		response.writeHead(409)
		response.end('Thread does not exist')
		return
	}

	let text = request.body.text

	if(!text) {
		response.writeHead(409)
		response.end('Post text required')
		return
	}
	
	text = text.trim()
	
	if(text.length > maxPostLength) {
		response.writeHead(409)
		response.end('Post too long')
		return
	}
	
	let postId = shortid.generate()
	
	// Save post
	yield arn.set('Posts', postId, {
		id: postId,
		authorId: user.id,
		threadId: thread.id,
		text,
		likes: [],
		created: (new Date()).toISOString()
	})
	
	// Update reply count
	if(!thread.replies)
		thread.replies = 0
	
	arn.set('Threads', thread.id, {
		replies: thread.replies + 1
	})
	
	response.end(postId)
}