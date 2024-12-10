import fs from "fs";

console.log("Hello AoC 2024!\n");

const test_input = fs.readFileSync("test_input.txt", "utf8");

function generateBinaryStates(n){
    let states = [];
    // Convert to decimal
    let maxDecimal = parseInt("1".repeat(n),2);
    // For every number between 0->decimal
    for(let i = 0; i <= maxDecimal; i++){
        // Convert to binary, pad with 0, and add to final results
        states.push(i.toString(2).padStart(n,'0'));
    }
    return states;
}

function generateTrinaryStates(n){
    let states = [];
    // Convert to decimal
    let maxDecimal = parseInt("2".repeat(n),3);
    // For every number between 0->decimal
    for(let i = 0; i <= maxDecimal; i++){
        // Convert to binary, pad with 0, and add to final results
        states.push(i.toString(3).padStart(n,'0'));
    }
    return states;
}

function day7_part1(string_input) {
    let final_result = 0;
    let lines = string_input.split("\r\n");

    //console.log(lines);

    const results = [];

    for (const line of lines) {
        const [eqn_result, eqn_rest] = line.split(": ");
        const eqn_parts = eqn_rest.split(" ");
        const nums = eqn_parts.map(part => parseInt(part));
        let states = generateBinaryStates(eqn_parts.length - 1);
        let matching_state = "";
        let matches = false;
        for (var i=0; i < states.length; i++)
        {
            let temp_result = nums[0];
            for (let j=1; j < eqn_parts.length; j++) {
                
                if (states[i][j-1] == "1") {
                    temp_result *= nums[j];
                }
                else {
                    temp_result += nums[j];
                }
            }
            if (temp_result == eqn_result) {
                matches = true;
                matching_state = states[i];
                break;
            }
        }
        if (matches) {
            //console.log(line, "with", matching_state, "matches!");
            final_result += parseInt(eqn_result);
        }
    }
    return final_result;
}


function day7_part2(string_input) {
    let final_result = 0;
    let lines = string_input.split("\r\n");

    //console.log(lines);

    const results = [];

    for (const line of lines) {
        const [eqn_result, eqn_rest] = line.split(": ");
        const eqn_parts = eqn_rest.split(" ");
        //const nums = eqn_parts.map(part => parseInt(part));
        let states = generateTrinaryStates(eqn_parts.length - 1);
        let matching_state = "";
        let matches = false;
        for (var i=0; i < states.length; i++)
        {
            let temp_result = eqn_parts[0];
            let temp_number = 0;
            for (let j=1; j < eqn_parts.length; j++) {

                if (states[i][j-1] == "2") {
                    temp_result = Math.pow(temp_result, nums[j]);
                }
                else if (states[i][j-1] == "1") {
                    temp_result *= nums[j];
                }
                else {
                    temp_result += nums[j];
                }
            }
            if (temp_result == eqn_result) {
                matches = true;
                matching_state = states[i];
                break;
            }
        }
        if (matches) {
            //console.log(line, "with", matching_state, "matches!");
            final_result += parseInt(eqn_result);
        }
    }
    return final_result;
}


console.log(day7_part1(test_input));

const my_input = fs.readFileSync("my_input.txt", "utf8");
console.log(day7_part1(my_input));
