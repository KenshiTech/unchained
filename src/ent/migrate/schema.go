// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AssetPricesColumns holds the columns for the "asset_prices" table.
	AssetPricesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "block", Type: field.TypeUint64, Unique: true},
		{Name: "signers_count", Type: field.TypeUint64, Nullable: true},
		{Name: "price", Type: field.TypeString},
		{Name: "signature", Type: field.TypeBytes},
	}
	// AssetPricesTable holds the schema information for the "asset_prices" table.
	AssetPricesTable = &schema.Table{
		Name:       "asset_prices",
		Columns:    AssetPricesColumns,
		PrimaryKey: []*schema.Column{AssetPricesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "assetprice_block",
				Unique:  true,
				Columns: []*schema.Column{AssetPricesColumns[1]},
			},
		},
	}
	// DataSetsColumns holds the columns for the "data_sets" table.
	DataSetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// DataSetsTable holds the schema information for the "data_sets" table.
	DataSetsTable = &schema.Table{
		Name:       "data_sets",
		Columns:    DataSetsColumns,
		PrimaryKey: []*schema.Column{DataSetsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "dataset_name",
				Unique:  true,
				Columns: []*schema.Column{DataSetsColumns[1]},
			},
		},
	}
	// SignersColumns holds the columns for the "signers" table.
	SignersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "key", Type: field.TypeBytes, Unique: true},
		{Name: "points", Type: field.TypeInt64},
	}
	// SignersTable holds the schema information for the "signers" table.
	SignersTable = &schema.Table{
		Name:       "signers",
		Columns:    SignersColumns,
		PrimaryKey: []*schema.Column{SignersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "signer_key",
				Unique:  true,
				Columns: []*schema.Column{SignersColumns[2]},
			},
		},
	}
	// AssetPriceDataSetColumns holds the columns for the "asset_price_DataSet" table.
	AssetPriceDataSetColumns = []*schema.Column{
		{Name: "asset_price_id", Type: field.TypeInt},
		{Name: "data_set_id", Type: field.TypeInt},
	}
	// AssetPriceDataSetTable holds the schema information for the "asset_price_DataSet" table.
	AssetPriceDataSetTable = &schema.Table{
		Name:       "asset_price_DataSet",
		Columns:    AssetPriceDataSetColumns,
		PrimaryKey: []*schema.Column{AssetPriceDataSetColumns[0], AssetPriceDataSetColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "asset_price_DataSet_asset_price_id",
				Columns:    []*schema.Column{AssetPriceDataSetColumns[0]},
				RefColumns: []*schema.Column{AssetPricesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "asset_price_DataSet_data_set_id",
				Columns:    []*schema.Column{AssetPriceDataSetColumns[1]},
				RefColumns: []*schema.Column{DataSetsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// AssetPriceSignersColumns holds the columns for the "asset_price_Signers" table.
	AssetPriceSignersColumns = []*schema.Column{
		{Name: "asset_price_id", Type: field.TypeInt},
		{Name: "signer_id", Type: field.TypeInt},
	}
	// AssetPriceSignersTable holds the schema information for the "asset_price_Signers" table.
	AssetPriceSignersTable = &schema.Table{
		Name:       "asset_price_Signers",
		Columns:    AssetPriceSignersColumns,
		PrimaryKey: []*schema.Column{AssetPriceSignersColumns[0], AssetPriceSignersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "asset_price_Signers_asset_price_id",
				Columns:    []*schema.Column{AssetPriceSignersColumns[0]},
				RefColumns: []*schema.Column{AssetPricesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "asset_price_Signers_signer_id",
				Columns:    []*schema.Column{AssetPriceSignersColumns[1]},
				RefColumns: []*schema.Column{SignersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AssetPricesTable,
		DataSetsTable,
		SignersTable,
		AssetPriceDataSetTable,
		AssetPriceSignersTable,
	}
)

func init() {
	AssetPriceDataSetTable.ForeignKeys[0].RefTable = AssetPricesTable
	AssetPriceDataSetTable.ForeignKeys[1].RefTable = DataSetsTable
	AssetPriceSignersTable.ForeignKeys[0].RefTable = AssetPricesTable
	AssetPriceSignersTable.ForeignKeys[1].RefTable = SignersTable
}
