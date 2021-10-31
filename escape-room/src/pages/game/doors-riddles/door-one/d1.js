
import store from "../../../../store/store";


function D1() {   
const player = store.getState();
const playerName = player.player.username;
  


    return (
      <main className='first-door-page'>
        <h1>welcome {playerName} to door 1</h1>
      </main>
      
    );
  };
  
  export default D1;