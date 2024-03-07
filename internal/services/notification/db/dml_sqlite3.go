// Code generated by go generate; DO NOT EDIT.
package db

import (
	stdsql "database/sql"
	"time"

	"github.com/sorintlab/errors"
	sq "github.com/huandu/go-sqlbuilder"

	"agola.io/agola/internal/sqlg/sql"

	types "agola.io/agola/services/notification/types"
)
var (
	runWebhookInsertSqlite3 = func(inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inPayload []byte, inProjectID string) *sq.InsertBuilder {
		ib:= sq.NewInsertBuilder()
		return ib.InsertInto("runwebhook").Cols("id", "revision", "creation_time", "update_time", "payload", "project_id").Values(inID, inRevision, inCreationTime, inUpdateTime, inPayload, inProjectID)
	}
	runWebhookUpdateSqlite3 = func(curRevision uint64, inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inPayload []byte, inProjectID string) *sq.UpdateBuilder {
		ub:= sq.NewUpdateBuilder()
		return ub.Update("runwebhook").Set(ub.Assign("id", inID), ub.Assign("revision", inRevision), ub.Assign("creation_time", inCreationTime), ub.Assign("update_time", inUpdateTime), ub.Assign("payload", inPayload), ub.Assign("project_id", inProjectID)).Where(ub.E("id", inID), ub.E("revision", curRevision))
	}

	runWebhookInsertRawSqlite3 = func(inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inPayload []byte, inProjectID string) *sq.InsertBuilder {
		ib:= sq.NewInsertBuilder()
		return ib.InsertInto("runwebhook").Cols("id", "revision", "creation_time", "update_time", "payload", "project_id").SQL("").Values(inID, inRevision, inCreationTime, inUpdateTime, inPayload, inProjectID)
	}
)

func (d *DB) insertRunWebhookSqlite3(tx *sql.Tx, runwebhook *types.RunWebhook) error {
	q := runWebhookInsertSqlite3(runwebhook.ID, runwebhook.Revision, runwebhook.CreationTime, runwebhook.UpdateTime, runwebhook.Payload, runwebhook.ProjectID)

	if _, err := d.exec(tx, q); err != nil {
		return errors.Wrap(err, "failed to insert runWebhook")
	}

	return nil
}

func (d *DB) updateRunWebhookSqlite3(tx *sql.Tx, curRevision uint64, runwebhook *types.RunWebhook) (stdsql.Result, error) {
	q := runWebhookUpdateSqlite3(curRevision, runwebhook.ID, runwebhook.Revision, runwebhook.CreationTime, runwebhook.UpdateTime, runwebhook.Payload, runwebhook.ProjectID)

	res, err := d.exec(tx, q)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update runWebhook")
	}

	return res, nil
}

func (d *DB) insertRawRunWebhookSqlite3(tx *sql.Tx, runwebhook *types.RunWebhook) error {
	q := runWebhookInsertRawSqlite3(runwebhook.ID, runwebhook.Revision, runwebhook.CreationTime, runwebhook.UpdateTime, runwebhook.Payload, runwebhook.ProjectID)

	if _, err := d.exec(tx, q); err != nil {
		return errors.Wrap(err, "failed to insert runWebhook")
	}

	return nil
}
var (
	runWebhookDeliveryInsertSqlite3 = func(inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inSequence uint64, inRunWebhookID string, inDeliveryStatus types.DeliveryStatus, inDeliveredAt *time.Time, inStatusCode int) *sq.InsertBuilder {
		ib:= sq.NewInsertBuilder()
		return ib.InsertInto("runwebhookdelivery").Cols("id", "revision", "creation_time", "update_time", "sequence", "run_webhook_id", "delivery_status", "delivered_at", "status_code").Values(inID, inRevision, inCreationTime, inUpdateTime, inSequence, inRunWebhookID, inDeliveryStatus, inDeliveredAt, inStatusCode)
	}
	runWebhookDeliveryUpdateSqlite3 = func(curRevision uint64, inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inRunWebhookID string, inDeliveryStatus types.DeliveryStatus, inDeliveredAt *time.Time, inStatusCode int) *sq.UpdateBuilder {
		ub:= sq.NewUpdateBuilder()
		return ub.Update("runwebhookdelivery").Set(ub.Assign("id", inID), ub.Assign("revision", inRevision), ub.Assign("creation_time", inCreationTime), ub.Assign("update_time", inUpdateTime), ub.Assign("run_webhook_id", inRunWebhookID), ub.Assign("delivery_status", inDeliveryStatus), ub.Assign("delivered_at", inDeliveredAt), ub.Assign("status_code", inStatusCode)).Where(ub.E("id", inID), ub.E("revision", curRevision))
	}

	runWebhookDeliveryInsertRawSqlite3 = func(inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inSequence uint64, inRunWebhookID string, inDeliveryStatus types.DeliveryStatus, inDeliveredAt *time.Time, inStatusCode int) *sq.InsertBuilder {
		ib:= sq.NewInsertBuilder()
		return ib.InsertInto("runwebhookdelivery").Cols("id", "revision", "creation_time", "update_time", "sequence", "run_webhook_id", "delivery_status", "delivered_at", "status_code").SQL("").Values(inID, inRevision, inCreationTime, inUpdateTime, inSequence, inRunWebhookID, inDeliveryStatus, inDeliveredAt, inStatusCode)
	}
)

func (d *DB) insertRunWebhookDeliverySqlite3(tx *sql.Tx, runwebhookdelivery *types.RunWebhookDelivery) error {
	q := runWebhookDeliveryInsertSqlite3(runwebhookdelivery.ID, runwebhookdelivery.Revision, runwebhookdelivery.CreationTime, runwebhookdelivery.UpdateTime, runwebhookdelivery.Sequence, runwebhookdelivery.RunWebhookID, runwebhookdelivery.DeliveryStatus, runwebhookdelivery.DeliveredAt, runwebhookdelivery.StatusCode)

	if _, err := d.exec(tx, q); err != nil {
		return errors.Wrap(err, "failed to insert runWebhookDelivery")
	}

	return nil
}

func (d *DB) updateRunWebhookDeliverySqlite3(tx *sql.Tx, curRevision uint64, runwebhookdelivery *types.RunWebhookDelivery) (stdsql.Result, error) {
	q := runWebhookDeliveryUpdateSqlite3(curRevision, runwebhookdelivery.ID, runwebhookdelivery.Revision, runwebhookdelivery.CreationTime, runwebhookdelivery.UpdateTime, runwebhookdelivery.RunWebhookID, runwebhookdelivery.DeliveryStatus, runwebhookdelivery.DeliveredAt, runwebhookdelivery.StatusCode)

	res, err := d.exec(tx, q)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update runWebhookDelivery")
	}

	return res, nil
}

func (d *DB) insertRawRunWebhookDeliverySqlite3(tx *sql.Tx, runwebhookdelivery *types.RunWebhookDelivery) error {
	q := runWebhookDeliveryInsertRawSqlite3(runwebhookdelivery.ID, runwebhookdelivery.Revision, runwebhookdelivery.CreationTime, runwebhookdelivery.UpdateTime, runwebhookdelivery.Sequence, runwebhookdelivery.RunWebhookID, runwebhookdelivery.DeliveryStatus, runwebhookdelivery.DeliveredAt, runwebhookdelivery.StatusCode)

	if _, err := d.exec(tx, q); err != nil {
		return errors.Wrap(err, "failed to insert runWebhookDelivery")
	}

	return nil
}
var (
	lastRunEventSequenceInsertSqlite3 = func(inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inValue uint64) *sq.InsertBuilder {
		ib:= sq.NewInsertBuilder()
		return ib.InsertInto("lastruneventsequence").Cols("id", "revision", "creation_time", "update_time", "value").Values(inID, inRevision, inCreationTime, inUpdateTime, inValue)
	}
	lastRunEventSequenceUpdateSqlite3 = func(curRevision uint64, inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inValue uint64) *sq.UpdateBuilder {
		ub:= sq.NewUpdateBuilder()
		return ub.Update("lastruneventsequence").Set(ub.Assign("id", inID), ub.Assign("revision", inRevision), ub.Assign("creation_time", inCreationTime), ub.Assign("update_time", inUpdateTime), ub.Assign("value", inValue)).Where(ub.E("id", inID), ub.E("revision", curRevision))
	}

	lastRunEventSequenceInsertRawSqlite3 = func(inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inValue uint64) *sq.InsertBuilder {
		ib:= sq.NewInsertBuilder()
		return ib.InsertInto("lastruneventsequence").Cols("id", "revision", "creation_time", "update_time", "value").SQL("").Values(inID, inRevision, inCreationTime, inUpdateTime, inValue)
	}
)

func (d *DB) insertLastRunEventSequenceSqlite3(tx *sql.Tx, lastruneventsequence *types.LastRunEventSequence) error {
	q := lastRunEventSequenceInsertSqlite3(lastruneventsequence.ID, lastruneventsequence.Revision, lastruneventsequence.CreationTime, lastruneventsequence.UpdateTime, lastruneventsequence.Value)

	if _, err := d.exec(tx, q); err != nil {
		return errors.Wrap(err, "failed to insert lastRunEventSequence")
	}

	return nil
}

func (d *DB) updateLastRunEventSequenceSqlite3(tx *sql.Tx, curRevision uint64, lastruneventsequence *types.LastRunEventSequence) (stdsql.Result, error) {
	q := lastRunEventSequenceUpdateSqlite3(curRevision, lastruneventsequence.ID, lastruneventsequence.Revision, lastruneventsequence.CreationTime, lastruneventsequence.UpdateTime, lastruneventsequence.Value)

	res, err := d.exec(tx, q)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update lastRunEventSequence")
	}

	return res, nil
}

func (d *DB) insertRawLastRunEventSequenceSqlite3(tx *sql.Tx, lastruneventsequence *types.LastRunEventSequence) error {
	q := lastRunEventSequenceInsertRawSqlite3(lastruneventsequence.ID, lastruneventsequence.Revision, lastruneventsequence.CreationTime, lastruneventsequence.UpdateTime, lastruneventsequence.Value)

	if _, err := d.exec(tx, q); err != nil {
		return errors.Wrap(err, "failed to insert lastRunEventSequence")
	}

	return nil
}
var (
	commitStatusInsertSqlite3 = func(inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inProjectID string, inState types.CommitState, inCommitSHA string, inRunCounter uint64, inDescription string, inContext string) *sq.InsertBuilder {
		ib:= sq.NewInsertBuilder()
		return ib.InsertInto("commitstatus").Cols("id", "revision", "creation_time", "update_time", "project_id", "state", "commit_sha", "run_counter", "description", "context").Values(inID, inRevision, inCreationTime, inUpdateTime, inProjectID, inState, inCommitSHA, inRunCounter, inDescription, inContext)
	}
	commitStatusUpdateSqlite3 = func(curRevision uint64, inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inProjectID string, inState types.CommitState, inCommitSHA string, inRunCounter uint64, inDescription string, inContext string) *sq.UpdateBuilder {
		ub:= sq.NewUpdateBuilder()
		return ub.Update("commitstatus").Set(ub.Assign("id", inID), ub.Assign("revision", inRevision), ub.Assign("creation_time", inCreationTime), ub.Assign("update_time", inUpdateTime), ub.Assign("project_id", inProjectID), ub.Assign("state", inState), ub.Assign("commit_sha", inCommitSHA), ub.Assign("run_counter", inRunCounter), ub.Assign("description", inDescription), ub.Assign("context", inContext)).Where(ub.E("id", inID), ub.E("revision", curRevision))
	}

	commitStatusInsertRawSqlite3 = func(inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inProjectID string, inState types.CommitState, inCommitSHA string, inRunCounter uint64, inDescription string, inContext string) *sq.InsertBuilder {
		ib:= sq.NewInsertBuilder()
		return ib.InsertInto("commitstatus").Cols("id", "revision", "creation_time", "update_time", "project_id", "state", "commit_sha", "run_counter", "description", "context").SQL("").Values(inID, inRevision, inCreationTime, inUpdateTime, inProjectID, inState, inCommitSHA, inRunCounter, inDescription, inContext)
	}
)

func (d *DB) insertCommitStatusSqlite3(tx *sql.Tx, commitstatus *types.CommitStatus) error {
	q := commitStatusInsertSqlite3(commitstatus.ID, commitstatus.Revision, commitstatus.CreationTime, commitstatus.UpdateTime, commitstatus.ProjectID, commitstatus.State, commitstatus.CommitSHA, commitstatus.RunCounter, commitstatus.Description, commitstatus.Context)

	if _, err := d.exec(tx, q); err != nil {
		return errors.Wrap(err, "failed to insert commitStatus")
	}

	return nil
}

func (d *DB) updateCommitStatusSqlite3(tx *sql.Tx, curRevision uint64, commitstatus *types.CommitStatus) (stdsql.Result, error) {
	q := commitStatusUpdateSqlite3(curRevision, commitstatus.ID, commitstatus.Revision, commitstatus.CreationTime, commitstatus.UpdateTime, commitstatus.ProjectID, commitstatus.State, commitstatus.CommitSHA, commitstatus.RunCounter, commitstatus.Description, commitstatus.Context)

	res, err := d.exec(tx, q)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update commitStatus")
	}

	return res, nil
}

func (d *DB) insertRawCommitStatusSqlite3(tx *sql.Tx, commitstatus *types.CommitStatus) error {
	q := commitStatusInsertRawSqlite3(commitstatus.ID, commitstatus.Revision, commitstatus.CreationTime, commitstatus.UpdateTime, commitstatus.ProjectID, commitstatus.State, commitstatus.CommitSHA, commitstatus.RunCounter, commitstatus.Description, commitstatus.Context)

	if _, err := d.exec(tx, q); err != nil {
		return errors.Wrap(err, "failed to insert commitStatus")
	}

	return nil
}
var (
	commitStatusDeliveryInsertSqlite3 = func(inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inSequence uint64, inCommitStatusID string, inDeliveryStatus types.DeliveryStatus, inDeliveredAt *time.Time) *sq.InsertBuilder {
		ib:= sq.NewInsertBuilder()
		return ib.InsertInto("commitstatusdelivery").Cols("id", "revision", "creation_time", "update_time", "sequence", "commit_status_id", "delivery_status", "delivered_at").Values(inID, inRevision, inCreationTime, inUpdateTime, inSequence, inCommitStatusID, inDeliveryStatus, inDeliveredAt)
	}
	commitStatusDeliveryUpdateSqlite3 = func(curRevision uint64, inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inCommitStatusID string, inDeliveryStatus types.DeliveryStatus, inDeliveredAt *time.Time) *sq.UpdateBuilder {
		ub:= sq.NewUpdateBuilder()
		return ub.Update("commitstatusdelivery").Set(ub.Assign("id", inID), ub.Assign("revision", inRevision), ub.Assign("creation_time", inCreationTime), ub.Assign("update_time", inUpdateTime), ub.Assign("commit_status_id", inCommitStatusID), ub.Assign("delivery_status", inDeliveryStatus), ub.Assign("delivered_at", inDeliveredAt)).Where(ub.E("id", inID), ub.E("revision", curRevision))
	}

	commitStatusDeliveryInsertRawSqlite3 = func(inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inSequence uint64, inCommitStatusID string, inDeliveryStatus types.DeliveryStatus, inDeliveredAt *time.Time) *sq.InsertBuilder {
		ib:= sq.NewInsertBuilder()
		return ib.InsertInto("commitstatusdelivery").Cols("id", "revision", "creation_time", "update_time", "sequence", "commit_status_id", "delivery_status", "delivered_at").SQL("").Values(inID, inRevision, inCreationTime, inUpdateTime, inSequence, inCommitStatusID, inDeliveryStatus, inDeliveredAt)
	}
)

func (d *DB) insertCommitStatusDeliverySqlite3(tx *sql.Tx, commitstatusdelivery *types.CommitStatusDelivery) error {
	q := commitStatusDeliveryInsertSqlite3(commitstatusdelivery.ID, commitstatusdelivery.Revision, commitstatusdelivery.CreationTime, commitstatusdelivery.UpdateTime, commitstatusdelivery.Sequence, commitstatusdelivery.CommitStatusID, commitstatusdelivery.DeliveryStatus, commitstatusdelivery.DeliveredAt)

	if _, err := d.exec(tx, q); err != nil {
		return errors.Wrap(err, "failed to insert commitStatusDelivery")
	}

	return nil
}

func (d *DB) updateCommitStatusDeliverySqlite3(tx *sql.Tx, curRevision uint64, commitstatusdelivery *types.CommitStatusDelivery) (stdsql.Result, error) {
	q := commitStatusDeliveryUpdateSqlite3(curRevision, commitstatusdelivery.ID, commitstatusdelivery.Revision, commitstatusdelivery.CreationTime, commitstatusdelivery.UpdateTime, commitstatusdelivery.CommitStatusID, commitstatusdelivery.DeliveryStatus, commitstatusdelivery.DeliveredAt)

	res, err := d.exec(tx, q)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update commitStatusDelivery")
	}

	return res, nil
}

func (d *DB) insertRawCommitStatusDeliverySqlite3(tx *sql.Tx, commitstatusdelivery *types.CommitStatusDelivery) error {
	q := commitStatusDeliveryInsertRawSqlite3(commitstatusdelivery.ID, commitstatusdelivery.Revision, commitstatusdelivery.CreationTime, commitstatusdelivery.UpdateTime, commitstatusdelivery.Sequence, commitstatusdelivery.CommitStatusID, commitstatusdelivery.DeliveryStatus, commitstatusdelivery.DeliveredAt)

	if _, err := d.exec(tx, q); err != nil {
		return errors.Wrap(err, "failed to insert commitStatusDelivery")
	}

	return nil
}
