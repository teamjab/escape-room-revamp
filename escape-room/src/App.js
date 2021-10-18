import userForm from "./components/userForm.js";
import {useSelector} from 'react-redux';
import { BrowserRouter as Router, Switch, Route} from "react-router-dom";
import './styling/App.scss';
import Landing from './pages/landing/landing.js';



function App() {
  const name = useSelector( state => state);
  console.log(userForm);
  console.log(name.player);
  
  return (
    <Router>
      <div className="App">
        <header className="App-header"></header>
        <main>
          <Switch>
            <Route exact path='/'>
              <Landing/>
            </Route>
          </Switch>
        </main>
        <footer>
          <h4>Team Jab</h4>
        </footer>
      </div>
    </Router>
  );
}

export default App;
