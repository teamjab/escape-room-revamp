const player = {
    username: '',
    score:0,
}


function playerReducer(state = player, action) {
    switch(action.type) {
        case 'name': {
            return {
                ...state,
                username:action.payload,
            }
        }
        case 'score': {
            return {
                ...state,
                score:action.payload,
            }
        }
        default:
    }    
    return state
}

export default playerReducer;