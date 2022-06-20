import { Card, Typography, Box, Stack } from "@mui/material";
import React from "react";
import AudioPlayer from "./AudioPlayer";

function AudioPlayerWithStaff({ file, onReset, phonenumber, staff }) {
  return (
    <Stack direction="row" alignItems="center" justifyContent="center">
      <Box sx={{ width: 600 }}>
        <AudioPlayer file={file} onReset={onReset} />
      </Box>
      <Stack
        spacing={2}
        direction="column"
        justifyContent="center"
        alignItems="center"
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
