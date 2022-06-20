import React, { useState, useEffect } from "react";
import Grid from "@mui/material/Grid";

import NavBar from "../components/NavBar/NavBar";
import Microphone from "../components/Microphone/Microphone";
import AudioPlayer from "../components/AudioPlayer/AudioPlayer";
import UploadFile from "../components/UploadFile/UploadFile";

function AllVoiceAnalystic() {
  const [files, setFiles] = useState([]);

  const pushFile = (file) => {
    setFiles([...files, file]);
  };
  useEffect(() => {
    if (reset) {
      setFiles([]);
      setReset(false);
      setSubmited(false);
    }
  }, [reset]);

  const onReset = () => {
    setReset(true);
  };

  const onSubmit = (file) => {
    let path = null;
    if (file) {
      if (file.blobURL) path = file.blobURL;
      else {
        path = URL.createObjectURL(file);
      }
    }
    setSubmited(true);
    uploadAudio(path)
      .then(() => {
        console.log("sending");
      })
      .catch((err) => {
        console.log(err.message);
        // setIsLoading(false);
      });
  };

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
export default AllVoiceAnalystic;
