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

// Archive is an object representing the database table.
type Archive struct {
	ID             string `boil:"id" json:"id" toml:"id" yaml:"id"`
	DeviceID       string `boil:"device_id" json:"device_id" toml:"device_id" yaml:"device_id"`
	ArchiveEventID string `boil:"archive_event_id" json:"archive_event_id" toml:"archive_event_id" yaml:"archive_event_id"`
	ContentType    string `boil:"content_type" json:"content_type" toml:"content_type" yaml:"content_type"`

	R *archiveR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L archiveL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ArchiveColumns = struct {
	ID             string
	DeviceID       string
	ArchiveEventID string
	ContentType    string
}{
	ID:             "id",
	DeviceID:       "device_id",
	ArchiveEventID: "archive_event_id",
	ContentType:    "content_type",
}

var ArchiveTableColumns = struct {
	ID             string
	DeviceID       string
	ArchiveEventID string
	ContentType    string
}{
	ID:             "archive.id",
	DeviceID:       "archive.device_id",
	ArchiveEventID: "archive.archive_event_id",
	ContentType:    "archive.content_type",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var ArchiveWhere = struct {
	ID             whereHelperstring
	DeviceID       whereHelperstring
	ArchiveEventID whereHelperstring
	ContentType    whereHelperstring
}{
	ID:             whereHelperstring{field: "`archive`.`id`"},
	DeviceID:       whereHelperstring{field: "`archive`.`device_id`"},
	ArchiveEventID: whereHelperstring{field: "`archive`.`archive_event_id`"},
	ContentType:    whereHelperstring{field: "`archive`.`content_type`"},
}

// ArchiveRels is where relationship names are stored.
var ArchiveRels = struct {
	Device string
}{
	Device: "Device",
}

// archiveR is where relationships are stored.
type archiveR struct {
	Device *Device `boil:"Device" json:"Device" toml:"Device" yaml:"Device"`
}

// NewStruct creates a new relationship struct
func (*archiveR) NewStruct() *archiveR {
	return &archiveR{}
}

func (r *archiveR) GetDevice() *Device {
	if r == nil {
		return nil
	}
	return r.Device
}

// archiveL is where Load methods for each relationship are stored.
type archiveL struct{}

var (
	archiveAllColumns            = []string{"id", "device_id", "archive_event_id", "content_type"}
	archiveColumnsWithoutDefault = []string{"id", "device_id", "archive_event_id", "content_type"}
	archiveColumnsWithDefault    = []string{}
	archivePrimaryKeyColumns     = []string{"id"}
	archiveGeneratedColumns      = []string{}
)

type (
	// ArchiveSlice is an alias for a slice of pointers to Archive.
	// This should almost always be used instead of []Archive.
	ArchiveSlice []*Archive
	// ArchiveHook is the signature for custom Archive hook methods
	ArchiveHook func(context.Context, boil.ContextExecutor, *Archive) error

	archiveQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	archiveType                 = reflect.TypeOf(&Archive{})
	archiveMapping              = queries.MakeStructMapping(archiveType)
	archivePrimaryKeyMapping, _ = queries.BindMapping(archiveType, archiveMapping, archivePrimaryKeyColumns)
	archiveInsertCacheMut       sync.RWMutex
	archiveInsertCache          = make(map[string]insertCache)
	archiveUpdateCacheMut       sync.RWMutex
	archiveUpdateCache          = make(map[string]updateCache)
	archiveUpsertCacheMut       sync.RWMutex
	archiveUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var archiveAfterSelectHooks []ArchiveHook

var archiveBeforeInsertHooks []ArchiveHook
var archiveAfterInsertHooks []ArchiveHook

var archiveBeforeUpdateHooks []ArchiveHook
var archiveAfterUpdateHooks []ArchiveHook

var archiveBeforeDeleteHooks []ArchiveHook
var archiveAfterDeleteHooks []ArchiveHook

var archiveBeforeUpsertHooks []ArchiveHook
var archiveAfterUpsertHooks []ArchiveHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Archive) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range archiveAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Archive) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range archiveBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Archive) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range archiveAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Archive) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range archiveBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Archive) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range archiveAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Archive) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range archiveBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Archive) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range archiveAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Archive) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range archiveBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Archive) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range archiveAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddArchiveHook registers your hook function for all future operations.
func AddArchiveHook(hookPoint boil.HookPoint, archiveHook ArchiveHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		archiveAfterSelectHooks = append(archiveAfterSelectHooks, archiveHook)
	case boil.BeforeInsertHook:
		archiveBeforeInsertHooks = append(archiveBeforeInsertHooks, archiveHook)
	case boil.AfterInsertHook:
		archiveAfterInsertHooks = append(archiveAfterInsertHooks, archiveHook)
	case boil.BeforeUpdateHook:
		archiveBeforeUpdateHooks = append(archiveBeforeUpdateHooks, archiveHook)
	case boil.AfterUpdateHook:
		archiveAfterUpdateHooks = append(archiveAfterUpdateHooks, archiveHook)
	case boil.BeforeDeleteHook:
		archiveBeforeDeleteHooks = append(archiveBeforeDeleteHooks, archiveHook)
	case boil.AfterDeleteHook:
		archiveAfterDeleteHooks = append(archiveAfterDeleteHooks, archiveHook)
	case boil.BeforeUpsertHook:
		archiveBeforeUpsertHooks = append(archiveBeforeUpsertHooks, archiveHook)
	case boil.AfterUpsertHook:
		archiveAfterUpsertHooks = append(archiveAfterUpsertHooks, archiveHook)
	}
}

// One returns a single archive record from the query.
func (q archiveQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Archive, error) {
	o := &Archive{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "schema: failed to execute a one query for archive")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Archive records from the query.
func (q archiveQuery) All(ctx context.Context, exec boil.ContextExecutor) (ArchiveSlice, error) {
	var o []*Archive

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "schema: failed to assign all query results to Archive slice")
	}

	if len(archiveAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Archive records in the query.
func (q archiveQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to count archive rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q archiveQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "schema: failed to check if archive exists")
	}

	return count > 0, nil
}

// Device pointed to by the foreign key.
func (o *Archive) Device(mods ...qm.QueryMod) deviceQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.DeviceID),
	}

	queryMods = append(queryMods, mods...)

	return Devices(queryMods...)
}

// LoadDevice allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (archiveL) LoadDevice(ctx context.Context, e boil.ContextExecutor, singular bool, maybeArchive interface{}, mods queries.Applicator) error {
	var slice []*Archive
	var object *Archive

	if singular {
		var ok bool
		object, ok = maybeArchive.(*Archive)
		if !ok {
			object = new(Archive)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeArchive)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeArchive))
			}
		}
	} else {
		s, ok := maybeArchive.(*[]*Archive)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeArchive)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeArchive))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &archiveR{}
		}
		args = append(args, object.DeviceID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &archiveR{}
			}

			for _, a := range args {
				if a == obj.DeviceID {
					continue Outer
				}
			}

			args = append(args, obj.DeviceID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`device`),
		qm.WhereIn(`device.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Device")
	}

	var resultSlice []*Device
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Device")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for device")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for device")
	}

	if len(deviceAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Device = foreign
		if foreign.R == nil {
			foreign.R = &deviceR{}
		}
		foreign.R.Archives = append(foreign.R.Archives, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.DeviceID == foreign.ID {
				local.R.Device = foreign
				if foreign.R == nil {
					foreign.R = &deviceR{}
				}
				foreign.R.Archives = append(foreign.R.Archives, local)
				break
			}
		}
	}

	return nil
}

// SetDevice of the archive to the related item.
// Sets o.R.Device to related.
// Adds o to related.R.Archives.
func (o *Archive) SetDevice(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Device) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `archive` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"device_id"}),
		strmangle.WhereClause("`", "`", 0, archivePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.DeviceID = related.ID
	if o.R == nil {
		o.R = &archiveR{
			Device: related,
		}
	} else {
		o.R.Device = related
	}

	if related.R == nil {
		related.R = &deviceR{
			Archives: ArchiveSlice{o},
		}
	} else {
		related.R.Archives = append(related.R.Archives, o)
	}

	return nil
}

// Archives retrieves all the records using an executor.
func Archives(mods ...qm.QueryMod) archiveQuery {
	mods = append(mods, qm.From("`archive`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`archive`.*"})
	}

	return archiveQuery{q}
}

// FindArchive retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindArchive(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Archive, error) {
	archiveObj := &Archive{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `archive` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, archiveObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "schema: unable to select from archive")
	}

	if err = archiveObj.doAfterSelectHooks(ctx, exec); err != nil {
		return archiveObj, err
	}

	return archiveObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Archive) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("schema: no archive provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(archiveColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	archiveInsertCacheMut.RLock()
	cache, cached := archiveInsertCache[key]
	archiveInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			archiveAllColumns,
			archiveColumnsWithDefault,
			archiveColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(archiveType, archiveMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(archiveType, archiveMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `archive` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `archive` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `archive` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, archivePrimaryKeyColumns))
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
		return errors.Wrap(err, "schema: unable to insert into archive")
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
		return errors.Wrap(err, "schema: unable to populate default values for archive")
	}

CacheNoHooks:
	if !cached {
		archiveInsertCacheMut.Lock()
		archiveInsertCache[key] = cache
		archiveInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Archive.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Archive) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	archiveUpdateCacheMut.RLock()
	cache, cached := archiveUpdateCache[key]
	archiveUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			archiveAllColumns,
			archivePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("schema: unable to update archive, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `archive` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, archivePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(archiveType, archiveMapping, append(wl, archivePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "schema: unable to update archive row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by update for archive")
	}

	if !cached {
		archiveUpdateCacheMut.Lock()
		archiveUpdateCache[key] = cache
		archiveUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q archiveQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to update all for archive")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to retrieve rows affected for archive")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ArchiveSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), archivePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `archive` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, archivePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to update all in archive slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to retrieve rows affected all in update all archive")
	}
	return rowsAff, nil
}

var mySQLArchiveUniqueColumns = []string{
	"id",
	"archive_event_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Archive) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("schema: no archive provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(archiveColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLArchiveUniqueColumns, o)

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

	archiveUpsertCacheMut.RLock()
	cache, cached := archiveUpsertCache[key]
	archiveUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			archiveAllColumns,
			archiveColumnsWithDefault,
			archiveColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			archiveAllColumns,
			archivePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("schema: unable to upsert archive, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`archive`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `archive` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(archiveType, archiveMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(archiveType, archiveMapping, ret)
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
		return errors.Wrap(err, "schema: unable to upsert for archive")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(archiveType, archiveMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "schema: unable to retrieve unique values for archive")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "schema: unable to populate default values for archive")
	}

CacheNoHooks:
	if !cached {
		archiveUpsertCacheMut.Lock()
		archiveUpsertCache[key] = cache
		archiveUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Archive record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Archive) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("schema: no Archive provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), archivePrimaryKeyMapping)
	sql := "DELETE FROM `archive` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to delete from archive")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by delete for archive")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q archiveQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("schema: no archiveQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to delete all from archive")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by deleteall for archive")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ArchiveSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(archiveBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), archivePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `archive` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, archivePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to delete all from archive slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by deleteall for archive")
	}

	if len(archiveAfterDeleteHooks) != 0 {
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
func (o *Archive) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindArchive(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ArchiveSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ArchiveSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), archivePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `archive`.* FROM `archive` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, archivePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "schema: unable to reload all in ArchiveSlice")
	}

	*o = slice

	return nil
}

// ArchiveExists checks if the Archive row exists.
func ArchiveExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `archive` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "schema: unable to check if archive exists")
	}

	return exists, nil
}

// Exists checks if the Archive row exists.
func (o *Archive) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ArchiveExists(ctx, exec, o.ID)
}