import { useState, useEffect } from "react"
import { BrowserRouter } from "react-router-dom";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faCog } from "@fortawesome/free-solid-svg-icons";
import { useAppContext } from "./AppContext";
import Loader from "./components/loader";
import SqsIcon from "./icons/sqs";
import SnsIcon from "./icons/sns";
import DynamoDbIcon from "./icons/dynamoDb";
import S3Icon from "./icons/s3";
import MenuOption from "./components/menuoption";
import RenderPage from "./pages";
import ErrorBoundary from "./ErrorBoundary";
import "./App.scss";

function App() {
    const { isLoading } = useAppContext();
    const [showLoader, setShowLoader] = useState(false);

    useEffect(() => {
        let timer;

        if (isLoading) {
            timer = setTimeout(() => {
                setShowLoader(true);
            }, 300);
        } else {
            clearTimeout(timer);
            setShowLoader(false);
        }

        return () => clearTimeout(timer);
    }, [isLoading]);

    return (
        <>
            {showLoader && <Loader />}
            <BrowserRouter>
                <div className="App">
                    <header className="App-header">
                        AWS Localstack manager / V1.0006
                    </header>
                    <div className="main-wrapper">
                        <div className="left-menu">
                            <ul>
                                <MenuOption to="/aws/s3">
                                    S3
                                    <br />
                                    <S3Icon />
                                </MenuOption>
                                <MenuOption to="/aws/sqs">
                                    SQS
                                    <br />
                                    <SqsIcon />
                                </MenuOption>
                                <MenuOption to="/aws/sns">
                                    SNS
                                    <br />
                                    <SnsIcon />
                                </MenuOption>
                                <MenuOption to="/aws/dynamodb">
                                    Dynamo DB
                                    <br />
                                    <DynamoDbIcon />
                                </MenuOption>
                                <MenuOption to="/aws/listeners_sns">
                                    SNS
                                    <br /> Listeners
                                    <br />
                                    <SnsIcon />
                                </MenuOption>
                            </ul>
                            <ul>
                                <MenuOption to="/aws/settings">
                                    Settings
                                    <br />
                                    <FontAwesomeIcon icon={faCog} />
                                </MenuOption>
                            </ul>
                        </div>
                        <ErrorBoundary>
                            <RenderPage />
                        </ErrorBoundary>
                    </div>
                </div>
            </BrowserRouter>
        </>
    );
}

export default App;
