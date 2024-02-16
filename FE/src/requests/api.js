import axios from "axios";

export const getConferenceDetails = async () => {
  try {
    return axios.get("http://localhost:8000/conference");
  } catch (error) {
    console.error("error getting details", error);
  }
};
