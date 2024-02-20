package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ConfigurationPoliciesItemAssignPostResponseable instead.
type ConfigurationPoliciesItemAssignResponse struct {
    ConfigurationPoliciesItemAssignPostResponse
}
// NewConfigurationPoliciesItemAssignResponse instantiates a new ConfigurationPoliciesItemAssignResponse and sets the default values.
func NewConfigurationPoliciesItemAssignResponse()(*ConfigurationPoliciesItemAssignResponse) {
    m := &ConfigurationPoliciesItemAssignResponse{
        ConfigurationPoliciesItemAssignPostResponse: *NewConfigurationPoliciesItemAssignPostResponse(),
    }
    return m
}
// CreateConfigurationPoliciesItemAssignResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConfigurationPoliciesItemAssignResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConfigurationPoliciesItemAssignResponse(), nil
}
// Deprecated: This class is obsolete. Use ConfigurationPoliciesItemAssignPostResponseable instead.
type ConfigurationPoliciesItemAssignResponseable interface {
    ConfigurationPoliciesItemAssignPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
