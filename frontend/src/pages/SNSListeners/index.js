import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import InteractiveTable from "../../components/interactiveTable";
import { useAppContext } from "../../AppContext";
import SaveBox from "../../components/savebox";
import FilterBox from "../../components/filterBox";
import Button from "../../components/button";

const SNSListenersPage = () => {
    const { get, del } = useAppContext();
    const [data, setData] = useState([]);
    const [filter, setFilter] = useState("");
    const [error, setError] = useState("");
    const [modalVisible, setModalVisible] = useState(false);
    const navigate = useNavigate();

    // APIS
    const listeners = async () => {
        return get("/api/sns/listeners");
    };

    const addListener = async (port) => {
        return get(`/api/sns/listener/${port}`);
    };

    const delListener = async (port) => {
        return del(`/api/sns/listener/${port}`);
    };
    // END APIS

    const onEvent = (e) => {
        if (e.name === "Delete") {
            delListener(e.i.port)
                .then(() => listeners().then((data) => setData(data)))
                .catch((err) => setError(err.message ?? "Error fetching data"));
        }

        if (e.name === "View") {
            navigate(`/aws/listeners_sns/${e.i.port}`);
        }
    };

    useEffect(() => {
        listeners()
            .then((data) => setData(data))
            .catch((err) => setError(err.message ?? "Error fetching data"));
    }, []);
    return (
        <div>
            <SaveBox
                isOpen={modalVisible}
                onClose={() => setModalVisible(false)}
                numOnly={true}
                title="New Port:"
                onSubmit={(message) => {
                    setModalVisible(false);
                    addListener(message)
                        .then((_) =>
                            listeners()
                                .then((data) => setData(data))
                                .catch((err) => setError(err.message ?? err))
                        )
                        .catch((err) => setError(err.message ?? err));
                }}
            />
            <Button
                label="Add new listener"
                margin={6}
                onClick={() => {
                    setModalVisible(true);
                }}
            />
            {error !== "" && <div className="errorLine">{error}</div>}
            <FilterBox onSubmit={(text) => setFilter(text)} />
            {data && data.ports && (
                <InteractiveTable
                    structInfo={{
                        initialSort: {
                            field: "port",
                            asc: true,
                        },
                        filterField: "port",
                        columns: [
                            {
                                field: "port",
                                title: "Port",
                                clickable: false,
                            },
                            {
                                field: "info",
                                title: "Status",
                                clickable: false,
                            },
                        ],
                        events: ["Delete", "View"],
                    }}
                    data={data.ports}
                    filter={filter}
                    onEvent={(e) => onEvent(e)}
                />
            )}
        </div>
    );
};

export default SNSListenersPage;
