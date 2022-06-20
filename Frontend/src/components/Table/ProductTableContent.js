import React from "react";
import { TableBody } from "@mui/material";

import ProductRow from "./ProductRow";

function ProductTableContent({
  content,
  colors,
  changes,
  setChanges,
  setError,
}) {
  return (
    <TableBody>
      {content.map((row, index) => (
        <ProductRow
          key={index}
          row={row}
          colors={colors}
          changes={changes}
          setChanges={setChanges}
          setError={setError}
        />
      ))}
    </TableBody>
  );
}

export default ProductTableContent;
