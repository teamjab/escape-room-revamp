import userForm from "./components/userForm.js";
import {useSelector} from 'react-redux';
import './styling/App.scss';




function App() {
  const name = useSelector( state => state);
  console.log(userForm);
  console.log(name.player);
  
  return (
    <div className="App">
      <header className="App-header"></header>
      <main>
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
          <input placeholder="enter username"></input>
        </form>
        <button type="submit">Submit</button>
      </main>
      <footer>
        <h4>Team Jab</h4>
      </footer>
    </div>
  );
}

export default App;
