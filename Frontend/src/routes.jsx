import React  from 'react';
import { useRoutes, Outlet  } from "react-router-dom";
// import Microphone from "./components/Microphone/Microphone";
import NavBar from "./components/NavBar/NavBar";
import FirstVoiceAnalystic from './page/FirstVoiceAnalystic';
import AllVoiceAnalystic from "./page/AllVoiceAnalystic";

// ----------------------------------------------------------------------

export default function Router() {
  return useRoutes([
    {
      path: "/*",
      element: (<> <NavBar /> <Outlet /></>),
      children: [
        { path: "*", element: <AllVoiceAnalystic /> },
        { path: "call", element: <AllVoiceAnalystic /> },
        { path: "voice", element: <FirstVoiceAnalystic /> },
      ],
    },
  ]);
}
