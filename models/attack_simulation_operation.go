package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttackSimulationOperation 
type AttackSimulationOperation struct {
    LongRunningOperation
    // The attack simulation operation type. Possible values are: createSimulation, updateSimulation, unknownFutureValue.
    TypeEscaped *AttackSimulationOperationType
}
// NewAttackSimulationOperation instantiates a new AttackSimulationOperation and sets the default values.
func NewAttackSimulationOperation()(*AttackSimulationOperation) {
    m := &AttackSimulationOperation{
        LongRunningOperation: *NewLongRunningOperation(),
    }
    return m
}
// CreateAttackSimulationOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttackSimulationOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAttackSimulationOperation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttackSimulationOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.LongRunningOperation.GetFieldDeserializers()
    res["percentageCompleted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPercentageCompleted(val)
        }
        return nil
    }
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAttackSimulationOperationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*AttackSimulationOperationType))
        }
        return nil
    }
    return res
}
// GetPercentageCompleted gets the percentageCompleted property value. Percentage of completion of the respective operation.
func (m *AttackSimulationOperation) GetPercentageCompleted()(*int32) {
    val, err := m.GetBackingStore().Get("percentageCompleted")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTenantId gets the tenantId property value. Tenant identifier.
func (m *AttackSimulationOperation) GetTenantId()(*string) {
    val, err := m.GetBackingStore().Get("tenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetType gets the type property value. The attack simulation operation type. Possible values are: createSimulation, updateSimulation, unknownFutureValue.
func (m *AttackSimulationOperation) GetType()(*AttackSimulationOperationType) {
    val, err := m.GetBackingStore().Get("typeEscaped")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AttackSimulationOperationType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AttackSimulationOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.LongRunningOperation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("percentageCompleted", m.GetPercentageCompleted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPercentageCompleted sets the percentageCompleted property value. Percentage of completion of the respective operation.
func (m *AttackSimulationOperation) SetPercentageCompleted(value *int32)() {
    err := m.GetBackingStore().Set("percentageCompleted", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantId sets the tenantId property value. Tenant identifier.
func (m *AttackSimulationOperation) SetTenantId(value *string)() {
    err := m.GetBackingStore().Set("tenantId", value)
    if err != nil {
        panic(err)
    }
}
// SetType sets the type property value. The attack simulation operation type. Possible values are: createSimulation, updateSimulation, unknownFutureValue.
func (m *AttackSimulationOperation) SetType(value *AttackSimulationOperationType)() {
    err := m.GetBackingStore().Set("typeEscaped", value)
    if err != nil {
        panic(err)
    }
}
// AttackSimulationOperationable 
type AttackSimulationOperationable interface {
    LongRunningOperationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPercentageCompleted()(*int32)
    GetTenantId()(*string)
    GetType()(*AttackSimulationOperationType)
    SetPercentageCompleted(value *int32)()
    SetTenantId(value *string)()
    SetType(value *AttackSimulationOperationType)()
}
