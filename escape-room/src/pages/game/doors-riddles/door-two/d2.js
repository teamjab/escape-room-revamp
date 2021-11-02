import store from "../../../../store/store";

function D2() {   
const player = store.getState();
const playerName = player.player.username;
const riddle2 = player.riddles.riddles[1].Question;
console.log(riddle2);


    return (
      <main className='second-door-page'>
        <h1>welcome {playerName} to door 2</h1>
        <p>{riddle2}</p>
      </main>
      
    );
  };
  
  export default D2;