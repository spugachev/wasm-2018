const fs = require('fs');

async function load(params) {
    var bytes = fs.readFileSync('./sum.wasm');
    const obj = await WebAssembly.instantiate(bytes)
    const res = obj.instance.exports.sum(2, 40);

    console.log("Result from WASM: ", res);          
}

load();