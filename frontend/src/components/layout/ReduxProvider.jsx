"use client"; 
import { persistor, store } from "/src/stores/redux";
import { ReactNode } from "react";
import { Provider } from "react-redux";
import { PersistGate } from "redux-persist/integration/react";

const ReduxProvider = ({ children }) => {
  return (
    <Provider store={store}>
        <PersistGate loading={null} persistor={persistor}>
            {children}
        </PersistGate>  
    </Provider>
  );
};

export default ReduxProvider;
