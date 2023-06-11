import logo from './logo.svg';
import './App.css';
import axios from 'axios'

function App() {

  const onConnectAPI = async() => {
    const res = await axios.get('http://localhost:8080/test')
    console.log(res)
  }

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
        <div>
          <button onClick={onConnectAPI}>
            API通信
          </button>
        </div>
      </header>
    </div>
  );
}

export default App;
