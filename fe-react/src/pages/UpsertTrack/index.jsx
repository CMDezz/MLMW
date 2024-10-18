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
import { normFile } from '../../utils/function';
import useUpsertTrackHook from './useUpsertTrackHook';
const UpsertTrack = () => {
  let { id } = useParams();
  const {
    onFinish,
    beforeUploadFileTrack,
    beforeUploadCoverImage,
    fileList,
    imageList,
    setFileList,
    form,
    setImageList,
  } = useUpsertTrackHook();

  return (
    <div>
      <PageTitle title={id ? 'UPDATE TRACK DETAILS' : 'CREATE A NEW TRACK'} />

      <Form
        // {...formItemLayout}
        layout='vertical'
        autoComplete='off'
        form={form}
        name='upsertTrack'
        onFinish={onFinish}
        style={{ maxWidth: 600 }}
        scrollToFirstError
        className='mx-auto border px-10 py-2 rounded-md'
      >
        <Form.Item
          name='title'
          label='Title'
          rules={[
            {
              required: true,
              message: `Please input track's title!`,
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
          name='artist'
          label='Artist'
          rules={[
            {
              required: true,
              message: `Please input track's artist!`,
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
          name='album'
          label='Album'
          rules={[
            {
              required: true,
              message: `Please input track's album!`,
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
          name='genre'
          label='Genre'
          rules={[
            {
              required: true,
              message: `Please input track's genre!`,
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
          name='release_year'
          label='Release Year'
          rules={[
            {
              required: true,
              message: `Please input track's release year!`,
            },
            {
              min: 1900,
              max: new Date().getFullYear(),
              transform: (value) => +value,
              message: 'Years from 1900~now',
            },
          ]}
        >
          <InputNumber rootClassName='w-full' autoComplete='off' />
        </Form.Item>
        <Form.Item
          name='duration'
          label='Duration (in seconds)'
          rules={[
            {
              required: true,
              message: `Please input track's duration!`,
            },
            {
              min: 0,
              max: 3600 * 3,
              transform: (value) => +value,
              message: 'Value should be less than 3 hours',
            },
          ]}
        >
          <InputNumber
            rootClassName='w-full'
            type='number'
            autoComplete='off'
          />
        </Form.Item>
        <Form.Item
          name='track_file'
          label='MP3 File'
          getValueFromEvent={normFile}
          valuePropName='track_file'
          rules={[
            {
              required: true,
              message: `Please choose track file!`,
            },
          ]}
        >
          <Upload
            beforeUpload={beforeUploadFileTrack}
            fileList={fileList}
            onChange={({ fileList: newFileList }) => setFileList(newFileList)}
            name='track_file'
            accept='audio/*'
          >
            <Button disabled={fileList.length >= 1} icon={<UploadOutlined />}>
              Click to upload
            </Button>
          </Upload>
        </Form.Item>
        <Form.Item
          name='cover_image'
          getValueFromEvent={normFile}
          label='Cover Image (16x9 resolution is prefer)'
          valuePropName='cover_image'
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
              name={'is_public'}
              valuePropName={'is_public'}
              label={'Is Plulic'}
              layout='horizontal'
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

export default UpsertTrack;
