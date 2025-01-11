var fs = require('fs');

try {  
    var data = fs.readFileSync('my_input.txt', 'utf8');
    //console.log(data.toString());    
} catch(e) {
    console.log('Error:', e.stack);
}

class fileObject {
    constructor(fileSize, fileNumber) {
        this.size = fileSize;
        this.fileNumber = fileNumber;
    }

    get fileNumber() {
        return this.fileNumber;
    }
}

let dataArray = data.split("");
const sum_digits = dataArray.reduce((partialSum, a) => partialSum + Number(a), 0);
let diskArray = Array(sum_digits).fill(-1);
var fileNumber = 0;
let diskPosition = 0;
for (let i=0; i<dataArray.length; i++) {
    if (i%2 != 0){
        diskPosition += Number(dataArray[i]);
    } else {
        for (let j=diskPosition; j < diskPosition+Number(dataArray[i]); j++) {
            diskArray[j] = fileNumber;
        }
        diskPosition += Number(dataArray[i]);
        fileNumber += 1;
    }
}
// initialize start pointer/position
var outputArray = [...diskArray];
var fillPosition = 0; // start of array
var fetchPosition = outputArray.length - 1; // last position in array
while (fillPosition < fetchPosition) {
    if (outputArray[fillPosition] >= 0) {
        fillPosition += 1;
        continue;
    }
    if (outputArray[fetchPosition] < 0) {
        fetchPosition -= 1;
        continue;
    }
    outputArray[fillPosition] = Number(outputArray[fetchPosition]);
    outputArray[fetchPosition] = -1;
    fillPosition++;
    fetchPosition -= 1;
}
console.log("Final array...");
console.log(outputArray);
let checkSum = 0;
for (let i = 0; i<outputArray.length; i++) {
    if (outputArray[i] >= 0) {
        checkSum += i*outputArray[i];
    }
}
console.log("Checksum =", checkSum);
