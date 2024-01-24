import SqsPage from './sqs';
import S3Page from './s3';
import SnsPage from './sns';
import DynamoDBPage from './DynamoDb';
import SettingsPage from './settings';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCog } from '@fortawesome/free-solid-svg-icons';
import SqsIcon from '../icons/sqs';
import SnsIcon from '../icons/sns';
import DynamoDbIcon from '../icons/dynamoDb';
import S3Icon from '../icons/s3';
// import withErrorBoundary from '../ErrorBoundary';

const RenderPage = ({selectedMenu = 5}) => {
    return (
        <div className="container">
          <div className="containerHeader">
            {selectedMenu === 1 && <><S3Icon /><span>S3</span></>}
            {selectedMenu === 2 && <><SqsIcon /><span>SQS</span></>}
            {selectedMenu === 3 && <><SnsIcon /><span>SNS</span></>}
            {selectedMenu === 4 && <><DynamoDbIcon /><span>Dynamo DB</span></>}
            {selectedMenu === 5 && <><FontAwesomeIcon icon={faCog} size="2x" /><span>Settings</span></>}
          </div>
          <div className="pageContainer">
            {selectedMenu === 1 && <S3Page />}
            {selectedMenu === 2 && <SqsPage />}
            {selectedMenu === 3 && <SnsPage />}
            {selectedMenu === 4 && <DynamoDBPage />}
            {selectedMenu === 5 && <SettingsPage />}
          </div>
        </div>
    )

}

// export default withErrorBoundary(RenderPage);
export default RenderPage;