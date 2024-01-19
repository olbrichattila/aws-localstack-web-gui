import './App.css';
import Sqs from './components/sqs';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        AWS Localstack manager
      </header>
      <Sqs />
      
    </div>
  );
}

export default App;
