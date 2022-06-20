import React from "react";
import { Link } from "react-router-dom";
import AppBar from "@mui/material/AppBar";
import { Tabs, Tab, Box } from "@mui/material";

export default function NavBar() {
  const [value, setValue] = React.useState(0);

  const handleChange = (event, newValue) => {
    setValue(newValue);
  };
  return (
    <div>
      <AppBar position="static">
        {/* <Toolbar>
          <IconButton edge="start" color="inherit" aria-label="menu">
            <MenuIcon />
          </IconButton>
          <Typography variant="h6">Voice Analystic</Typography>
        </Toolbar> */}
        <Box sx={{ width: "100%", bgcolor: "background.paper" }}>
          <Tabs value={value} onChange={handleChange} centered>
            <Tab label="Call Analystic" component={Link} to="/call" />
            <Tab
              label="First Voice Call Analystic"
              component={Link}
              to="/voice"
            />
          </Tabs>
        </Box>
      </AppBar>
    </div>
  );
}
