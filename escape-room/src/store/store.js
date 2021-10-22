import playerReducer from './userReducer.js';
import {createStore} from 'redux';


const store = createStore(playerReducer);

export default store;