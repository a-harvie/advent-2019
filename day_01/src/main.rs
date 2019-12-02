use std::io::{BufRead, BufReader};
use std::fs::File;

fn main() {
    let tests = [
        [12, 2],
        [14, 2],
        [1969, 654],
        [100756, 33583],
    ];

    let tests2 = [
        [12, 2],
        [14, 2],
        [1969, 966],
        [100756, 50346],
    ];

    let input = get_input();
    let mut sum = 0;

    for test in tests.iter() {
        println!("Test {} should equal {}", test[0], test[1]);
        println!("Actual: {}", fuel_required(test[0]));
    }

    for i in input.iter() {
        sum += fuel_required(*i);
    }
    println!("Final answer part 1: {}", sum);

    for test in tests2.iter() {
        println!("Test {} should equal {}", test[0], test[1]);
        println!("Actual: {}", fuel_required_recursive(test[0]));
    }

    sum = 0;
    for i in input.iter() {
        sum += fuel_required_recursive(*i);
    }

    println!("Final answer part 2: {}", sum);
}

fn get_input() -> Vec<i32> {
    let f = File::open("input.txt").unwrap();
    let reader = BufReader::new(f);
    let mut values = Vec::new();
    
    for line in reader.lines() {
        let input : i32 = line.unwrap().parse().unwrap();
        values.push(input);
    }

    return values;
}

fn fuel_required_recursive(mass: i32) -> i32 {
    let fuel = fuel_required(mass);
    if fuel <= 0 {
        return 0;
    } else {
        return fuel + fuel_required_recursive(fuel);
    }
}

fn fuel_required(mass: i32) -> i32 {
    let div = mass as f64 / 3.0;
    return div as i32 - 2; 
}
