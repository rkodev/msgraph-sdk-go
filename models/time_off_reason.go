package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TimeOffReason 
type TimeOffReason struct {
    ChangeTrackedEntity
    // The name of the timeOffReason. Required.
    displayName *string
    // Supported icon types: none; car; calendar; running; plane; firstAid; doctor; notWorking; clock; juryDuty; globe; cup; phone; weather; umbrella; piggyBank; dog; cake; trafficCone; pin; sunny. Required.
    iconType *TimeOffReasonIconType
    // Indicates whether the timeOffReason can be used when creating new entities or updating existing ones. Required.
    isActive *bool
}
// NewTimeOffReason instantiates a new TimeOffReason and sets the default values.
func NewTimeOffReason()(*TimeOffReason) {
    m := &TimeOffReason{
        ChangeTrackedEntity: *NewChangeTrackedEntity(),
    }
    odataTypeValue := "#microsoft.graph.timeOffReason";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTimeOffReasonFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTimeOffReasonFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTimeOffReason(), nil
}
// GetDisplayName gets the displayName property value. The name of the timeOffReason. Required.
func (m *TimeOffReason) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TimeOffReason) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ChangeTrackedEntity.GetFieldDeserializers()
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["iconType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseTimeOffReasonIconType , m.SetIconType)
    res["isActive"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsActive)
    return res
}
// GetIconType gets the iconType property value. Supported icon types: none; car; calendar; running; plane; firstAid; doctor; notWorking; clock; juryDuty; globe; cup; phone; weather; umbrella; piggyBank; dog; cake; trafficCone; pin; sunny. Required.
func (m *TimeOffReason) GetIconType()(*TimeOffReasonIconType) {
    return m.iconType
}
// GetIsActive gets the isActive property value. Indicates whether the timeOffReason can be used when creating new entities or updating existing ones. Required.
func (m *TimeOffReason) GetIsActive()(*bool) {
    return m.isActive
}
// Serialize serializes information the current object
func (m *TimeOffReason) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ChangeTrackedEntity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetIconType() != nil {
        cast := (*m.GetIconType()).String()
        err = writer.WriteStringValue("iconType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isActive", m.GetIsActive())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The name of the timeOffReason. Required.
func (m *TimeOffReason) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIconType sets the iconType property value. Supported icon types: none; car; calendar; running; plane; firstAid; doctor; notWorking; clock; juryDuty; globe; cup; phone; weather; umbrella; piggyBank; dog; cake; trafficCone; pin; sunny. Required.
func (m *TimeOffReason) SetIconType(value *TimeOffReasonIconType)() {
    m.iconType = value
}
// SetIsActive sets the isActive property value. Indicates whether the timeOffReason can be used when creating new entities or updating existing ones. Required.
func (m *TimeOffReason) SetIsActive(value *bool)() {
    m.isActive = value
}
