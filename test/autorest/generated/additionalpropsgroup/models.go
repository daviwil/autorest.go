// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package additionalpropsgroup

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CatApTrue struct {
	PetApTrue
	Friendly *bool `json:"friendly,omitempty"`
}

func (c CatApTrue) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if c.ID != nil {
		objectMap["id"] = c.ID
	}
	if c.Name != nil {
		objectMap["name"] = c.Name
	}
	if c.Status != nil {
		objectMap["status"] = c.Status
	}
	if c.Friendly != nil {
		objectMap["friendly"] = c.Friendly
	}
	if c.AdditionalProperties != nil {
		for k, v := range *c.AdditionalProperties {
			objectMap[k] = v
		}
	}
	return json.Marshal(objectMap)
}

func (c *CatApTrue) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for k, v := range rawMsg {
		var err error
		switch k {
		case "id":
			if v != nil {
				err = json.Unmarshal(*v, &c.ID)
			}
		case "name":
			if v != nil {
				err = json.Unmarshal(*v, &c.Name)
			}
		case "status":
			if v != nil {
				err = json.Unmarshal(*v, &c.Status)
			}
		case "friendly":
			if v != nil {
				err = json.Unmarshal(*v, &c.Friendly)
			}
		default:
			if c.AdditionalProperties == nil {
				c.AdditionalProperties = &map[string]interface{}{}
			}
			if v != nil {
				var aux interface{}
				err = json.Unmarshal(*v, &aux)
				(*c.AdditionalProperties)[k] = aux
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// CatApTrueResponse is the response envelope for operations that return a CatApTrue type.
type CatApTrueResponse struct {
	CatApTrue *CatApTrue

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type Error struct {
	Message *string `json:"message,omitempty"`
	Status  *int32  `json:"status,omitempty"`
}

func (e Error) Error() string {
	msg := ""
	if e.Message != nil {
		msg += fmt.Sprintf("Message: %v\n", *e.Message)
	}
	if e.Status != nil {
		msg += fmt.Sprintf("Status: %v\n", *e.Status)
	}
	if msg == "" {
		msg = "missing error info"
	}
	return msg
}

type PetApInProperties struct {
	// Dictionary of <number>
	AdditionalProperties *map[string]float32 `json:"additionalProperties,omitempty"`
	ID                   *int32              `json:"id,omitempty"`
	Name                 *string             `json:"name,omitempty"`
	Status               *bool               `json:"status,omitempty"`
}

// PetApInPropertiesResponse is the response envelope for operations that return a PetApInProperties type.
type PetApInPropertiesResponse struct {
	PetApInProperties *PetApInProperties

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type PetApInPropertiesWithApstring struct {
	// Contains additional key/value pairs not defined in the schema.
	AdditionalProperties *map[string]string

	// Dictionary of <number>
	AdditionalProperties1 *map[string]float32 `json:"additionalProperties,omitempty"`
	ID                    *int32              `json:"id,omitempty"`
	Name                  *string             `json:"name,omitempty"`
	OdataLocation         *string             `json:"@odata.location,omitempty"`
	Status                *bool               `json:"status,omitempty"`
}

func (p PetApInPropertiesWithApstring) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if p.AdditionalProperties1 != nil {
		objectMap["additionalProperties"] = p.AdditionalProperties1
	}
	if p.ID != nil {
		objectMap["id"] = p.ID
	}
	if p.Name != nil {
		objectMap["name"] = p.Name
	}
	if p.OdataLocation != nil {
		objectMap["@odata.location"] = p.OdataLocation
	}
	if p.Status != nil {
		objectMap["status"] = p.Status
	}
	if p.AdditionalProperties != nil {
		for k, v := range *p.AdditionalProperties {
			objectMap[k] = v
		}
	}
	return json.Marshal(objectMap)
}

func (p *PetApInPropertiesWithApstring) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for k, v := range rawMsg {
		var err error
		switch k {
		case "additionalProperties":
			if v != nil {
				err = json.Unmarshal(*v, &p.AdditionalProperties1)
			}
		case "id":
			if v != nil {
				err = json.Unmarshal(*v, &p.ID)
			}
		case "name":
			if v != nil {
				err = json.Unmarshal(*v, &p.Name)
			}
		case "@odata.location":
			if v != nil {
				err = json.Unmarshal(*v, &p.OdataLocation)
			}
		case "status":
			if v != nil {
				err = json.Unmarshal(*v, &p.Status)
			}
		default:
			if p.AdditionalProperties == nil {
				p.AdditionalProperties = &map[string]string{}
			}
			if v != nil {
				var aux string
				err = json.Unmarshal(*v, &aux)
				(*p.AdditionalProperties)[k] = aux
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// PetApInPropertiesWithApstringResponse is the response envelope for operations that return a PetApInPropertiesWithApstring
// type.
type PetApInPropertiesWithApstringResponse struct {
	PetApInPropertiesWithApstring *PetApInPropertiesWithApstring

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type PetApObject struct {
	// Contains additional key/value pairs not defined in the schema.
	AdditionalProperties *map[string]interface{}
	ID                   *int32  `json:"id,omitempty"`
	Name                 *string `json:"name,omitempty"`
	Status               *bool   `json:"status,omitempty"`
}

func (p PetApObject) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if p.ID != nil {
		objectMap["id"] = p.ID
	}
	if p.Name != nil {
		objectMap["name"] = p.Name
	}
	if p.Status != nil {
		objectMap["status"] = p.Status
	}
	if p.AdditionalProperties != nil {
		for k, v := range *p.AdditionalProperties {
			objectMap[k] = v
		}
	}
	return json.Marshal(objectMap)
}

func (p *PetApObject) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for k, v := range rawMsg {
		var err error
		switch k {
		case "id":
			if v != nil {
				err = json.Unmarshal(*v, &p.ID)
			}
		case "name":
			if v != nil {
				err = json.Unmarshal(*v, &p.Name)
			}
		case "status":
			if v != nil {
				err = json.Unmarshal(*v, &p.Status)
			}
		default:
			if p.AdditionalProperties == nil {
				p.AdditionalProperties = &map[string]interface{}{}
			}
			if v != nil {
				var aux interface{}
				err = json.Unmarshal(*v, &aux)
				(*p.AdditionalProperties)[k] = aux
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// PetApObjectResponse is the response envelope for operations that return a PetApObject type.
type PetApObjectResponse struct {
	PetApObject *PetApObject

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type PetApString struct {
	// Contains additional key/value pairs not defined in the schema.
	AdditionalProperties *map[string]string
	ID                   *int32  `json:"id,omitempty"`
	Name                 *string `json:"name,omitempty"`
	Status               *bool   `json:"status,omitempty"`
}

func (p PetApString) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if p.ID != nil {
		objectMap["id"] = p.ID
	}
	if p.Name != nil {
		objectMap["name"] = p.Name
	}
	if p.Status != nil {
		objectMap["status"] = p.Status
	}
	if p.AdditionalProperties != nil {
		for k, v := range *p.AdditionalProperties {
			objectMap[k] = v
		}
	}
	return json.Marshal(objectMap)
}

func (p *PetApString) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for k, v := range rawMsg {
		var err error
		switch k {
		case "id":
			if v != nil {
				err = json.Unmarshal(*v, &p.ID)
			}
		case "name":
			if v != nil {
				err = json.Unmarshal(*v, &p.Name)
			}
		case "status":
			if v != nil {
				err = json.Unmarshal(*v, &p.Status)
			}
		default:
			if p.AdditionalProperties == nil {
				p.AdditionalProperties = &map[string]string{}
			}
			if v != nil {
				var aux string
				err = json.Unmarshal(*v, &aux)
				(*p.AdditionalProperties)[k] = aux
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// PetApStringResponse is the response envelope for operations that return a PetApString type.
type PetApStringResponse struct {
	PetApString *PetApString

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type PetApTrue struct {
	// Contains additional key/value pairs not defined in the schema.
	AdditionalProperties *map[string]interface{}
	ID                   *int32  `json:"id,omitempty"`
	Name                 *string `json:"name,omitempty"`
	Status               *bool   `json:"status,omitempty"`
}

func (p PetApTrue) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if p.ID != nil {
		objectMap["id"] = p.ID
	}
	if p.Name != nil {
		objectMap["name"] = p.Name
	}
	if p.Status != nil {
		objectMap["status"] = p.Status
	}
	if p.AdditionalProperties != nil {
		for k, v := range *p.AdditionalProperties {
			objectMap[k] = v
		}
	}
	return json.Marshal(objectMap)
}

func (p *PetApTrue) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for k, v := range rawMsg {
		var err error
		switch k {
		case "id":
			if v != nil {
				err = json.Unmarshal(*v, &p.ID)
			}
		case "name":
			if v != nil {
				err = json.Unmarshal(*v, &p.Name)
			}
		case "status":
			if v != nil {
				err = json.Unmarshal(*v, &p.Status)
			}
		default:
			if p.AdditionalProperties == nil {
				p.AdditionalProperties = &map[string]interface{}{}
			}
			if v != nil {
				var aux interface{}
				err = json.Unmarshal(*v, &aux)
				(*p.AdditionalProperties)[k] = aux
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// PetApTrueResponse is the response envelope for operations that return a PetApTrue type.
type PetApTrueResponse struct {
	PetApTrue *PetApTrue

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}