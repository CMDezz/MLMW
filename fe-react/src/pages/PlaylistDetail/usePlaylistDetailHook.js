import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import PlaylistApis from '../../apis/PlaylistApis';

const usePlaylistDetailHook = () => {
  const { id } = useParams();
  const [data, setData] = useState({});
  const [loading, setLoading] = useState(false);
  useEffect(() => {
    if (id) {
      getDetailPlaylist(id);
    }
  }, [id]);

  const getDetailPlaylist = async (id) => {
    setLoading(true);
    const res = await PlaylistApis.GetFullPlaylistDetail(id);
    if (res.Data) {
      setData(res.Data);
    }
    setLoading(false);
  };
  return {
    data,
    loading,
  };
};

export default usePlaylistDetailHook;
