import React, { useEffect, useRef, useState } from "react";
import WaveSurfer from "wavesurfer";
import uuidv4 from "uuid/v4";

// import Avatar from "@mui/material/Avatar";
import Card from "@mui/material/Card";
// import CardMedia from "@mui/material/CardMedia";
// import CardContent from "@mui/material/CardContent";
// import Typography from "@mui/material/Typography";
import IconButton from "@mui/material/IconButton";
import PlayArrowIcon from "@mui/icons-material/PlayArrow";
import StopIcon from "@mui/icons-material/Stop";
// import ChatBubbleIcon from "@mui/icons-material/ChatBubble";
import BackspaceIcon from "@mui/icons-material/Backspace";
// import ShareIcon from "@mui/icons-material/Share";
// import FavoriteIcon from "@mui/icons-material/Favorite";
import { red } from "@mui/material/colors";

import PauseIcon from "@mui/icons-material/Pause";
import Grid from "@mui/material/Grid";

// import List from "@mui/material/List";
// import ListItem from "@mui/material/ListItem";
// import ListItemText from "@mui/material/ListItemText";
// import ListItemAvatar from "@mui/material/ListItemAvatar";

// const faces = [
//   "http://i.pravatar.cc/300?img=1",
//   "http://i.pravatar.cc/300?img=2",
//   "http://i.pravatar.cc/300?img=3",
//   "http://i.pravatar.cc/300?img=4",
// ];

function AudioPlayer({ file, onReset }) {
  const wavesurfer = useRef(null);

  // const [playerReady, setPlayerReady] = useState(false);
  const [isPlaying, setIsPlaying] = useState(false);
  const wavesurferId = `wavesurfer--${uuidv4()}`;

  useEffect(() => {
    wavesurfer.current = WaveSurfer.create({
      container: `#${wavesurferId}`,
      waveColor: "grey",
      progressColor: "tomato",
      height: 70,
      cursorWidth: 1,
      cursorColor: "lightgray",
      barWidth: 2,
      normalize: true,
      responsive: true,
      fillParent: true,
    });

    // const wav = require("../../static/12346 3203.ogg");

    // console.log("wav", wav);
    // wavesurfer.current.load(wav);

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
    // eslint-disable-next-line
  }, []);

  useEffect(() => {
    if (file) {
      if (file.blobURL) wavesurfer.current.load(file.blobURL);
      else {
        const fileUrl = URL.createObjectURL(file);
        wavesurfer.current.load(fileUrl);
      }
    }
  }, [file]);

  const togglePlayback = () => {
    if (!isPlaying) {
      wavesurfer.current.play();
    } else {
      wavesurfer.current.pause();
    }
  };

  const stopPlayback = () => wavesurfer.current.stop();
  const deleteFile = () => {
    onReset();
  };

  let transportPlayButton;

  if (!isPlaying) {
    transportPlayButton = (
      <IconButton onClick={togglePlayback}>
        <PlayArrowIcon />
      </IconButton>
    );
  } else {
    transportPlayButton = (
      <IconButton onClick={togglePlayback}>
        <PauseIcon />
      </IconButton>
    );
  }

  return (
    <>
      <Card
        sx={{
          maxWidth: 600,
          minWidth: 240,
          margin: "auto",
          transition: "0.3s",
          boxShadow: "0 8px 40px -12px rgba(0,0,0,0.3)",
          "&:hover": {
            boxShadow: "0 16px 70px -12.125px rgba(0,0,0,0.3)",
          },
        }}
      >
        <Grid container direction="column">
          <Grid item id={wavesurferId} />
          <Grid item container>
            <Grid item xs={5}>
              {transportPlayButton}
              <IconButton onClick={stopPlayback}>
                <StopIcon />
              </IconButton>
            </Grid>
            <Grid item xs={7} container direction="row-reverse">
              {/* <Grid item>
                <IconButton>
                  <FavoriteIcon
                    style={{ color: blue[500] }}
                    className={classes.icon}
                  />
                </IconButton>
              </Grid>
              <Grid item>
                <IconButton>
                  <ShareIcon
                    style={{ color: red[500] }}
                    className={classes.icon}
                  />
                </IconButton>
              </Grid>
              <Grid item>
                <IconButton>
                  <ChatBubbleIcon
                    style={{ color: green[500] }}
                    className={classes.icon}
                  />
                </IconButton>
              </Grid> */}
              <Grid item>
                <IconButton>
                  <BackspaceIcon
                    onClick={deleteFile}
                    style={{ color: red[500] }}
                  />
                </IconButton>
              </Grid>
            </Grid>
          </Grid>
        </Grid>
      </Card>
    </>
  );
}

export default AudioPlayer;
