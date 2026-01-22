package types

import "cosmossdk.io/collections"

const (
	// ModuleName defines the module name
	ModuleName = "projectbit"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// GovModuleName duplicates the gov module's name to avoid a dependency with x/gov.
	// It should be synced with the gov module's name if it is ever changed.
	// See: https://github.com/cosmos/cosmos-sdk/blob/v0.52.0-beta.2/x/gov/types/keys.go#L9
	GovModuleName = "gov"

	// Event Types
    PostCreatedEventType = "post_created"
    PostUpdatedEventType = "post_updated"
    PostDeletedEventType = "post_deleted"

	// Attribute Keys
    PostCreatorAttribute = "creator"
    PostIdAttribute      = "post_id"
    PostTitleAttribute   = "title"
)

// ParamsKey is the prefix to retrieve all Params
var ParamsKey = collections.NewPrefix("p_projectbit")

var (
	PostKey      = collections.NewPrefix("post/value/")
	PostCountKey = collections.NewPrefix("post/count/")
)
