import { useEffect, useState } from 'react';
import TrackApis from '../../apis/TrackApis';
import PlaylistApis from '../../apis/PlaylistApis';

const useDashBoardHook = () => {
  const [isLoadingPlaylist, setIsLoadingPlaylist] = useState(false);
  const [isLoadingTrack, setIsLoadingTrack] = useState(false);
  const [dataPlaylist, setDataPlaylist] = useState([]);
  const [dataTrack, setDataTrack] = useState([]);

  useEffect(() => {
    Promise.all([getDataPublicTrack(), getDataPublicPlaylist()]);
    console.log('mounted');

    return () => {
      setDataTrack([]);
      setDataPlaylist([]);
      setIsLoadingPlaylist(false);
      setIsLoadingTrack(false);
    };
  }, []);

  const getDataPublicTrack = async () => {
    setIsLoadingTrack(true);
    const res = await TrackApis.GetAllPublicsTrack();
    if (res.Data) {
      console.log('res.Data.Tracks ', res.Data.Tracks);

      setDataTrack(res.Data.Tracks);
    }
    setIsLoadingTrack(false);
  };

  const getDataPublicPlaylist = async () => {
    setIsLoadingPlaylist(true);
    const res = await PlaylistApis.GetAllPublicsPlaylist();
    if (res.Data) {
      console.log('res.Data.Playlists ', res.Data.Playlists);

      setDataPlaylist(res.Data.Playlists);
    }
    setIsLoadingPlaylist(false);
  };

  return {
    isLoadingPlaylist,
    isLoadingTrack,
    dataPlaylist,
    dataTrack,
  };
};

export default useDashBoardHook;
