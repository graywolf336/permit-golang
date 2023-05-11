/*
Permit.io API

 Authorization as a service

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"time"
)

// RoleRead struct for RoleRead
type RoleRead struct {
	// The name of the role
	Name string `json:"name"`
	// optional description string explaining what this role represents, or what permissions are granted to it.
	Description *string `json:"description,omitempty"`
	// list of action keys that define what actions this resource role is permitted to do
	Permissions []string `json:"permissions,omitempty"`
	// list of role keys that define what roles this role extends. In other words: this role will automatically inherit all the permissions of the given roles in this list.
	Extends []string `json:"extends,omitempty"`
	// optional dictionary of key-value pairs that can be used to store arbitrary metadata about this role. This metadata can be used to filter role using query parameters with attr_ prefix, currently supports only 'equals' operator
	Attributes map[string]string `json:"attributes,omitempty"`
	// A URL-friendly name of the role (i.e: slug). You will be able to query later using this key instead of the id (UUID) of the role.
	Key string `json:"key"`
	// Unique id of the role
	Id string `json:"id"`
	// Unique id of the organization that the role belongs to.
	OrganizationId string `json:"organization_id"`
	// Unique id of the project that the role belongs to.
	ProjectId string `json:"project_id"`
	// Unique id of the environment that the role belongs to.
	EnvironmentId string `json:"environment_id"`
	// Date and time when the role was created (ISO_8601 format).
	CreatedAt time.Time `json:"created_at"`
	// Date and time when the role was last updated/modified (ISO_8601 format).
	UpdatedAt time.Time `json:"updated_at"`
}

// NewRoleRead instantiates a new RoleRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleRead(name string, key string, id string, organizationId string, projectId string, environmentId string, createdAt time.Time, updatedAt time.Time) *RoleRead {
	this := RoleRead{}
	this.Name = name
	this.Key = key
	this.Id = id
	this.OrganizationId = organizationId
	this.ProjectId = projectId
	this.EnvironmentId = environmentId
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewRoleReadWithDefaults instantiates a new RoleRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleReadWithDefaults() *RoleRead {
	this := RoleRead{}
	return &this
}

// GetName returns the Name field value
func (o *RoleRead) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RoleRead) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RoleRead) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RoleRead) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}
func (o *RoleRead) GetAttributes() map[string]string {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]string
		return ret
	}
	return o.Attributes
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleRead) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RoleRead) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RoleRead) SetDescription(v string) {
	o.Description = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *RoleRead) GetPermissions() []string {
	if o == nil || IsNil(o.Permissions) {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleRead) GetPermissionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *RoleRead) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *RoleRead) SetPermissions(v []string) {
	o.Permissions = v
}

// GetExtends returns the Extends field value if set, zero value otherwise.
func (o *RoleRead) GetExtends() []string {
	if o == nil || IsNil(o.Extends) {
		var ret []string
		return ret
	}
	return o.Extends
}

// GetExtendsOk returns a tuple with the Extends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleRead) GetExtendsOk() ([]string, bool) {
	if o == nil || IsNil(o.Extends) {
		return nil, false
	}
	return o.Extends, true
}

// HasExtends returns a boolean if a field has been set.
func (o *RoleRead) HasExtends() bool {
	if o != nil && !IsNil(o.Extends) {
		return true
	}

	return false
}

// SetExtends gets a reference to the given []string and assigns it to the Extends field.
func (o *RoleRead) SetExtends(v []string) {
	o.Extends = v
}

// GetKey returns the Key field value
func (o *RoleRead) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *RoleRead) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *RoleRead) SetKey(v string) {
	o.Key = v
}

// GetId returns the Id field value
func (o *RoleRead) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RoleRead) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RoleRead) SetId(v string) {
	o.Id = v
}

// GetOrganizationId returns the OrganizationId field value
func (o *RoleRead) GetOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *RoleRead) GetOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *RoleRead) SetOrganizationId(v string) {
	o.OrganizationId = v
}

// GetProjectId returns the ProjectId field value
func (o *RoleRead) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *RoleRead) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *RoleRead) SetProjectId(v string) {
	o.ProjectId = v
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *RoleRead) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *RoleRead) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *RoleRead) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *RoleRead) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *RoleRead) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *RoleRead) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *RoleRead) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *RoleRead) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *RoleRead) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o RoleRead) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if !IsNil(o.Extends) {
		toSerialize["extends"] = o.Extends
	}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if true {
		toSerialize["project_id"] = o.ProjectId
	}
	if true {
		toSerialize["environment_id"] = o.EnvironmentId
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableRoleRead struct {
	value *RoleRead
	isSet bool
}

func (v NullableRoleRead) Get() *RoleRead {
	return v.value
}

func (v *NullableRoleRead) Set(val *RoleRead) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleRead) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleRead(val *RoleRead) *NullableRoleRead {
	return &NullableRoleRead{value: val, isSet: true}
}

func (v NullableRoleRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
