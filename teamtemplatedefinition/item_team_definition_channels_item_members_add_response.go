package teamtemplatedefinition

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemTeamDefinitionChannelsItemMembersAddResponse 
// Deprecated: This class is obsolete. Use addPostResponse instead.
type ItemTeamDefinitionChannelsItemMembersAddResponse struct {
    ItemTeamDefinitionChannelsItemMembersAddPostResponse
}
// NewItemTeamDefinitionChannelsItemMembersAddResponse instantiates a new ItemTeamDefinitionChannelsItemMembersAddResponse and sets the default values.
func NewItemTeamDefinitionChannelsItemMembersAddResponse()(*ItemTeamDefinitionChannelsItemMembersAddResponse) {
    m := &ItemTeamDefinitionChannelsItemMembersAddResponse{
        ItemTeamDefinitionChannelsItemMembersAddPostResponse: *NewItemTeamDefinitionChannelsItemMembersAddPostResponse(),
    }
    return m
}
// CreateItemTeamDefinitionChannelsItemMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemTeamDefinitionChannelsItemMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamDefinitionChannelsItemMembersAddResponse(), nil
}
// ItemTeamDefinitionChannelsItemMembersAddResponseable 
// Deprecated: This class is obsolete. Use addPostResponse instead.
type ItemTeamDefinitionChannelsItemMembersAddResponseable interface {
    ItemTeamDefinitionChannelsItemMembersAddPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
