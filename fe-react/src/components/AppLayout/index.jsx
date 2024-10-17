import React, { useState } from 'react';

import { Layout, theme } from 'antd';
import { AppHeader, AppMenu, Logo } from '..';
import { Outlet } from 'react-router-dom';

const { Sider, Content } = Layout;

const AppLayout = () => {
  const [collapsed, setCollapsed] = useState(false);
  const {
    token: { colorBgContainer, borderRadiusLG },
  } = theme.useToken();
  return (
    <Layout className='h-[100vh]'>
      <Sider trigger={null} collapsible collapsed={collapsed}>
        <Logo collapsed={collapsed} />
        <AppMenu />
      </Sider>
      <Layout>
        <AppHeader
          colorBgContainer={colorBgContainer}
          collapsed={collapsed}
          setCollapsed={setCollapsed}
        />
        <Content
          style={{
            margin: '24px 16px',
            padding: 24,
            minHeight: 280,
            background: colorBgContainer,
            borderRadius: borderRadiusLG,
          }}
        >
          <Outlet />
        </Content>
      </Layout>
    </Layout>
  );
};

export default AppLayout;
