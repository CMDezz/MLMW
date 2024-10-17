import React from 'react';
import PageTitle from '../../components/Base/PageTitle';
import { Button, Checkbox, Form, Input } from 'antd';
import ModalRegister from './modalRegister';
import useLoginHook from './useLoginHook';

const Login = () => {
  const { onLogin, setIsOpenModalRegister, onRegister, isOpenModalRegister } =
    useLoginHook();
  const onFinish = (values) => {
    console.log('Success:2', values);
    console.log('process.env.REACT_HOST_API; ', process.env.REACT_APP_HOST_API);

    onLogin(values);
  };

  const onFinishFailed = (errorInfo) => {
    console.log('Failed:', errorInfo);
  };

  return (
    <div className=''>
      <PageTitle title={'Login'} />

      <div className='w-full flex justify-center items-center'>
        <Form
          name='login'
          labelCol={{ span: 8 }}
          wrapperCol={{ span: 16 }}
          style={{ maxWidth: 600 }}
          className='w-[70%]'
          initialValues={{ remember: true }}
          onFinish={onFinish}
          onFinishFailed={onFinishFailed}
          autoComplete='off'
        >
          <Form.Item
            label='Username'
            name='username'
            rules={[{ required: true, message: 'Please input your username!' }]}
          >
            <Input />
          </Form.Item>

          <Form.Item
            label='Password'
            name='password'
            rules={[
              {
                required: true,
                message: 'Please input your password!',
              },
            ]}
          >
            <Input.Password />
          </Form.Item>

          <Form.Item className='w-full' wrapperCol={{ offset: 8, span: 16 }}>
            <div className='flex justify-between'>
              <Button onClick={() => setIsOpenModalRegister(true)} type='link'>
                Don't have an account? Register now!
              </Button>
              <Button type='primary' htmlType='submit'>
                Submit
              </Button>
            </div>
          </Form.Item>
        </Form>
      </div>
      <ModalRegister
        open={isOpenModalRegister}
        onOk={onRegister}
        onCancel={() => setIsOpenModalRegister(false)}
      />
    </div>
  );
};

export default Login;
