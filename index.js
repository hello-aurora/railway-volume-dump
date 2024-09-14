const { join } = require("node:path");
const express = require("express");

const { generateZip } = require("./generate-zip");

const PORT = process.env.PORT || 3000;
const { RAILWAY_VOLUME_MOUNT_PATH, RAILWAY_VOLUME_NAME } = process.env;
const ZIP_FILE_NAME = `${RAILWAY_VOLUME_NAME}.zip`;

const app = express();

app.get("/", async (req, res) => {
  const password = req.headers["password"];

  if (!password || !process.env.PASSWORD || password !== process.env.PASSWORD) {
    return res.status(401).json({ error: "Unauthorized" });
  }

  if (!RAILWAY_VOLUME_MOUNT_PATH) {
    return res.status(500).json({
      error: "No volume mounted to this service, please mount a volume first.",
    });
  }

  const path = join(__dirname, ZIP_FILE_NAME);
  const zipPath = await generateZip(RAILWAY_VOLUME_MOUNT_PATH, path);

  return res.download(zipPath, ZIP_FILE_NAME);
});

app.listen(PORT, () => {
  console.log(`Server running at http://localhost:${PORT}`);
});
