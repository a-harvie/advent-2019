use std::io::{BufRead, BufReader};
use std::fs::File;
    
fn main() {
    let tests = [
        [
            vec!["R8","U5","L5","D3"],
            vec!["U7","R6","D4","L4"],
            vec!["6"]
        ],
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

    let input = get_input();
    println!("Final answer part 1: {}", closest_intersection(&input[0], &input[1]));
    
}

fn get_input() -> [Vec<&'static str>; 2] {
    let f = File::open("input/input.txt").unwrap();
    let reader = BufReader::new(f);
    let mut values1 = Vec::new();
    let mut values2 = Vec::new();
    
    for (index,line) in reader.lines().enumerate() {
        if index == 0 {
            values1 = line.unwrap().split(",").collect();
        }
        else {
            values2 = line.unwrap().split(",").collect();
        }
    }
    println!("{:?}",values1);
    return [values1, values2];
}

fn closest_intersection(wire_1: &Vec<&str>, wire_2: &Vec<&str>) -> i32 {
    let path_1 = calculate_path(wire_1);
    let path_2 = calculate_path(wire_2);
    let intersections = get_intersections(path_1, path_2);
    return get_min_distance(intersections);
}

fn calculate_path(wire: &Vec<&str>) -> Vec<[i32;2]> {
    let mut path = Vec::new();
    let mut x = 0;
    let mut y = 0;
    path.push([x,y]);
    for step in wire.iter() {
        let direction = step.chars().nth(0).unwrap();
        let distance: i32 = step[1..].parse().unwrap();
        
        match direction {
            'U' => {
                for yy in (y+1)..=(y+distance) {
                    y = yy;    
                    path.push([x,y]);
                }
            },
            'D' => {
                for yy in ((y-distance)..y).rev() {
                    y = yy;
                    path.push([x,y]);
                }
            },
            'L' => {
                for xx in ((x-distance)..x).rev() {
                    x = xx;
                    path.push([x,y]);
                }
            },
            'R' => {
                for xx in (x+1)..=(x+distance) {
                    x = xx;
                    path.push([x,y]);
                }
            },
            _ => {
                panic!(":c");
            }

        }
    }
    println!("{:?}",path);
    return path;
}

fn get_intersections(path_1: Vec<[i32;2]>, path_2: Vec<[i32;2]>) -> Vec<[i32;2]> {
    let mut intersections = Vec::new();
    for point in path_1.iter() {
        if path_2.contains(point) {
            intersections.push(point.clone());
        }
    }
    for point in path_2.iter() {
        if (path_1.contains(point) && !intersections.contains(point)) {
            intersections.push(point.clone());
        }
    }
    println!("{:?}",intersections);
    return intersections;
}

fn get_min_distance(intersections: Vec<[i32;2]>) -> i32 {
    let mut min = std::i32::MAX;
    for point in intersections.iter() {
        if point[0] == 0 && point[1] == 0 {
            continue;
        }
        let dist = get_manattan_distance([0,0], *point);
        println!("{:?} {}", point, dist);
        if  dist < min {
            min = dist; 
        }
    }
    return min;
}

fn get_manattan_distance(p1: [i32;2], p2: [i32;2]) -> i32 {
    return (p1[0] - p2[0]).abs() + (p1[1] - p2[1]).abs();
}
