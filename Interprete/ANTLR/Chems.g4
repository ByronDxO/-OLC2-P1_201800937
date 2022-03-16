parser grammar Chems;



options {
  tokenVocab = ChemsLexer;
}


@header {
    import "OLC2/Interprete/interfaces"
    import "OLC2/Interprete/expresion"
    import "OLC2/Interprete/instruction"
    import arrayList "github.com/colegno/arraylist"
    
}


start returns [*arrayList.List lista]
  : instrucciones {$lista = $instrucciones.l}
;

instrucciones returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e +=instruccion*  {
        listInt := localctx.(*InstruccionesContext).GetE()
        for _, e := range listInt {
            $l.Add(e.GetInstr())
        }
    }
;

instruccion returns [interfaces.Instruction instr]
  : R_PRINTLN TK_PARA expression TK_PARC TK_PUNTOCOMA { $instr =  instruction.PRINTLN($expression.p)}
;

expression returns [interfaces.Expresion p]
  : exp_arit    {$p = $exp_arit.p}
;



exp_arit returns [interfaces.Expresion p]
  : left = exp_arit op=('*'|'/') right = exp_arit            { $p = expresion.NewOperacion($left.p,$op.text,$right.p,false)}
  | left = exp_arit op=('+'|'-') right = exp_arit            { $p = expresion.NewOperacion($left.p,$op.text,$right.p,false)}     
  | left = exp_arit op=('<'|'<='|'>='|'>') right = exp_arit  { $p = expresion.NewOperacion($left.p,$op.text,$right.p,false)}     
  | primitivo                                               { $p = $primitivo.p}
  | TK_PARA expression TK_PARC                              { $p = $expression.p}
;


primitivo returns[interfaces.Expresion p]
    :NUMBER {
            	num,err := strconv.Atoi($NUMBER.text)
                if err!= nil{
                    fmt.Println(err)
                }
            $p = expresion.PRIMITIVO(num,interfaces.INTEGER)
       } 
    | DOUBLE  {  
                num,err := strconv.ParseFloat($DOUBLE.text, 64)
                if err!= nil{
                    fmt.Println(err)
                }
            $p = expresion.PRIMITIVO(num,interfaces.FLOAT)
              }
    | STRING { 
      str:= $STRING.text[1:len($STRING.text)-1]
      $p = expresion.PRIMITIVO(str,interfaces.STRING)}
    
;