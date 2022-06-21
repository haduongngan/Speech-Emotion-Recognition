import React from "react";
import { styled } from "@mui/material/styles";
import { Card, Grid, Stack, Container, Paper, Box } from "@mui/material";
// import ProductTableHeader from "./ProductTableHeader";
import Title from "../Items/Title";

const Item = styled(Paper)(({ theme }) => ({
  backgroundColor: theme.palette.mode === "dark" ? "#1A2027" : "#fff",
  ...theme.typography.body2,
  padding: theme.spacing(2),
  minHeight: theme.spacing(5),
  fontSize: theme.spacing(2.5),
  textAlign: "center",
  justifyContent: "center",
  color: theme.palette.text.secondary,
}));

function Template({ title }) {
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
                <Grid item xs={4}>
                  <Item>Template 1</Item>
                </Grid>
                <Grid item xs={4}>
                  <Item>Template 2</Item>
                </Grid>
                <Grid item xs={4}>
                  <Item>Template 3</Item>
                </Grid>
              </Grid>
            </Box>
          </Card>
        </Stack>
      </Container>
    </div>
  );
}

export default Template;
