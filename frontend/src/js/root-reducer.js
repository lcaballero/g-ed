export function RootReducer(state, action) {
    switch (action.type) {
        case "ADD_POST":
            var message = action.message.trim();
            if (message) {
                return {
                    ...state,
                    currentMessage: "",
                    posts: (state.posts || []).concat({
                        type: "message",
                        message: message
                   })
                };
            } else {
                return state;
            }
        case "current-message":
            return {...state, currentMessage: action.value };
        default:
            return state;
    }
}
