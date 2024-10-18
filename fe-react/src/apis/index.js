import axios from 'axios';
import { message } from 'antd';

const handleError = (err) => {
  if (err?.response?.data?.Message) {
    message.error(err.response.data.Message);
  } else {
    message.error('Something went wrong! Status: ' + err.code);
  }
};

const getHeader = (isFormData = false) => {
  const auth = localStorage.getItem('fe-react-auth') || '{}';
  const token = JSON.parse(auth).token;

  const header = {
    Authorization: `Bearer ${token}`, // Attach the token to the header
  };
  if (!isFormData) header['Content-Type'] = 'application/json';
  return header;
};

const Fetch = {
  GET: async (url) => {
    return await axios
      .get(url, { headers: getHeader() })
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
      .post(url, data, { headers: getHeader() })
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
      .post(url, data, { headers: getHeader(true) })
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
      .put(url, data, { headers: getHeader() })
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
      .put(url, data, { headers: getHeader(true) })
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
      .delete(url, { data, headers: getHeader() })
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
