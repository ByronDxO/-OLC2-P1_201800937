parser grammar Chems;



options {
  tokenVocab = ChemsLexer;
}


@header {
    import "OLC2/Interprete/interfaces"
    import "OLC2/Interprete/expresion"
    import "OLC2/Interprete/instruction"
    import "OLC2/Interprete/instruction/variable"
    import "OLC2/Interprete/instruction/control"
    import arrayList "github.com/colegno/arraylist"
    
}


start returns [*arrayList.List lista]
  : instrucciones {$lista = $instrucciones.l}
;

instrucciones returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : list +=instruccion+   {
        listInt := localctx.(*InstruccionesContext).GetList()
        for _, e := range listInt {
            $l.Add(e.GetInstr())
        }
    }
;

instruccion returns [interfaces.Instruction instr]
  : R_PRINTLN TK_PARA expression TK_PARC TK_PUNTOCOMA                 { $instr = instruction.PRINTLN($expression.p, "-1") }
  | R_PRINTLN TK_PARA STRING TK_COMA expression TK_PARC TK_PUNTOCOMA  { $instr = instruction.PRINTLN($expression.p, $STRING.text[1:len($STRING.text)-1]) }
  | instr_declaracion                                                 { $instr = $instr_declaracion.instr }
  | instr_asignacion                                                  { $instr = $instr_asignacion.instr  }
  | instr_if                                                          { $instr = $instr_if.instr } 
;

/******************************** [DECLARACION][VARIABLE] ********************************/

instr_declaracion returns [interfaces.Instruction instr]
  : R_LET R_MUT ID TK_IGUAL block_decla TK_PUNTOCOMA                           { $instr = variable.NewDeclaration($ID.text, interfaces.NULL,      $block_decla.p, true, false, false,  $R_LET.line, localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn()) }
  | R_LET R_MUT ID TK_DOSPUNTOS instr_tipo TK_IGUAL block_decla TK_PUNTOCOMA   { $instr = variable.NewDeclaration($ID.text, $instr_tipo.tipo_exp, $block_decla.p, true, false, false,  $R_LET.line, localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn()) }
  | R_LET ID TK_IGUAL block_decla TK_PUNTOCOMA                                 { $instr = variable.NewDeclaration($ID.text, interfaces.NULL,      $block_decla.p, false, false, false, $R_LET.line, localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn()) }
  | R_LET ID TK_DOSPUNTOS instr_tipo TK_IGUAL block_decla TK_PUNTOCOMA         { $instr = variable.NewDeclaration($ID.text, $instr_tipo.tipo_exp, $block_decla.p, false, false, false, $R_LET.line, localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn()) }
  
  // | R_LET R_MUT ID TK_IGUAL instr_ternario TK_PUNTOCOMA                           { $instr = variable.NewDeclaration($ID.text, interfaces.NULL,      $instr_ternario.p, true, false, false,  $R_LET.line, localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn()) }
  // | R_LET R_MUT ID TK_DOSPUNTOS instr_tipo TK_IGUAL instr_ternario TK_PUNTOCOMA   { $instr = variable.NewDeclaration($ID.text, $instr_tipo.tipo_exp, $instr_ternario.p, true, false, false,  $R_LET.line, localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn()) }
  // | R_LET ID TK_IGUAL instr_ternario TK_PUNTOCOMA                                 { $instr = variable.NewDeclaration($ID.text, interfaces.NULL,      $instr_ternario.p, false, false, false, $R_LET.line, localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn()) }
  // | R_LET ID TK_DOSPUNTOS instr_tipo TK_IGUAL instr_ternario TK_PUNTOCOMA         { $instr = variable.NewDeclaration($ID.text, $instr_tipo.tipo_exp, $instr_ternario.p, false, false, false, $R_LET.line, localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn()) }
  
  
;

block_decla returns [interfaces.Expresion p]
  : expression          {$p = $expression.p }
  | instr_ternario      {$p = $instr_ternario.p }
;

/******************************** [ASIGNACION][VARIABLE] ********************************/

instr_asignacion returns [interfaces.Instruction instr]
  : ID TK_IGUAL expression TK_PUNTOCOMA                                       { $instr = variable.NewAssignment($ID.text, $expression.p, $ID.line, localctx.(*Instr_asignacionContext).Get_ID().GetColumn()) }
;

/******************************** [CONTROL][IF] ********************************/
instr_if returns [interfaces.Instruction instr]
  : R_IF expression TK_LLAVEA left_instr=instrucciones TK_LLAVEC                                                       { $instr = control.NewIf($expression.p, $left_instr.l, nil, nil,               $R_IF.line) }
  | R_IF expression TK_LLAVEA left_instr=instrucciones TK_LLAVEC R_ELSE TK_LLAVEA right_intr=instrucciones TK_LLAVEC   { $instr = control.NewIf($expression.p, $left_instr.l, $right_intr.l, nil,     $R_IF.line) }
  | R_IF expression TK_LLAVEA left_instr=instrucciones TK_LLAVEC R_ELSE instr_else_if                                  { $instr = control.NewIf($expression.p, $left_instr.l, nil, $instr_else_if.l,   $R_IF.line) }
;

instr_else_if returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e +=instr_if*  {
        listInt := localctx.(*Instr_else_ifContext).GetE()
        for _, e := range listInt {
            $l.Add(e.GetInstr())
        }
    }
;


/******************************** [CONTROL][TERNARIO] ********************************/
instr_ternario returns [interfaces.Expresion p]
  : R_IF block=expression TK_LLAVEA left_instr=expression TK_LLAVEC                                                       { $p = control.NewTernario($block.p, $left_instr.p, nil, nil,               $R_IF.line) }
  | R_IF block=expression TK_LLAVEA left_instr=expression TK_LLAVEC R_ELSE TK_LLAVEA right_intr=expression TK_LLAVEC      { $p = control.NewTernario($block.p, $left_instr.p, $right_intr.p, nil,     $R_IF.line) }
  | R_IF block=expression TK_LLAVEA left_instr=expression TK_LLAVEC R_ELSE instr_else_if_ternario                         { $p = control.NewTernario($block.p, $left_instr.p, nil, $instr_else_if_ternario.l,   $R_IF.line) }
;

instr_else_if_ternario returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e +=instr_ternario*  {
        listInt := localctx.(*Instr_else_if_ternarioContext).GetE()
        for _, e := range listInt {
            $l.Add(e.GetP())
        }
    }
;



/******************************** [TIPO] ********************************/
instr_tipo returns [interfaces.TipoExpresion tipo_exp]
  : R_INT       {$tipo_exp = interfaces.INTEGER}
  | R_FLOAT     {$tipo_exp = interfaces.FLOAT}
  | R_STRING    {$tipo_exp = interfaces.STRING}
  | R_STR       {$tipo_exp = interfaces.STRING}
  | R_BOOL      {$tipo_exp = interfaces.BOOLEAN}
;


expression returns [interfaces.Expresion p]
  : exp_arit            {$p = $exp_arit.p}
;



exp_arit returns [interfaces.Expresion p]
  : left = exp_arit op=('*'|'/'|'%') right = exp_arit                                                                                                           { $p = expresion.NewOperacion($left.p, $op.text, $right.p, false, "-1", "-1",                        $op.line, localctx.(*Exp_aritContext).GetOp().GetColumn()) }
  | TK_PARA left = exp_arit tipo_left=('as f64'|'as i64') TK_PARC op=('*'|'/'|'%') right = exp_arit                                                             { $p = expresion.NewOperacion($left.p, $op.text, $right.p, false, $tipo_left.text, "-1",             $op.line, localctx.(*Exp_aritContext).GetOp().GetColumn()) }
  | left = exp_arit op=('*'|'/'|'%') TK_PARA right = exp_arit tipo_right=('as f64'|'as i64') TK_PARC                                                            { $p = expresion.NewOperacion($left.p, $op.text, $right.p, false, "-1", $tipo_right.text,            $op.line, localctx.(*Exp_aritContext).GetOp().GetColumn()) }
  | TK_PARA left = exp_arit tipo_left=('as f64'|'as i64') TK_PARC op=('*'|'/'|'%') TK_PARA right = exp_arit tipo_right=('as f64'|'as i64') TK_PARC              { $p = expresion.NewOperacion($left.p, $op.text, $right.p, false, $tipo_left.text, $tipo_right.text, $op.line, localctx.(*Exp_aritContext).GetOp().GetColumn()) }
  | left = exp_arit op=('+'|'-') right = exp_arit                                                                                                               { $p = expresion.NewOperacion($left.p, $op.text, $right.p, false, "-1", "-1",                        $op.line, localctx.(*Exp_aritContext).GetOp().GetColumn()) }   
  | TK_PARA left = exp_arit tipo_left=('as f64'|'as i64') TK_PARC op=('+'|'-') right = exp_arit                                                                 { $p = expresion.NewOperacion($left.p, $op.text, $right.p, false, $tipo_left.text, "-1",             $op.line, localctx.(*Exp_aritContext).GetOp().GetColumn()) }
  | left = exp_arit op=('+'|'-') TK_PARA right = exp_arit tipo_right=('as f64'|'as i64') TK_PARC                                                                { $p = expresion.NewOperacion($left.p, $op.text, $right.p, false, "-1", $tipo_right.text,            $op.line, localctx.(*Exp_aritContext).GetOp().GetColumn()) }
  | TK_PARA left = exp_arit tipo_left=('as f64'|'as i64') TK_PARC op=('+'|'-') TK_PARA right = exp_arit tipo_right=('as f64'|'as i64') TK_PARC                  { $p = expresion.NewOperacion($left.p, $op.text, $right.p, false, $tipo_left.text, $tipo_right.text, $op.line, localctx.(*Exp_aritContext).GetOp().GetColumn()) }
  | left = exp_arit op=('<'|'<='|'>='|'>'|'!='|'==') right = exp_arit                                                                                                { $p = expresion.NewOperacion($left.p, $op.text, $right.p, false, "-1", "-1",                        $op.line, localctx.(*Exp_aritContext).GetOp().GetColumn()) }      
  | TK_PARA left = exp_arit tipo_left=('as f64'|'as i64') TK_PARC op=('<'|'<='|'>='|'>'|'!='|'==') right = exp_arit                                                  { $p = expresion.NewOperacion($left.p, $op.text, $right.p, false, $tipo_left.text, "-1",             $op.line, localctx.(*Exp_aritContext).GetOp().GetColumn()) }
  | left = exp_arit op=('<'|'<='|'>='|'>'|'!='|'==') TK_PARA right = exp_arit tipo_right=('as f64'|'as i64') TK_PARC                                                 { $p = expresion.NewOperacion($left.p, $op.text, $right.p, false, "-1", $tipo_right.text,            $op.line, localctx.(*Exp_aritContext).GetOp().GetColumn()) }
  | TK_PARA left = exp_arit tipo_left=('as f64'|'as i64') TK_PARC op=('<'|'<='|'>='|'>'|'!='|'==') TK_PARA right = exp_arit tipo_right=('as f64'|'as i64') TK_PARC   { $p = expresion.NewOperacion($left.p, $op.text, $right.p, false, $tipo_left.text, $tipo_right.text, $op.line, localctx.(*Exp_aritContext).GetOp().GetColumn()) }
  | left = exp_arit op=('&&'|'||') right = exp_arit                                                                                                             { $p = expresion.NewOperacion($left.p, $op.text, $right.p, false, "-1", "-1",                        $op.line, localctx.(*Exp_aritContext).GetOp().GetColumn()) }   

  | op=('!'|'-') expression                                                                                                                                     { $p = expresion.NewOperacion($expression.p, $op.text, nil,true, "-1", "-1",                         $op.line, localctx.(*Exp_aritContext).GetOp().GetColumn()) }      
  | op=('!'|'-') TK_PARA left = exp_arit tipo_left=('as f64'|'as i64') TK_PARC                                                                                  { $p = expresion.NewOperacion($left.p,       $op.text, nil,true, $tipo_left.text, "-1",              $op.line, localctx.(*Exp_aritContext).GetOp().GetColumn()) }      
  

  | primitivo                                                                                                                                                   { $p = $primitivo.p}
  | TK_PARA expression TK_PARC                                                                                                                                  { $p = $expression.p}
;


primitivo returns[interfaces.Expresion p]
    :NUMBER {
            	num,err := strconv.Atoi($NUMBER.text)
                if err!= nil{
                    fmt.Println(err)
                }

            $p = expresion.PRIMITIVO(num,interfaces.INTEGER, $NUMBER.line, localctx.(*PrimitivoContext).Get_NUMBER().GetColumn())
       } 
    |DOUBLE  {  
                num,err := strconv.ParseFloat($DOUBLE.text, 64)
                if err!= nil{
                    fmt.Println(err)
                }
            $p = expresion.PRIMITIVO(num,interfaces.FLOAT, $DOUBLE.line, localctx.(*PrimitivoContext).Get_DOUBLE().GetColumn())
              }
    |STRING { 
              str:= $STRING.text[1:len($STRING.text)-1]
              $p = expresion.PRIMITIVO(str,interfaces.STRING, $STRING.line, localctx.(*PrimitivoContext).Get_STRING().GetColumn())
            
            }
            
    |BOOLEAN { 
              // str:= $BOOLEAN.text[1:len($BOOLEAN.text)-1]
              exp,_ := strconv.ParseBool($BOOLEAN.text)
              $p = expresion.PRIMITIVO(exp,interfaces.BOOLEAN, $BOOLEAN.line, localctx.(*PrimitivoContext).Get_BOOLEAN().GetColumn())
            }

    |ID       { $p = variable.NewIdentifier($ID.text, $ID.line, localctx.(*PrimitivoContext).Get_ID().GetColumn()) }
    
;