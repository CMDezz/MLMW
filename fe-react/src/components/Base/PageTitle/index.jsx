import React from 'react';

const PageTitle = (props) => {
  const { title } = props;
  return (
    <div className=' mb-5 px-3 py-2 w-full bg-cyan-50 uppercase font-semibold text-lg'>
      <h5 className='truncate'>{title}</h5>
    </div>
  );
};

export default PageTitle;
