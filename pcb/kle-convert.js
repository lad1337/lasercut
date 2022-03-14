#!/usr/bin/env node

const fs = require('fs');
const path = require('path');
var kle = require("@ijprest/kle-serial");

function convert(argv) {
    const srcFile = argv.slice(-1)[0];
    try {
        const data = fs.readFileSync(srcFile, 'utf8');
        var keyboard = kle.Serial.parse(data);

        const extName = path.extname(srcFile);
        const fileName = path.basename(srcFile, extName);
        const dirName = path.dirname(srcFile);
        const dest = path.join(dirName, fileName + "-converted" + extName)
        try {
            console.log("Writing file to " + dest);
            fs.writeFileSync(dest, JSON.stringify(keyboard))
        } catch (err) {
            console.error(err)
        }
    } catch (err) {
        console.error(err)
    }
}

convert(process.argv);
