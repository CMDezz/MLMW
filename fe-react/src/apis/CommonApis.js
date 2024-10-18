import Fetch from './index';
import HOST_API from './host';

const ControllerUri = HOST_API;

const CommonApis = {
  GetDataTrackPlaylistById: async (id) => {
    const resp = await Fetch.GET(
      `${ControllerUri}/trackplaylist/getDataTrackPlaylistById/${id}`
    );
    return resp;
  },
  UpsertTracksPlaylists: async (data) => {
    const resp = await Fetch.POST(
      `${ControllerUri}/trackplaylist/upsert`,
      data
    );
    return resp;
  },
  Search: async (keyword) => {
    const resp = await Fetch.GET(`${ControllerUri}/search?keyword=${keyword}`);
    return resp;
  },
};
export default CommonApis;
