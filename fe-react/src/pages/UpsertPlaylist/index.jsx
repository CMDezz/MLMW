import React from 'react';
import PageTitle from '../../components/Base/PageTitle';
import { useParams } from 'react-router-dom';
import {
  Button,
  Col,
  Form,
  Input,
  InputNumber,
  Row,
  Switch,
  Upload,
} from 'antd';
import { UploadOutlined } from '@ant-design/icons';
import useUpsertPlaylistHook from './useUpsertPlaylistHook';
import { normFile } from '../../utils/function';
const UpsertPlaylist = () => {
  let { id } = useParams();
  const {
    onFinish,
    beforeUploadCoverImage,
    onRemoveCoverImage,
    imageList,
    form,
    setImageList,
  } = useUpsertPlaylistHook();

  return (
    <div>
      <PageTitle
        title={id ? 'UPDATE PLAYLIST DETAILS' : 'CREATE A NEW PLAYLIST'}
      />

      <Form
        // {...formItemLayout}
        // labelCol={{
        //   md: { span: 24 },
        // }}
        // wrapperCol={{
        //   md: { span: 24 },
        // }}
        layout='vertical'
        autoComplete='off'
        form={form}
        name='upsertPlaylit'
        onFinish={onFinish}
        style={{ maxWidth: 600 }}
        scrollToFirstError
        className='mx-auto border px-10 py-2 rounded-md'
      >
        <Form.Item
          name='playlist_name'
          label='Playlist Name'
          // labelCol={{
          //   offset: 0,
          // }}
          // wrapperCol={{
          //   offset: 12,
          // }}
          rules={[
            {
              required: true,
              message: `Please input playlist's name!`,
              whitespace: true,
            },
            {
              max: 150,
              message: 'Maximum 150 characters',
            },
          ]}
        >
          <Input autoComplete='off' />
        </Form.Item>

        <Form.Item
          name='description'
          label='Description (optional)'
          rules={[
            {
              max: 250,
              message: 'Maximum 150 characters',
            },
          ]}
        >
          <Input.TextArea rows={4} autoComplete='off' />
        </Form.Item>

        <Form.Item
          name='cover_image'
          label='Cover Image (16x9 resolution is prefer)'
          valuePropName='cover_image'
          getValueFromEvent={normFile}
          rules={[
            {
              required: true,
              message: `Please choose track cover image!`,
            },
          ]}
        >
          <Upload
            fileList={imageList}
            onChange={({ fileList: newFileList }) => setImageList(newFileList)}
            limit
            beforeUpload={beforeUploadCoverImage}
            onRemove={onRemoveCoverImage}
            name='cover_image'
            accept='image/*'
          >
            <Button disabled={imageList.length >= 1} icon={<UploadOutlined />}>
              Click to upload
            </Button>
          </Upload>
        </Form.Item>
        <Row justify={'space-between'}>
          <Col>
            <Form.Item
              // wrapperCol={{ offset: 8, span: 8 }}
              // labelCol={{
              //   md: { span: 24 },
              // }}
              // wrapperCol={{
              //   md: { span: 24 },
              // }}
              name={'is_public'}
              valuePropName={'is_public'}
              layout='horizontal'
              label={'Is Plulic'}
              // labelCol={{
              //   md: { span: 24 },
              // }}
              // wrapperCol={{
              //   md: { span: 24 },
              // }}
            >
              <Switch name={'is_public'} defaultChecked />
            </Form.Item>
          </Col>
          <Col>
            <Form.Item className='w-full' wrapperCol={{ offset: 8, span: 16 }}>
              <Button type='primary' htmlType='submit'>
                Submit
              </Button>
            </Form.Item>
          </Col>
        </Row>
      </Form>
    </div>
  );
};

export default UpsertPlaylist;
