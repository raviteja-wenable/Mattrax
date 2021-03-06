// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/mattrax/Mattrax/pkg/null"
)

type DeviceState string

const (
	DeviceStateDeploying      DeviceState = "deploying"
	DeviceStateManaged        DeviceState = "managed"
	DeviceStateUserUnenrolled DeviceState = "user_unenrolled"
	DeviceStateMissing        DeviceState = "missing"
)

func (e *DeviceState) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = DeviceState(s)
	case string:
		*e = DeviceState(s)
	default:
		return fmt.Errorf("unsupported scan type for DeviceState: %T", src)
	}
	return nil
}

type ManagementProtocol string

const (
	ManagementProtocolWindows ManagementProtocol = "windows"
	ManagementProtocolAgent   ManagementProtocol = "agent"
	ManagementProtocolApple   ManagementProtocol = "apple"
)

func (e *ManagementProtocol) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ManagementProtocol(s)
	case string:
		*e = ManagementProtocol(s)
	default:
		return fmt.Errorf("unsupported scan type for ManagementProtocol: %T", src)
	}
	return nil
}

type Operation string

const (
	OperationINSERT Operation = "INSERT"
	OperationUPDATE Operation = "UPDATE"
	OperationDELETE Operation = "DELETE"
)

func (e *Operation) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Operation(s)
	case string:
		*e = Operation(s)
	default:
		return fmt.Errorf("unsupported scan type for Operation: %T", src)
	}
	return nil
}

type UserPermissionLevel string

const (
	UserPermissionLevelUser          UserPermissionLevel = "user"
	UserPermissionLevelAdministrator UserPermissionLevel = "administrator"
)

func (e *UserPermissionLevel) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserPermissionLevel(s)
	case string:
		*e = UserPermissionLevel(s)
	default:
		return fmt.Errorf("unsupported scan type for UserPermissionLevel: %T", src)
	}
	return nil
}

type Application struct {
	ID          string      `json:"id"`
	TenantID    string      `json:"tenant_id"`
	Name        null.String `json:"name"`
	Description null.String `json:"description"`
	Publisher   null.String `json:"publisher"`
}

type ApplicationTarget struct {
	AppID    string      `json:"app_id"`
	TenantID string      `json:"tenant_id"`
	MsiFile  null.String `json:"msi_file"`
	StoreID  null.String `json:"store_id"`
}

type Certificate struct {
	ID   string `json:"id"`
	Cert []byte `json:"cert"`
	Key  []byte `json:"key"`
}

type Device struct {
	ID          string             `json:"id"`
	TenantID    string             `json:"tenant_id"`
	Protocol    ManagementProtocol `json:"protocol"`
	Udid        string             `json:"udid"`
	Name        string             `json:"name"`
	Description null.String        `json:"description"`
	Model       null.String        `json:"model"`
	State       DeviceState        `json:"state"`
	Owner       null.String        `json:"owner"`
	AzureDid    null.String        `json:"azure_did"`
	EnrolledAt  time.Time          `json:"enrolled_at"`
}

type EventLog struct {
	EventID       int32           `json:"event_id"`
	TableName     string          `json:"table_name"`
	ResourceID    null.String     `json:"resource_id"`
	TenantID      null.String     `json:"tenant_id"`
	UserUpn       null.String     `json:"user_upn"`
	Time          time.Time       `json:"time"`
	Operation     Operation       `json:"operation"`
	ExistingValue json.RawMessage `json:"existing_value"`
}

type Group struct {
	ID          string      `json:"id"`
	TenantID    string      `json:"tenant_id"`
	Name        string      `json:"name"`
	Description null.String `json:"description"`
}

type GroupDevice struct {
	GroupID  string `json:"group_id"`
	DeviceID string `json:"device_id"`
}

type GroupPolicy struct {
	GroupID  string `json:"group_id"`
	PolicyID string `json:"policy_id"`
}

type Object struct {
	ID       string      `json:"id"`
	TenantID string      `json:"tenant_id"`
	Filename null.String `json:"filename"`
	Data     []byte      `json:"data"`
}

type Policy struct {
	ID          string          `json:"id"`
	TenantID    string          `json:"tenant_id"`
	Name        string          `json:"name"`
	Type        string          `json:"type"`
	Payload     json.RawMessage `json:"payload"`
	Description null.String     `json:"description"`
}

type Tenant struct {
	ID            string      `json:"id"`
	DisplayName   string      `json:"display_name"`
	PrimaryDomain string      `json:"primary_domain"`
	Email         null.String `json:"email"`
	Phone         null.String `json:"phone"`
	Description   null.String `json:"description"`
}

type TenantDomain struct {
	TenantID    string `json:"tenant_id"`
	Domain      string `json:"domain"`
	LinkingCode string `json:"linking_code"`
	Verified    bool   `json:"verified"`
}

type TenantUser struct {
	UserUpn         string              `json:"user_upn"`
	TenantID        string              `json:"tenant_id"`
	PermissionLevel UserPermissionLevel `json:"permission_level"`
}

type User struct {
	UPN        string      `json:"upn"`
	Fullname   string      `json:"fullname"`
	Disabled   bool        `json:"disabled"`
	Password   null.String `json:"password"`
	MfaToken   null.String `json:"mfa_token"`
	AzureadOid null.String `json:"azuread_oid"`
	TenantID   null.String `json:"tenant_id"`
}
