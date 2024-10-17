import React from 'react';

import { UserOutlined } from '@ant-design/icons';
import { Menu } from 'antd';
import HomeIcon from '../Icons/HomeIcon';
import LibraryIcon from '../Icons/LibraryIcon';
import AddTrackIcon from '../Icons/AddTrackIcon';
import CreatePlaylistIcon from '../Icons/CreatePlaylistIcon';
import { Link } from 'react-router-dom';

const AppMenu = () => {
  return (
    <Menu theme='dark' mode='inline' defaultSelectedKeys={['1']}>
      <Menu.Item key={'1'} icon={<HomeIcon />} label={'Home'}>
        <Link to='/'>Home</Link>
      </Menu.Item>
      <Menu.Item key={'2'} icon={<LibraryIcon />} label={'Your Library'}>
        <Link to='/myLibrary'> Your Library</Link>
      </Menu.Item>
      <Menu.Item key={'3'} icon={<AddTrackIcon />} label={'Add Track'}>
        <Link to='/upsertTrack'> Add Track</Link>
      </Menu.Item>
      <Menu.Item
        key={'4'}
        icon={<CreatePlaylistIcon />}
        label={'Create Playlist'}
      >
        <Link to='/upsertPlaylist'>Create Playlist</Link>
      </Menu.Item>
    </Menu>
  );
};

export default AppMenu;
