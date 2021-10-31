import store from "../../../../store/store";

function D3() {   
const player = store.getState();
const playerName = player.player.username;
  


    return (
      <main className='third-door-page'>
        <h1>welcome {playerName} to door 3</h1>
      </main>
      
    );
  };
  
  export default D3;