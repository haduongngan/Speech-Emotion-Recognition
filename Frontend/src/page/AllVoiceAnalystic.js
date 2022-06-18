import React, { useState, useEffect } from "react";
import { Button, Stack, TextField, Box } from "@mui/material";

import Microphone from "../components/Microphone/Microphone";
import AudioPlayer from "../components/AudioPlayer/AudioPlayer";
import UploadFile from "../components/UploadFile/UploadFile";
import { uploadAudio } from "../apis/voiceProcessing";
import AudioPlayerWithStaff from "../components/AudioPlayer/AudioPlayerWithStaff";

function AllVoiceAnalystic() {
  const [files, setFiles] = useState([]);
  const [submited, setSubmited] = useState(false);
  const [reset, setReset] = useState(false);
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
      <Stack
        spacing={3}
        container
        justifyContent="center"
        alignItems="center"
        direction="row"
        sx={{ m: 4 }}
      >
        <Microphone pushFile={pushFile} onReset={onReset} />
        <UploadFile pushFile={pushFile} onReset={onReset} />
      </Stack>
      <Stack container direction="column" spacing={3}>
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
      </Stack>
    </>
  );
}
export default AllVoiceAnalystic;
