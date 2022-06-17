import React, { useEffect, useState } from "react";

import PropTypes from "prop-types";
import { styled } from "@mui/material/styles";
import Button from "@mui/material/Button";
import Grid from "@mui/material/Grid";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
// import DoneIcon from "@mui/icons-material/Done";
import CloseIcon from "@mui/icons-material/Close";
import IconButton from "@mui/material/IconButton";
import CloudUploadIcon from "@mui/icons-material/CloudUpload";
// import DialogContentText from "@mui/material/DialogContentText";
import DialogTitle from "@mui/material/DialogTitle";
// import { green } from "@mui/material/colors";
import FileUpload from "react-material-file-upload";

const BootstrapDialog = styled(Dialog)(({ theme }) => ({
  "& .MuiDialogContent-root": {
    padding: theme.spacing(2),
  },
  "& .MuiDialogActions-root": {
    padding: theme.spacing(1),
  },
}));

const BootstrapDialogTitle = (props) => {
  const { children, onClose, ...other } = props;

  return (
    <DialogTitle sx={{ m: 0, p: 2 }} {...other}>
      {children}
      {onClose ? (
        <IconButton
          aria-label="close"
          onClick={onClose}
          sx={{
            position: "absolute",
            right: 8,
            top: 8,
            color: (theme) => theme.palette.grey[500],
          }}
        >
          <CloseIcon />
        </IconButton>
      ) : null}
    </DialogTitle>
  );
};

BootstrapDialogTitle.propTypes = {
  children: PropTypes.node,
  onClose: PropTypes.func.isRequired,
};

function UploadFile({ pushFile, onReset }) {
  const [open, setOpen] = useState(false);
  const [files, setFiles] = useState([]);
  const [disable, setDisable] = useState(true);
  const [tempFile, setTempFile] = useState(null);

  const handleClickOpen = () => {
    setOpen(true);
    onReset();
  };
  const handleCancel = () => {
    setOpen(false);
  };
  const handleDone = () => {
    if (tempFile) {
      pushFile(tempFile);
      setOpen(false);
    }
  };

  useEffect(() => {
    // console.log("file", file);

    setTempFile(files[0]);
    if (tempFile) {
      setDisable(false);
    } else {
      setDisable(true);
    }
  }, [files, tempFile]);
  return (
    <>
      <Grid>
        <IconButton onClick={handleClickOpen}>
          <CloudUploadIcon fontSize="large" />
        </IconButton>
      </Grid>
      <BootstrapDialog maxWidth="sm" open={open} onClose={handleCancel}>
        <BootstrapDialogTitle
          id="customized-dialog-title"
          onClose={handleCancel}
        >
          Upload file audio
        </BootstrapDialogTitle>

        <DialogContent>
          <FileUpload value={files} onChange={setFiles} />
        </DialogContent>
        <DialogActions>
          <Grid container justifyContent="flex-end">
            <Button onClick={handleDone} variant="contained" disabled={disable}>
              save
            </Button>
          </Grid>
        </DialogActions>
      </BootstrapDialog>
    </>
  );
}
export default UploadFile;
