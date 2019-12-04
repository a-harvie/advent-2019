use std::io::{BufRead, BufReader};
use std::fs::File;
    
fn main() {
    let tests = [
        [
            vec!["R75","D30","R83","U83","L12","D49","R71","U7","L72"],
            vec!["U62","R66","U55","R34","D71","R55","D58","R83"],
            vec!["159"],
        ],
        [
            vec!["R98","U47","R26","D63","R33","U87","L62","D20","R33","U53","R51"],
            vec!["U98","R91","D20","R16","D67","R40","U7","R15","U6","R7"],
            vec!["135"]
        ],
    ];

    for test in tests.iter() {
        println!("Test\n {:?} {:?}\nshould equal\n {:?}", test[0], test[1], test[2]);
        println!("Actual:\n {:?}\n", closest_intersection(&test[0], &test[1]));
    }

    // let input = get_input();
    // println!("Final answer part 1: {}", intcode_compute_input(&input, 12, 2)[0]);
    
}

// fn get_input() -> Vec<i32> {
//     let f = File::open("input/input.txt").unwrap();
//     let reader = BufReader::new(f);
//     let mut values = Vec::new();
    
//     for line in reader.lines() {
//         for symbol in line.unwrap().split(",") {
//             let input: i32 = symbol.parse().unwrap();
//             values.push(input);
//         }
//     }

//     return values;
// }

fn closest_intersection(wire_1: &Vec<&str>, wire_2: &Vec<&str>) -> i32 {
    let path_1 = calculate_path(wire_1);
    let path_2 = calculate_path(wire_2);
    let intersections = get_intersections(path_1, path_2);
    return get_min_distance(intersections);
}

fn calculate_path(wire: &Vec<&str>) -> Vec<[i32;2]> {
    let path = Vec::new();
    for step in wire.iter() {
        let direction = step.chars().nth(0).unwrap();
        let distance: i32 = step[1..].parse().unwrap();
        println!("{:?}", distance);
        match direction {
            'U' => {
                println!("Up!");
            },
            'D' => {
                println!("Down!");
            },
            'L' => {
                println!("Left!");
            },
            'R' => {
                println!("Right!");
            },
            _ => {
                panic!(":c");
            }

        }
    }
    return path;
}

fn get_intersections(path_1: Vec<[i32;2]>, path_2: Vec<[i32;2]>) -> Vec<[i32;2]> {
    let mut intersections = Vec::new();
    return intersections;
}

fn get_min_distance(intersections: Vec<[i32;2]>) -> i32 {
    let min = std::i32::MAX;
    return min;
}

// fn intcode_compute_input(start: &Vec<i32>, noun: i32, verb: i32) -> Vec<i32> {
//     let mut input = start.clone();
//     input[1] = noun;
//     input[2] = verb;
//     return intcode_compute(&input);
// }

// fn intcode_compute(start: &Vec<i32>) -> Vec<i32> {
//     let mut input = start.clone();
//     let mut i = 0;
//     loop {
//         // println!("at input {}",i);
//         let cmd = input[i];
//         if cmd == 99 {
//             // println!("{} means stop", cmd);
//             break;
//         }
        
//         let op_a = input[input[i+1] as usize];
//         let op_b = input[input[i+2] as usize];
//         let output_index = input[i+3] as usize;
//         let mut result: i32;

//         match cmd {
//             1 => {
//                 result = op_a + op_b;
//             },
//             2 => {
//                 result = op_a * op_b;
//             },
//             _ => {
//                 panic!("oh noes, can't do {}", cmd);
//             }
//         }

//         input[output_index] = result;
//         i += 4;
//     }

//     return input;
// }
