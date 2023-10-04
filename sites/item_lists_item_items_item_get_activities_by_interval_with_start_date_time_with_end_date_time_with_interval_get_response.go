package sites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ItemListsItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse 
type ItemListsItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponse
}
// NewItemListsItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse instantiates a new ItemListsItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse and sets the default values.
func NewItemListsItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse()(*ItemListsItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse) {
    m := &ItemListsItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse{
        BaseCollectionPaginationCountResponse: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateItemListsItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemListsItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemListsItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemListsItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateItemActivityStatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemActivityStatable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemActivityStatable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ItemListsItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse) GetValue()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemActivityStatable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemActivityStatable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ItemListsItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemListsItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse) SetValue(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemActivityStatable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// ItemListsItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponseable 
type ItemListsItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponseable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemActivityStatable)
    SetValue(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemActivityStatable)()
}
