const user = {
    username: 'Brendon',
    score:0,
}

export default function useReducer( state = user, action) {
    const {type, payload} = action;
    switch(type) {
        case 'INITIALIZE':
            return{...state,user:payload}
            default:return state
    }
}

export const User = (payload) => {
    return {
        type:'INITIALIZE',
        payload:payload,
    }
}