import React, { useState, useEffect } from "react";
import Grid from "@mui/material/Grid";

import NavBar from "../components/NavBar/NavBar";
import Microphone from "../components/Microphone/Microphone";
import AudioPlayer from "../components/AudioPlayer/AudioPlayer";
import UploadFile from "../components/UploadFile/UploadFile";


function VoiceAllAnalystic() {
  const [files, setFiles] = useState([]);

  const pushFile = (file) => {
    setFiles([...files, file]);
  };
  useEffect(() => {
    // console.log("filea", files);
  }, [files]);

  return (
    <>
      <NavBar />
      <Grid
        item
        spacing={3}
        container 
        justifyContent="center"
        alignItems="center"
        sx={{ mt: 2 }}
      >
        <Microphone pushFile={pushFile} />
        <UploadFile pushFile={pushFile} />
      </Grid>
      <Grid container direction="column" spacing={3}>
        {files.map(
          (file, index) =>
            file && (
              <Grid key={index} item>
                <AudioPlayer file={file} />
              </Grid>
            )
        )}
      </Grid>
    </>
  );
}
export default VoiceAllAnalystic;
