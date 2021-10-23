
import Landing from './pages/landing/landing.js';
import Game from './pages/game/game.js';
import { BrowserRouter as Router, Switch, Route} from "react-router-dom";
import './styling/App.scss';
import store from './store/store.js';





function App() {
console.log('initial State', store.getState());
  
  return (
    <Router>
      <Switch>
        <Route path='/' exact component={Landing}/>
        <Route path='/game'exact component={Game}/>
      </Switch>
    </Router>
  );
}

export default App;
