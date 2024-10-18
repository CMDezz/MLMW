import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { Form, message } from 'antd';
import { createSampleFileFormUrl } from '../../utils/function';
import PlaylistApis from '../../apis/PlaylistApis';

const useUpsertPlaylistHook = () => {
  let { id } = useParams();
  const [imageList, setImageList] = useState([]);
  const [form] = Form.useForm();
  const [isLoading, setIsLoading] = useState(false);

  useEffect(() => {
    if (id) {
      getCurrentPlaylist(id);
    }
    return () => {
      setIsLoading(false);
      setImageList([]);
    };
  }, [id]);

  const resetFieldFormData = (resData) => {
    const coverImage = createSampleFileFormUrl(
      resData.cover_image,
      'image/jpeg'
    );

    form.setFieldsValue({
      playlist_name: resData.playlist_name,
      description: resData.description,
      cover_image: coverImage,
      is_public: resData.is_public,
    });
    setImageList([coverImage]);
  };

  const getCurrentPlaylist = async (id) => {
    setIsLoading(true);
    const res = await PlaylistApis.GetPlaylistById(id);
    if (res.Data) {
      resetFieldFormData(res.Data);
    }
    setIsLoading(false);
  };

  const beforeUploadCoverImage = () => {
    return false;
  };

  const onFinish = async (values) => {
    const formData = new FormData();
    formData.append('playlist_name', values.playlist_name);
    formData.append('description', values.description);
    formData.append('is_public', values.is_public);

    const _cover_image = values.cover_image?.[0]?.originFileObj;

    // formData.delete('_cover_image');
    if (_cover_image?.size) {
      formData.append('cover_image', values.cover_image?.[0]?.originFileObj);
    }
    let res = {};
    if (id) {
      formData.append('id', id);

      res = await PlaylistApis.UpdatePlaylist(formData);
    } else {
      res = await PlaylistApis.CreatePlaylist(formData);
    }

    if (res.Data) {
      message.success(res.Message);
      if (id) {
        resetFieldFormData(res.Data);
      } else {
        form.resetFields();
        setImageList([]);
      }
    }
    // onOk(values);
  };

  const onRemoveCoverImage = () => {
    setImageList([]);
    form.setFieldValue('cover_image', { file: {}, fileList: [] });
  };

  return {
    beforeUploadCoverImage,
    onRemoveCoverImage,
    imageList,
    setImageList,
    form,
    onFinish,
    isLoading,
  };
};

export default useUpsertPlaylistHook;
