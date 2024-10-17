import { Spin } from 'antd';
import React from 'react';
import PlayIcon from '../Icons/PlayIcon';
import { CONSTANT_LIST_TYPE } from '../../utils/constant';
import { Link } from 'react-router-dom';
import Empty from '../Empty';

const ItemCoverImage = ({ image }) => {
  return (
    <div className='w-[275px]'>
      <div className='relative w-full pt-[56.25%] overflow-hidden'>
        <div className='bg-black hover:cursor-pointer bg-opacity-0 hover:bg-opacity-50 transition-all hover:opacity-100 opacity-0 duration-500 flex item-center justify-center z-10 top-0 left-0 w-full h-full absolute object-fit-cover '>
          <PlayIcon />
        </div>
        <img
          src={image}
          className='top-0 left-0 w-full h-full absolute object-cover'
        />
      </div>
    </div>
  );
};

const ItemPlaylist = ({ item }) => {
  return (
    <div>
      <ItemCoverImage image={item.cover_image} />
      <div>
        <Link to='#'>
          <h5 className='font-semibold text-lg mb-2'>{item.playlist_name}</h5>
        </Link>
      </div>
    </div>
  );
};

const ItemTrack = ({ item }) => {
  return (
    <div>
      <ItemCoverImage image={item.cover_image} />
      <div>
        <Link to={'upsertTrack/' + item.id}>
          <h5 className='font-semibold text-lg mb-1'>{item.title}</h5>
        </Link>
        <p>{item.artist}</p>
      </div>
    </div>
  );
};

const ListItem = (props) => {
  const { data = [], loading = false, type } = props;

  const renderItem = (data = []) => {
    if (type === CONSTANT_LIST_TYPE.Playlist) {
      return data.map((it, k) => {
        return <ItemPlaylist item={it} key={it?.id || k} />;
      });
    }
    return data.map((it, k) => {
      return <ItemTrack item={it} key={it?.id || k} />;
    });
  };
  return (
    <Spin spinning={loading}>
      <div className='w-full min-h-[15vh] flex gap-3 flex-wrap'>
        {data.length > 0 ? renderItem(data) : <Empty />}
      </div>
    </Spin>
  );
};

export default ListItem;
