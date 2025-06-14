import React, { createContext, useContext, useState } from "react";

// 1. Create the context
const AppContext = createContext();

// 2. Create a provider component
export const AppProvider = ({ children }) => {
    const [isLoading, setIsLoading] = useState(false);

    const get = async (path) => {
        return request("GET", path);
    };

    const post = async (path, payload) => {
        return request("POST", path, payload);
    };

    const del = async (path, payload) => {
        return request("DELETE", path, payload);
    };

    const request = async (type, path, payload = null) => {
        const options = {
            method: type,
            headers: {
                "Content-Type": "application/json",
            },
        };

        if (payload) {
            options.body = JSON.stringify(payload);
        }

        try {
            setIsLoading(true);
            const response = await fetch(
                `${process.env.REACT_APP_API_URL}${path}`,
                options
            );

            const rawText = await response.text(); // Read body once

            if (!response.ok) {
                let message = "Whoops, something went wrong";

                try {
                    const json = JSON.parse(rawText); // Try parsing it manually

                    if (json.error) {
                        message = json.error;
                    } else if (json.errors) {
                        message = json.errors;
                    }
                } catch {
                    message = rawText; // Use plain text as fallback
                }

                throw new Error(message);
            }

            // If ok, parse the original text as JSON
            const data = JSON.parse(rawText);
            setIsLoading(false);
            return data;
        } catch (error) {
            setIsLoading(false);
            throw error;
        }
    };

    const value = {
        isLoading,
        get,
        post,
        del,
    };

    return <AppContext.Provider value={value}>{children}</AppContext.Provider>;
};

// 3. Custom hook for consuming the context
export const useAppContext = () => {
    const context = useContext(AppContext);
    if (!context) {
        throw new Error("useAppContext must be used within an AppProvider");
    }
    return context;
};
