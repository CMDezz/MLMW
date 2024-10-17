import React from 'react';
import PageTitle from '../../components/Base/PageTitle';
import useDashBoardHook from './useDashboardHook';
import { ListItem } from '../../components';
import { CONSTANT_LIST_TYPE } from '../../utils/constant';

const Dashboard = () => {
  const { isLoadingPlaylist, isLoadingTrack, dataPlaylist, dataTrack } =
    useDashBoardHook();

  return (
    <div>
      <PageTitle title='EXPLORE MORE PUBLIC PLAYLIST' />
      <ListItem
        loading={isLoadingPlaylist}
        data={dataPlaylist}
        type={CONSTANT_LIST_TYPE.Playlist}
      />
      <PageTitle title='EXPLORE MORE PUBLIC TRACKS' />
      <ListItem
        loading={isLoadingTrack}
        data={dataTrack}
        type={CONSTANT_LIST_TYPE.Track}
      />
    </div>
  );
};

export default Dashboard;
