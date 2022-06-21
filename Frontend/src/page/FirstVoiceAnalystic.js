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
import { uploadFirstVoice } from "../apis/voiceProcessing";
import AudioPlayerWithStaff from "../components/AudioPlayer/AudioPlayerWithStaff";
import Emotion from "../components/Table/Emotion";
import VoicesHistory from "../components/Table/VoicesHistory";
import { getHistoryByPhone } from "../apis/voiceProcessing";
import Template from "../components/Table/Template";
// import VoicesHistory from "../components/Table/VoicesHistory";
const fakeData = {
  customer: {
    emo: {
      happy: 2,
      sad: 2,
      angry: 0,
      calm: 0,
      surprise: 0,
      disgust: 0,
      neutral: 2,
      fear: 2,
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
      surprisE: 0,
      disgust: 0,
      neutral: 2,
      fear: 2,
    },
    gender: "male",
    feel: "happy",
  },
  dur: "12",
};
const INIT_DATA = [
  {
    id: 2,
    phone: "0394691908",
    staffId: 1,
    startTime: "2022-06-20T10:45:00Z",
    duration: "1m",
    staffEmotion: "happy",
    emotion: "bored",
    segments: [],
  },
];
function FirstVoiceAnalystic() {
  const [files, setFiles] = useState([]);
  const [submited, setSubmited] = useState(false);
  const [reset, setReset] = useState(false);
  const [data, setData] = useState(fakeData);
  const [loading, setLoading] = useState(false);
  const [loading2, setLoading2] = useState(false);
  const [phone, setPhone] = useState("0394691908");
  const [staffId, setStaffId] = useState(1);
  const [history, setHistory] = useState(INIT_DATA);
  const pushFile = (file) => {
    setFiles([...files, file]);
  };
  useEffect(() => {
    if (reset) {
      setFiles([]);
      setReset(false);
      setSubmited(false);
      setLoading(false);
      setLoading2(false);
    }
  }, [reset]);

  const onReset = () => {
    setReset(true);
  };

  const onChangePhone = (e) => {
    setPhone(e.target.value);
  };
  const onChangeStaffId = (e) => {
    setStaffId(e.target.value);
  };

  const onSubmit = (file) => {
    // let path = null;
    let path = new FormData();
    if (file) {
      path.append("file", file);
    }
    setLoading(true);
    setSubmited(true);
    uploadFirstVoice(path)
      .then((res) => {
        console.log("res", res.data);
        if (res.data) {
          setData(res.data);
        }
        setLoading(false);
      })
      .catch((err) => {
        console.log(err.message);
        setLoading(false);
      });
    const param = {
      phone: phone,
    };
    setLoading2(true);
    getHistoryByPhone(param)
      .then((res) => {
        setLoading2(false);
        setHistory(res.data.data);
        console.log(res.data.data);
      })
      .catch((err) => {
        console.log(err);
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
              direction="column"
              sx={{ m: 4 }}
            >
              {submited ? (
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
                      <Stack direction="column" justifyContent="center">
                        <Emotion
                          title={"Customer Emotion Analystics"}
                          data={data.customer}
                        />
                        <Template title="Recomendation" />
                      </Stack>
                      <VoicesHistory data={history} isLoading={loading2} />
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
                      // defaultValue={"0987654321"}
                      value={phone}
                      onChange={onChangePhone}
                    />
                    <TextField
                      id="staff"
                      label="staff"
                      // defaultValue={staffId}
                      value={staffId}
                      onChange={onChangeStaffId}
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
export default FirstVoiceAnalystic;
