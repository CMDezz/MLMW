import * as React from 'react';
import Icon from '@ant-design/icons';
const SvgComponent = (props) => {
  const { color = '#fff' } = props;
  return (
    <svg
      xmlns='http://www.w3.org/2000/svg'
      fill='none'
      viewBox='0 0 24 24'
      width={36}
      height={36}
      {...props}
    >
      <path
        stroke={color}
        strokeWidth={1.5}
        d='M8 9.5c0-.466 0-.699.076-.883a1 1 0 0 1 .541-.54C8.801 8 9.034 8 9.5 8s.699 0 .883.076a1 1 0 0 1 .54.541c.077.184.077.417.077.883v5c0 .466 0 .699-.076.883a1 1 0 0 1-.541.54C10.199 16 9.966 16 9.5 16s-.699 0-.883-.076a1 1 0 0 1-.54-.541C8 15.199 8 14.966 8 14.5v-5ZM13 9.5c0-.466 0-.699.076-.883a1 1 0 0 1 .541-.54C13.801 8 14.034 8 14.5 8s.699 0 .883.076a1 1 0 0 1 .54.541c.077.184.077.417.077.883v5c0 .466 0 .699-.076.883a1 1 0 0 1-.541.54c-.184.077-.417.077-.883.077s-.699 0-.883-.076a1 1 0 0 1-.54-.541C13 15.199 13 14.966 13 14.5v-5Z'
      />
      <path
        stroke={color}
        strokeLinecap='round'
        strokeWidth={1.5}
        d='M7 3.338A9.954 9.954 0 0 1 12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12c0-1.821.487-3.53 1.338-5'
      />
    </svg>
  );
};
const PauseIcon = (props) => {
  return <Icon component={() => <SvgComponent {...props} />} />;
};

export default PauseIcon;
