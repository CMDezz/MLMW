import { useContext, useState } from 'react';
import UserApis from '../../apis/UserApis';
import { AuthContext } from '../../contexts/AuthProvider';
import { message } from 'antd';
import { useNavigate } from 'react-router-dom';

const useLoginHook = () => {
  const { auth, updateAuth } = useContext(AuthContext);

  const natigate = useNavigate();

  const [isOpenModalRegister, setIsOpenModalRegister] = useState(false);

  const onLogin = async (loginFormData) => {
    const res = await UserApis.Login(loginFormData);
    if (res.Data) {
      message.success(res.Message);
      updateAuth({
        token: res.Data.token,
        email: res.Data.email,
        username: res.Data.username,
        expiredAt: res.Data.expired_at,
      });
      natigate('/');
    }
  };

  const onRegister = async (registerFormData) => {
    const _data = {
      username: registerFormData.username,
      password: registerFormData.password,
      full_name: registerFormData.fullname,
      email: registerFormData.email,
    };
    const res = await UserApis.Create(_data);

    if (res.Data) {
      message.success(res.Message);
      setIsOpenModalRegister(false);
    }
  };

  return {
    onLogin,
    onRegister,
    isOpenModalRegister,
    setIsOpenModalRegister,
  };
};

export default useLoginHook;
