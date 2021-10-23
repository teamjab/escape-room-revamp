import store from '../../store/store';
import '../../styling/game/game.scss';
function Game() {

    console.log(' Game Page State', store.getState());
    return (
      <main className='game-page'>
        <h1>Game page</h1>
      </main>
      
    );
  };
  
  export default Game;