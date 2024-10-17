import React, { createContext, useEffect, useState } from 'react';

const AuthContext = createContext();

const AuthProvider = ({ children }) => {
  const [auth, setAuth] = useState({
    token: '',
    expiredAt: '',
    username: '',
    email: '',
  });

  useEffect(() => {
    let authStorage = localStorage.getItem('fe-react-auth') || '{}';
    authStorage = JSON.parse(authStorage);

    //check expired
    if (isTokenExpired(authStorage)) {
      updateAuth({});
    } else {
      setAuth({ ...authStorage });
    }
  }, []);

  const isTokenExpired = (auth = {}) => {
    if (!auth.token) return true;
    let now = new Date();
    let expiredTime = new Date(auth.expiredAt);
    if (now > expiredTime) return true;
    return false;
  };

  const updateAuth = (data) => {
    setAuth(data);
    localStorage.setItem('fe-react-auth', JSON.stringify(data));
  };

  return (
    <AuthContext.Provider value={{ auth, updateAuth, isTokenExpired }}>
      {children}
    </AuthContext.Provider>
  );
};

export { AuthContext, AuthProvider };
