import Fetch from './index';
import HOST_API from './host';

const ControllerName = 'user';
const ControllerUri = HOST_API + '/' + ControllerName;

const UserApis = {
  Login: async (data) => {
    const resp = await Fetch.POST(`${ControllerUri}/login`, data);
    return resp;
  },
  Create: async (data) => {
    const resp = await Fetch.POST(`${ControllerUri}/createUser`, data);
    return resp;
  },
};
export default UserApis;
