import React from 'react';
import useMyLibraryHook from './useMyLibraryHook';
import PageTitle from '../../components/Base/PageTitle';
import { ListItem } from '../../components';
import { CONSTANT_LIST_TYPE } from '../../utils/constant';

const MyLibrary = () => {
  const { isLoadingPlaylist, isLoadingTrack, dataPlaylist, dataTrack } =
    useMyLibraryHook();

  return (
    <div>
      <PageTitle title='YOUR PLAYLIST' />
      <ListItem
        loading={isLoadingPlaylist}
        data={dataPlaylist}
        type={CONSTANT_LIST_TYPE.Playlist}
      />
      <PageTitle title='YOUR TRACK' />
      <ListItem
        loading={isLoadingTrack}
        data={dataTrack}
        type={CONSTANT_LIST_TYPE.Track}
      />
    </div>
  );
};

export default MyLibrary;
