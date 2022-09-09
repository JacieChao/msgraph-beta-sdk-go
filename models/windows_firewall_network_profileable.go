package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsFirewallNetworkProfileable 
type WindowsFirewallNetworkProfileable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthorizedApplicationRulesFromGroupPolicyMerged()(*bool)
    GetAuthorizedApplicationRulesFromGroupPolicyNotMerged()(*bool)
    GetConnectionSecurityRulesFromGroupPolicyMerged()(*bool)
    GetConnectionSecurityRulesFromGroupPolicyNotMerged()(*bool)
    GetFirewallEnabled()(*StateManagementSetting)
    GetGlobalPortRulesFromGroupPolicyMerged()(*bool)
    GetGlobalPortRulesFromGroupPolicyNotMerged()(*bool)
    GetInboundConnectionsBlocked()(*bool)
    GetInboundConnectionsRequired()(*bool)
    GetInboundNotificationsBlocked()(*bool)
    GetInboundNotificationsRequired()(*bool)
    GetIncomingTrafficBlocked()(*bool)
    GetIncomingTrafficRequired()(*bool)
    GetOdataType()(*string)
    GetOutboundConnectionsBlocked()(*bool)
    GetOutboundConnectionsRequired()(*bool)
    GetPolicyRulesFromGroupPolicyMerged()(*bool)
    GetPolicyRulesFromGroupPolicyNotMerged()(*bool)
    GetSecuredPacketExemptionAllowed()(*bool)
    GetSecuredPacketExemptionBlocked()(*bool)
    GetStealthModeBlocked()(*bool)
    GetStealthModeRequired()(*bool)
    GetUnicastResponsesToMulticastBroadcastsBlocked()(*bool)
    GetUnicastResponsesToMulticastBroadcastsRequired()(*bool)
    SetAuthorizedApplicationRulesFromGroupPolicyMerged(value *bool)()
    SetAuthorizedApplicationRulesFromGroupPolicyNotMerged(value *bool)()
    SetConnectionSecurityRulesFromGroupPolicyMerged(value *bool)()
    SetConnectionSecurityRulesFromGroupPolicyNotMerged(value *bool)()
    SetFirewallEnabled(value *StateManagementSetting)()
    SetGlobalPortRulesFromGroupPolicyMerged(value *bool)()
    SetGlobalPortRulesFromGroupPolicyNotMerged(value *bool)()
    SetInboundConnectionsBlocked(value *bool)()
    SetInboundConnectionsRequired(value *bool)()
    SetInboundNotificationsBlocked(value *bool)()
    SetInboundNotificationsRequired(value *bool)()
    SetIncomingTrafficBlocked(value *bool)()
    SetIncomingTrafficRequired(value *bool)()
    SetOdataType(value *string)()
    SetOutboundConnectionsBlocked(value *bool)()
    SetOutboundConnectionsRequired(value *bool)()
    SetPolicyRulesFromGroupPolicyMerged(value *bool)()
    SetPolicyRulesFromGroupPolicyNotMerged(value *bool)()
    SetSecuredPacketExemptionAllowed(value *bool)()
    SetSecuredPacketExemptionBlocked(value *bool)()
    SetStealthModeBlocked(value *bool)()
    SetStealthModeRequired(value *bool)()
    SetUnicastResponsesToMulticastBroadcastsBlocked(value *bool)()
    SetUnicastResponsesToMulticastBroadcastsRequired(value *bool)()
}