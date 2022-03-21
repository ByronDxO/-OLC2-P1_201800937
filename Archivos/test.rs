fn main() {
    // Declaraciones, mutaciones y casteos (numéricos)
    let a = 1.2 + 5.4;
    let b;
    b = (5 as f64) + 3.4;
    let mut c: i64 = 2 + 3;

    // Impresiones
    println!("b: {} c: {}", c, b);
    c = (7.8 as i64) + 1;
    println!("La suma de 1.2 + 5.4 es {}", a);
    println!("c mutado es: {} ", c);

    // Strings
    let string1: String = "hello".to_string();
    let string2: &str = "world";
    let string3 = string1 + string2;
    println!("{}", string3);

    // Operaciones aritmeticas
    let sub1: f64 = 1.2 - 5.4;
    let mult1 = 3 * 2;
    let div1 = 1.2/5.4;
    let mod1 = 1.0%5.0;
    println!("sub:{} mul:{} div:{} mod:{}", sub1, mult1, div1, mod1);

    // Casteo con potencia
    let pot = f64::powf(2.5, 2.6);
    println!("{}", (pot as i64));

    // Operaciones Logicas
    let ftrue: bool = 4.3 <= 4.3;
    let ifalse: bool = 4 > 4;
    let sfalse: bool = "hola" > "hola que tal";
    println!("f: {}, i: {}, s: {}", ftrue, ifalse, sfalse);

    let logicOr = ftrue || ifalse;
    let logicAnd = ftrue && ifalse;
    let logicNot = !ftrue;
    println!("or: {}", logicOr);
    println!("and: {}", logicAnd);
    println!("not: {}", logicNot);

    // Condicionales
    // IF
    let x: i64 = 60;
    if x < 61 {
    println!("Reprobado con nota de {}", x);
    } else {
    println!("Aprobado con nota de {}", x);
    }

    // if expresion
    let mut n = 10;
    let mut operacion = if n < 10 {
        10 * n
    } else if n == 10 {
        2 * n
    } else {
        n / 2
    };
    println!("{}", operacion);

    // If expresion 2
    let g = "15".to_string();
    let f = if true { "Numero ".to_string() + "30" } else { g };
    println!("{}", f);


    // If expresion 3
    let algo = if true { false && true } else { true };
    println!("{}", algo);
    
    // Match 
    n = 15;
    match n {
        1 | 2 | 3 => {
        let x = 100;
        println!("Rango de 1 a 3");
        }
        6 | 7 | 8 => println!("Rango 6 a 7"),
        _ => println!("Otro"),
    }
    
    // Match expression
    let p = 4;
    let numString = match p + 1 {
        3 => "Tres",
        4 => "Cuatro",
        5 => "Cinco"
    };
   	println!("{}", numString);
   	
   	// Match expression 2
    let booleano = true;
    let binario = match booleano {
    false => 0,
    true => 1,
    };
    
    println!("{}", binario);

    // While
    let mut var = 2;
    while (var <= 3) {
    println!("hola");
    var = var + 1;
    }

    // While con If
    let mut k = 5;
    let mut completado = false;
    while !completado {
    k = k - 3;
    print!("{}", k);
    if k % 5 == 0 {
    completado = true;
    }
    }

    // Llamada a funciones
    let y = factorial(2);
    println!("{}", y);
    
    println!("true: {}", par(4));
    
    // Declaracion y Asignacion de Arreglos
    let mut arr3: [[[i64; 4];2]; 2] = [
    [ [ 1, 3, 5, 7], [ 9, 11, 13, 15] ],
    [ [ 2, 4, 6, 8], [10, 12, 14, 16] ]
    ];
    println!("{:?}", arr3[0][1][3]);
    arr3[0][1][3] = 50;
    println!("{:?}", arr3[0][1][3]);
    println!("{:?}", arr3);
    
    // Arrays de str y String
    let arr1: [&str; 2] = ["Hola", "Mundo"];
    println!("{:?}", arr1[0]);
    let arr2: [String; 2] = ["Hola".to_string(), "Mundo".to_string()];
    println!("{:?}", arr2[1]);
    
    // Vectores
    let mut v: Vec<i64> = Vec::new();
    v.push(45);
    v.push(20);
    v.push(74);
    println!("{:?}", v);
    println!("{:?}", v[2]);
    
    v = vec![0;10];
    println!("{:?}", v);
}

fn factorial(x: i64) -> i64 {
  if x == 0 {
    1
  } else {
    x * factorial(x-1)
  }
}

fn par(n: i64) -> bool {
  if n == 0 {
    true
  }
  impar(n-1)
}

fn impar(n: i64) -> bool {
  if n == 0 {
    false
  }
  par(n-1)
}


/*

Salidas:
b: 5 c: 8.4
La suma de 1.2 + 5.4 es 6.6000000000000005
c mutado es: 8
helloworld
sub:-4.2 mul:6 div:0.2222222222222222 mod:1
10
f: true, i: false, s: false
or: true
and: false
not: false
Reprobado con nota de 60
20
Numero 30
false
hola
hola
2-1-4-7-102
15
50
[[[1 3 5 7] [9 11 13 50]] [[2 4 6 8] [10 12 14 16]]]
*/
