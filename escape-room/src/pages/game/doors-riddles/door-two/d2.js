import store from "../../../../store/store";

function D2() {   
const player = store.getState();
const playerName = player.player.username;
  


    return (
      <main className='second-door-page'>
        <h1>welcome {playerName} to door 2</h1>
      </main>
      
    );
  };
  
  export default D2;