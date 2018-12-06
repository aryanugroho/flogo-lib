// Code generated by gocc; DO NOT EDIT.

package parser

import "github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/direction"

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : Flogo	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Flogo : BaseExpr	<<  >>`,
		Id:         "Flogo",
		NTType:     1,
		Index:      1,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Flogo : TernaryExpr	<<  >>`,
		Id:         "Flogo",
		NTType:     1,
		Index:      2,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Flogo : IfExpr	<<  >>`,
		Id:         "Flogo",
		NTType:     1,
		Index:      3,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `BaseExpr : Func	<<  >>`,
		Id:         "BaseExpr",
		NTType:     2,
		Index:      4,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `BaseExpr : Expr	<<  >>`,
		Id:         "BaseExpr",
		NTType:     2,
		Index:      5,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `BaseExpr : Param	<<  >>`,
		Id:         "BaseExpr",
		NTType:     2,
		Index:      6,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Func : function_name "(" Args ")"	<< direction.NewFunction(X[0], X[2]) >>`,
		Id:         "Func",
		NTType:     3,
		Index:      7,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewFunction(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Func : function_name "()"	<< direction.NewFunction(X[0], "") >>`,
		Id:         "Func",
		NTType:     3,
		Index:      8,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewFunction(X[0], "")
		},
	},
	ProdTabEntry{
		String: `Args : ExprParam	<< direction.NewArgument(X[0]) >>`,
		Id:         "Args",
		NTType:     4,
		Index:      9,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewArgument(X[0])
		},
	},
	ProdTabEntry{
		String: `Args : Args "," Args	<< direction.NewArguments(X[0], X[2]) >>`,
		Id:         "Args",
		NTType:     4,
		Index:      10,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewArguments(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : ExprArg Operator ExprArg	<< direction.NewExpression(X[0], X[1], X[2]) >>`,
		Id:         "Expr",
		NTType:     5,
		Index:      11,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewExpression(X[0], X[1], X[2])
		},
	},
	ProdTabEntry{
		String: `ExprArg : "(" Expr ")"	<< direction.NewExpressionField(X[1]) >>`,
		Id:         "ExprArg",
		NTType:     6,
		Index:      12,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewExpressionField(X[1])
		},
	},
	ProdTabEntry{
		String: `ExprArg : Expr	<< direction.NewExpressionField(X[0]) >>`,
		Id:         "ExprArg",
		NTType:     6,
		Index:      13,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewExpressionField(X[0])
		},
	},
	ProdTabEntry{
		String: `ExprArg : ExprParam	<< direction.NewExpressionField(X[0]) >>`,
		Id:         "ExprArg",
		NTType:     6,
		Index:      14,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewExpressionField(X[0])
		},
	},
	ProdTabEntry{
		String: `Operator : operator_charactor	<<  >>`,
		Id:         "Operator",
		NTType:     7,
		Index:      15,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `TernaryExpr : BaseExpr "?" BaseExpr ":" BaseExpr	<< direction.NewTernaryExpression(X[0], X[2], X[4]) >>`,
		Id:         "TernaryExpr",
		NTType:     8,
		Index:      16,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewTernaryExpression(X[0], X[2], X[4])
		},
	},
	ProdTabEntry{
		String: `IfExpr : "if" BaseExpr Block	<< direction.NewIFF(X[1], X[2]) >>`,
		Id:         "IfExpr",
		NTType:     9,
		Index:      17,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewIFF(X[1], X[2])
		},
	},
	ProdTabEntry{
		String: `IfExpr : "if" BaseExpr Block ElseStatement	<< direction.NewIFFElse(X[1], X[2], X[3]) >>`,
		Id:         "IfExpr",
		NTType:     9,
		Index:      18,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewIFFElse(X[1], X[2], X[3])
		},
	},
	ProdTabEntry{
		String: `Block : "{" Statement "}"	<< direction.NewStatement(X[1]) >>`,
		Id:         "Block",
		NTType:     10,
		Index:      19,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewStatement(X[1])
		},
	},
	ProdTabEntry{
		String: `ElseStatement : "else" IfExpr	<< direction.NewStatement(X[1]) >>`,
		Id:         "ElseStatement",
		NTType:     11,
		Index:      20,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewStatement(X[1])
		},
	},
	ProdTabEntry{
		String: `ElseStatement : "else" Block	<< direction.NewStatement(X[1]) >>`,
		Id:         "ElseStatement",
		NTType:     11,
		Index:      21,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewStatement(X[1])
		},
	},
	ProdTabEntry{
		String: `Statement : IfExpr	<<  >>`,
		Id:         "Statement",
		NTType:     12,
		Index:      22,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statement : BaseExpr	<<  >>`,
		Id:         "Statement",
		NTType:     12,
		Index:      23,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statement : TernaryExpr	<<  >>`,
		Id:         "Statement",
		NTType:     12,
		Index:      24,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ExprParam : Param	<< direction.NewExpressionField(X[0]) >>`,
		Id:         "ExprParam",
		NTType:     13,
		Index:      25,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewExpressionField(X[0])
		},
	},
	ProdTabEntry{
		String: `ExprParam : Func	<< direction.NewExpressionField(X[0]) >>`,
		Id:         "ExprParam",
		NTType:     13,
		Index:      26,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewExpressionField(X[0])
		},
	},
	ProdTabEntry{
		String: `Param : Int	<<  >>`,
		Id:         "Param",
		NTType:     14,
		Index:      27,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Param : Float	<<  >>`,
		Id:         "Param",
		NTType:     14,
		Index:      28,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Param : Bool	<<  >>`,
		Id:         "Param",
		NTType:     14,
		Index:      29,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Param : DoubleQString	<<  >>`,
		Id:         "Param",
		NTType:     14,
		Index:      30,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Param : SingleQString	<<  >>`,
		Id:         "Param",
		NTType:     14,
		Index:      31,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Param : MappingRef	<<  >>`,
		Id:         "Param",
		NTType:     14,
		Index:      32,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Param : Nil	<<  >>`,
		Id:         "Param",
		NTType:     14,
		Index:      33,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `DoubleQString : doublequotes_string	<< direction.NewDoubleQuoteStringLit(X[0]) >>`,
		Id:         "DoubleQString",
		NTType:     15,
		Index:      34,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewDoubleQuoteStringLit(X[0])
		},
	},
	ProdTabEntry{
		String: `SingleQString : singlequote_string	<< direction.NewSingleQuoteStringLit(X[0]) >>`,
		Id:         "SingleQString",
		NTType:     16,
		Index:      35,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewSingleQuoteStringLit(X[0])
		},
	},
	ProdTabEntry{
		String: `Int : number	<< direction.NewIntLit(X[0]) >>`,
		Id:         "Int",
		NTType:     17,
		Index:      36,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewIntLit(X[0])
		},
	},
	ProdTabEntry{
		String: `MappingRef : argument	<< direction.NewMappingRef(X[0]) >>`,
		Id:         "MappingRef",
		NTType:     18,
		Index:      37,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewMappingRef(X[0])
		},
	},
	ProdTabEntry{
		String: `Bool : "true"	<< direction.NewBool(X[0]) >>`,
		Id:         "Bool",
		NTType:     19,
		Index:      38,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewBool(X[0])
		},
	},
	ProdTabEntry{
		String: `Bool : "false"	<< direction.NewBool(X[0]) >>`,
		Id:         "Bool",
		NTType:     19,
		Index:      39,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewBool(X[0])
		},
	},
	ProdTabEntry{
		String: `Float : float	<< direction.NewFloatLit(X[0]) >>`,
		Id:         "Float",
		NTType:     20,
		Index:      40,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewFloatLit(X[0])
		},
	},
	ProdTabEntry{
		String: `Nil : "nil"	<< direction.NewNilLit(X[0]) >>`,
		Id:         "Nil",
		NTType:     21,
		Index:      41,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewNilLit(X[0])
		},
	},
	ProdTabEntry{
		String: `Nil : "null"	<< direction.NewNilLit(X[0]) >>`,
		Id:         "Nil",
		NTType:     21,
		Index:      42,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewNilLit(X[0])
		},
	},
}
