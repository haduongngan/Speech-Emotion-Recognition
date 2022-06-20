import React from "react";
import {
  Card,
  Table,
  Stack,
  Container,
  CircularProgress,
  TableBody,
  TableRow,
  TableCell,
} from "@mui/material";
import ProductTableHeader from "./ProductTableHeader";
import Title from "../Items/Title";
// import ProductTableContent from "../components/ProductTableContent";
// import { TABLE_HEAD } from "../types/product";
// import { IProduct } from "../types/product";
// import { INIT_DATA } from "../constants/product";

const TABLE_HEAD = [
  { key: "id", label: "ID" },
  { key: "startTime", label: "Start time" },
  { key: "duration", label: "Duration" },
  { key: "emotion", label: "Emotion" },
  { key: "staffId", label: "Staff Id" },
  { key: "staffEmotion", label: "Staff Emotion" },
];
// const INIT_DATA = {
//   color: 0,
//   errorDescription: "",
//   id: "",
//   image: "",
//   name: "",
//   sku: "",
// };
function VoicesHistory() {
  // const [isLoading, setIsLoading] = useState(false);
  // const [colors, setColors] = useState<IProductColor[]>([]);
  // const [data, setData] = useState<IProduct[]>([]);
  // const [show, setShow] = useState(false);

  return (
    <div>
      <Container sx={{ marginTop: 4 }}>
        <Stack mb={5}>
          <Card
            sx={{
              maxWidth: 500,
              minWidth: 240,
              margin: "auto",
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
              <TableBody sx={{ background: "#FFF !important" }}>
                <TableRow>
                  <TableCell align="center" colSpan={12} sx={{ py: 3 }}>
                    <CircularProgress color="secondary" />
                  </TableCell>
                </TableRow>
              </TableBody>
            </Table>
          </Card>
        </Stack>
      </Container>
    </div>
  );
}

export default VoicesHistory;
