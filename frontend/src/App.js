import { useState } from 'react';
import { BrowserRouter, useLocation } from 'react-router-dom';
import './App.scss';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCog } from '@fortawesome/free-solid-svg-icons';
import SqsIcon from './icons/sqs';
import SnsIcon from './icons/sns';
import DynamoDbIcon from './icons/dynamoDb';
import S3Icon from './icons/s3';
import MenuOption from './components/menuoption';
import RenderPage from './pages'
import ErrorBoundary from './ErrorBoundary';

function App() {
  const [selectedMenu, setSelectedMenu] = useState(5);
  

  return (
    <BrowserRouter>
    <div className="App">
      <header className="App-header">
        AWS Localstack manager / V1.0001
      </header>
      <div className="main-wrapper">
        <div className="left-menu">
          <ul>
            <MenuOption to="/s3" active={selectedMenu === 1} >
              S3<br /><S3Icon />
            </MenuOption>
            <MenuOption to="/sqs" active={selectedMenu === 2} >
              SQS<br /><SqsIcon />
            </MenuOption>
            <MenuOption to="/sns"active={selectedMenu === 3} >
              SNS<br /><SnsIcon />
            </MenuOption>
            <MenuOption to="/dynamodb" active={selectedMenu === 4} >
              Dynamo DB<br /><DynamoDbIcon />
            </MenuOption>
          </ul>
          <ul>
            <MenuOption to="/settings" active={selectedMenu === 5} >
              Settings<br /><FontAwesomeIcon icon={faCog} />
            </MenuOption>
          </ul>
        </div>
        <ErrorBoundary>
          <RenderPage selectedMenu={selectedMenu} />
        </ErrorBoundary>
      </div>
    </div>
    </BrowserRouter>
  );
}

export default App;
