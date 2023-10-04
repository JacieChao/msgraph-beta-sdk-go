package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/networkaccess"
)

// ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeTransactionSummariesWithStartDateTimeWithEndDateTimeGetResponse 
type ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeTransactionSummariesWithStartDateTimeWithEndDateTimeGetResponse struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponse
}
// NewReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeTransactionSummariesWithStartDateTimeWithEndDateTimeGetResponse instantiates a new ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeTransactionSummariesWithStartDateTimeWithEndDateTimeGetResponse and sets the default values.
func NewReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeTransactionSummariesWithStartDateTimeWithEndDateTimeGetResponse()(*ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeTransactionSummariesWithStartDateTimeWithEndDateTimeGetResponse) {
    m := &ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeTransactionSummariesWithStartDateTimeWithEndDateTimeGetResponse{
        BaseCollectionPaginationCountResponse: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeTransactionSummariesWithStartDateTimeWithEndDateTimeGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeTransactionSummariesWithStartDateTimeWithEndDateTimeGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeTransactionSummariesWithStartDateTimeWithEndDateTimeGetResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeTransactionSummariesWithStartDateTimeWithEndDateTimeGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.CreateTransactionSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.TransactionSummaryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.TransactionSummaryable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeTransactionSummariesWithStartDateTimeWithEndDateTimeGetResponse) GetValue()([]i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.TransactionSummaryable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.TransactionSummaryable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeTransactionSummariesWithStartDateTimeWithEndDateTimeGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeTransactionSummariesWithStartDateTimeWithEndDateTimeGetResponse) SetValue(value []i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.TransactionSummaryable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeTransactionSummariesWithStartDateTimeWithEndDateTimeGetResponseable 
type ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeTransactionSummariesWithStartDateTimeWithEndDateTimeGetResponseable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.TransactionSummaryable)
    SetValue(value []i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.TransactionSummaryable)()
}
