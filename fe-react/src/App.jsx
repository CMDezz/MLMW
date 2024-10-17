import React, { useContext } from 'react';
import { AppLayout } from './components';
import {
  BrowserRouter,
  Navigate,
  redirect,
  Route,
  Routes,
} from 'react-router-dom';
import Dashboard from './pages/Dashboard';
import Login from './pages/Login';
import MyLibrary from './pages/MyLibrary';
import PlaylistDetail from './pages/PlaylistDetail';
import SearchResult from './pages/SearchResult';
import TrackDetail from './pages/TrackDetail';
import UpsertPlaylist from './pages/UpsertPlaylist';
import UpsertTrack from './pages/UpsertTrack';
import { AuthContext, AuthProvider } from './contexts/AuthProvider';

const PrivateRoute = ({ element }) => {
  const { auth } = useContext(AuthContext);
  const localAuth = JSON.parse(localStorage.getItem('fe-react-auth') || '{}');

  return auth.token || localAuth.token ? element : <Navigate to='/login' />;
};

function App() {
  return (
    <div className='App'>
      <AuthProvider>
        <BrowserRouter>
          <Routes>
            <Route path='/' element={<AppLayout />}>
              <Route index element={<Dashboard />} />
              <Route path='/login' element={<Login />}></Route>
              <Route path='/searchResult' element={<SearchResult />}></Route>
              <Route
                path='myLibrary'
                element={<PrivateRoute element={<MyLibrary />} />}
              ></Route>
              <Route
                path='/playlistDetail'
                element={<PrivateRoute element={<PlaylistDetail />} />}
              ></Route>
              <Route
                path='/trackDetail'
                element={<PrivateRoute element={<TrackDetail />} />}
              ></Route>
              <Route
                path='/upsertPlaylist/:id'
                element={<PrivateRoute element={<UpsertPlaylist />} />}
              ></Route>
              <Route
                path='/upsertTrack'
                element={<PrivateRoute element={<UpsertTrack />} />}
              ></Route>
              <Route
                path='/upsertTrack/:id'
                element={<PrivateRoute element={<UpsertTrack />} />}
              ></Route>
            </Route>
          </Routes>
        </BrowserRouter>
      </AuthProvider>
    </div>
  );
}

export default App;
