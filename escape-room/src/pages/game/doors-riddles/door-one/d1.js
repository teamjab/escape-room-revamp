
import store from "../../../../store/store";


function D1() {   
const player = store.getState();
const playerName = player.player.username;
const riddle1 = player.riddles.riddles[0].Question;
console.log(riddle1);
  


    return (
      <main className='first-door-page'>
        <h1>welcome {playerName} to door 1</h1>
        <p>{riddle1}</p>
      </main>
      
    );
  };
  
  export default D1;