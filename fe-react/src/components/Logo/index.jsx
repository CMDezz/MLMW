import React from 'react';
import { Link } from 'react-router-dom';
import LogoIcon from '../Icons/PlayIcon';

const Logo = ({ collapsed }) => {
  return (
    <div className='my-2'>
      <Link to='/'>
        <div className='pl-5 flex items-center gap-5'>
          <div>
            <LogoIcon />
          </div>
          <h5
            className={
              (collapsed ? 'opacity-0 ' : 'opacity-100 ') +
              ' transition-all duration-300  text-white font-extrabold text-xl'
            }
          >
            MLMW
          </h5>
        </div>
      </Link>
    </div>
  );
};

export default Logo;
