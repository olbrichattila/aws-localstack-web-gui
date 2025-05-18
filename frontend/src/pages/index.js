import { Route, Routes } from 'react-router-dom';
import SqsPage from './sqs';
import S3Bucket from './s3/bucket';
import S3BucketContent from './s3/content';
import TopicPage from './sns/topic';
import SubscriptionPage from './sns/subscription';
import SNSListenersPage from './SNSListeners';
import SNSListPage from './SNSListeners/list';
import DynamoDbTables from './DynamoDb/tables';
import DynamoDbContent from './DynamoDb/content';
import SettingsPage from './settings';
import SqsIcon from '../icons/sqs';
import SnsIcon from '../icons/sns';
import DynamoDbIcon from '../icons/dynamoDb';
import S3Icon from '../icons/s3';
import SettingsIcon from '../icons/settings';
import Layout from './layout';
// import withErrorBoundary from '../ErrorBoundary';

const RenderPage = () => {
  return (
    <div className="container">
      <Routes>
        <Route path='/s3' element={<Layout Component={S3Bucket} Icon={S3Icon} title="S3" />} />
        <Route path='/s3/:bucketName' element={<Layout Component={S3BucketContent} Icon={S3Icon} title="S3" />} />
        <Route path='/sqs' element={<Layout Component={SqsPage} Icon={SqsIcon} title="SQS" />} />
        <Route path='/sns' element={<Layout Component={TopicPage} Icon={SnsIcon} title="SNS" />} />
        <Route path='/sns/:topicArn' element={<Layout Component={SubscriptionPage} Icon={SnsIcon} title="SNS Subscriptions" />} />
        <Route path='/dynamodb' element={<Layout Component={DynamoDbTables} Icon={DynamoDbIcon} title="DynamoDB" />} />
        <Route path='/dynamodb/:tableName' element={<Layout Component={DynamoDbContent} Icon={DynamoDbIcon} title="DynamoDB content" />} />
        <Route path='/listeners_sns' element={<Layout Component={SNSListenersPage} Icon={SnsIcon} title="SNS Listeners" />} />
        <Route path='/listeners_sns/:portNum' element={<Layout Component={SNSListPage} Icon={SnsIcon} title="SNS Listeners" />} />
        <Route path='/settings' element={<Layout Component={SettingsPage} Icon={SettingsIcon} title="Settings" />} />
      </Routes>
    </div>
  )
}

// export default withErrorBoundary(RenderPage);
export default RenderPage;
