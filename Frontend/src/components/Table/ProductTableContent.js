import React from "react";
import { TableBody } from "@mui/material";

import ProductRow from "./ProductRow";

function ProductTableContent({ data }) {
  return (
    <TableBody sx={{ background: "#FFF !important" }}>
      {data.map((row, index) => (
        <ProductRow key={index} row={row} />
      ))}
    </TableBody>
  );
}

export default ProductTableContent;
