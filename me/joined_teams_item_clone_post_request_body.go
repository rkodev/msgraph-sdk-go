package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// JoinedTeamsItemClonePostRequestBody provides operations to call the clone method.
type JoinedTeamsItemClonePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The classification property
    classification *string
    // The description property
    description *string
    // The displayName property
    displayName *string
    // The mailNickname property
    mailNickname *string
    // The partsToClone property
    partsToClone *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ClonableTeamParts
    // The visibility property
    visibility *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamVisibilityType
}
// NewJoinedTeamsItemClonePostRequestBody instantiates a new JoinedTeamsItemClonePostRequestBody and sets the default values.
func NewJoinedTeamsItemClonePostRequestBody()(*JoinedTeamsItemClonePostRequestBody) {
    m := &JoinedTeamsItemClonePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateJoinedTeamsItemClonePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateJoinedTeamsItemClonePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewJoinedTeamsItemClonePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *JoinedTeamsItemClonePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetClassification gets the classification property value. The classification property
func (m *JoinedTeamsItemClonePostRequestBody) GetClassification()(*string) {
    return m.classification
}
// GetDescription gets the description property value. The description property
func (m *JoinedTeamsItemClonePostRequestBody) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *JoinedTeamsItemClonePostRequestBody) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *JoinedTeamsItemClonePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["classification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassification(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["mailNickname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailNickname(val)
        }
        return nil
    }
    res["partsToClone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseClonableTeamParts)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartsToClone(val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ClonableTeamParts))
        }
        return nil
    }
    res["visibility"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseTeamVisibilityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisibility(val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamVisibilityType))
        }
        return nil
    }
    return res
}
// GetMailNickname gets the mailNickname property value. The mailNickname property
func (m *JoinedTeamsItemClonePostRequestBody) GetMailNickname()(*string) {
    return m.mailNickname
}
// GetPartsToClone gets the partsToClone property value. The partsToClone property
func (m *JoinedTeamsItemClonePostRequestBody) GetPartsToClone()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ClonableTeamParts) {
    return m.partsToClone
}
// GetVisibility gets the visibility property value. The visibility property
func (m *JoinedTeamsItemClonePostRequestBody) GetVisibility()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamVisibilityType) {
    return m.visibility
}
// Serialize serializes information the current object
func (m *JoinedTeamsItemClonePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("classification", m.GetClassification())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mailNickname", m.GetMailNickname())
        if err != nil {
            return err
        }
    }
    if m.GetPartsToClone() != nil {
        cast := (*m.GetPartsToClone()).String()
        err := writer.WriteStringValue("partsToClone", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetVisibility() != nil {
        cast := (*m.GetVisibility()).String()
        err := writer.WriteStringValue("visibility", &cast)
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
func (m *JoinedTeamsItemClonePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetClassification sets the classification property value. The classification property
func (m *JoinedTeamsItemClonePostRequestBody) SetClassification(value *string)() {
    m.classification = value
}
// SetDescription sets the description property value. The description property
func (m *JoinedTeamsItemClonePostRequestBody) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *JoinedTeamsItemClonePostRequestBody) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetMailNickname sets the mailNickname property value. The mailNickname property
func (m *JoinedTeamsItemClonePostRequestBody) SetMailNickname(value *string)() {
    m.mailNickname = value
}
// SetPartsToClone sets the partsToClone property value. The partsToClone property
func (m *JoinedTeamsItemClonePostRequestBody) SetPartsToClone(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ClonableTeamParts)() {
    m.partsToClone = value
}
// SetVisibility sets the visibility property value. The visibility property
func (m *JoinedTeamsItemClonePostRequestBody) SetVisibility(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamVisibilityType)() {
    m.visibility = value
}