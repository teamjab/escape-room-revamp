import useReducer from './userReducer.js';
import {createStore, combineReducers} from 'redux';

const reducer = combineReducers({
    user:useReducer,
})

const store = () => {
    return createStore(reducer);
}

export default store();