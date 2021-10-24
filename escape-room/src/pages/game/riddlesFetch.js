

const url = 'https://escape-room-revamp-api.herokuapp.com/riddles';

const fetchRiddles = () => {
    fetch(url)
    .then(response => response.json())
    // .then(data => console.log(data))
    .catch(error => console.error(error));
}

export default fetchRiddles;