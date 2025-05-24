import { useEffect, useState } from "react";
import { useAppContext } from '../../AppContext';
import Button from "../../components/button";
// import { getSettings, saveSettings } from "../../api/settings";
import "./index.scss";

const initialFormState = {
    key: "",
    secret: "",
    endpoint: "",
    region: "",
};

const SettingsPage = () => {
    const { get, post } = useAppContext();
    const [error, setError] = useState("");
    const [formState, setFormState] = useState(initialFormState);
    const [message, setMessage] = useState("");

    const getSettings = async () => {
        return get("/api/settings");
    };

    const saveSettings = async (request) => {
        return post("/api/settings", request);
    };

    const save = () => {
        saveSettings(formState).then((settings) => {
            setMessage("Settings saved successfully");
            assignSettingsToForm(settings);
        });
    };

    const assignSettingsToForm = (settings) => {
        setFormState({
            key: settings.credentials.key ?? "",
            secret: settings.credentials.secret ?? "",
            endpoint: settings.endpoint ?? "",
            region: settings.region ?? "",
        });
    };

    useEffect(() => {
        getSettings()
            .then((settings) => assignSettingsToForm(settings))
            .catch((err) => setError("failed to get settings"));
    }, []);

    useEffect(() => {
        let timeoutId = -1;
        if (error !== "") {
            timeoutId = setTimeout(() => {
                setError("");
            }, 6000);
        }

        return () => {
            if (timeoutId !== -1) {
                clearTimeout(timeoutId);
            }
        };
    }, [error]);

    useEffect(() => {
        let messageDisplayTimeoutId = -1;
        if (message !== "") {
            messageDisplayTimeoutId = setTimeout(() => {
                setMessage("");
            }, 2000);
        }

        return () => {
            if (messageDisplayTimeoutId !== -1) {
                clearTimeout(messageDisplayTimeoutId);
            }
        };
    }, [message]);

    return (
        <div className="settings">
            {error !== "" && <div className="errorLine">{error}</div>}
            {message !== "" && <div className="messageLine">{message}</div>}

            <label>
                Region:
                <input
                    type="text"
                    value={formState.region}
                    placeholder="us-east-1"
                    onChange={(e) =>
                        setFormState({ ...formState, region: e.target.value })
                    }
                />
            </label>
            <label>
                Endpoint:
                <input
                    type="text"
                    value={formState.endpoint}
                    placeholder="http://localhost:4566"
                    onChange={(e) =>
                        setFormState({ ...formState, endpoint: e.target.value })
                    }
                />
            </label>
            <label>
                Key:
                <input
                    type="text"
                    value={formState.key}
                    placeholder="your-access-key-id"
                    onChange={(e) =>
                        setFormState({ ...formState, key: e.target.value })
                    }
                />
            </label>
            <label>
                Secret:
                <input
                    type="text"
                    value={formState.secret}
                    placeholder="your-secret-access-key"
                    onChange={(e) =>
                        setFormState({ ...formState, secret: e.target.value })
                    }
                />
            </label>
            <label>
                <div></div>
                <Button label="save" onClick={() => save()} />
            </label>
        </div>
    );
};

export default SettingsPage;
