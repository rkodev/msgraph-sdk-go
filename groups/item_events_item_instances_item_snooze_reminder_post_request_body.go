package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemEventsItemInstancesItemSnoozeReminderPostRequestBody provides operations to call the snoozeReminder method.
type ItemEventsItemInstancesItemSnoozeReminderPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The NewReminderTime property
    newReminderTime iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DateTimeTimeZoneable
}
// NewItemEventsItemInstancesItemSnoozeReminderPostRequestBody instantiates a new ItemEventsItemInstancesItemSnoozeReminderPostRequestBody and sets the default values.
func NewItemEventsItemInstancesItemSnoozeReminderPostRequestBody()(*ItemEventsItemInstancesItemSnoozeReminderPostRequestBody) {
    m := &ItemEventsItemInstancesItemSnoozeReminderPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemEventsItemInstancesItemSnoozeReminderPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemEventsItemInstancesItemSnoozeReminderPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemEventsItemInstancesItemSnoozeReminderPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemEventsItemInstancesItemSnoozeReminderPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemEventsItemInstancesItemSnoozeReminderPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["newReminderTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewReminderTime(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DateTimeTimeZoneable))
        }
        return nil
    }
    return res
}
// GetNewReminderTime gets the newReminderTime property value. The NewReminderTime property
func (m *ItemEventsItemInstancesItemSnoozeReminderPostRequestBody) GetNewReminderTime()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DateTimeTimeZoneable) {
    return m.newReminderTime
}
// Serialize serializes information the current object
func (m *ItemEventsItemInstancesItemSnoozeReminderPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("newReminderTime", m.GetNewReminderTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemEventsItemInstancesItemSnoozeReminderPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetNewReminderTime sets the newReminderTime property value. The NewReminderTime property
func (m *ItemEventsItemInstancesItemSnoozeReminderPostRequestBody) SetNewReminderTime(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DateTimeTimeZoneable)() {
    m.newReminderTime = value
}