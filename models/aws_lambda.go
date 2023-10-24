package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AwsLambda 
type AwsLambda struct {
    AwsIdentity
}
// NewAwsLambda instantiates a new awsLambda and sets the default values.
func NewAwsLambda()(*AwsLambda) {
    m := &AwsLambda{
        AwsIdentity: *NewAwsIdentity(),
    }
    odataTypeValue := "#microsoft.graph.awsLambda"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAwsLambdaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAwsLambdaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAwsLambda(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AwsLambda) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AwsIdentity.GetFieldDeserializers()
    res["resource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAwsAuthorizationSystemResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val.(AwsAuthorizationSystemResourceable))
        }
        return nil
    }
    return res
}
// GetResource gets the resource property value. The resource property
func (m *AwsLambda) GetResource()(AwsAuthorizationSystemResourceable) {
    val, err := m.GetBackingStore().Get("resource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AwsAuthorizationSystemResourceable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AwsLambda) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AwsIdentity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetResource sets the resource property value. The resource property
func (m *AwsLambda) SetResource(value AwsAuthorizationSystemResourceable)() {
    err := m.GetBackingStore().Set("resource", value)
    if err != nil {
        panic(err)
    }
}
// AwsLambdaable 
type AwsLambdaable interface {
    AwsIdentityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetResource()(AwsAuthorizationSystemResourceable)
    SetResource(value AwsAuthorizationSystemResourceable)()
}
