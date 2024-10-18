import React, { useContext, useEffect, useRef, useState } from 'react';
import { MediaContext } from '../../contexts/MediaProvider';
import { Button, message } from 'antd';
import LogoIcon from '../Icons/PlayIcon';
import PauseIcon from '../Icons/PauseIcon';

const MediaPlayer = () => {
  const { openMedia, onCloseMedia, currentTrack } = useContext(MediaContext);
  //   const [volume, setVolume] = useState(1);
  const [isPlaying, setIsPlaying] = useState(false);
  const audioRef = useRef(null);
  const handlePlay = () => {
    if (audioRef.current) {
      audioRef.current.play();
    }
  };
  const handlePause = () => {
    if (audioRef.current) {
      audioRef.current.pause();
    }
  };
  const onPlaying = () => {
    setIsPlaying(true);
  };

  const onPause = () => {
    setIsPlaying(false);
  };

  const closeMedia = () => {
    onCloseMedia();
    handlePause();
    setIsPlaying(false);
  };

  useEffect(() => {
    if (currentTrack?.id) {
      //pause
      handlePause();
      setIsPlaying(false);
    }
  }, [currentTrack?.id]);

  useEffect(() => {
    if (currentTrack.url) {
      handlePlay();
    }
  }, [currentTrack.url]);
  return (
    <div
      className={
        'w-full h-[100px] bg-stone-100 sticky ' +
        (openMedia ? 'visible' : 'hidden')
      }
    >
      <div className='h-full py-2 px-3 flex  items-center gap-3'>
        <div className='flex gap-3 w-[45%]'>
          <img
            className='w-[70px] h-70px object-cover'
            src={process.env.REACT_APP_HOST + currentTrack.cover_image}
          />
          <div className='w-[70%] '>
            <h5 className='pb-1 text-lg font-semibold truncate'>
              {currentTrack.title}
            </h5>
            <p className='truncate'>{currentTrack.artist}</p>
          </div>
        </div>
        <div className='flex items-center justify-center'>
          {isPlaying ? (
            <Button onClick={handlePause} type='text' className='h-[50px]'>
              <PauseIcon width={46} height={46} color={'#1C274C'} />
            </Button>
          ) : (
            <Button onClick={handlePlay} type='text' className='h-[50px]'>
              <LogoIcon width={46} height={46} color={'#1C274C'} />
            </Button>
          )}
          <audio
            onPlay={onPlaying} // Listen to play event
            onPause={onPause} // Listen to pause event
            ref={audioRef}
            src={process.env.REACT_APP_HOST + currentTrack.url}
            preload='auto'
          />
        </div>

        <div>
          {/* <input
            type='range'
            min={0}
            max={1}
            step={0.02}
            value={volume}
            onChange={(event) => {
              setVolume(event.target.valueAsNumber);
            }}
          /> */}
        </div>
      </div>
      <Button onClick={closeMedia} className='absolute top-0 right-0'>
        X
      </Button>
    </div>
  );
};
export default MediaPlayer;
