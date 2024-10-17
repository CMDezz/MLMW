import Fetch from './index';
import HOST_API from './host';

const ControllerName = 'playlist';
const ControllerUri = HOST_API + '/' + ControllerName;

const PlaylistApis = {
  GetAllPublicsPlaylist: async () => {
    const resp = await Fetch.GET(`${ControllerUri}/getAllPublicPlaylists`);
    return resp;
  },
  GetAllPlaylistsByUserId: async () => {
    const resp = await Fetch.GET(`${ControllerUri}/getAllPlaylistsByUserId`);
    return resp;
  },
};
export default PlaylistApis;
