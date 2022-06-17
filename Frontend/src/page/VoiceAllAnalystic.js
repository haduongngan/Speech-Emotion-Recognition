import React, { useState, useEffect } from "react";
import { Grid, Button, Stack, TextField, Box } from "@mui/material";

import NavBar from "../components/NavBar/NavBar";
import Microphone from "../components/Microphone/Microphone";
import AudioPlayer from "../components/AudioPlayer/AudioPlayer";
import UploadFile from "../components/UploadFile/UploadFile";
import { uploadAudio } from "../apis/voiceProcessing";

function VoiceAllAnalystic() {
  const [files, setFiles] = useState([]);

  const pushFile = (file) => {
    setFiles([...files, file]);
  };
  useEffect(() => {
    // console.log("filea", files);
  }, [files]);

  const onSubmit = (file) => {
    let path = null;
    if (file) {
      if (file.blobURL) path = file.blobURL;
      else {
        path = URL.createObjectURL(file);
      }
    }
    console.log(file);
    console.log(path);
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
        direction="column"
        justifyContent="center"
        alignItems="center"
        sx={{ m: 4 }}
      >
        <Stack direction="row" spacing={3}>
          <Microphone pushFile={pushFile} />
          <UploadFile pushFile={pushFile} />
        </Stack>
        <Stack container direction="column" spacing={3}>
          {files.map(
            (file, index) =>
              file && (
                <Stack key={index} spacing={2}>
                  <Box sx={{ width: "100%" }}>
                    <AudioPlayer file={file} />
                  </Box>
                  <Stack
                    direction="row"
                    justifyContent="center"
                    alignItems="center"
                    spacing={2}
                    sx={{ m: 2 }}
                  >
                    <TextField id="phonenumber" label="Phone Number" />
                    <TextField id="staff" label="staff" />
                  </Stack>
                </Stack>
              )
          )}
        </Stack>
        {files[0] && (
          <Grid justifyContent="center" alignItems="center" sx={{ m: 2 }}>
            <Button variant="contained" onClick={() => onSubmit(files[0])}>
              Submit
            </Button>
          </Grid>
        )}
      </Grid>
    </>
  );
}
export default VoiceAllAnalystic;
