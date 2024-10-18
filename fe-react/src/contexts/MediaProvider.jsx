import React, { createContext, useEffect, useState } from 'react';

const MediaContext = createContext();

const MediaProvider = ({ children }) => {
  const [openMedia, setOpenMedia] = useState(false);
  const [currentTrack, setCurrentTrack] = useState({});
  // const [currentPlaylist,setCurrentPlatlist] = useState([])

  const onCloseMedia = () => {
    setOpenMedia(false);
  };
  const onOpenMedia = () => {
    setOpenMedia(true);
  };
  const changeCurrentTrack = (track = {}) => {
    setCurrentTrack(track);
  };

  // const changeCurrentPlaylist= ()=>{}

  return (
    <MediaContext.Provider
      value={{
        currentTrack,
        openMedia,
        changeCurrentTrack,
        onOpenMedia,
        onCloseMedia,
      }}
    >
      {children}
    </MediaContext.Provider>
  );
};

export { MediaProvider, MediaContext };
