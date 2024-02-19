import "./App.css";
import React from "react";
import { Route, Routes, BrowserRouter as Router } from "react-router-dom";
import ConferenceDetails from "./Components/ConferenceDetails";
import OrderTickets from "./Components/OrderTickets";

function App() {
  return (
    <>
      <Router>
        <Routes>
          <Route path="/" element={<ConferenceDetails />} />
          <Route path="/order" element={<OrderTickets />} />
        </Routes>
      </Router>
    </>
  );
}

export default App;
