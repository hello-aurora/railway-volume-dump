const { createWriteStream } = require("node:fs");
const archiver = require("archiver");

const generateZip = async (volumePath, path) =>
  new Promise((resolve, reject) => {
    const output = createWriteStream(path);
    const archive = archiver("zip");

    output.on("close", () => {
      console.info(`${archive.pointer()} total bytes`);
      console.info("Zip file created successfully.");

      resolve(path);
    });

    archive.on("error", (error) => {
      reject(error);
    });

    archive.pipe(output);
    archive.directory(volumePath, volumePath);
    archive.finalize();
  });

module.exports = { generateZip };
