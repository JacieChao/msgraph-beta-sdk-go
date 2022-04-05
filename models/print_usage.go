package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrintUsage 
type PrintUsage struct {
    Entity
    // The completedBlackAndWhiteJobCount property
    completedBlackAndWhiteJobCount *int64;
    // The completedColorJobCount property
    completedColorJobCount *int64;
    // The incompleteJobCount property
    incompleteJobCount *int64;
    // The usageDate property
    usageDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly;
}
// NewPrintUsage instantiates a new printUsage and sets the default values.
func NewPrintUsage()(*PrintUsage) {
    m := &PrintUsage{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrintUsageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrintUsageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrintUsage(), nil
}
// GetCompletedBlackAndWhiteJobCount gets the completedBlackAndWhiteJobCount property value. The completedBlackAndWhiteJobCount property
func (m *PrintUsage) GetCompletedBlackAndWhiteJobCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.completedBlackAndWhiteJobCount
    }
}
// GetCompletedColorJobCount gets the completedColorJobCount property value. The completedColorJobCount property
func (m *PrintUsage) GetCompletedColorJobCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.completedColorJobCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintUsage) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["completedBlackAndWhiteJobCount"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedBlackAndWhiteJobCount(val)
        }
        return nil
    }
    res["completedColorJobCount"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedColorJobCount(val)
        }
        return nil
    }
    res["incompleteJobCount"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncompleteJobCount(val)
        }
        return nil
    }
    res["usageDate"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsageDate(val)
        }
        return nil
    }
    return res
}
// GetIncompleteJobCount gets the incompleteJobCount property value. The incompleteJobCount property
func (m *PrintUsage) GetIncompleteJobCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.incompleteJobCount
    }
}
// GetUsageDate gets the usageDate property value. The usageDate property
func (m *PrintUsage) GetUsageDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.usageDate
    }
}
// Serialize serializes information the current object
func (m *PrintUsage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("completedBlackAndWhiteJobCount", m.GetCompletedBlackAndWhiteJobCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("completedColorJobCount", m.GetCompletedColorJobCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("incompleteJobCount", m.GetIncompleteJobCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("usageDate", m.GetUsageDate())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCompletedBlackAndWhiteJobCount sets the completedBlackAndWhiteJobCount property value. The completedBlackAndWhiteJobCount property
func (m *PrintUsage) SetCompletedBlackAndWhiteJobCount(value *int64)() {
    if m != nil {
        m.completedBlackAndWhiteJobCount = value
    }
}
// SetCompletedColorJobCount sets the completedColorJobCount property value. The completedColorJobCount property
func (m *PrintUsage) SetCompletedColorJobCount(value *int64)() {
    if m != nil {
        m.completedColorJobCount = value
    }
}
// SetIncompleteJobCount sets the incompleteJobCount property value. The incompleteJobCount property
func (m *PrintUsage) SetIncompleteJobCount(value *int64)() {
    if m != nil {
        m.incompleteJobCount = value
    }
}
// SetUsageDate sets the usageDate property value. The usageDate property
func (m *PrintUsage) SetUsageDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    if m != nil {
        m.usageDate = value
    }
}
