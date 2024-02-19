import "./App.css";
import React from "react";
import { Route, Routes, BrowserRouter as Router } from "react-router-dom";
import ConferenceDetails from "./Components/ConferenceDetails";
import OrderTickets from "./Components/OrderTickets";
import OrderConfirmation from "./Components/OrderConfirmation";

function App() {
  return (
    <>
      <Router>
        <Routes>
          <Route path="/" element={<ConferenceDetails />} />
          <Route path="/order" element={<OrderTickets />} />
          <Route path="/confirmation" element={<OrderConfirmation />} />
        </Routes>
      </Router>
    </>
  );
}

export default App;
