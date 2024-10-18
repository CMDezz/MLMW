import React, { useContext } from 'react';
import { MenuFoldOutlined, MenuUnfoldOutlined } from '@ant-design/icons';
import { Button, Dropdown, Space, Layout, Input } from 'antd';
import { AuthContext } from '../../contexts/AuthProvider';
import { redirect, useNavigate } from 'react-router-dom';
import { DownOutlined, SmileOutlined } from '@ant-design/icons';
const AppHeader = (props) => {
  const { colorBgContainer, collapsed, setCollapsed } = props;
  const { auth, updateAuth } = useContext(AuthContext);
  const navigate = useNavigate();
  const items = [
    {
      key: '4',
      danger: true,
      label: 'Logout',
      onClick: () => {
        updateAuth({});
        navigate('/');
      },
    },
  ];

  const onSearch = (value) => {
    navigate('/?keyword=' + value);
  };

  return (
    <Layout.Header className='px-4' style={{ background: colorBgContainer }}>
      <div className='flex justify-between items-center h-full'>
        <Button
          type='text'
          icon={collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
          onClick={() => setCollapsed(!collapsed)}
          style={{
            fontSize: '16px',
          }}
        />
        <div className='flex justify-between items-center h-full'>
          <Input.Search
            placeholder='Find playlist/ tracks by artis,name'
            onSearch={onSearch}
            style={{ width: 350 }}
            maxLength={50}
          />
        </div>
        <div>
          {!auth.token ? (
            <Button onClick={() => navigate('/login')}>Login</Button>
          ) : (
            <Dropdown menu={{ items }}>
              <a
                className='font-semibold text-lg'
                onClick={(e) => e.preventDefault()}
              >
                {auth.username}
              </a>
            </Dropdown>
          )}
        </div>
      </div>
    </Layout.Header>
  );
};

export default AppHeader;
