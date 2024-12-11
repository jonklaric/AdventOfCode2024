const arr = Array(100).fill(0);

var fs = require('fs');

try {  
    var data = fs.readFileSync('MyInput.txt', 'utf8');
    //console.log(data.toString());    
} catch(e) {
    console.log('Error:', e.stack);
}

const data_rows = data.split("\r\n")
let left_array = Array(data_rows.length).fill(0);
let right_array = Array(data_rows.length).fill(0);
for (let i=0; i < data_rows.length; i++ ){
    let temp_row = data_rows[i].split("   ");
    left_array[i] = Number(temp_row[0]);
    right_array[i] = Number(temp_row[1]);
}
left_array = left_array.sort();
right_array = right_array.sort();
// console.log(left_array);
// console.log(right_array);
let running_sum = 0
for (let i=0; i < left_array.length; i++){
    running_sum += Math.abs(left_array[i] - right_array[i]);
}
console.log("Total distance between arrays:", running_sum)
