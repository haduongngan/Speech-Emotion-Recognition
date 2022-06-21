import React from "react";
import { styled } from "@mui/material/styles";
import { Card, Grid, Stack, Container, Paper, Box } from "@mui/material";

import FemaleIcon from "@mui/icons-material/Female";
import MaleIcon from "@mui/icons-material/Male";
// import ProductTableHeader from "./ProductTableHeader";
import Title from "../Items/Title";

const Item = styled(Paper)(({ theme }) => ({
  backgroundColor: theme.palette.mode === "dark" ? "#1A2027" : "#fff",
  ...theme.typography.body2,
  padding: theme.spacing(1),
  minHeight: theme.spacing(4),
  fontSize: theme.spacing(2.5),
  textAlign: "left",
  paddingLeft: theme.spacing(4),
  justifyContent: "center",
  paddingTop: theme.spacing(2),
  color: theme.palette.text.secondary,
}));
const Item2 = styled(Paper)(({ theme }) => ({
  backgroundColor: theme.palette.mode === "dark" ? "#1A2027" : "#fff",
  ...theme.typography.body2,
  padding: theme.spacing(1),
  minHeight: theme.spacing(4),
  fontSize: theme.spacing(2.5),
  textAlign: "center",
  // paddingLeft: theme.spacing(4),
  justifyContent: "center",
  paddingTop: theme.spacing(2),
  color: theme.palette.text.secondary,
}));
function Emotion({ data, title }) {
  return (
    <div>
      <Container sx={{ marginTop: 2 }}>
        <Stack mb={5}>
          <Card
            sx={{
              maxWidth: 400,
              minWidth: 240,
              transition: "0.3s",
              boxShadow: "0 8px 40px -12px rgba(0,0,0,0.3)",
              "&:hover": {
                boxShadow: "0 16px 70px -12.125px rgba(0,0,0,0.3)",
              },
              background: "#ffe7f4 !important",
            }}
          >
            <Title title={title} />

            <Box p={1}>
              <Grid
                container
                rowSpacing={1}
                columnSpacing={{ xs: 1, sm: 1, md: 1 }}
              >
                <Grid item xs={6}>
                  <Item>Neutral: {data.emo.neutral}</Item>
                </Grid>
                <Grid item xs={6}>
                  <Item>Calm : {data.emo.calm} </Item>
                </Grid>
                <Grid item xs={6}>
                  <Item>Happy : {data.emo.happy}</Item>
                </Grid>
                <Grid item xs={6}>
                  <Item>Sad : {data.emo.sad}</Item>
                </Grid>
                <Grid item xs={6}>
                  <Item>Angry : {data.emo.angry}</Item>
                </Grid>
                <Grid item xs={6}>
                  <Item>Fear : {data.emo.fear}</Item>
                </Grid>
                <Grid item xs={6}>
                  <Item>Disgust : {data.emo.disgust}</Item>
                </Grid>
                <Grid item xs={6}>
                  <Item>surprise : {data.emo.surprise}</Item>
                </Grid>
                <Grid item xs={2}>
                  <Item2>
                    {" "}
                    {data.gender === "male" ? (
                      <MaleIcon color="primary" />
                    ) : (
                      <FemaleIcon color="secondary" />
                    )}
                  </Item2>
                </Grid>
                <Grid item xs={10}>
                  <Item2>Feeling: {data.feel}</Item2>
                </Grid>
              </Grid>
            </Box>
          </Card>
        </Stack>
      </Container>
    </div>
  );
}

export default Emotion;
