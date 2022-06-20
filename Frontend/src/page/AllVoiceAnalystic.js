import React, { useState, useEffect } from "react";
import {
  Button,
  Stack,
  TextField,
  Box,
  Typography,
  Card,
  // getNativeSelectUtilityClasses,
  CircularProgress,
} from "@mui/material";

// import NavBar from "../components/NavBar/NavBar";
import Microphone from "../components/Microphone/Microphone";
import AudioPlayer from "../components/AudioPlayer/AudioPlayer";
import UploadFile from "../components/UploadFile/UploadFile";
import { uploadAudio } from "../apis/voiceProcessing";
import AudioPlayerWithStaff from "../components/AudioPlayer/AudioPlayerWithStaff";
import Emotion from "../components/Table/Emotion";
// import VoicesHistory from "../components/Table/VoicesHistory";
const fakeData = {
  customer: {
    emo: {
      happy: 2,
      sad: 2,
      angry: 0,
      calm: 0,
      surprised: 0,
      disgusted: 0,
    },
    gender: "male",
    feel: "happy",
  },
  staff: {
    emo: {
      happy: 2,
      sad: 2,
      angry: 0,
      calm: 0,
      surprised: 0,
      disgusted: 0,
    },
    gender: "male",
    feel: "happy",
  },
  dur: "12",
};
function AllVoiceAnalystic() {
  const [files, setFiles] = useState([]);
  const [submited, setSubmited] = useState(false);
  const [reset, setReset] = useState(false);
  const [data, setData] = useState(fakeData);
  const [loading, setLoading] = useState(false);
  const pushFile = (file) => {
    setFiles([...files, file]);
  };
  useEffect(() => {
    if (reset) {
      setFiles([]);
      setReset(false);
      setSubmited(false);
      setLoading(false);
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
    setLoading(true);
    setSubmited(true);
    uploadAudio(path)
      .then((res) => {
        console.log("res", res.data.data);
        if (res.data.data) {
          setData(res.data.data);
        }
        setLoading(false);
      })
      .catch((err) => {
        console.log(err.message);
        setLoading(false);
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
      <Stack container direction="column" spacing={1}>
        {files[0] && (
          <>
            <Stack
              spacing={3}
              container
              justifyContent="center"
              alignItems="center"
              direction="row"
              sx={{ m: 4 }}
            >
              {!submited ? (
                loading ? (
                  <Box align="center" colSpan={12} sx={{ py: 3 }}>
                    <CircularProgress color="secondary" />
                  </Box>
                ) : (
                  <Stack container direction="column" spacing={1}>
                    <Stack
                      container
                      direction="row"
                      justifyContent="center"
                      spacing={3}
                    >
                      <AudioPlayerWithStaff
                        file={files[0]}
                        phonenumber={"0987654321"}
                        staff={"Hat Nho"}
                        onReset={onReset}
                      />
                    </Stack>
                    <Stack container direction="row" justifyContent="center">
                      <Card sx={{ width: 845 }}>
                        <Typography textAlign="center" p={2}>
                          {" "}
                          Duration: {data.dur}{" "}
                        </Typography>
                      </Card>
                    </Stack>
                    <Stack container direction="row" justifyContent="center">
                      <Emotion
                        title={"Customer Emotion Analystics"}
                        data={data.customer}
                      />
                      <Emotion
                        title={"Staff Emotion Analystics"}
                        data={data.staff}
                      />
                    </Stack>
                  </Stack>
                )
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