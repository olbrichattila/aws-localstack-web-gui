import './App.scss';
import SqsPage from './pages/sqs';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCog } from '@fortawesome/free-solid-svg-icons';
import SqsIcon from './icons/sqs';
import SnsIcon from './icons/sns';
import DynamoDbIcon from './icons/dynamoDb';
import S3Icon from './icons/s3';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        AWS Localstack manager / V1.0001
      </header>
      <div className="main-wrapper">
        <div className="left-menu">
          <ul>
            <li>S3<br /><S3Icon /></li>
            <li className='active'>SQS<br /><SqsIcon /></li>
            <li>SNS<br /><SnsIcon /></li>
            <li>Dynamo DB<br /><DynamoDbIcon /></li>
          </ul>
          <ul>
            <li>Settings<br /><FontAwesomeIcon icon={faCog} /></li>
          </ul>
          
        </div>
        <div className="container">
          <SqsPage />
        </div>
      
      </div>
      
      
      
    </div>
  );
}

export default App;
