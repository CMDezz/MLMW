import { Button, message, Spin, Tooltip } from 'antd';
import React, { useContext, useEffect, useState } from 'react';
import PlayIcon from '../Icons/PlayIcon';
import { CONSTANT_LIST_TYPE } from '../../utils/constant';
import { Link, useNavigate } from 'react-router-dom';
import Empty from '../Empty';
import CreatePlaylistIcon from '../Icons/CreatePlaylistIcon';
import ModalTrackPlaylist from './modalTrackPlaylist';
import { AuthContext } from '../../contexts/AuthProvider';
import { MediaContext } from '../../contexts/MediaProvider';

const ItemCoverImage = ({ image, id }) => {
  const { onChange } = useContext(MediaContext);
  const navigate = useNavigate();
  return (
    <div className='' onClick={() => navigate('/playlistDetail/' + id)}>
      <div className='relative w-full pt-[56.25%] overflow-hidden'>
        <div className='bg-black hover:cursor-pointer bg-opacity-0 hover:bg-opacity-50 transition-all hover:opacity-100 opacity-0 duration-500 flex item-center justify-center z-10 top-0 left-0 w-full h-full absolute object-fit-cover '>
          <PlayIcon />
        </div>
        <img
          src={process.env.REACT_APP_HOST + image}
          className='top-0 left-0 w-full h-full absolute object-cover'
        />
      </div>
    </div>
  );
};
const ItemCoverImageTrack = ({ item, id, image, onClickBtn }) => {
  const { changeCurrentTrack, onOpenMedia } = useContext(MediaContext);

  const handleOnClick = () => {
    onOpenMedia();
    changeCurrentTrack(item);
  };
  return (
    <div className=''>
      <div className='relative w-full pt-[56.25%] overflow-hidden'>
        <div
          onClick={handleOnClick}
          className='bg-black hover:cursor-pointer bg-opacity-0 hover:bg-opacity-50 transition-all hover:opacity-100 opacity-0 duration-500 flex item-center justify-center z-10 top-0 left-0 w-full h-full absolute object-fit-cover '
        >
          <div className='flex items-center justify-center'>
            <PlayIcon />
          </div>
          <div className='top-0 right-0 absolute '>
            <Tooltip title='Add/Remove to/from playlist'>
              <Button
                onClick={(e) => {
                  e.stopPropagation();
                  onClickBtn(id);
                }}
              >
                <CreatePlaylistIcon width={16} height={16} color={'#000'} />
              </Button>
            </Tooltip>
          </div>
        </div>
        <img
          src={process.env.REACT_APP_HOST + image}
          className='top-0 left-0 w-full h-full absolute object-cover'
        />
      </div>
    </div>
  );
};

const ItemPlaylist = ({ item }) => {
  return (
    <div className='w-[275px]'>
      <ItemCoverImage id={item.id} image={item.cover_image} />
      <div className=''>
        <Link to='#'>
          <h5 className='font-semibold text-lg mb-2 truncate'>
            {item.playlist_name}
          </h5>
        </Link>
      </div>
    </div>
  );
};

const ItemTrack = ({ item, onClickBtn = () => {} }) => {
  return (
    <div className='w-[275px]'>
      <ItemCoverImageTrack
        id={item.id}
        image={item.cover_image}
        item={item}
        onClickBtn={onClickBtn}
      />
      <div>
        <Link to={'upsertTrack/' + item.id}>
          <h5 className='font-semibold text-lg mb-1 truncate'>{item.title}</h5>
        </Link>
        <p className='truncate'>{item.artist}</p>
      </div>
    </div>
  );
};

const ListItem = (props) => {
  const { data = [], loading = false, type } = props;
  const [isOpen, setIsOpen] = useState(false);
  const [trackId, setTrackId] = useState('');
  const { auth } = useContext(AuthContext);

  useEffect(
    () => () => {
      setIsOpen(false);
      setTrackId('');
    },
    []
  );

  const onClickBtn = (id) => {
    if (!auth.token) {
      message.error('Please login before using this feature!');
      return;
    }
    setIsOpen(true);
    setTrackId(id);
  };

  const renderItem = (data = []) => {
    if (type === CONSTANT_LIST_TYPE.Playlist) {
      return data.map((it, k) => {
        return <ItemPlaylist item={it} key={it?.id || k} />;
      });
    }
    return data.map((it, k) => {
      return <ItemTrack item={it} key={it?.id || k} onClickBtn={onClickBtn} />;
    });
  };
  return (
    <Spin spinning={loading}>
      <div className='w-full min-h-[15vh] flex gap-3 flex-wrap'>
        {data.length > 0 ? renderItem(data) : <Empty />}
      </div>
      <ModalTrackPlaylist
        id={trackId}
        open={isOpen}
        onCancel={() => setIsOpen(false)}
      />
    </Spin>
  );
};

export default ListItem;
