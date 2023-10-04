package reports

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetRelyingPartyDetailedSummaryWithPeriodResponse 
// Deprecated: This class is obsolete. Use getRelyingPartyDetailedSummaryWithPeriodGetResponse instead.
type GetRelyingPartyDetailedSummaryWithPeriodResponse struct {
    GetRelyingPartyDetailedSummaryWithPeriodGetResponse
}
// NewGetRelyingPartyDetailedSummaryWithPeriodResponse instantiates a new GetRelyingPartyDetailedSummaryWithPeriodResponse and sets the default values.
func NewGetRelyingPartyDetailedSummaryWithPeriodResponse()(*GetRelyingPartyDetailedSummaryWithPeriodResponse) {
    m := &GetRelyingPartyDetailedSummaryWithPeriodResponse{
        GetRelyingPartyDetailedSummaryWithPeriodGetResponse: *NewGetRelyingPartyDetailedSummaryWithPeriodGetResponse(),
    }
    return m
}
// CreateGetRelyingPartyDetailedSummaryWithPeriodResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetRelyingPartyDetailedSummaryWithPeriodResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetRelyingPartyDetailedSummaryWithPeriodResponse(), nil
}
// GetRelyingPartyDetailedSummaryWithPeriodResponseable 
// Deprecated: This class is obsolete. Use getRelyingPartyDetailedSummaryWithPeriodGetResponse instead.
type GetRelyingPartyDetailedSummaryWithPeriodResponseable interface {
    GetRelyingPartyDetailedSummaryWithPeriodGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
