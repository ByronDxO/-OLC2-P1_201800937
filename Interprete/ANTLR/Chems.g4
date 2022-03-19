parser grammar Chems;



options {
  tokenVocab = ChemsLexer;
}


@header {
    import "OLC2/Interprete/interfaces"
    import "OLC2/Interprete/expresion"
    import "OLC2/Interprete/instruction"
    import "OLC2/Interprete/instruction/variable"
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
  : R_PRINTLN TK_PARA expression TK_PARC TK_PUNTOCOMA                 { $instr = instruction.PRINTLN($expression.p, "-1") }
  | R_PRINTLN TK_PARA STRING TK_COMA expression TK_PARC TK_PUNTOCOMA  { $instr = instruction.PRINTLN($expression.p, $STRING.text[1:len($STRING.text)-1]) }
  | instr_declaracion                                                 { $instr = $instr_declaracion.instr}
  | instr_asignacion                                                  { $instr = $instr_asignacion.instr } 
;

/******************************** [DECLARACION][VARIABLE] ********************************/

instr_declaracion returns [interfaces.Instruction instr]
  : R_LET R_MUT ID TK_IGUAL expression TK_PUNTOCOMA                           { $instr = variable.NewDeclaration($ID.text, interfaces.NULL,      $expression.p, true, false, false)   }
  | R_LET R_MUT ID TK_DOSPUNTOS instr_tipo TK_IGUAL expression TK_PUNTOCOMA   { $instr = variable.NewDeclaration($ID.text, $instr_tipo.tipo_exp, $expression.p, true, false, false)   }
  | R_LET ID TK_IGUAL expression TK_PUNTOCOMA                                 { $instr = variable.NewDeclaration($ID.text, interfaces.NULL,      $expression.p, false, false, false)   }
  | R_LET ID TK_DOSPUNTOS instr_tipo TK_IGUAL expression TK_PUNTOCOMA         { $instr = variable.NewDeclaration($ID.text, $instr_tipo.tipo_exp, $expression.p, false, false, false)   }

;

/******************************** [ASIGNACION][VARIABLE] ********************************/

instr_asignacion returns [interfaces.Instruction instr]
  : ID TK_IGUAL expression TK_PUNTOCOMA { $instr = variable.NewAssignment($ID.text,$expression.p) }
;

/******************************** [TIPO] ********************************/
instr_tipo returns [interfaces.TipoExpresion tipo_exp]
  : R_INT       {$tipo_exp = interfaces.INTEGER}
  | R_FLOAT     {$tipo_exp = interfaces.FLOAT}
  | R_STRING    {$tipo_exp = interfaces.STRING}
;


expression returns [interfaces.Expresion p]
  : exp_arit    {$p = $exp_arit.p}
;



exp_arit returns [interfaces.Expresion p]
  : left = exp_arit op=('*'|'/'|'%') right = exp_arit                                                                                                               { $p = expresion.NewOperacion($left.p,$op.text,$right.p,false, "-1", "-1") }
  | TK_PARA left = exp_arit tipo_left=('as f64'|'as i64') TK_PARC op=('*'|'/'|'%') right = exp_arit                                                                 { $p = expresion.NewOperacion($left.p,$op.text,$right.p,false, $tipo_left.text, "-1") }
  | left = exp_arit op=('*'|'/'|'%') TK_PARA right = exp_arit tipo_right=('as f64'|'as i64') TK_PARC                                                                { $p = expresion.NewOperacion($left.p,$op.text,$right.p,false, "-1", $tipo_right.text) }
  | TK_PARA left = exp_arit tipo_left=('as f64'|'as i64') TK_PARC op=('*'|'/'|'%') TK_PARA right = exp_arit tipo_right=('as f64'|'as i64') TK_PARC                  { $p = expresion.NewOperacion($left.p,$op.text,$right.p,false, $tipo_left.text, $tipo_right.text) }
  
  | left = exp_arit op=('+'|'-') right = exp_arit                                                                                                               { $p = expresion.NewOperacion($left.p,$op.text,$right.p,false, "-1", "-1") }   
  | TK_PARA left = exp_arit tipo_left=('as f64'|'as i64') TK_PARC op=('+'|'-') right = exp_arit                                                                 { $p = expresion.NewOperacion($left.p,$op.text,$right.p,false, $tipo_left.text, "-1") }
  | left = exp_arit op=('+'|'-') TK_PARA right = exp_arit tipo_right=('as f64'|'as i64') TK_PARC                                                                { $p = expresion.NewOperacion($left.p,$op.text,$right.p,false, "-1", $tipo_right.text) }
  | TK_PARA left = exp_arit tipo_left=('as f64'|'as i64') TK_PARC op=('+'|'-') TK_PARA right = exp_arit tipo_right=('as f64'|'as i64') TK_PARC                  { $p = expresion.NewOperacion($left.p,$op.text,$right.p,false, $tipo_left.text, $tipo_right.text) }
  
  | left = exp_arit op=('<'|'<='|'>='|'>'|'!=') right = exp_arit                                                                                                { $p = expresion.NewOperacion($left.p,$op.text,$right.p,false, "-1", "-1") }      
  | TK_PARA left = exp_arit tipo_left=('as f64'|'as i64') TK_PARC op=('<'|'<='|'>='|'>'|'!=') right = exp_arit                                                  { $p = expresion.NewOperacion($left.p,$op.text,$right.p,false, $tipo_left.text, "-1") }
  | left = exp_arit op=('<'|'<='|'>='|'>'|'!=') TK_PARA right = exp_arit tipo_right=('as f64'|'as i64') TK_PARC                                                 { $p = expresion.NewOperacion($left.p,$op.text,$right.p,false, "-1", $tipo_right.text) }
  | TK_PARA left = exp_arit tipo_left=('as f64'|'as i64') TK_PARC op=('<'|'<='|'>='|'>'|'!=') TK_PARA right = exp_arit tipo_right=('as f64'|'as i64') TK_PARC   { $p = expresion.NewOperacion($left.p,$op.text,$right.p,false, $tipo_left.text, $tipo_right.text) }
  
  | left = exp_arit op=('&&'|'||') right = exp_arit                                                                                                             { $p = expresion.NewOperacion($left.p,$op.text,$right.p,false, "-1", "-1") }      
  | op=('!'|'-') expression                                                                                                                                     { $p = expresion.NewOperacion($expression.p,$op.text, nil,true, "-1", "-1") }      
  | op=('!'|'-') TK_PARA left = exp_arit tipo_left=('as f64'|'as i64') TK_PARC                                                                                  { $p = expresion.NewOperacion($left.p,$op.text, nil,true, $tipo_left.text, "-1") }      
  

  | primitivo                                                                                                                                                   { $p = $primitivo.p}
  | TK_PARA expression TK_PARC                                                                                                                                  { $p = $expression.p}
;


primitivo returns[interfaces.Expresion p]
    :NUMBER {
            	num,err := strconv.Atoi($NUMBER.text)
                if err!= nil{
                    fmt.Println(err)
                }
            $p = expresion.PRIMITIVO(num,interfaces.INTEGER)
       } 
    |DOUBLE  {  
                num,err := strconv.ParseFloat($DOUBLE.text, 64)
                if err!= nil{
                    fmt.Println(err)
                }
            $p = expresion.PRIMITIVO(num,interfaces.FLOAT)
              }
    |STRING { 
              str:= $STRING.text[1:len($STRING.text)-1]
              $p = expresion.PRIMITIVO(str,interfaces.STRING)}
            
    |BOOLEAN { 
              // str:= $BOOLEAN.text[1:len($BOOLEAN.text)-1]
              $p = expresion.PRIMITIVO($BOOLEAN.text,interfaces.BOOLEAN)
            }

    |ID       { $p = variable.NewIdentifier($ID.text)}
;