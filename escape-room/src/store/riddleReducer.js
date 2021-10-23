const riddles = [];

function riddleReducer(state = riddles, action) {
    switch(action.type) {
        case 'initialState': {
            return {
                ...state
            }
        }
        default:
    }    
    return state
}

export default riddleReducer;