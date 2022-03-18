fn main() {
    /*
    println!(" ***************** Aritméticas *****************");
    println!("\t\t+");
    println!("int");
    println!("{}", 4 + 3 );
    println!("{}", (1 as i64) + (1 as i64));
    println!("{}", (1 as i64) + (1.8 as i64));
    println!("{}", (1.8 as i64) + (1.2 as i64));
    // println!("{}", 15.2 + 4); // error
    println!("float");
    // println!("{}", 1 + (1.22 as f64));
    // println!("{}", (1.22 as f64) + 1);
    println!("{}", (1 as f64) + (1 as f64));
    println!("{}", (1 as f64) + (1.8 as f64));
    println!("{}", (1.8 as f64) + (1.2 as f64));

    println!("\t\t-");
    println!("int");
    println!("{}", 4 + 3 );
    println!("{}", (1 as i64) - (1 as i64));
    println!("{}", (1 as i64) - (1.8 as i64));
    println!("{}", (1.8 as i64) - (1.2 as i64));
    // println!("{}", 15.2 + 4); // error
    println!("float");
    // println!("{}", 1 + (1.22 as f64));
    // println!("{}", (1.22 as f64) + 1);
    println!("{}", (1 as f64) - (1 as f64));
    println!("{}", (1 as f64) - (1.8 as f64));
    println!("{}", (1.8 as f64) - (1.2 as f64));


    println!("\t\t*");
    println!("int");
    println!("{}", 4 * 3 );
    println!("{}", (1 as i64) * (1 as i64));
    println!("{}", (1 as i64) * (1.8 as i64));
    println!("{}", (1.8 as i64) * (1.2 as i64));
    // println!("{}", 15.2 + 4); // error
    println!("float");
    // println!("{}", 1 + (1.22 as f64));
    // println!("{}", (1.22 as f64) + 1);
    println!("{}", (1 as f64) * (1 as f64));
    println!("{}", (1 as f64) * (1.8 as f64));
    println!("{}", (1.8 as f64) * (1.2 as f64));

    println!("\t\t/");
    println!("int");
    println!("{}", 4 / 3 );
    println!("{}", (1 as i64) / (1 as i64));
    println!("{}", (1 as i64) / (1.8 as i64));
    println!("{}", (1.8 as i64) / (1.2 as i64));
    // println!("{}", 15.2 + 4); // error
    println!("float");
    // println!("{}", 1 + (1.22 as f64));
    // println!("{}", (1.22 as f64) + 1);
    println!("{}", (1 as f64) / (1 as f64));
    
    
    println!(" ***************** RELACIONAL *****************");
    println!("\t\t ************************* < *************************");
    println!("\tint");
    println!("{}", 4 < 3 );
    println!("{}", (1 as i64) < (1 as i64));
    println!("aa {}", (1 as i64) < (1.8 as i64));
    println!("{}", (1.8 as i64) < (1.2 as i64));
    
    println!("\tfloat");
    println!("{}", (1 as f64) < (1 as f64));
    println!("{}", (1 as f64) < (1.8 as f64));
    println!("{}", (1.8 as f64) < (1.2 as f64));

    println!("\t\t ************************* > *************************");
    println!("\tint");
    println!("{}", 4 > 3 );
    println!("{}", (1 as i64) > (1 as i64));
    println!("{}", (1 as i64) > (1.8 as i64));
    println!("{}", (1.8 as i64) > (1.2 as i64));
    // println!("{}", 15.2 + 4); // error
    println!("\tfloat");
    // println!("{}", 1 + (1.22 as f64));
    // println!("{}", (1.22 as f64) + 1);
    println!("{}", (1 as f64) > (1 as f64));
    println!("{}", (1 as f64) > (1.8 as f64));
    println!("{}", (1.8 as f64) > (1.2 as f64));

    println!("\t\t ************************* <= *************************");
    println!("\tint");
    println!("{}", 4 < 3 );
    println!("{}", (1 as i64) <= (1 as i64));
    println!("{}", (1 as i64) <= (1.8 as i64));
    println!("{}", (1.8 as i64) <= (1.2 as i64));
    // println!("{}", 15.2 + 4); // error
    println!("\tfloat");
    // println!("{}", 1 + (1.22 as f64));
    // println!("{}", (1.22 as f64) + 1);
    println!("{}", (1 as f64) <= (1 as f64));
    println!("{}", (1 as f64) <= (1.8 as f64));
    println!("{}", (1.8 as f64) <= (1.2 as f64));

    println!("\t\t ************************* >= *************************");
    println!("int");
    println!("{}", 4 * 3 );
    println!("{}", (1 as i64) >= (1 as i64));
    println!("{}", (1 as i64) >= (1.8 as i64));
    println!("{}", (1.8 as i64) >= (1.2 as i64));
    // println!("{}", 15.2 + 4); // error
    println!("float");
    // println!("{}", 1 + (1.22 as f64));
    // println!("{}", (1.22 as f64) + 1);
    println!("{}", (1 as f64) >= (1 as f64));
    println!("{}", (1 as f64) >= (1.8 as f64));
    println!("{}", (1.8 as f64) >= (1.2 as f64));

    println!("\t\t ************************* != *************************");
    println!("int");
    println!("{}", 4 != 3 );
    println!("{}", (1 as i64) != (1 as i64));
    println!("{}", (1 as i64) != (1.8 as i64));
    println!("{}", (1.8 as i64) != (1.2 as i64));
    // println!("{}", 15.2 + 4); // error
    println!("float");
    // println!("{}", 1 + (1.22 as f64));
    // println!("{}", (1.22 as f64) + 1);
    println!("{}", (1 as f64) != (1 as f64));
    println!("{}", (1 as f64) != (1.8 as f64));
    println!("{}", (1.8 as f64) != (1.2 as f64));

    */
    

    println!(" ***************** MODULO *****************");
    println!("\t\t ************************* % *************************");
    println!("\tint");
    println!("{}", 1.0 % (5 as f64));
    println!("{}", (50 as i64) % (5 as i64));
    println!("{}", (140 as i64) % (12 as i64));
    println!("{}", (3 as i64) % (90 as i64));
    
    println!("\tfloat");
    println!("{}", (360 as f64) % (6 as f64));
    println!("{}", (600 as f64) % (2 as f64));
    println!("{}", (56 as f64) % (3 as f64));

}