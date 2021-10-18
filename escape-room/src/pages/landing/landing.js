import { BrowserRouter as Router, Link } from "react-router-dom";



const Landing = () => {
    return (
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
          <Link to="/game">
            <button type='button'>Play!</button>
          </Link>
        </main>
    )
  };
  
  export default Landing;