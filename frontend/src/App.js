import { useState } from 'react';
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
    <div className="App">
      <header className="App-header">
        AWS Localstack manager / V1.0001
      </header>
      <div className="main-wrapper">
        <div className="left-menu">
          <ul>
            <MenuOption active={selectedMenu === 1} onClick={() => setSelectedMenu(1)}>
              S3<br /><S3Icon />
            </MenuOption>
            <MenuOption active={selectedMenu === 2} onClick={() => setSelectedMenu(2)}>
              SQS<br /><SqsIcon />
            </MenuOption>
            <MenuOption active={selectedMenu === 3} onClick={() => setSelectedMenu(3)}>
              SNS<br /><SnsIcon />
            </MenuOption>
            <MenuOption active={selectedMenu === 4} onClick={() => setSelectedMenu(4)}>
              Dynamo DB<br /><DynamoDbIcon />
            </MenuOption>
          </ul>
          <ul>
            <MenuOption active={selectedMenu === 5} onClick={() => setSelectedMenu(5)}>
              Settings<br /><FontAwesomeIcon icon={faCog} />
            </MenuOption>
          </ul>

        </div>
        <ErrorBoundary>
          <RenderPage selectedMenu={selectedMenu} />
        </ErrorBoundary>
      </div>
    </div>
  );
}

export default App;
