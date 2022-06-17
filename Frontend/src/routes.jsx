import React  from 'react';
import { useRoutes, Outlet  } from "react-router-dom";
import Microphone from "./components/Microphone/Microphone";
import NavBar from "./components/NavBar/NavBar";
import VoiceAllAnalystic from "./page/VoiceAllAnalystic";

// ----------------------------------------------------------------------

export default function Router() {
  return useRoutes([
    {
      path: "/*",
      element: (<> <NavBar /> <Outlet /></>),
      children: [
        { path: "*", element: <VoiceAllAnalystic /> },
        { path: "call", element: <VoiceAllAnalystic /> },
        { path: "voice", element: <Microphone /> },
      ],
    },
    // {
    //   path: "/call", element: <VoiceAllAnalystic /> 
    // },
  ]);
}
