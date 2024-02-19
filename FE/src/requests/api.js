import axios from "axios";

export const getConferenceDetails = async () => {
  try {
    return axios.get("http://localhost:8000/conference");
  } catch (error) {
    console.error("error getting conference details", error);
  }
};

export const getBookings = async () => {
  try {
    return axios.get("http://localhost:8000/bookings");
  } catch (error) {
    console.error("errors getting all booking details", error);
  }
};
