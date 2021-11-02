import store from "../../../../store/store";


function D4() {   
const player = store.getState();
const playerName = player.player.username;
const riddle4 = player.riddles.riddles[3].Question;
console.log(riddle4);

  


    return (
      <main className='fourth-door-page'>
        <h1>welcome {playerName} to door 4</h1>
        <p>{riddle4}</p>
      </main>
      
    );
  };
  
  export default D4;