// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/KenshiTech/unchained/ent/correctnessreport"
)

// CorrectnessReport is the model entity for the CorrectnessReport schema.
type CorrectnessReport struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// SignersCount holds the value of the "signersCount" field.
	SignersCount uint64 `json:"signersCount,omitempty"`
	// Timestamp holds the value of the "timestamp" field.
	Timestamp uint64 `json:"timestamp,omitempty"`
	// Signature holds the value of the "signature" field.
	Signature []byte `json:"signature,omitempty"`
	// Hash holds the value of the "hash" field.
	Hash []byte `json:"hash,omitempty"`
	// Topic holds the value of the "topic" field.
	Topic []byte `json:"topic,omitempty"`
	// Correct holds the value of the "correct" field.
	Correct bool `json:"correct,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CorrectnessReportQuery when eager-loading is set.
	Edges        CorrectnessReportEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CorrectnessReportEdges holds the relations/edges for other nodes in the graph.
type CorrectnessReportEdges struct {
	// Signers holds the value of the signers edge.
	Signers []*Signer `json:"signers,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedSigners map[string][]*Signer
}

// SignersOrErr returns the Signers value or an error if the edge
// was not loaded in eager-loading.
func (e CorrectnessReportEdges) SignersOrErr() ([]*Signer, error) {
	if e.loadedTypes[0] {
		return e.Signers, nil
	}
	return nil, &NotLoadedError{edge: "signers"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CorrectnessReport) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case correctnessreport.FieldSignature, correctnessreport.FieldHash, correctnessreport.FieldTopic:
			values[i] = new([]byte)
		case correctnessreport.FieldCorrect:
			values[i] = new(sql.NullBool)
		case correctnessreport.FieldID, correctnessreport.FieldSignersCount, correctnessreport.FieldTimestamp:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CorrectnessReport fields.
func (cr *CorrectnessReport) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case correctnessreport.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cr.ID = int(value.Int64)
		case correctnessreport.FieldSignersCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field signersCount", values[i])
			} else if value.Valid {
				cr.SignersCount = uint64(value.Int64)
			}
		case correctnessreport.FieldTimestamp:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field timestamp", values[i])
			} else if value.Valid {
				cr.Timestamp = uint64(value.Int64)
			}
		case correctnessreport.FieldSignature:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field signature", values[i])
			} else if value != nil {
				cr.Signature = *value
			}
		case correctnessreport.FieldHash:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field hash", values[i])
			} else if value != nil {
				cr.Hash = *value
			}
		case correctnessreport.FieldTopic:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field topic", values[i])
			} else if value != nil {
				cr.Topic = *value
			}
		case correctnessreport.FieldCorrect:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field correct", values[i])
			} else if value.Valid {
				cr.Correct = value.Bool
			}
		default:
			cr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the CorrectnessReport.
// This includes values selected through modifiers, order, etc.
func (cr *CorrectnessReport) Value(name string) (ent.Value, error) {
	return cr.selectValues.Get(name)
}

// QuerySigners queries the "signers" edge of the CorrectnessReport entity.
func (cr *CorrectnessReport) QuerySigners() *SignerQuery {
	return NewCorrectnessReportClient(cr.config).QuerySigners(cr)
}

// Update returns a builder for updating this CorrectnessReport.
// Note that you need to call CorrectnessReport.Unwrap() before calling this method if this CorrectnessReport
// was returned from a transaction, and the transaction was committed or rolled back.
func (cr *CorrectnessReport) Update() *CorrectnessReportUpdateOne {
	return NewCorrectnessReportClient(cr.config).UpdateOne(cr)
}

// Unwrap unwraps the CorrectnessReport entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cr *CorrectnessReport) Unwrap() *CorrectnessReport {
	_tx, ok := cr.config.driver.(*txDriver)
	if !ok {
		panic("ent: CorrectnessReport is not a transactional entity")
	}
	cr.config.driver = _tx.drv
	return cr
}

// String implements the fmt.Stringer.
func (cr *CorrectnessReport) String() string {
	var builder strings.Builder
	builder.WriteString("CorrectnessReport(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cr.ID))
	builder.WriteString("signersCount=")
	builder.WriteString(fmt.Sprintf("%v", cr.SignersCount))
	builder.WriteString(", ")
	builder.WriteString("timestamp=")
	builder.WriteString(fmt.Sprintf("%v", cr.Timestamp))
	builder.WriteString(", ")
	builder.WriteString("signature=")
	builder.WriteString(fmt.Sprintf("%v", cr.Signature))
	builder.WriteString(", ")
	builder.WriteString("hash=")
	builder.WriteString(fmt.Sprintf("%v", cr.Hash))
	builder.WriteString(", ")
	builder.WriteString("topic=")
	builder.WriteString(fmt.Sprintf("%v", cr.Topic))
	builder.WriteString(", ")
	builder.WriteString("correct=")
	builder.WriteString(fmt.Sprintf("%v", cr.Correct))
	builder.WriteByte(')')
	return builder.String()
}

// NamedSigners returns the Signers named value or an error if the edge was not
// loaded in eager-loading with this name.
func (cr *CorrectnessReport) NamedSigners(name string) ([]*Signer, error) {
	if cr.Edges.namedSigners == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := cr.Edges.namedSigners[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (cr *CorrectnessReport) appendNamedSigners(name string, edges ...*Signer) {
	if cr.Edges.namedSigners == nil {
		cr.Edges.namedSigners = make(map[string][]*Signer)
	}
	if len(edges) == 0 {
		cr.Edges.namedSigners[name] = []*Signer{}
	} else {
		cr.Edges.namedSigners[name] = append(cr.Edges.namedSigners[name], edges...)
	}
}

// CorrectnessReports is a parsable slice of CorrectnessReport.
type CorrectnessReports []*CorrectnessReport
