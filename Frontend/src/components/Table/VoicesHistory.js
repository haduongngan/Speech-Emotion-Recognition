import React from "react";
import {
  Card,
  Table,
  Stack,
  Container,
  CircularProgress,
  TableRow,
  TableCell,
} from "@mui/material";
import ProductTableHeader from "./ProductTableHeader";
import Title from "../Items/Title";
import ProductTableContent from "./ProductTableContent";

const TABLE_HEAD = [
  { key: "id", label: "ID" },
  { key: "startTime", label: "Start time" },
  { key: "duration", label: "Duration" },
  { key: "emotion", label: "Emotion" },
  { key: "staffId", label: "Staff Id" },
  { key: "staffEmotion", label: "Staff Emotion" },
];
function VoicesHistory({ data, isLoading }) {

  return (
    <div>
      <Container sx={{ marginTop: 2 }}>
        <Stack mb={5}>
          <Card
            sx={{
              maxWidth: 600,
              minWidth: 240,
              transition: "0.3s",
              boxShadow: "0 8px 40px -12px rgba(0,0,0,0.3)",
              "&:hover": {
                boxShadow: "0 16px 70px -12.125px rgba(0,0,0,0.3)",
              },
              background: "#ffe7f4 !important",
            }}
          >
            <Title title={"Statistical data"} />

            <Table>
              <ProductTableHeader head={TABLE_HEAD} />
              {!isLoading ? (
                <ProductTableContent data={data} />
              ) : (
                <TableRow>
                  <TableCell align="center" colSpan={12} sx={{ py: 3 }}>
                    <CircularProgress color="secondary" />
                  </TableCell>
                </TableRow>
              )}
            </Table>
          </Card>
        </Stack>
      </Container>
    </div>
  );
}

export default VoicesHistory;
