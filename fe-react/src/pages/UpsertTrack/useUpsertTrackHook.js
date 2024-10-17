import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import TrackApis from '../../apis/TrackApis';
import { Form, message } from 'antd';
import { createSampleFileFormUrl } from '../../utils/function';

const useUpsertTrackHook = () => {
  let { id } = useParams();
  const [fileList, setFileList] = useState([]);
  const [imageList, setImageList] = useState([]);
  const [form] = Form.useForm();
  const [isLoading, setIsLoading] = useState(false);

  useEffect(() => {
    if (id) {
      getCurrentTrack(id);
    }
    return () => {
      setIsLoading(false);
      setImageList([]);
      setFileList([]);
    };
  }, [id]);

  const resetFieldFormData = (resData) => {
    const trackFile = createSampleFileFormUrl(resData.url, 'audio/mpeg');
    const coverImage = createSampleFileFormUrl(
      resData.cover_image,
      'image/jpeg'
    );

    form.setFieldsValue({
      title: resData.title,
      artist: resData.artist,
      album: resData.album,
      genre: resData.genre,
      release_year: resData.release_year,
      duration: resData.duration,
      track_file: trackFile,
      cover_image: coverImage,
    });
    setFileList([trackFile]);
    setImageList([coverImage]);
  };

  const getCurrentTrack = async (id) => {
    setIsLoading(true);
    const res = await TrackApis.GetTrackById(id);
    if (res.Data) {
      resetFieldFormData(res.Data);
    }
    setIsLoading(false);
  };

  const beforeUploadFileTrack = () => {
    return false;
  };
  const beforeUploadCoverImage = () => {
    return false;
  };

  const onFinish = async (values) => {
    console.log('Received values of form: ', values);
    console.log('');

    const formData = new FormData();
    formData.append('album', values.album);
    formData.append('artist', values.artist);
    formData.append('duration', values.duration);
    formData.append('genre', values.genre);
    formData.append('release_year', values.release_year);
    formData.append('title', values.title);

    const _track_file = values.cover_image?.fileList?.[0]?.originFileObj;
    const _cover_image = values.cover_image?.fileList?.[0]?.originFileObj;
    if (_cover_image?.size) {
      formData.append(
        'cover_image',
        values.cover_image?.fileList[0]?.originFileObj
      );
    }
    if (_track_file?.size) {
      formData.append(
        'track_file',
        values.track_file?.fileList[0]?.originFileObj
      );
    }

    let res = {};
    if (id) {
      formData.append('id', id);
      res = await TrackApis.UpdateTrack(formData);
    } else {
      res = await TrackApis.CreateTrack(formData);
    }

    if (res.Data) {
      message.success(res.Message);
      if (id) {
        resetFieldFormData(res.Data);
      } else {
        form.resetFields();
        setFileList([]);
        setImageList([]);
      }
    }
    // onOk(values);
  };

  return {
    beforeUploadFileTrack,
    beforeUploadCoverImage,
    fileList,
    imageList,
    setFileList,
    setImageList,
    form,
    onFinish,
    isLoading,
  };
};

export default useUpsertTrackHook;
