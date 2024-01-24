import React, { useEffect, useState } from "react";
import Button from '../../components/button';
import { getSettings, saveSettings } from "../../api/settings";
import './index.scss';

const initialFormState = {
    key: '',
    secret: '',
    endpoint: '',
    region: '',
}

const SettingsPage = () => {
    const [formState, setFormState] = useState(initialFormState);

    const save = () => {
        saveSettings(formState).then(settings => assignSettingsToForm(settings));
    }

    const assignSettingsToForm = (settings) => {
        setFormState({
            key: settings.credentials.key ?? '',
            secret: settings.credentials.secret ?? '',
            endpoint: settings.endpoint ?? '',
            region: settings.region ?? '',
        });
    }

    useEffect(() => {
        getSettings().then(settings => assignSettingsToForm(settings));
    }, []);


    return (
        <div className="settings">
            <label>
                Region
                <input
                    type='text'
                    value={formState.region}
                    onChange={(e) => setFormState({...formState, region: e.target.value})}
                />
            </label>
            <label>
                Endpoint
                <input
                    type='text'
                    value={formState.endpoint}
                    onChange={(e) => setFormState({...formState, endpoint: e.target.value})}
                />
            </label>
            <label>
                Key
                <input
                    type='text'
                    value={formState.key}
                    onChange={(e) => setFormState({...formState, key: e.target.value})}
                />
            </label>
            <label>
                secret
                <input
                    type='text'
                    value={formState.secret}
                    onChange={(e) => setFormState({...formState, secret: e.target.value})}
                />
            </label>
            <label>
                <div></div>
                <Button label="save" onClick={() => save()} />
            </label>

        </div>

    );
}

export default SettingsPage;
