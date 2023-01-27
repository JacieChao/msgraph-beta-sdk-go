package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsDeviceTimelineEventsable 
type UserExperienceAnalyticsDeviceTimelineEventsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceId()(*string)
    GetEventAdditionalInformation()(*string)
    GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEventDetails()(*string)
    GetEventLevel()(*DeviceEventLevel)
    GetEventName()(*string)
    GetEventSource()(*string)
    SetDeviceId(value *string)()
    SetEventAdditionalInformation(value *string)()
    SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEventDetails(value *string)()
    SetEventLevel(value *DeviceEventLevel)()
    SetEventName(value *string)()
    SetEventSource(value *string)()
}