import Fetch from './index';
import HOST_API from './host';

const ControllerName = 'playlist';
const ControllerUri = HOST_API + '/' + ControllerName;

const PlaylistApis = {
  GetPlaylistById: async (id) => {
    const resp = await Fetch.GET(`${ControllerUri}/getPlaylistById/${id}`);
    return resp;
  },
  CreatePlaylist: async (data) => {
    const resp = await Fetch.POSTFORM(`${ControllerUri}/createPlaylist`, data);
    return resp;
  },
  UpdatePlaylist: async (data) => {
    const resp = await Fetch.PUTFORM(`${ControllerUri}/updatePlaylist`, data);
    return resp;
  },
  GetAllPublicsPlaylist: async () => {
    const resp = await Fetch.GET(`${ControllerUri}/getAllPublicPlaylists`);
    return resp;
  },
  GetAllPlaylistsByUserId: async () => {
    const resp = await Fetch.GET(`${ControllerUri}/getAllPlaylistsByUserId`);
    return resp;
  },
  GetFullPlaylistDetail: async (id) => {
    const resp = await Fetch.GET(
      `${ControllerUri}/getFullPlaylistDetail/${id}`
    );
    return resp;
  },
};
export default PlaylistApis;
