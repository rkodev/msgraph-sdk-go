package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedAppConfiguration 
type ManagedAppConfiguration struct {
    ManagedAppPolicy
    // A set of string key and string value pairs to be sent to apps for users to whom the configuration is scoped, unalterned by this service
    customSettings []KeyValuePairable;
}
// NewManagedAppConfiguration instantiates a new managedAppConfiguration and sets the default values.
func NewManagedAppConfiguration()(*ManagedAppConfiguration) {
    m := &ManagedAppConfiguration{
        ManagedAppPolicy: *NewManagedAppPolicy(),
    }
    return m
}
// CreateManagedAppConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedAppConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedAppConfiguration(), nil
}
// GetCustomSettings gets the customSettings property value. A set of string key and string value pairs to be sent to apps for users to whom the configuration is scoped, unalterned by this service
func (m *ManagedAppConfiguration) GetCustomSettings()([]KeyValuePairable) {
    if m == nil {
        return nil
    } else {
        return m.customSettings
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedAppConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ManagedAppPolicy.GetFieldDeserializers()
    res["customSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(KeyValuePairable)
            }
            m.SetCustomSettings(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ManagedAppConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ManagedAppPolicy.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCustomSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomSettings()))
        for i, v := range m.GetCustomSettings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("customSettings", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCustomSettings sets the customSettings property value. A set of string key and string value pairs to be sent to apps for users to whom the configuration is scoped, unalterned by this service
func (m *ManagedAppConfiguration) SetCustomSettings(value []KeyValuePairable)() {
    if m != nil {
        m.customSettings = value
    }
}