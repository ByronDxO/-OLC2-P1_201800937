lexer grammar ChemsLexer;


// Tokens

R_PRINTLN:  'println!';
P_NUMBER:   'number';
P_STRING:   'string';
// P_IF:       'if';
// P_WHILE:    'while';

NUMBER: [0-9]+;
STRING: '"'~["]*'"';
ID: ([a-zA-Z_])[a-zA-Z0-9_]*;

TK_PUNTO:        '.';
TK_PUNTOCOMA:    ';';
TK_COMA:         ',';
// TK_DIFERENTE:    '!';
TK_IGUAL:        '=';
TK_MAYORIGUAL:   '>=';
TK_MENORIGUAL:   '<=';
TK_MAYOR:        '>';
TK_MENOR:        '<';
TK_MULT:         '*';
TK_DIV:          '/';
TK_MAS:          '+';
TK_MENOS:        '-';
TK_PARA:         '(';
TK_PARC:         ')';
TK_LLAVEA:       '{';
TK_LLAVEC:       '}';
TK_CORA:         '[';
TK_CORC:         ']';


WHITESPACE: [ \\\r\n\t]+ -> skip;

fragment
ESC_SEQ
    :   '\\' ('\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'!'|':'|' ')
    ;

