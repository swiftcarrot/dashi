import React from "react";
import get from "lodash/get";

export default function Table({ keyExtractor, columns, data, ...props }) {
  return (
    <table {...props}>
      <thead>
        <tr>
          {columns.map(({ key, title, render, ...props }) => (
            <th key={key} {...props}>
              {title}
            </th>
          ))}
        </tr>
      </thead>

      <tbody>
        {data.map((row) => (
          <tr key={keyExtractor(row)}>
            {columns.map((col) => (
              <td key={col.key}>{renderCell(col, row)}</td>
            ))}
          </tr>
        ))}
      </tbody>
    </table>
  );
}

function renderCell(col, row) {
  if (col.render) {
    return col.render(row);
  }

  return get(row, col.key);
}
