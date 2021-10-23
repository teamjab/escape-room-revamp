import store from '../../store/store';
import '../../styling/game/game.scss';
function Game() {
    
    const player = store.getState();
    console.log(player.username)
    return (
      <main className='game-page'>
        <h1>welcome {player.username}</h1>
        <h3>here is your riddle???</h3>
        <p></p>
      </main>
      
    );
  };
  
  export default Game;