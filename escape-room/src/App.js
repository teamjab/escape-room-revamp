import userForm from "./components/userForm.js";
import {useSelector} from 'react-redux';
import Landing from './pages/landing/landing.js';
import Game from './pages/game/game.js';
import { BrowserRouter as Router, Switch, Route} from "react-router-dom";




function App() {
  const name = useSelector( state => state);
  console.log(userForm);
  console.log(name.player);
  
  return (
    <Router>
      <Switch>
        <Route path='/' exact component={Landing}/>
        <Route path='/game'exact component={Game}/>
      </Switch>
        <div className="App">
          <header className="App-header"></header>
          <footer id='JAB'>
            <h4>Team Jab</h4>
          </footer>
        </div>
    </Router>
  );
}

export default App;
