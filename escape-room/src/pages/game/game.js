import store from '../../store/store';
import '../../styling/game/game.scss';
function Game() {
    
  const player = store.getState();
  const playerName = player.player.username;
  const riddle = store.getState();


    return (
      <main className='game-page'>
        <h1>welcome {playerName}</h1>
        <h3>here is your riddle???</h3>
        <p></p>
      </main>
      
    );
  };
  
  export default Game;