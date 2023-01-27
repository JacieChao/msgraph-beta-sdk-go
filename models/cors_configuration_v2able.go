package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CorsConfiguration_v2able 
type CorsConfiguration_v2able interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedHeaders()([]string)
    GetAllowedMethods()([]string)
    GetAllowedOrigins()([]string)
    GetMaxAgeInSeconds()(*int32)
    GetResource()(*string)
    SetAllowedHeaders(value []string)()
    SetAllowedMethods(value []string)()
    SetAllowedOrigins(value []string)()
    SetMaxAgeInSeconds(value *int32)()
    SetResource(value *string)()
}