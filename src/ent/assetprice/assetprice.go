// Code generated by ent, DO NOT EDIT.

package assetprice

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the assetprice type in the database.
	Label = "asset_price"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBlock holds the string denoting the block field in the database.
	FieldBlock = "block"
	// FieldSignersCount holds the string denoting the signerscount field in the database.
	FieldSignersCount = "signers_count"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldSignature holds the string denoting the signature field in the database.
	FieldSignature = "signature"
	// FieldAsset holds the string denoting the asset field in the database.
	FieldAsset = "asset"
	// FieldChain holds the string denoting the chain field in the database.
	FieldChain = "chain"
	// FieldPair holds the string denoting the pair field in the database.
	FieldPair = "pair"
	// EdgeSigners holds the string denoting the signers edge name in mutations.
	EdgeSigners = "signers"
	// Table holds the table name of the assetprice in the database.
	Table = "asset_prices"
	// SignersTable is the table that holds the signers relation/edge. The primary key declared below.
	SignersTable = "asset_price_signers"
	// SignersInverseTable is the table name for the Signer entity.
	// It exists in this package in order to avoid circular dependency with the "signer" package.
	SignersInverseTable = "signers"
)

// Columns holds all SQL columns for assetprice fields.
var Columns = []string{
	FieldID,
	FieldBlock,
	FieldSignersCount,
	FieldPrice,
	FieldSignature,
	FieldAsset,
	FieldChain,
	FieldPair,
}

var (
	// SignersPrimaryKey and SignersColumn2 are the table columns denoting the
	// primary key for the signers relation (M2M).
	SignersPrimaryKey = []string{"asset_price_id", "signer_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// SignatureValidator is a validator for the "signature" field. It is called by the builders before save.
	SignatureValidator func([]byte) error
)

// OrderOption defines the ordering options for the AssetPrice queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByBlock orders the results by the block field.
func ByBlock(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBlock, opts...).ToFunc()
}

// BySignersCountField orders the results by the signersCount field.
func BySignersCountField(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSignersCount, opts...).ToFunc()
}

// ByPrice orders the results by the price field.
func ByPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrice, opts...).ToFunc()
}

// ByAsset orders the results by the asset field.
func ByAsset(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAsset, opts...).ToFunc()
}

// ByChain orders the results by the chain field.
func ByChain(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChain, opts...).ToFunc()
}

// ByPair orders the results by the pair field.
func ByPair(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPair, opts...).ToFunc()
}

// BySignersCount orders the results by signers count.
func BySignersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSignersStep(), opts...)
	}
}

// BySigners orders the results by signers terms.
func BySigners(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSignersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newSignersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SignersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, SignersTable, SignersPrimaryKey...),
	)
}
