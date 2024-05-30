import Registration from './Pages/Registration'
import Log from './Pages/Log'
import NoPage from './Pages/NoPage';
import './App.css'
import { BrowserRouter, Routes, Route } from "react-router-dom";
import MainP from './forUser/MainP';
import Profile from './forUser/Profile';
import ProfileOrganizer from './forOrganizer/ProfileOrganizer'
import CreateEvent from './forOrganizer/CreateEvent';
import CreateFeedback from './forUser/CreateFeedback'
import Feedback from './forOrganizer/Feedback'

function App(){
{
    return (
    <div className='App'>

      <BrowserRouter>
            <Routes>
                <Route path="/" element={<Registration/>}/>
                <Route path="/log" element={<Log/>}/>
                <Route path="/mainpage" element={<MainP/>}/>
                <Route path="/user-profile" element={<Profile/>}/>
                <Route path="/organizer-profile" element={<ProfileOrganizer/>}/>
                <Route path="create-event" element={<CreateEvent/>}/>
                <Route path="create-feedback" element={<CreateFeedback/>}/>
                <Route path="/feedback" element={<Feedback/>} />
                <Route path="*" element={<NoPage/>}/>
            </Routes>
       </BrowserRouter>

    </div>
    );
}
}
export default App;
