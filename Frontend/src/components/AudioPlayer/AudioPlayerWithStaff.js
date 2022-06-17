import { Card, Typography, Box, Stack } from "@mui/material";
import React from "react";
import AudioPlayer from "./AudioPlayer";

function AudioPlayerWithStaff({ file, onReset, phonenumber, staff }) {
  return (
    <Stack
      spacing={2}
      direction="row"
      alignItems="center"
      sx={{ minWidth: 600 }}
    >
      <Box sx={{ width: "100%" }}>
        <AudioPlayer file={file} onReset={() => onReset} />
      </Box>
      <Stack
        direction="column"
        justifyContent="center"
        alignItems="center"
        spacing={2}
        sx={{ m: 2 }}
      >
        <Card sx={{ minWidth: 180, minHeight: 30 }}>
          <Typography sx={{ m: 1 }}>Phone: {phonenumber}</Typography>
        </Card>
        <Card sx={{ minWidth: 180, minHeight: 30 }}>
          <Typography sx={{ m: 1 }}>Staff: {staff}</Typography>
        </Card>
      </Stack>
    </Stack>
  );
}

export default AudioPlayerWithStaff;
