package models
// Android For Work VPN connection type.
type AndroidForWorkVpnConnectionType int

const (
    // Cisco AnyConnect.
    CISCOANYCONNECT_ANDROIDFORWORKVPNCONNECTIONTYPE AndroidForWorkVpnConnectionType = iota
    // Pulse Secure.
    PULSESECURE_ANDROIDFORWORKVPNCONNECTIONTYPE
    // F5 Edge Client.
    F5EDGECLIENT_ANDROIDFORWORKVPNCONNECTIONTYPE
    // Dell SonicWALL Mobile Connection.
    DELLSONICWALLMOBILECONNECT_ANDROIDFORWORKVPNCONNECTIONTYPE
    // Check Point Capsule VPN.
    CHECKPOINTCAPSULEVPN_ANDROIDFORWORKVPNCONNECTIONTYPE
    // Citrix
    CITRIX_ANDROIDFORWORKVPNCONNECTIONTYPE
)

func (i AndroidForWorkVpnConnectionType) String() string {
    return []string{"ciscoAnyConnect", "pulseSecure", "f5EdgeClient", "dellSonicWallMobileConnect", "checkPointCapsuleVpn", "citrix"}[i]
}
func ParseAndroidForWorkVpnConnectionType(v string) (any, error) {
    result := CISCOANYCONNECT_ANDROIDFORWORKVPNCONNECTIONTYPE
    switch v {
        case "ciscoAnyConnect":
            result = CISCOANYCONNECT_ANDROIDFORWORKVPNCONNECTIONTYPE
        case "pulseSecure":
            result = PULSESECURE_ANDROIDFORWORKVPNCONNECTIONTYPE
        case "f5EdgeClient":
            result = F5EDGECLIENT_ANDROIDFORWORKVPNCONNECTIONTYPE
        case "dellSonicWallMobileConnect":
            result = DELLSONICWALLMOBILECONNECT_ANDROIDFORWORKVPNCONNECTIONTYPE
        case "checkPointCapsuleVpn":
            result = CHECKPOINTCAPSULEVPN_ANDROIDFORWORKVPNCONNECTIONTYPE
        case "citrix":
            result = CITRIX_ANDROIDFORWORKVPNCONNECTIONTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAndroidForWorkVpnConnectionType(values []AndroidForWorkVpnConnectionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AndroidForWorkVpnConnectionType) isMultiValue() bool {
    return false
}
