import { Stack, Typography } from "@mui/material";
import React from "react";

function Title({ title }) {
  return (
    <Stack
      sx={{ width: "100%" }}
      justifyContent={"center"}
      alignItems={"center"}
      spacing={1}
    >
      <Typography variant="h6">{title}</Typography>
    </Stack>
  );
}

export default Title;
