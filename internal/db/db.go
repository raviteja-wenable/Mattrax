// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createRawCertStmt, err = db.PrepareContext(ctx, createRawCert); err != nil {
		return nil, fmt.Errorf("error preparing query CreateRawCert: %w", err)
	}
	if q.deleteDeviceCacheNodeStmt, err = db.PrepareContext(ctx, deleteDeviceCacheNode); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteDeviceCacheNode: %w", err)
	}
	if q.deviceCheckinStatusStmt, err = db.PrepareContext(ctx, deviceCheckinStatus); err != nil {
		return nil, fmt.Errorf("error preparing query DeviceCheckinStatus: %w", err)
	}
	if q.deviceUserUnenrollmentStmt, err = db.PrepareContext(ctx, deviceUserUnenrollment); err != nil {
		return nil, fmt.Errorf("error preparing query DeviceUserUnenrollment: %w", err)
	}
	if q.getDeviceStmt, err = db.PrepareContext(ctx, getDevice); err != nil {
		return nil, fmt.Errorf("error preparing query GetDevice: %w", err)
	}
	if q.getDeviceByUDIDStmt, err = db.PrepareContext(ctx, getDeviceByUDID); err != nil {
		return nil, fmt.Errorf("error preparing query GetDeviceByUDID: %w", err)
	}
	if q.getDevicesDetachedPayloadsStmt, err = db.PrepareContext(ctx, getDevicesDetachedPayloads); err != nil {
		return nil, fmt.Errorf("error preparing query GetDevicesDetachedPayloads: %w", err)
	}
	if q.getDevicesPayloadsStmt, err = db.PrepareContext(ctx, getDevicesPayloads); err != nil {
		return nil, fmt.Errorf("error preparing query GetDevicesPayloads: %w", err)
	}
	if q.getDevicesPayloadsAwaitingDeploymentStmt, err = db.PrepareContext(ctx, getDevicesPayloadsAwaitingDeployment); err != nil {
		return nil, fmt.Errorf("error preparing query GetDevicesPayloadsAwaitingDeployment: %w", err)
	}
	if q.getPoliciesPayloadsStmt, err = db.PrepareContext(ctx, getPoliciesPayloads); err != nil {
		return nil, fmt.Errorf("error preparing query GetPoliciesPayloads: %w", err)
	}
	if q.getPolicyStmt, err = db.PrepareContext(ctx, getPolicy); err != nil {
		return nil, fmt.Errorf("error preparing query GetPolicy: %w", err)
	}
	if q.getRawCertStmt, err = db.PrepareContext(ctx, getRawCert); err != nil {
		return nil, fmt.Errorf("error preparing query GetRawCert: %w", err)
	}
	if q.getUserStmt, err = db.PrepareContext(ctx, getUser); err != nil {
		return nil, fmt.Errorf("error preparing query GetUser: %w", err)
	}
	if q.newAzureADUserStmt, err = db.PrepareContext(ctx, newAzureADUser); err != nil {
		return nil, fmt.Errorf("error preparing query NewAzureADUser: %w", err)
	}
	if q.newDeviceStmt, err = db.PrepareContext(ctx, newDevice); err != nil {
		return nil, fmt.Errorf("error preparing query NewDevice: %w", err)
	}
	if q.newDeviceCacheNodeStmt, err = db.PrepareContext(ctx, newDeviceCacheNode); err != nil {
		return nil, fmt.Errorf("error preparing query NewDeviceCacheNode: %w", err)
	}
	if q.newDeviceReplacingExistingStmt, err = db.PrepareContext(ctx, newDeviceReplacingExisting); err != nil {
		return nil, fmt.Errorf("error preparing query NewDeviceReplacingExisting: %w", err)
	}
	if q.newDeviceReplacingExistingResetCacheStmt, err = db.PrepareContext(ctx, newDeviceReplacingExistingResetCache); err != nil {
		return nil, fmt.Errorf("error preparing query NewDeviceReplacingExistingResetCache: %w", err)
	}
	if q.newDeviceReplacingExistingResetInventoryStmt, err = db.PrepareContext(ctx, newDeviceReplacingExistingResetInventory); err != nil {
		return nil, fmt.Errorf("error preparing query NewDeviceReplacingExistingResetInventory: %w", err)
	}
	if q.setDeviceStateStmt, err = db.PrepareContext(ctx, setDeviceState); err != nil {
		return nil, fmt.Errorf("error preparing query SetDeviceState: %w", err)
	}
	if q.settingsStmt, err = db.PrepareContext(ctx, settings); err != nil {
		return nil, fmt.Errorf("error preparing query Settings: %w", err)
	}
	if q.updateDeviceInventoryNodeStmt, err = db.PrepareContext(ctx, updateDeviceInventoryNode); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateDeviceInventoryNode: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createRawCertStmt != nil {
		if cerr := q.createRawCertStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createRawCertStmt: %w", cerr)
		}
	}
	if q.deleteDeviceCacheNodeStmt != nil {
		if cerr := q.deleteDeviceCacheNodeStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteDeviceCacheNodeStmt: %w", cerr)
		}
	}
	if q.deviceCheckinStatusStmt != nil {
		if cerr := q.deviceCheckinStatusStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deviceCheckinStatusStmt: %w", cerr)
		}
	}
	if q.deviceUserUnenrollmentStmt != nil {
		if cerr := q.deviceUserUnenrollmentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deviceUserUnenrollmentStmt: %w", cerr)
		}
	}
	if q.getDeviceStmt != nil {
		if cerr := q.getDeviceStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getDeviceStmt: %w", cerr)
		}
	}
	if q.getDeviceByUDIDStmt != nil {
		if cerr := q.getDeviceByUDIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getDeviceByUDIDStmt: %w", cerr)
		}
	}
	if q.getDevicesDetachedPayloadsStmt != nil {
		if cerr := q.getDevicesDetachedPayloadsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getDevicesDetachedPayloadsStmt: %w", cerr)
		}
	}
	if q.getDevicesPayloadsStmt != nil {
		if cerr := q.getDevicesPayloadsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getDevicesPayloadsStmt: %w", cerr)
		}
	}
	if q.getDevicesPayloadsAwaitingDeploymentStmt != nil {
		if cerr := q.getDevicesPayloadsAwaitingDeploymentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getDevicesPayloadsAwaitingDeploymentStmt: %w", cerr)
		}
	}
	if q.getPoliciesPayloadsStmt != nil {
		if cerr := q.getPoliciesPayloadsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPoliciesPayloadsStmt: %w", cerr)
		}
	}
	if q.getPolicyStmt != nil {
		if cerr := q.getPolicyStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPolicyStmt: %w", cerr)
		}
	}
	if q.getRawCertStmt != nil {
		if cerr := q.getRawCertStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getRawCertStmt: %w", cerr)
		}
	}
	if q.getUserStmt != nil {
		if cerr := q.getUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserStmt: %w", cerr)
		}
	}
	if q.newAzureADUserStmt != nil {
		if cerr := q.newAzureADUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing newAzureADUserStmt: %w", cerr)
		}
	}
	if q.newDeviceStmt != nil {
		if cerr := q.newDeviceStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing newDeviceStmt: %w", cerr)
		}
	}
	if q.newDeviceCacheNodeStmt != nil {
		if cerr := q.newDeviceCacheNodeStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing newDeviceCacheNodeStmt: %w", cerr)
		}
	}
	if q.newDeviceReplacingExistingStmt != nil {
		if cerr := q.newDeviceReplacingExistingStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing newDeviceReplacingExistingStmt: %w", cerr)
		}
	}
	if q.newDeviceReplacingExistingResetCacheStmt != nil {
		if cerr := q.newDeviceReplacingExistingResetCacheStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing newDeviceReplacingExistingResetCacheStmt: %w", cerr)
		}
	}
	if q.newDeviceReplacingExistingResetInventoryStmt != nil {
		if cerr := q.newDeviceReplacingExistingResetInventoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing newDeviceReplacingExistingResetInventoryStmt: %w", cerr)
		}
	}
	if q.setDeviceStateStmt != nil {
		if cerr := q.setDeviceStateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing setDeviceStateStmt: %w", cerr)
		}
	}
	if q.settingsStmt != nil {
		if cerr := q.settingsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing settingsStmt: %w", cerr)
		}
	}
	if q.updateDeviceInventoryNodeStmt != nil {
		if cerr := q.updateDeviceInventoryNodeStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateDeviceInventoryNodeStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                                           DBTX
	tx                                           *sql.Tx
	createRawCertStmt                            *sql.Stmt
	deleteDeviceCacheNodeStmt                    *sql.Stmt
	deviceCheckinStatusStmt                      *sql.Stmt
	deviceUserUnenrollmentStmt                   *sql.Stmt
	getDeviceStmt                                *sql.Stmt
	getDeviceByUDIDStmt                          *sql.Stmt
	getDevicesDetachedPayloadsStmt               *sql.Stmt
	getDevicesPayloadsStmt                       *sql.Stmt
	getDevicesPayloadsAwaitingDeploymentStmt     *sql.Stmt
	getPoliciesPayloadsStmt                      *sql.Stmt
	getPolicyStmt                                *sql.Stmt
	getRawCertStmt                               *sql.Stmt
	getUserStmt                                  *sql.Stmt
	newAzureADUserStmt                           *sql.Stmt
	newDeviceStmt                                *sql.Stmt
	newDeviceCacheNodeStmt                       *sql.Stmt
	newDeviceReplacingExistingStmt               *sql.Stmt
	newDeviceReplacingExistingResetCacheStmt     *sql.Stmt
	newDeviceReplacingExistingResetInventoryStmt *sql.Stmt
	setDeviceStateStmt                           *sql.Stmt
	settingsStmt                                 *sql.Stmt
	updateDeviceInventoryNodeStmt                *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                                       tx,
		tx:                                       tx,
		createRawCertStmt:                        q.createRawCertStmt,
		deleteDeviceCacheNodeStmt:                q.deleteDeviceCacheNodeStmt,
		deviceCheckinStatusStmt:                  q.deviceCheckinStatusStmt,
		deviceUserUnenrollmentStmt:               q.deviceUserUnenrollmentStmt,
		getDeviceStmt:                            q.getDeviceStmt,
		getDeviceByUDIDStmt:                      q.getDeviceByUDIDStmt,
		getDevicesDetachedPayloadsStmt:           q.getDevicesDetachedPayloadsStmt,
		getDevicesPayloadsStmt:                   q.getDevicesPayloadsStmt,
		getDevicesPayloadsAwaitingDeploymentStmt: q.getDevicesPayloadsAwaitingDeploymentStmt,
		getPoliciesPayloadsStmt:                  q.getPoliciesPayloadsStmt,
		getPolicyStmt:                            q.getPolicyStmt,
		getRawCertStmt:                           q.getRawCertStmt,
		getUserStmt:                              q.getUserStmt,
		newAzureADUserStmt:                       q.newAzureADUserStmt,
		newDeviceStmt:                            q.newDeviceStmt,
		newDeviceCacheNodeStmt:                   q.newDeviceCacheNodeStmt,
		newDeviceReplacingExistingStmt:           q.newDeviceReplacingExistingStmt,
		newDeviceReplacingExistingResetCacheStmt: q.newDeviceReplacingExistingResetCacheStmt,
		newDeviceReplacingExistingResetInventoryStmt: q.newDeviceReplacingExistingResetInventoryStmt,
		setDeviceStateStmt:                           q.setDeviceStateStmt,
		settingsStmt:                                 q.settingsStmt,
		updateDeviceInventoryNodeStmt:                q.updateDeviceInventoryNodeStmt,
	}
}
