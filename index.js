const { existsSync } = require("node:fs");
const { join } = require("node:path");
const express = require("express");

const { generateZip } = require("./generate-zip");

const PORT = process.env.PORT || 3000;
const VOLUME_NAME = "data";
const VOLUME_PATH = `/${VOLUME_NAME}`;
const ZIP_FILE_NAME = `${VOLUME_NAME}.zip`;

const app = express();

app.get("/download", async (req, res) => {
  const password = req.headers["password"];

  if (!password || !process.env.PASSWORD || password !== process.env.PASSWORD) {
    return res.status(401).json({ error: "Unauthorized" });
  }

  const hasVolume = existsSync(VOLUME_PATH);

  if (!hasVolume) {
    return res
      .status(500)
      .json({ error: `No data found or no volume mounted at ${VOLUME_PATH}.` });
  }

  const path = join(__dirname, ZIP_FILE_NAME);
  const zipPath = await generateZip(VOLUME_PATH, path);

  return res.download(zipPath, ZIP_FILE_NAME);
});

app.listen(PORT, () => {
  console.log(`Server running at http://localhost:${PORT}`);
});
