import * as React from 'react';
import Icon from '@ant-design/icons';
const SvgComponent = (props) => {
  const { color = '#fff' } = props;
  return (
    <svg
      xmlns='http://www.w3.org/2000/svg'
      fill='none'
      viewBox='0 0 24 24'
      width={24}
      height={24}
      {...props}
    >
      <path
        stroke={color}
        strokeLinecap='round'
        strokeWidth={1.5}
        d='M11 14H3M11 18H3'
      />
      <path
        stroke={color}
        strokeWidth={1.5}
        d='M18.875 14.118c1.654.955 2.48 1.433 2.602 2.121a1.5 1.5 0 0 1 0 .521c-.121.69-.948 1.167-2.602 2.121-1.654.955-2.48 1.433-3.138 1.194a1.499 1.499 0 0 1-.451-.261c-.536-.45-.536-1.404-.536-3.314 0-1.91 0-2.865.536-3.314a1.5 1.5 0 0 1 .451-.26c.657-.24 1.484.238 3.138 1.192Z'
      />
      <path
        stroke={color}
        strokeLinecap='round'
        strokeWidth={1.5}
        d='M3 6h10.5M20 6h-2.25M20 10H9.5M3 10h2.25'
      />
    </svg>
  );
};
const CreatePlaylistIcon = (props) => {
  return <Icon component={() => <SvgComponent {...props} />} />;
};

export default CreatePlaylistIcon;
