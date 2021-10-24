const riddles = {
    riddles:[],
};

function riddleReducer(state = riddles, action) {
    switch(action.type) {
        case 'initialState': {
            return {
                ...state
            }
        }
        case 'getriddles': {
            return {
                ...state,
                riddles:action.payload,
            }
        }
        default:
    }    
    return state
}

export default riddleReducer;