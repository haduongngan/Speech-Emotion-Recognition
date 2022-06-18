import React from "react";
import { TableRow, TableCell, TextField } from "@mui/material";

function ProductRow({ data }) {
  const { errorDescription, id, image } = data || {};
  return (
    <TableRow>
      <TableCell align={"center"}>{id}</TableCell>
      <TableCell align={"center"}>{errorDescription}</TableCell>
      <TableCell align={"left"}>
        <TextField
          sx={{
            "& legend": {
              display: "none",
            },
          }}
          fullWidth
          placeholder="name"
          {...getFieldProps("name")}
          type="text"
          helperText={errors.name}
          error={Boolean(errors.name)}
        />
      </TableCell>
      <TableCell align={"center"}>
        <TextField
          sx={{
            width: "15ch",
            "& legend": {
              display: "none",
            },
          }}
          placeholder="sku"
          {...getFieldProps("sku")}
          helperText={errors.sku}
          error={Boolean(errors.sku)}
        />
      </TableCell>
      <TableCell align={"center"}>
        <TextField
          className="custom"
          placeholder="color"
          sx={{
            width: "100px",
            "& legend": {
              display: "none",
            },
          }}
          select
          value={getFieldProps("color").value || ""}
          onChange={handleColorChange}
        >

        </TextField>
      </TableCell>
    </TableRow>
  );
}
export default ProductRow;
