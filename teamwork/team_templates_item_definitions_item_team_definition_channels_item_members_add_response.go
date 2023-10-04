package teamwork

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddResponse 
// Deprecated: This class is obsolete. Use addPostResponse instead.
type TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddResponse struct {
    TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddPostResponse
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddResponse instantiates a new TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddResponse and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddResponse()(*TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddResponse) {
    m := &TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddResponse{
        TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddPostResponse: *NewTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddPostResponse(),
    }
    return m
}
// CreateTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddResponse(), nil
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddResponseable 
// Deprecated: This class is obsolete. Use addPostResponse instead.
type TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddPostResponseable
}
