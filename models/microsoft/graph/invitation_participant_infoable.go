package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// InvitationParticipantInfoable 
type InvitationParticipantInfoable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetEndpointType()(*EndpointType)
    GetHidden()(*bool)
    GetIdentity()(IdentitySetable)
    GetParticipantId()(*string)
    GetRemoveFromDefaultAudioRoutingGroup()(*bool)
    GetReplacesCallId()(*string)
    SetEndpointType(value *EndpointType)()
    SetHidden(value *bool)()
    SetIdentity(value IdentitySetable)()
    SetParticipantId(value *string)()
    SetRemoveFromDefaultAudioRoutingGroup(value *bool)()
    SetReplacesCallId(value *string)()
}