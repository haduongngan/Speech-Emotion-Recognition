import React, { useState, useRef, useEffect } from "react";
import { ReactMic } from "react-mic";
import WaveSurfer from "wavesurfer";

import MicIcon from "@mui/icons-material/Mic";
import IconButton from "@mui/material/IconButton";
import StopIcon from "@mui/icons-material/Stop";
import ReplayIcon from "@mui/icons-material/Replay";
import FiberManualRecordIcon from "@mui/icons-material/FiberManualRecord";
import DoneIcon from "@mui/icons-material/Done";
import CancelIcon from "@mui/icons-material/Cancel";
import PlayArrowIcon from "@mui/icons-material/PlayArrow";
import PauseIcon from "@mui/icons-material/Pause";

// import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
// import DialogContentText from "@mui/material/DialogContentText";
import DialogTitle from "@mui/material/DialogTitle";
import Grid from "@mui/material/Grid";
import { green, red } from "@mui/material/colors";

import "./microphone.css";

export default function Microphone({ pushFile, onReset }) {
  const [record, setRecord] = useState(false);
  const [open, setOpen] = React.useState(false);
  const [tempFile, setTempFile] = React.useState(null);

  // const [playerReady, setPlayerReady] = useState(false);
  const [isPlaying, setIsPlaying] = useState(false);
  const wavesurfer = useRef(null);

  useEffect(() => {
    if (!open || (open && !tempFile)) return;

    wavesurfer.current = WaveSurfer.create({
      container: "#wavesurfer-id",
      waveColor: "grey",
      progressColor: "tomato",
      height: 140,
      cursorWidth: 1,
      cursorColor: "lightgrey",
      barWidth: 2,
      normalize: true,
      responsive: true,
      fillParent: true,
    });

    // wavesurfer.current.on("ready", () => {
    //   setPlayerReady(true);
    // });

    const handleResize = wavesurfer.current.util.debounce(() => {
      wavesurfer.current.empty();
      wavesurfer.current.drawBuffer();
    }, 150);

    wavesurfer.current.on("play", () => setIsPlaying(true));
    wavesurfer.current.on("pause", () => setIsPlaying(false));
    window.addEventListener("resize", handleResize, false);
  }, [open, tempFile]);

  useEffect(() => {
    console.log("tempFile", tempFile);
    if (tempFile) {
      wavesurfer.current.load(tempFile.blobURL);
    }
  }, [tempFile]);

  const togglePlayback = () => {
    if (!isPlaying) {
      wavesurfer.current.play();
    } else {
      wavesurfer.current.pause();
    }
  };
  const stopPlayback = () => wavesurfer.current.stop();

  const handleClickOpen = () => {
    setOpen(true);
    onReset();
  };

  const handleDone = () => {
    if (tempFile) {
      pushFile(tempFile);
      setTempFile(null);
      setRecord(false);
      setOpen(false);
      console.log(tempFile);
    }
  };

  const handleCancel = () => {
    setRecord(false);
    setTempFile(null);
    setOpen(false);
  };

  const startRecording = () => {
    setTempFile(null);
    setRecord(true);
  };

  const stopRecording = () => {
    setRecord(false);
  };

  const onData = (recordedBlob) => {
    //console.log("chunk of real-time data is: ", recordedBlob);
  };

  const onStop = (recordedBlob) => {
    setTempFile(recordedBlob);
  };

  return (
    <>
      <Grid>
        <IconButton onClick={handleClickOpen}>
          <MicIcon fontSize="large" />
        </IconButton>
      </Grid>
      <Dialog maxWidth="sm" open={open} onClose={handleCancel}>
        <DialogTitle>Record</DialogTitle>
        <DialogContent
          sx={{
            ".childClass": {
              width: "100%",
            },
          }}
        >
          {tempFile ? (
            <div style={{ width: "100%" }} id="wavesurfer-id" />
          ) : (
            <ReactMic
              record={record}
              className="childClass"
              onStop={onStop}
              onData={onData}
              strokeColor="grey"
              backgroundColor="white"
            />
          )}
        </DialogContent>
        <DialogActions>
          <Grid container justify="center" alignItems="center">
            {tempFile && (
              <Grid container justify="center" alignItems="center" xs={12}>
                {!isPlaying ? (
                  <IconButton onClick={togglePlayback}>
                    <PlayArrowIcon fontSize="large" />
                  </IconButton>
                ) : (
                  <IconButton onClick={togglePlayback}>
                    <PauseIcon fontSize="large" />
                  </IconButton>
                )}
                <IconButton onClick={stopPlayback}>
                  <StopIcon fontSize="large" />
                </IconButton>
              </Grid>
            )}
            <Grid container justify="center" alignItems="center">
              {!record && !tempFile && (
                <IconButton onClick={startRecording}>
                  <FiberManualRecordIcon
                    fontSize="large"
                    style={{ color: red[500] }}
                  />
                </IconButton>
              )}

              {!record && tempFile && (
                <IconButton onClick={startRecording}>
                  <ReplayIcon fontSize="large" />
                </IconButton>
              )}

              {record && (
                <IconButton onClick={stopRecording}>
                  <StopIcon fontSize="large" />
                </IconButton>
              )}

              <IconButton onClick={handleDone}>
                <DoneIcon
                  style={tempFile && !record ? { color: green[500] } : {}}
                />
              </IconButton>
              <IconButton onClick={handleCancel}>
                <CancelIcon
                  style={tempFile && !record ? { color: red[500] } : {}}
                />
              </IconButton>
            </Grid>
          </Grid>
        </DialogActions>
      </Dialog>
    </>
  );
}
