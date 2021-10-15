import userForm from "./components/userForm.js";
import "./App.css";

function App() {
  console.log(userForm);
  // simple fetch
  fetch('http://localhost:8080/')
  .then(response => response.json())
  .then(data => console.log(data))
  .catch(err => console.error(err));
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
