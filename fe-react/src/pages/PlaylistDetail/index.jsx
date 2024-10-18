import React, { useContext } from 'react';
import usePlaylistDetailHook from './usePlaylistDetailHook';
import PageTitle from '../../components/Base/PageTitle';
import { useParams } from 'react-router-dom';
import { Spin, Tooltip } from 'antd';
import { Empty } from '../../components';
import { MediaContext } from '../../contexts/MediaProvider';

const TrackAudio = ({ item }) => {
  const { onOpenMedia, changeCurrentTrack } = useContext(MediaContext);
  return (
    <Tooltip title='Click to play'>
      <div
        onClick={() => {
          onOpenMedia();
          changeCurrentTrack(item);
        }}
        className='px-2 py-3 truncate hover:cursor-pointer relative'
      >
        <div className='rounded-md bg-black hover:cursor-pointer bg-opacity-0 hover:bg-opacity-20 transition-all hover:opacity-100 opacity-0 duration-200 flex item-center justify-center z-10 top-0 left-0 w-full h-full absolute object-fit-cover '></div>
        <div className='flex gap-5'>
          <img
            className='w-[70px] h-70px object-cover'
            src={process.env.REACT_APP_HOST + item.cover_image}
          />
          <div className='w-[70%] '>
            <h5 className='pb-1 text-lg font-semibold truncate'>
              {item.title}
            </h5>
            <p className='truncate'>{item.artist}</p>
          </div>
        </div>
      </div>
    </Tooltip>
  );
};

const PlaylistDetail = () => {
  const { data, loading } = usePlaylistDetailHook();

  const renderTrackSong = () => {
    if (!data?.Track?.length) {
      return <Empty />;
    }

    return data?.Track?.map((item, i) => {
      return <TrackAudio item={item} key={i} />;
    });
  };
  return (
    <Spin spinning={loading}>
      <PageTitle title={data.playlist?.playlist_name} />

      <div className='flex gap-5 py-3'>
        <div className='w-[333px]'>
          <div className='relative w-full pt-[56.25%] overflow-hidden'>
            <img
              src={process.env.REACT_APP_HOST + data?.playlist?.cover_image}
              className='top-0 left-0 w-full h-full absolute object-cover'
            />
          </div>
        </div>
        <div className='flex-1 max-w-[80%]'>
          <p className='line-clamp-6'>
            {data?.playlist?.description}
            alksdjlaks djalksdjlaksdjalksdjlaksdjalksdjla ksdjalksd jlaksdjal
            alksdjlaksdjalksdjlaksdj alksdj laksdjalksdjlaksdja lksdjlaksdjalk
            sdjlaksdjalksdjlaksd jalksdjla ksdjalksdjlaksd jalksdjlaksdja
            lksdjlaksdjalksdjlaksdjalksdjl aksdjalksdjl
            aksdjalksdjlaksdjalksdjlaks djalksdjlaksdjalk sdjlaksdj
            alksdjlaksdjalksdjlaksdj alksdjla ksdjalksdjlaksdja
            lksdjlaksdjalksdj laksdjalk sdjlaksdja lksdjlaks djalks
          </p>
        </div>
      </div>

      <h5 className='text-xl font-semibold'>Tracks song in this playlist: </h5>

      {renderTrackSong()}
    </Spin>
  );
};

export default PlaylistDetail;
