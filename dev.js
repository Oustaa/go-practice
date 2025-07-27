import { watch } from "node:fs";
import { exec } from "node:child_process";

exec("go run .");

watch("./", { encoding: "buffer" }, (eventType, filename) => {
  exec("go run .");
  console.log(filename.toString());
});
