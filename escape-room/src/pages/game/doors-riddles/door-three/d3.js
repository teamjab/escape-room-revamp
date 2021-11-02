import store from "../../../../store/store";

function D3() {   
const player = store.getState();
const playerName = player.player.username;
const riddle3 = player.riddles.riddles[2].Question;
console.log(riddle3);



    return (
      <main className='third-door-page'>
        <h1>welcome {playerName} to door 3</h1>
        <p>{riddle3}</p>
      </main>
      
    );
  };
  
  export default D3;