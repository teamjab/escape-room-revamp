import store from "../../store/store";

const url = 'https://escape-room-revamp-api.herokuapp.com/riddles';


const fetchRiddles = async () => {

   await fetch(url)
    .then(res => res.json())
    .then(json => {
      store.dispatch({type:'getriddles', payload:json });
    })
}
console.log(store.getState())
export default fetchRiddles;