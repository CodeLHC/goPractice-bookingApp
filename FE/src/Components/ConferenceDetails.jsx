import React, { useState, useEffect } from "react";
import { getConferenceDetails } from "../requests/api";

import axios from "axios";

function ConferenceDetails() {
  const [conferenceInfo, setConferenceInfo] = useState(null);

  useEffect(() => {
    getConferenceDetails().then((response) => {
      setConferenceInfo(response.data);
    });
  }, []);

  console.log(conferenceInfo);

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
    </>
  );
}

export default ConferenceDetails;
