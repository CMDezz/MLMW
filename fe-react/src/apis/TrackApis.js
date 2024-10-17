import Fetch from './index';
import HOST_API from './host';

const ControllerName = 'track';
const ControllerUri = HOST_API + '/' + ControllerName;

const TrackApis = {
  GetTrackById: async (id) => {
    const resp = await Fetch.GET(`${ControllerUri}/getTrackById/${id}`);
    return resp;
  },
  GetAllPublicsTrack: async () => {
    const resp = await Fetch.GET(`${ControllerUri}/getAllPublicTracks`);
    return resp;
  },
  GetAllTracksByUserId: async () => {
    const resp = await Fetch.GET(`${ControllerUri}/getAllTracksByUserId`);
    return resp;
  },
  CreateTrack: async (data) => {
    const resp = await Fetch.POSTFORM(`${ControllerUri}/createTrack`, data);
    return resp;
  },
  UpdateTrack: async (data) => {
    const resp = await Fetch.PUTFORM(`${ControllerUri}/updateTrack`, data);
    return resp;
  },
};
export default TrackApis;
