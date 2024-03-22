package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"
	"encoding/hex"

	"github.com/KenshiTech/unchained/ent"
	"github.com/KenshiTech/unchained/ent/correctnessreport"
	"github.com/KenshiTech/unchained/ent/signer"
)

// Topic is the resolver for the topic field.
func (r *correctnessReportWhereInputResolver) Topic(ctx context.Context, obj *ent.CorrectnessReportWhereInput, data *string) error {
	if obj == nil || data == nil {
		return nil
	}

	bytes, err := hex.DecodeString(*data)

	if err != nil {
		return err
	}

	obj.AddPredicates(correctnessreport.TopicEQ(bytes))

	return nil
}

// Hash is the resolver for the hash field.
func (r *correctnessReportWhereInputResolver) Hash(ctx context.Context, obj *ent.CorrectnessReportWhereInput, data *string) error {
	if obj == nil || data == nil {
		return nil
	}

	bytes, err := hex.DecodeString(*data)

	if err != nil {
		return err
	}

	obj.AddPredicates(correctnessreport.HashEQ(bytes))

	return nil
}

// Key is the resolver for the key field.
func (r *signerWhereInputResolver) Key(ctx context.Context, obj *ent.SignerWhereInput, data *string) error {
	if obj == nil || data == nil {
		return nil
	}

	bytes, err := hex.DecodeString(*data)

	if err != nil {
		return err
	}

	obj.AddPredicates(signer.KeyEQ(bytes))

	return nil
}
