import reducers from './rootReducer.js';
import {createStore} from 'redux';


const store = createStore(reducers);

export default store;