import store from "../../../../store/store";


function D4() {   
const player = store.getState();
const playerName = player.player.username;
  


    return (
      <main className='fourth-door-page'>
        <h1>welcome {playerName} to door 4</h1>
      </main>
      
    );
  };
  
  export default D4;