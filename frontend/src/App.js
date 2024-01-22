import { useState } from 'react';
import './App.scss';
import SqsPage from './pages/sqs';
import S3Page from './pages/s3';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCog } from '@fortawesome/free-solid-svg-icons';
import SqsIcon from './icons/sqs';
import SnsIcon from './icons/sns';
import DynamoDbIcon from './icons/dynamoDb';
import S3Icon from './icons/s3';
import MenuOption from './components/menuoption';

function App() {
  const [selectedMenu, setSelectedMenu] = useState(1);

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
            <li>Settings<br /><FontAwesomeIcon icon={faCog} /></li>
          </ul>

        </div>
        <div className="container">
          <div className="containerHeader">
            {selectedMenu === 1 && <><S3Icon /><span>S3</span></>}
            {selectedMenu === 2 && <><SqsIcon /><span>SQS</span></>}
            {selectedMenu === 3 && <><SnsIcon /><span>SNS</span></>}
            {selectedMenu === 4 && <><DynamoDbIcon /><span>Dynamo DB</span></>}
          </div>
          <div className="pageContainer">
            {selectedMenu === 1 && <S3Page />}
            {selectedMenu === 2 && <SqsPage />}
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
