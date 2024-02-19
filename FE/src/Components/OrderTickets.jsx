import React, { useState } from "react";

function OrderTickets() {
  const [firstNameValue, setFirstNameValue] = useState("");
  const [lastNameValue, setLastNameValue] = useState("");
  const [emailValue, setEmailValue] = useState("");
  const [ticketValue, setTicketValue] = useState(0);

  function handleSubmit(event) {
    event.preventDefault();
  }
  return (
    <form onSubmit={handleSubmit}>
      <label>
        First name:
        <input
          type="text"
          value={firstNameValue}
          onChange={(e) => setFirstNameValue(e.target.value)}
        />
      </label>
      <label>
        Last name:
        <input
          type="text"
          value={lastNameValue}
          onChange={(e) => setLastNameValue(e.target.value)}
        />
      </label>
      <label>
        Email:
        <input
          type="text"
          value={emailValue}
          onChange={(e) => setEmailValue(e.target.value)}
        />
      </label>
      <label>
        How many tickets:
        <input
          type="text"
          value={ticketValue}
          onChange={(e) => setTicketValue(e.target.value)}
        />
      </label>
      <button type="submit">Confirm your booking</button>
    </form>
  );
}

export default OrderTickets;
