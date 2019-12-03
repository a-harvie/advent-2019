use std::io::{BufRead, BufReader};
use std::fs::File;

fn main() {
    let tests = [
        [
            vec![1,9,10,3,2,3,11,0,99,30,40,50],
            vec![3500,9,10,70,2,3,11,0,99,30,40,50],
        ],
        [
            vec![1,0,0,0,99], 
            vec![2,0,0,0,99]
        ],
        [
            vec![2,3,0,3,99],
            vec![2,3,0,6,99]
            ],
        [
            vec![2,4,4,5,99,0],
            vec![2,4,4,5,99,9801]
        ],
        [
            vec![1,1,1,4,99,5,6,0,99], 
            vec![30,1,1,4,2,5,6,0,99]
        ],
    ];

    for test in tests.iter() {
        println!("Test\n {:?}\nshould equal\n {:?}", test[0], test[1]);
        println!("Actual:\n {:?}\n", intcode_compute(&test[0]));
    }

    let input = get_input();
    // println!("input: {:?}",input);
    println!("Final answer part 1: {}", intcode_compute_input(&input, 12, 2)[0]);

    for noun in 0..100 {
        for verb in 0..100 {
            let result = intcode_compute_input(&input, noun, verb);
            if result[0] == 19690720 {
                println!("Final answer part 2: {}", 100*noun+verb);
                break;
            }
        }
    }
    
}

fn get_input() -> Vec<i32> {
    let f = File::open("input/input.txt").unwrap();
    let reader = BufReader::new(f);
    let mut values = Vec::new();
    
    for line in reader.lines() {
        for symbol in line.unwrap().split(",") {
            let input: i32 = symbol.parse().unwrap();
            values.push(input);
        }
    }

    return values;
}

fn intcode_compute_input(start: &Vec<i32>, noun: i32, verb: i32) -> Vec<i32> {
    let mut input = start.clone();
    input[1] = noun;
    input[2] = verb;
    return intcode_compute(&input);
}

fn intcode_compute(start: &Vec<i32>) -> Vec<i32> {
    let mut input = start.clone();
    let mut i = 0;
    loop {
        // println!("at input {}",i);
        let cmd = input[i];
        if cmd == 99 {
            // println!("{} means stop", cmd);
            break;
        }
        
        let op_a = input[input[i+1] as usize];
        let op_b = input[input[i+2] as usize];
        let output_index = input[i+3] as usize;
        let mut result: i32;

        match cmd {
            1 => {
                result = op_a + op_b;
            },
            2 => {
                result = op_a * op_b;
            },
            _ => {
                panic!("oh noes, can't do {}", cmd);
            }
        }

        input[output_index] = result;
        i += 4;
    }

    return input;
}
