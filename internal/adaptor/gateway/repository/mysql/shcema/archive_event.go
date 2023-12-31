// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package schema

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// ArchiveEvent is an object representing the database table.
type ArchiveEvent struct {
	ID          string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	DeviceID    string    `boil:"device_id" json:"device_id" toml:"device_id" yaml:"device_id"`
	ClientID    string    `boil:"client_id" json:"client_id" toml:"client_id" yaml:"client_id"`
	RequestedAt time.Time `boil:"requested_at" json:"requested_at" toml:"requested_at" yaml:"requested_at"`

	R *archiveEventR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L archiveEventL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ArchiveEventColumns = struct {
	ID          string
	DeviceID    string
	ClientID    string
	RequestedAt string
}{
	ID:          "id",
	DeviceID:    "device_id",
	ClientID:    "client_id",
	RequestedAt: "requested_at",
}

var ArchiveEventTableColumns = struct {
	ID          string
	DeviceID    string
	ClientID    string
	RequestedAt string
}{
	ID:          "archive_event.id",
	DeviceID:    "archive_event.device_id",
	ClientID:    "archive_event.client_id",
	RequestedAt: "archive_event.requested_at",
}

// Generated where

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var ArchiveEventWhere = struct {
	ID          whereHelperstring
	DeviceID    whereHelperstring
	ClientID    whereHelperstring
	RequestedAt whereHelpertime_Time
}{
	ID:          whereHelperstring{field: "`archive_event`.`id`"},
	DeviceID:    whereHelperstring{field: "`archive_event`.`device_id`"},
	ClientID:    whereHelperstring{field: "`archive_event`.`client_id`"},
	RequestedAt: whereHelpertime_Time{field: "`archive_event`.`requested_at`"},
}

// ArchiveEventRels is where relationship names are stored.
var ArchiveEventRels = struct {
}{}

// archiveEventR is where relationships are stored.
type archiveEventR struct {
}

// NewStruct creates a new relationship struct
func (*archiveEventR) NewStruct() *archiveEventR {
	return &archiveEventR{}
}

// archiveEventL is where Load methods for each relationship are stored.
type archiveEventL struct{}

var (
	archiveEventAllColumns            = []string{"id", "device_id", "client_id", "requested_at"}
	archiveEventColumnsWithoutDefault = []string{"id", "device_id", "client_id", "requested_at"}
	archiveEventColumnsWithDefault    = []string{}
	archiveEventPrimaryKeyColumns     = []string{"id"}
	archiveEventGeneratedColumns      = []string{}
)

type (
	// ArchiveEventSlice is an alias for a slice of pointers to ArchiveEvent.
	// This should almost always be used instead of []ArchiveEvent.
	ArchiveEventSlice []*ArchiveEvent
	// ArchiveEventHook is the signature for custom ArchiveEvent hook methods
	ArchiveEventHook func(context.Context, boil.ContextExecutor, *ArchiveEvent) error

	archiveEventQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	archiveEventType                 = reflect.TypeOf(&ArchiveEvent{})
	archiveEventMapping              = queries.MakeStructMapping(archiveEventType)
	archiveEventPrimaryKeyMapping, _ = queries.BindMapping(archiveEventType, archiveEventMapping, archiveEventPrimaryKeyColumns)
	archiveEventInsertCacheMut       sync.RWMutex
	archiveEventInsertCache          = make(map[string]insertCache)
	archiveEventUpdateCacheMut       sync.RWMutex
	archiveEventUpdateCache          = make(map[string]updateCache)
	archiveEventUpsertCacheMut       sync.RWMutex
	archiveEventUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var archiveEventAfterSelectHooks []ArchiveEventHook

var archiveEventBeforeInsertHooks []ArchiveEventHook
var archiveEventAfterInsertHooks []ArchiveEventHook

var archiveEventBeforeUpdateHooks []ArchiveEventHook
var archiveEventAfterUpdateHooks []ArchiveEventHook

var archiveEventBeforeDeleteHooks []ArchiveEventHook
var archiveEventAfterDeleteHooks []ArchiveEventHook

var archiveEventBeforeUpsertHooks []ArchiveEventHook
var archiveEventAfterUpsertHooks []ArchiveEventHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ArchiveEvent) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range archiveEventAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ArchiveEvent) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range archiveEventBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ArchiveEvent) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range archiveEventAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ArchiveEvent) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range archiveEventBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ArchiveEvent) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range archiveEventAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ArchiveEvent) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range archiveEventBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ArchiveEvent) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range archiveEventAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ArchiveEvent) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range archiveEventBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ArchiveEvent) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range archiveEventAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddArchiveEventHook registers your hook function for all future operations.
func AddArchiveEventHook(hookPoint boil.HookPoint, archiveEventHook ArchiveEventHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		archiveEventAfterSelectHooks = append(archiveEventAfterSelectHooks, archiveEventHook)
	case boil.BeforeInsertHook:
		archiveEventBeforeInsertHooks = append(archiveEventBeforeInsertHooks, archiveEventHook)
	case boil.AfterInsertHook:
		archiveEventAfterInsertHooks = append(archiveEventAfterInsertHooks, archiveEventHook)
	case boil.BeforeUpdateHook:
		archiveEventBeforeUpdateHooks = append(archiveEventBeforeUpdateHooks, archiveEventHook)
	case boil.AfterUpdateHook:
		archiveEventAfterUpdateHooks = append(archiveEventAfterUpdateHooks, archiveEventHook)
	case boil.BeforeDeleteHook:
		archiveEventBeforeDeleteHooks = append(archiveEventBeforeDeleteHooks, archiveEventHook)
	case boil.AfterDeleteHook:
		archiveEventAfterDeleteHooks = append(archiveEventAfterDeleteHooks, archiveEventHook)
	case boil.BeforeUpsertHook:
		archiveEventBeforeUpsertHooks = append(archiveEventBeforeUpsertHooks, archiveEventHook)
	case boil.AfterUpsertHook:
		archiveEventAfterUpsertHooks = append(archiveEventAfterUpsertHooks, archiveEventHook)
	}
}

// One returns a single archiveEvent record from the query.
func (q archiveEventQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ArchiveEvent, error) {
	o := &ArchiveEvent{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "schema: failed to execute a one query for archive_event")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ArchiveEvent records from the query.
func (q archiveEventQuery) All(ctx context.Context, exec boil.ContextExecutor) (ArchiveEventSlice, error) {
	var o []*ArchiveEvent

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "schema: failed to assign all query results to ArchiveEvent slice")
	}

	if len(archiveEventAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ArchiveEvent records in the query.
func (q archiveEventQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to count archive_event rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q archiveEventQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "schema: failed to check if archive_event exists")
	}

	return count > 0, nil
}

// ArchiveEvents retrieves all the records using an executor.
func ArchiveEvents(mods ...qm.QueryMod) archiveEventQuery {
	mods = append(mods, qm.From("`archive_event`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`archive_event`.*"})
	}

	return archiveEventQuery{q}
}

// FindArchiveEvent retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindArchiveEvent(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*ArchiveEvent, error) {
	archiveEventObj := &ArchiveEvent{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `archive_event` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, archiveEventObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "schema: unable to select from archive_event")
	}

	if err = archiveEventObj.doAfterSelectHooks(ctx, exec); err != nil {
		return archiveEventObj, err
	}

	return archiveEventObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ArchiveEvent) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("schema: no archive_event provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(archiveEventColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	archiveEventInsertCacheMut.RLock()
	cache, cached := archiveEventInsertCache[key]
	archiveEventInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			archiveEventAllColumns,
			archiveEventColumnsWithDefault,
			archiveEventColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(archiveEventType, archiveEventMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(archiveEventType, archiveEventMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `archive_event` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `archive_event` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `archive_event` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, archiveEventPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "schema: unable to insert into archive_event")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "schema: unable to populate default values for archive_event")
	}

CacheNoHooks:
	if !cached {
		archiveEventInsertCacheMut.Lock()
		archiveEventInsertCache[key] = cache
		archiveEventInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ArchiveEvent.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ArchiveEvent) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	archiveEventUpdateCacheMut.RLock()
	cache, cached := archiveEventUpdateCache[key]
	archiveEventUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			archiveEventAllColumns,
			archiveEventPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("schema: unable to update archive_event, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `archive_event` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, archiveEventPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(archiveEventType, archiveEventMapping, append(wl, archiveEventPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to update archive_event row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by update for archive_event")
	}

	if !cached {
		archiveEventUpdateCacheMut.Lock()
		archiveEventUpdateCache[key] = cache
		archiveEventUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q archiveEventQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to update all for archive_event")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to retrieve rows affected for archive_event")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ArchiveEventSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("schema: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), archiveEventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `archive_event` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, archiveEventPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to update all in archiveEvent slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to retrieve rows affected all in update all archiveEvent")
	}
	return rowsAff, nil
}

var mySQLArchiveEventUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ArchiveEvent) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("schema: no archive_event provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(archiveEventColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLArchiveEventUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	archiveEventUpsertCacheMut.RLock()
	cache, cached := archiveEventUpsertCache[key]
	archiveEventUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			archiveEventAllColumns,
			archiveEventColumnsWithDefault,
			archiveEventColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			archiveEventAllColumns,
			archiveEventPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("schema: unable to upsert archive_event, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`archive_event`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `archive_event` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(archiveEventType, archiveEventMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(archiveEventType, archiveEventMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "schema: unable to upsert for archive_event")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(archiveEventType, archiveEventMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "schema: unable to retrieve unique values for archive_event")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "schema: unable to populate default values for archive_event")
	}

CacheNoHooks:
	if !cached {
		archiveEventUpsertCacheMut.Lock()
		archiveEventUpsertCache[key] = cache
		archiveEventUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ArchiveEvent record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ArchiveEvent) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("schema: no ArchiveEvent provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), archiveEventPrimaryKeyMapping)
	sql := "DELETE FROM `archive_event` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to delete from archive_event")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by delete for archive_event")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q archiveEventQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("schema: no archiveEventQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to delete all from archive_event")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by deleteall for archive_event")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ArchiveEventSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(archiveEventBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), archiveEventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `archive_event` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, archiveEventPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to delete all from archiveEvent slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by deleteall for archive_event")
	}

	if len(archiveEventAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ArchiveEvent) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindArchiveEvent(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ArchiveEventSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ArchiveEventSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), archiveEventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `archive_event`.* FROM `archive_event` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, archiveEventPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "schema: unable to reload all in ArchiveEventSlice")
	}

	*o = slice

	return nil
}

// ArchiveEventExists checks if the ArchiveEvent row exists.
func ArchiveEventExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `archive_event` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "schema: unable to check if archive_event exists")
	}

	return exists, nil
}

// Exists checks if the ArchiveEvent row exists.
func (o *ArchiveEvent) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ArchiveEventExists(ctx, exec, o.ID)
}
