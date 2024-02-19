import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import { getConferenceDetails } from "../requests/api";

import axios from "axios";

const bookTicketsButton = <button>Book Now</button>;

function ConferenceDetails() {
  const [conferenceInfo, setConferenceInfo] = useState(null);

  useEffect(() => {
    getConferenceDetails().then((response) => {
      setConferenceInfo(response.data);
    });
  }, []);

  return (
    <>
      {conferenceInfo ? (
        <div>
          <h1>{conferenceInfo.ConferenceName}</h1>
          <p>Total Tickets: {conferenceInfo.TotalTickets}</p>
          <p>Tickets Left: {conferenceInfo.TicketsLeft}</p>
        </div>
      ) : (
        <p>Loading...</p>
      )}
      <Link to="/order">{bookTicketsButton}</Link>
    </>
  );
}

export default ConferenceDetails;
