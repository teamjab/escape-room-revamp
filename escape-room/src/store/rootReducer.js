import player from './userReducer.js';
import riddles from './riddleReducer.js';
import { combineReducers } from 'redux';

export default combineReducers ({
    player,
    riddles
});