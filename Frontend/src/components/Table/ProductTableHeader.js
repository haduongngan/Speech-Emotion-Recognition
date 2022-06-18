import React from "react";
import {
  TableRow,
  Stack,
  TableCell,
  TableHead,
  Typography,
} from "@mui/material";
// import { IProductLabelHeader } from "../types/product";

function ProductTableHeader({ head }) {
  return (
    <TableHead align="center">
      <TableRow sx={{ background: "#ffe7f4 !important" }}>
        {head?.map((headCell) => (
          <TableCell key={headCell.key} align={"center"}>
            <b> {headCell.label}</b>
          </TableCell>
        ))}
      </TableRow>
    </TableHead>
  );
}
export default ProductTableHeader;
