import { Stack, TextField } from "@mui/material";
import React from "react";

function InputField() {
  return (
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
      <TextField id="staff" label="staff" defaultValue={"Hat Nho"} />
    </Stack>
  );
}

export default InputField;
