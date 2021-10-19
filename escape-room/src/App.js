import userForm from "./components/userForm.js";
import {useSelector} from 'react-redux';
import Landing from './pages/landing/landing.js';
import Game from './pages/game/game.js';
import { BrowserRouter as Router, Switch, Route} from "react-router-dom";
import './styling/App.scss';




function App() {
  const name = useSelector( state => state);
  console.log(userForm);
  console.log(name.player);
  
  return (
    <Router>
      <Switch>
        <Route path='/' exact component={Landing} className='landing'/>
        <Route path='/game'exact component={Game}/>
      </Switch>
    </Router>
  );
}

export default App;
