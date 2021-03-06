import store from '../../store/store';
import {Link} from 'react-router-dom';
import '../../styling/landing/landing.scss';
import fetchRiddles from '../game/riddlesFetch.js';
const Landing = () => {

  const getName = (event) => {
    // console.log(event.target.value);
    store.dispatch({ type: 'name', payload: event.target.value })
  }

    return (
        <main className="landing-page">
            <h1>
            <span>E</span>
            <span>s</span>
            <span>c</span>
            <span>a</span>
            <span>p</span>
            <span>e</span>
            <span>R</span>
            <span>o</span>
            <span>o</span>
            <span>m!</span>
            </h1>
            <form>
              <label>Username: </label>
              <input onChange={getName} placeholder="enter username"></input>
            </form>
            <Link to="/game">
              <button onClick={fetchRiddles} type='submit'>Play!</button>
            </Link>
          </main>
    )
  };
  
  export default Landing;