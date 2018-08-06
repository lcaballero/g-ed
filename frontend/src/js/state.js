

export function initial() {
    return {
        posts: [
            { type: "message", message: "Hello, welcome to the chatting app." }
        ],
        currentMessage:""
    }
}
