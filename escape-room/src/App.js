import userForm from "./components/userForm.js";
import {useSelector} from 'react-redux';


function App() {
  const name = useSelector( state => state);
  console.log(userForm);
  console.log(name.player);
  
  return (
    <div className="App">
      <header className="App-header"></header>
      <main>
        <h1>Escape-Room</h1>
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
