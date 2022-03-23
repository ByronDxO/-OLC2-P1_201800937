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
    import "OLC2/Interprete/instruction/loops"
    import "OLC2/Interprete/instruction/transferencia"
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

end_instr returns [int v]
  : TK_PUNTOCOMA            { $v = 1}
  |                         { $v = 0}
;

instruccion returns [interfaces.Instruction instr]
  : instr_println end_instr       { $instr = $instr_println.instr       }
  | instr_declaracion             { $instr = $instr_declaracion.instr   }
  | instr_asignacion              { $instr = $instr_asignacion.instr    }
  | instr_if                      { $instr = $instr_if.instr            } 
  | instr_match                   { $instr = $instr_match.instr         } 
  | instr_while                   { $instr = $instr_while.instr         }
  | instr_break                   { $instr = $instr_break.instr         }
  | instr_continue                { $instr = $instr_continue.instr      }
  | instr_loop                    { $instr = $instr_loop.instr          }
  | instr_for_in                  { $instr = $instr_for_in.instr          }
;


/******************************** [PRINTLN!] ********************************/
instr_println returns [interfaces.Instruction instr]
  : R_PRINTLN TK_PARA expression TK_PARC TK_PUNTOCOMA                 { $instr = instruction.PRINTLN($expression.p, "-1") }
  | R_PRINTLN TK_PARA STRING TK_COMA expression TK_PARC TK_PUNTOCOMA  { $instr = instruction.PRINTLN($expression.p, $STRING.text[1:len($STRING.text)-1]) }
;  

/******************************** [DECLARACION][VARIABLE] ********************************/
instr_declaracion returns [interfaces.Instruction instr]
  : R_LET R_MUT ID TK_IGUAL expression TK_PUNTOCOMA                           { $instr = variable.NewDeclaration($ID.text, interfaces.NULL,      $expression.p, true, false, false,  $R_LET.line, localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn()) }
  | R_LET R_MUT ID TK_DOSPUNTOS instr_tipo TK_IGUAL expression TK_PUNTOCOMA   { $instr = variable.NewDeclaration($ID.text, $instr_tipo.tipo_exp, $expression.p, true, false, false,  $R_LET.line, localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn()) }
  | R_LET ID TK_IGUAL expression TK_PUNTOCOMA                                 { $instr = variable.NewDeclaration($ID.text, interfaces.NULL,      $expression.p, false, false, false, $R_LET.line, localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn()) }
  | R_LET ID TK_DOSPUNTOS instr_tipo TK_IGUAL expression TK_PUNTOCOMA         { $instr = variable.NewDeclaration($ID.text, $instr_tipo.tipo_exp, $expression.p, false, false, false, $R_LET.line, localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn()) }
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


/******************************** [CONTROL][MATCH] ********************************/

instr_match returns [interfaces.Instruction instr]
  : R_MATCH expression TK_LLAVEA list_case block_default TK_LLAVEC        { $instr = control.NewMatch($expression.p, $list_case.l, $block_default.l, $R_MATCH.line ) }
  | R_MATCH expression TK_LLAVEA block_default TK_LLAVEC                  { $instr = control.NewMatch($expression.p, nil, $block_default.l, $R_MATCH.line) }
;

/*  CASE  */
instr_case returns [interfaces.Expresion instr]
  : list_expre_case TK_IGUALMAYOR TK_LLAVEA instrucciones TK_LLAVEC     { $instr = control.NewCase(nil, $list_expre_case.l, $instrucciones.l) }
  | list_expre_case TK_IGUALMAYOR block_instr_match TK_COMA             { $instr = control.NewCase(nil, $list_expre_case.l, $block_instr_match.l) }
;



block_instr_match returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : list +=instruccion   {
        listInt := localctx.(*Block_instr_matchContext).GetList()
        for _, e := range listInt {
            $l.Add(e.GetInstr())
        }
    }
;

list_case returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e += instr_case+  {
        listInt := localctx.(*List_caseContext).GetE()
        for _, e := range listInt {
            $l.Add(e.GetInstr())
        }
    }
;

/* List Expression Case */
list_expre_case returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e += block_case+  {
        listInt := localctx.(*List_expre_caseContext).GetE()
        for _, e := range listInt {
            $l.Add(e.GetInstr())
        }
    }
;

block_case returns [interfaces.Expresion instr]
  : expression TK_BARRA                                                         { $instr =  control.NewCase($expression.p, nil, nil)}
  | expression                                                                  { $instr =  control.NewCase($expression.p, nil, nil)}
;

/*  DEFAULT  */
instr_default returns [interfaces.Instruction instr]
  : ID TK_IGUALMAYOR TK_LLAVEA instrucciones TK_LLAVEC           { $instr = control.NewDefault($instrucciones.l) }
  | ID TK_IGUALMAYOR block_instr_match TK_COMA                   { $instr = control.NewDefault($block_instr_match.l) }
;

block_default returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e += instr_default+  {
        listInt := localctx.(*Block_defaultContext).GetE()
        for _, e := range listInt {
            $l.Add(e.GetInstr())
        }
    }
;

/******************************** [CONTROL][MATCH][TERNARIO] ********************************/
instr_match_ter returns [interfaces.Expresion instr]
  : R_MATCH expression TK_LLAVEA list_case_ternario instr_default_ter TK_LLAVEC    { $instr = control.NewTerMatch($expression.p, $list_case_ternario.l, $instr_default_ter.instr, $R_MATCH.line ) }
  | R_MATCH expression TK_LLAVEA instr_default_ter TK_LLAVEC                       { $instr = control.NewTerMatch($expression.p, nil, $instr_default_ter.instr, $R_MATCH.line) }
;

/*  CASE  */
instr_case_ter returns [interfaces.Expresion instr]
  : list_expre_case_ter TK_IGUALMAYOR expression TK_COMA                               { $instr = control.NewCaseTer(nil, $list_expre_case_ter.l, $expression.p) }
;


list_case_ternario returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e += instr_case_ter+  {
        listInt := localctx.(*List_case_ternarioContext).GetE()
        for _, e := range listInt {
            $l.Add(e.GetInstr())
        }
    }
;

/* List Expression Case */
list_expre_case_ter returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e += block_case_ter+  {
        listInt := localctx.(*List_expre_case_terContext).GetE()
        for _, e := range listInt {
            $l.Add(e.GetInstr())
        }
    }
;

block_case_ter returns [interfaces.Expresion instr]
  : expression TK_BARRA                                                         { $instr =  control.NewCaseTer($expression.p, nil, nil)}
  | expression                                                                  { $instr =  control.NewCaseTer($expression.p, nil, nil)}
;


/*  DEFAULT  */
instr_default_ter returns [interfaces.Expresion instr]
  : ID TK_IGUALMAYOR expression TK_COMA                   { $instr = control.NewDefaultTer($expression.p) }
;

/******************************** [LOOP][WHILE] ********************************/
instr_while returns [interfaces.Instruction instr]
  : R_WHILE expression TK_LLAVEA instrucciones TK_LLAVEC                           { $instr = loops.NewWhile($expression.p, $instrucciones.l) }
;

/******************************** [LOOP][LOOP] ********************************/
instr_loop returns [interfaces.Instruction instr]
  : R_LOOP TK_LLAVEA instrucciones TK_LLAVEC                           { $instr = loops.NewLoop($instrucciones.l) }
;

/******************************** [LOOP][LOOP][TERNARIO] ********************************/
instr_loop_ternario returns [interfaces.Expresion instr]
  : R_LOOP TK_LLAVEA instrucciones TK_LLAVEC                           { $instr = loops.NewLoopTernario($instrucciones.l) }
;

instr_for_in returns [interfaces.Instruction instr]
  : R_FOR ID R_IN left=expression TK_DOBLEPUNTO right=expression TK_LLAVEA instrucciones TK_LLAVEC     { $instr = loops.NewFor($ID.text, interfaces.INTEGER, $left.p, $right.p, $instrucciones.l, nil, $R_FOR.line, localctx.(*Instr_for_inContext).Get_R_FOR().GetColumn()) }
  | R_FOR ID R_IN left=expression TK_LLAVEA instrucciones TK_LLAVEC                                    { $instr = loops.NewFor($ID.text, interfaces.STRING,  $left.p, nil, $instrucciones.l, nil,      $R_FOR.line, localctx.(*Instr_for_inContext).Get_R_FOR().GetColumn()) }

;



/******************************** [TRANSFERENCIA][BREAK]    ********************************/
instr_break returns [interfaces.Instruction instr]
  : R_BREAK end_instr                               { $instr = transferencia.NewBreak(nil, $R_BREAK.line, localctx.(*Instr_breakContext).Get_R_BREAK().GetColumn()) }
  | R_BREAK expression end_instr                    { $instr = transferencia.NewBreak($expression.p, $R_BREAK.line, localctx.(*Instr_breakContext).Get_R_BREAK().GetColumn()) }
;

/******************************** [TRANSFERENCIA][CONTINUE]  ********************************/
instr_continue returns [interfaces.Instruction instr]
  : R_CONTINUE end_instr                            { $instr = transferencia.NewContinue($R_CONTINUE.line, localctx.(*Instr_continueContext).Get_R_CONTINUE().GetColumn()) }
;

/******************************** [TRANSFERENCIA][RETURN]  ********************************/
instr_return returns [interfaces.Instruction instr]
  : R_RETURN expression end_instr                            { $instr = transferencia.NewReturn($expression.p, $R_RETURN.line, localctx.(*Instr_returnContext).Get_R_RETURN().GetColumn()) }
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
  | TK_PARA expression TK_PARC                                                                                                                          { $p = $expression.p}
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

    | instr_ternario      {$p = $instr_ternario.p } 
    | instr_match_ter     {$p = $instr_match_ter.instr }
    | instr_loop_ternario {$p = $instr_loop_ternario.instr }
    
;