import React, { useState, useEffect } from "react";
import { getBookings } from "../requests/api";

function OrderConfirmation() {
  const [bookings, setBookings] = useState(null);

  useEffect(() => {
    getBookings().then((response) => {
      setBookings(response.data);
    });
  }, []);

  return (
    <>
      <h1>Whoop whoop</h1>
      <p>You are booked on for our event</p>
      <p>You`ll be joining...</p>
      <div>
        {bookings &&
          bookings.map((booking) => {
            return <p key={booking.FirstName}>{booking.FirstName}</p>;
          })}
        {!bookings && <p>Loading...</p>}
      </div>
    </>
  );
}

export default OrderConfirmation;
