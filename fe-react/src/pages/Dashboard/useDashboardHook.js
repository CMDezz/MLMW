import { useEffect, useState } from 'react';
import TrackApis from '../../apis/TrackApis';
import PlaylistApis from '../../apis/PlaylistApis';
import { useLocation } from 'react-router-dom';
import CommonApis from '../../apis/CommonApis';

const useDashBoardHook = () => {
  const [isLoadingPlaylist, setIsLoadingPlaylist] = useState(false);
  const [isLoadingTrack, setIsLoadingTrack] = useState(false);
  const [dataPlaylist, setDataPlaylist] = useState([]);
  const [dataTrack, setDataTrack] = useState([]);
  const [keyword, setKeyword] = useState('');
  const location = useLocation();

  const getKeywordFromSearchParams = () => {
    const queryParams = new URLSearchParams(location.search);
    return queryParams.get('keyword') || '';
  };

  useEffect(() => {
    // Promise.all([getDataPublicTrack(), getDataPublicPlaylist()]);
    const kw = getKeywordFromSearchParams();
    setKeyword(kw); // Update state when keyword changes

    if (kw) {
      getDataSearch(kw);
    } else {
      getDataPublicTrack();
      getDataPublicPlaylist();
    }

    return () => {
      setDataTrack([]);
      setDataPlaylist([]);
      setIsLoadingPlaylist(false);
      setIsLoadingTrack(false);
    };
  }, [location.search]);

  const getDataPublicTrack = async () => {
    setIsLoadingTrack(true);
    const res = await TrackApis.GetAllPublicsTrack();
    if (res.Data) {
      setDataTrack(res.Data.Tracks);
    }
    setIsLoadingTrack(false);
  };

  const getDataPublicPlaylist = async () => {
    setIsLoadingPlaylist(true);
    const res = await PlaylistApis.GetAllPublicsPlaylist();
    if (res.Data) {
      setDataPlaylist(res.Data.Playlists);
    }
    setIsLoadingPlaylist(false);
  };

  const getDataSearch = async (keyword) => {
    setIsLoadingPlaylist(true);
    setIsLoadingTrack(true);
    const res = await CommonApis.Search(keyword);
    if (res.Data) {
      setDataPlaylist(res.Data.Playlists);
    }
    setIsLoadingPlaylist(false);
    setIsLoadingTrack(false);
  };

  return {
    isLoadingPlaylist,
    isLoadingTrack,
    dataPlaylist,
    dataTrack,
    keyword,
  };
};

export default useDashBoardHook;
