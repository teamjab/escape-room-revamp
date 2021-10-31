import store from '../../store/store';
import '../../styling/game/game.scss';
import {Link} from 'react-router-dom';

function Game() {   
const player = store.getState();
const playerName = player.player.username;
  
  const getRiddles = () => {
    player.player.score += 100;
    console.log(store.getState());

  }


    return (
      <main className='game-page'>
        <h1>welcome {playerName}</h1>
        <p>Upon entering the House you notice a dim lit hallway teaming with doors all around you. You take note that the room inside looks nothing comparable to the house as it looked outside.</p>
        <p>Even stranger, you notice numbers on the doors, almost beckoning you to enter through one.</p>
        <Link to='/d1'>
          <button onClick={getRiddles}>First Door</button>
        </Link>
        <Link to='/d2'>
          <button onClick={getRiddles}>Second Door</button>
        </Link>
        <Link to='/d3'>
          <button onClick={getRiddles}>Third Door</button>
        </Link>
        <Link to='/d4'>
          <button onClick={getRiddles}>Fourth Door</button>
        </Link>
        
        
      </main>
      
    );
  };
  
  export default Game;