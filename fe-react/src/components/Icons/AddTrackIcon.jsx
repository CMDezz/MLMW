import * as React from 'react';
import Icon from '@ant-design/icons';
const SvgComponent = (props) => (
  <svg
    xmlns='http://www.w3.org/2000/svg'
    fill='none'
    viewBox='0 0 24 24'
    width={24}
    height={24}
    {...props}
  >
    <path
      stroke='#fff'
      strokeWidth={1.5}
      d='M12 15a2 2 0 1 1-4 0 2 2 0 0 1 4 0ZM12 15V9'
    />
    <path
      stroke='#fff'
      strokeLinecap='round'
      strokeWidth={1.5}
      d='m14.058 11.03-1.316-.659c-.183-.091-.274-.137-.35-.194a1 1 0 0 1-.374-.607C12 9.477 12 9.375 12 9.171c0-.486 0-.728.06-.893a1 1 0 0 1 1.056-.653c.174.02.391.129.826.346l1.316.658c.183.092.274.137.35.195a1 1 0 0 1 .374.606c.018.093.018.195.018.4 0 .485 0 .728-.06.893a1 1 0 0 1-1.056.652c-.174-.02-.391-.129-.826-.346Z'
    />
    <path
      stroke='#fff'
      strokeLinecap='round'
      strokeWidth={1.5}
      d='M14 21.8c-.646.131-1.315.2-2 .2-5.523 0-10-4.477-10-10 0-1.821.487-3.53 1.338-5M21.8 14c.131-.646.2-1.315.2-2 0-5.523-4.477-10-10-10-1.821 0-3.53.487-5 1.338'
    />
    <path
      stroke='#fff'
      strokeLinecap='round'
      strokeLinejoin='round'
      strokeWidth={1.5}
      d='M18 22v-7m0 0 2.5 2.5M18 15l-2.5 2.5'
    />
  </svg>
);
const AddTrackIcon = (props) => {
  return <Icon component={() => <SvgComponent />} {...props} />;
};

export default AddTrackIcon;
