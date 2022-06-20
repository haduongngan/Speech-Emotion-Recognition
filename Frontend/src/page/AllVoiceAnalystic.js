import React, { useState, useEffect } from "react";
import Grid from "@mui/material/Grid";

import NavBar from "../components/NavBar/NavBar";
import Microphone from "../components/Microphone/Microphone";
import AudioPlayer from "../components/AudioPlayer/AudioPlayer";
import UploadFile from "../components/UploadFile/UploadFile";
import { uploadAudio } from "../apis/voiceProcessing";
import AudioPlayerWithStaff from "../components/AudioPlayer/AudioPlayerWithStaff";
import Emotion from "../components/Table/Emotion";
import VoicesHistory from "../components/Table/VoicesHistory";

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
    // let path = null;
    let path = new FormData();
    if (file) {
      path.append("file", file);
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
        <Microphone pushFile={pushFile} onReset={onReset} />
        <UploadFile pushFile={pushFile} onReset={onReset} />
      </Stack>
      <Stack container direction="column" spacing={1}>
        <Stack container direction="row" justifyContent="center" spacing={3}>
          <Emotion title={"Emotion Analystics"} />
          <VoicesHistory />
        </Stack>
        {files[0] && (
          <>
            <Stack
              spacing={2}
              direction="column"
              justifyContent="center"
              alignItems="center"
              sx={{ minWidth: 600 }}
            >
              {submited ? (
                <AudioPlayerWithStaff
                  file={files[0]}
                  phonenumber={"0987654321"}
                  staff={"Hat Nho"}
                  onReset={onReset}
                />
              ) : (
                <>
                  <Box sx={{ width: "100%" }}>
                    <AudioPlayer file={files[0]} onReset={onReset} />
                  </Box>
                  <Stack
                    direction="row"
                    justifyContent="center"
                    alignItems="center"
                    spacing={2}
                    sx={{ m: 2 }}
                  >
                    <TextField
                      id="phonenumber"
                      label="Phone Number"
                      defaultValue={"0987654321"}
                    />
                    <TextField
                      id="staff"
                      label="staff"
                      defaultValue={"Hat Nho"}
                    />
                  </Stack>
                  <Stack
                    justifyContent="center"
                    alignItems="center"
                    sx={{ m: 2 }}
                  >
                    <Button
                      sx={{ maxWidth: 100 }}
                      variant="contained"
                      onClick={() => onSubmit(files[0])}
                    >
                      Submit
                    </Button>
                  </Stack>
                </>
              )}
            </Stack>
          </>
        )}
      </Grid>
    </>
  );
}
export default AllVoiceAnalystic;
