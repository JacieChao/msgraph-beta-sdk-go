package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewsDecisionsFilterByCurrentUserWithOnResponse 
// Deprecated: This class is obsolete. Use filterByCurrentUserWithOnGetResponse instead.
type AccessReviewsDecisionsFilterByCurrentUserWithOnResponse struct {
    AccessReviewsDecisionsFilterByCurrentUserWithOnGetResponse
}
// NewAccessReviewsDecisionsFilterByCurrentUserWithOnResponse instantiates a new AccessReviewsDecisionsFilterByCurrentUserWithOnResponse and sets the default values.
func NewAccessReviewsDecisionsFilterByCurrentUserWithOnResponse()(*AccessReviewsDecisionsFilterByCurrentUserWithOnResponse) {
    m := &AccessReviewsDecisionsFilterByCurrentUserWithOnResponse{
        AccessReviewsDecisionsFilterByCurrentUserWithOnGetResponse: *NewAccessReviewsDecisionsFilterByCurrentUserWithOnGetResponse(),
    }
    return m
}
// CreateAccessReviewsDecisionsFilterByCurrentUserWithOnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewsDecisionsFilterByCurrentUserWithOnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReviewsDecisionsFilterByCurrentUserWithOnResponse(), nil
}
// AccessReviewsDecisionsFilterByCurrentUserWithOnResponseable 
// Deprecated: This class is obsolete. Use filterByCurrentUserWithOnGetResponse instead.
type AccessReviewsDecisionsFilterByCurrentUserWithOnResponseable interface {
    AccessReviewsDecisionsFilterByCurrentUserWithOnGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
