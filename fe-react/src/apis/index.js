import axios from 'axios';
import { message } from 'antd';

const auth = localStorage.getItem('fe-react-auth') || '{}';
const token = JSON.parse(auth).token;
const headerJSON = {
  Authorization: `Bearer ${token}`, // Attach the token to the header
  'Content-Type': 'application/json',
};
const headerFormData = {
  Authorization: `Bearer ${token}`, // Attach the token to the header
};

const handleError = (err) => {
  if (err?.response?.data?.Message) {
    message.error(err.response.data.Message);
  } else {
    message.error('Something went wrong! Status: ' + err.code);
  }
};
const Fetch = {
  GET: async (url) => {
    return await axios
      .get(url, { headers: headerJSON })
      .then((resp) => {
        const data = resp?.data;
        return data;
      })
      .catch((err) => {
        handleError(err);
        return err.response || {};
      });
  },
  POST: async (url, data) => {
    return await axios
      .post(url, data, { headers: headerJSON })
      .then((resp) => {
        const data = resp?.data;
        return data;
      })
      .catch((err) => {
        handleError(err);
        return err.response || {};
      });
  },
  POSTFORM: async (url, data) => {
    return await axios
      .post(url, data, { headers: headerFormData })
      .then((resp) => {
        const data = resp?.data;
        return data;
      })
      .catch((err) => {
        handleError(err);
        return err.response || {};
      });
  },
  PUT: async (url, data) => {
    return await axios
      .put(url, data, { headers: headerJSON })
      .then((resp) => {
        const data = resp?.data;
        return data;
      })
      .catch((err) => {
        handleError(err);
        return err.response || {};
      });
  },
  PUTFORM: async (url, data) => {
    return await axios
      .put(url, data, { headers: headerFormData })
      .then((resp) => {
        const data = resp?.data;
        return data;
      })
      .catch((err) => {
        handleError(err);
        return err.response || {};
      });
  },
  DELETE: async (url, data) => {
    return await axios
      .delete(url, { data, headers: headerJSON })
      .then((resp) => {
        const data = resp?.data;
        return data;
      })
      .catch((err) => {
        handleError(err);
        return err.response || {};
      });
  },
};

export default Fetch;
