import React from "react";
import { TableRow, TableCell } from "@mui/material";

function ProductRow({ row }) {
  const { id, startTime, duration, emotion, staffId, staffEmotion } = row || {};
  const convert = (date) => {
    if (date) {
      // 2000-02-08T02:00:00Z
      // 2000-02-08T02:00:00.000Z
      // 2021-11-30T17:00:31.000Z
      let result = date.substring(0, 11);
      result += "02:00:00.000Z";
      return result;
    } else return null;
  };
  const start = new Date(convert(startTime));
  let d = start;
  let dformat =
    [d.getHours(), d.getMinutes(), d.getSeconds()].join(":") +
    " " +
    [d.getMonth() + 1, d.getDate(), d.getFullYear()].join("/");
  console.log(dformat);
  return (
    <TableRow>
      <TableCell align={"center"}>{id}</TableCell>
      <TableCell align={"center"}>{dformat}</TableCell>
      <TableCell align={"center"}>{duration}</TableCell>
      <TableCell align={"center"}>{emotion}</TableCell>
      <TableCell align={"center"}>{staffId}</TableCell>
      <TableCell align={"center"}>{staffEmotion}</TableCell>
    </TableRow>
  );
}
export default ProductRow;
