import { Checkbox, Form, message, Modal, Spin } from 'antd';
import { useForm } from 'antd/es/form/Form';
import React, { useEffect, useState } from 'react';
import CommonApis from '../../apis/CommonApis';
import PlaylistApis from '../../apis/PlaylistApis';

const ModalTrackPlaylist = (props) => {
  const { open, onCancel, id } = props;

  const [isLoading, setIsLoading] = useState(false);
  const [trackBelongTo, setTrackBelongTo] = useState({});
  const [dataMyPlaylist, setDataMyPlaylist] = useState([]);

  useEffect(() => {
    if (id) {
      setIsLoading(true);
      Promise.all([getDataTrackPlaylistById(id), getMyPlaylist()])
        .then(([dataPlId, dataPlList]) => {
          let mapData = {};
          if (dataPlId) {
            //map to map obj for better query
            mapData = dataPlId.reduce((o, i) => {
              return {
                ...o,
                [i]: true,
              };
            }, {});
            // setTrackBelongTo({ 1: 1 });
            setTrackBelongTo(mapData);
          }

          if (dataPlList) {
            setDataMyPlaylist(dataPlList);
          }
          setIsLoading(false);
        })
        .catch(() => setIsLoading(false));
    }
  }, [id]);

  const getDataTrackPlaylistById = async (id) => {
    const res = await CommonApis.GetDataTrackPlaylistById(id);
    if (res.Data) {
      return res.Data.playlist_id;
    }
    return [];
  };
  const onChangeCheckBox = (value, id) => {
    setTrackBelongTo((prev) => ({ ...prev, [id]: value }));
  };

  const getMyPlaylist = async () => {
    const res = await PlaylistApis.GetAllPlaylistsByUserId();
    if (res.Data) {
      return res.Data.Playlists;
    }
    return [];
  };

  const handleSubmitTrackPlaylist = async () => {
    setIsLoading(true);
    //map to api request struct
    const request = {
      track_id: id,
      playlist_id: [],
    };
    Object.keys(trackBelongTo).map((plId) => {
      if (trackBelongTo[plId]) {
        request.playlist_id.push(+plId);
      }
    });
    const res = await CommonApis.UpsertTracksPlaylists(request);
    if (res.Data) {
      message.success(res.Message);
    }

    setIsLoading(false);
  };

  return (
    <Modal
      destroyOnClose
      title='Add track to your playlist'
      open={open}
      onOk={handleSubmitTrackPlaylist}
      onCancel={onCancel}
      width={350}
    >
      <Spin spinning={isLoading}>
        {dataMyPlaylist.map((item) => {
          return (
            <div className='pb-2 truncate'>
              <Checkbox
                className=''
                checked={trackBelongTo[item.id] || false}
                onChange={(e) => onChangeCheckBox(e.target.checked, item.id)}
                rootClassName='d-block'
              >
                <p className=''>{item.playlist_name}</p>
              </Checkbox>
            </div>
          );
        })}
      </Spin>
    </Modal>
  );
};

export default ModalTrackPlaylist;
