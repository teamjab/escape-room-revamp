
import Landing from './pages/landing/landing.js';
import Game from './pages/game/game.js';
import D1 from './pages/game/doors-riddles/door-one/d1.js';
import D2 from './pages/game/doors-riddles/door-two/d2.js';
import D3 from './pages/game/doors-riddles/door-three/d3.js';
import D4 from './pages/game/doors-riddles/door-four/d4.js';

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
        <Route path='/d1'exact component={D1}/>
        <Route path='/d2'exact component={D2}/>
        <Route path='/d3'exact component={D3}/>
        <Route path='/d4'exact component={D4}/>
      </Switch>
    </Router>
  );
}

export default App;
