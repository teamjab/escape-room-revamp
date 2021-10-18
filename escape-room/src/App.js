import userForm from "./components/userForm.js";
import {useSelector} from 'react-redux';
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
import './styling/App.scss';
import Landing from './pages/landing/landing.js';
import Game from './pages/game/game.js';



function App() {
  const name = useSelector( state => state);
  console.log(userForm);
  console.log(name.player);
  
  return (
    <Router>
      <div className="App">
        <header className="App-header"></header>
        <main>
          <Landing></Landing>
          <form>
              <label>Username: </label>
              <input placeholder="enter username"></input>
          </form>
          <button type="submit">Submit</button>
        </main>
        <footer>
          <h4>Team Jab</h4>
        </footer>
      </div>
    </Router>
  );
}

export default App;
