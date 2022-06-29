//go:build linux && cgo && !agent

package cluster

// The code below was generated by lxd-generate - DO NOT EDIT!

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/lxc/lxd/lxd/db/query"
	"github.com/lxc/lxd/shared/api"
)

var _ = api.ServerEnvironment{}

var instanceObjects = RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  ORDER BY projects.id, instances.name
`)

var instanceObjectsByID = RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE instances.id = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByProject = RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE project = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByProjectAndType = RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE project = ? AND instances.type = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByProjectAndTypeAndNode = RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE project = ? AND instances.type = ? AND node = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByProjectAndTypeAndNodeAndName = RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE project = ? AND instances.type = ? AND node = ? AND instances.name = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByProjectAndTypeAndName = RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE project = ? AND instances.type = ? AND instances.name = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByProjectAndName = RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE project = ? AND instances.name = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByProjectAndNameAndNode = RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE project = ? AND instances.name = ? AND node = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByProjectAndNode = RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE project = ? AND node = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByType = RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE instances.type = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByTypeAndName = RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE instances.type = ? AND instances.name = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByTypeAndNameAndNode = RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE instances.type = ? AND instances.name = ? AND node = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByTypeAndNode = RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE instances.type = ? AND node = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByNode = RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE node = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByNodeAndName = RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE node = ? AND instances.name = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByName = RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE instances.name = ? ORDER BY projects.id, instances.name
`)

var instanceID = RegisterStmt(`
SELECT instances.id FROM instances JOIN projects ON instances.project_id = projects.id
  WHERE projects.name = ? AND instances.name = ?
`)

var instanceCreate = RegisterStmt(`
INSERT INTO instances (project_id, name, node_id, type, architecture, ephemeral, creation_date, stateful, last_use_date, description, expiry_date)
  VALUES ((SELECT projects.id FROM projects WHERE projects.name = ?), ?, (SELECT nodes.id FROM nodes WHERE nodes.name = ?), ?, ?, ?, ?, ?, ?, ?, ?)
`)

var instanceRename = RegisterStmt(`
UPDATE instances SET name = ? WHERE project_id = (SELECT projects.id FROM projects WHERE projects.name = ?) AND name = ?
`)

var instanceDeleteByProjectAndName = RegisterStmt(`
DELETE FROM instances WHERE project_id = (SELECT projects.id FROM projects WHERE projects.name = ?) AND name = ?
`)

var instanceUpdate = RegisterStmt(`
UPDATE instances
  SET project_id = (SELECT id FROM projects WHERE name = ?), name = ?, node_id = (SELECT id FROM nodes WHERE name = ?), type = ?, architecture = ?, ephemeral = ?, creation_date = ?, stateful = ?, last_use_date = ?, description = ?, expiry_date = ?
 WHERE id = ?
`)

// GetInstances returns all available instances.
// generator: instance GetMany
func GetInstances(ctx context.Context, tx *sql.Tx, filter InstanceFilter) ([]Instance, error) {
	var err error

	// Result slice.
	objects := make([]Instance, 0)

	// Pick the prepared statement and arguments to use based on active criteria.
	var sqlStmt *sql.Stmt
	var args []any

	if filter.Project != nil && filter.Type != nil && filter.Node != nil && filter.Name != nil && filter.ID == nil {
		sqlStmt = stmt(tx, instanceObjectsByProjectAndTypeAndNodeAndName)
		args = []any{
			filter.Project,
			filter.Type,
			filter.Node,
			filter.Name,
		}
	} else if filter.Project != nil && filter.Type != nil && filter.Node != nil && filter.ID == nil && filter.Name == nil {
		sqlStmt = stmt(tx, instanceObjectsByProjectAndTypeAndNode)
		args = []any{
			filter.Project,
			filter.Type,
			filter.Node,
		}
	} else if filter.Project != nil && filter.Type != nil && filter.Name != nil && filter.ID == nil && filter.Node == nil {
		sqlStmt = stmt(tx, instanceObjectsByProjectAndTypeAndName)
		args = []any{
			filter.Project,
			filter.Type,
			filter.Name,
		}
	} else if filter.Type != nil && filter.Name != nil && filter.Node != nil && filter.ID == nil && filter.Project == nil {
		sqlStmt = stmt(tx, instanceObjectsByTypeAndNameAndNode)
		args = []any{
			filter.Type,
			filter.Name,
			filter.Node,
		}
	} else if filter.Project != nil && filter.Name != nil && filter.Node != nil && filter.ID == nil && filter.Type == nil {
		sqlStmt = stmt(tx, instanceObjectsByProjectAndNameAndNode)
		args = []any{
			filter.Project,
			filter.Name,
			filter.Node,
		}
	} else if filter.Project != nil && filter.Type != nil && filter.ID == nil && filter.Name == nil && filter.Node == nil {
		sqlStmt = stmt(tx, instanceObjectsByProjectAndType)
		args = []any{
			filter.Project,
			filter.Type,
		}
	} else if filter.Type != nil && filter.Node != nil && filter.ID == nil && filter.Project == nil && filter.Name == nil {
		sqlStmt = stmt(tx, instanceObjectsByTypeAndNode)
		args = []any{
			filter.Type,
			filter.Node,
		}
	} else if filter.Type != nil && filter.Name != nil && filter.ID == nil && filter.Project == nil && filter.Node == nil {
		sqlStmt = stmt(tx, instanceObjectsByTypeAndName)
		args = []any{
			filter.Type,
			filter.Name,
		}
	} else if filter.Project != nil && filter.Node != nil && filter.ID == nil && filter.Name == nil && filter.Type == nil {
		sqlStmt = stmt(tx, instanceObjectsByProjectAndNode)
		args = []any{
			filter.Project,
			filter.Node,
		}
	} else if filter.Project != nil && filter.Name != nil && filter.ID == nil && filter.Node == nil && filter.Type == nil {
		sqlStmt = stmt(tx, instanceObjectsByProjectAndName)
		args = []any{
			filter.Project,
			filter.Name,
		}
	} else if filter.Node != nil && filter.Name != nil && filter.ID == nil && filter.Project == nil && filter.Type == nil {
		sqlStmt = stmt(tx, instanceObjectsByNodeAndName)
		args = []any{
			filter.Node,
			filter.Name,
		}
	} else if filter.Type != nil && filter.ID == nil && filter.Project == nil && filter.Name == nil && filter.Node == nil {
		sqlStmt = stmt(tx, instanceObjectsByType)
		args = []any{
			filter.Type,
		}
	} else if filter.Project != nil && filter.ID == nil && filter.Name == nil && filter.Node == nil && filter.Type == nil {
		sqlStmt = stmt(tx, instanceObjectsByProject)
		args = []any{
			filter.Project,
		}
	} else if filter.Node != nil && filter.ID == nil && filter.Project == nil && filter.Name == nil && filter.Type == nil {
		sqlStmt = stmt(tx, instanceObjectsByNode)
		args = []any{
			filter.Node,
		}
	} else if filter.Name != nil && filter.ID == nil && filter.Project == nil && filter.Node == nil && filter.Type == nil {
		sqlStmt = stmt(tx, instanceObjectsByName)
		args = []any{
			filter.Name,
		}
	} else if filter.ID != nil && filter.Project == nil && filter.Name == nil && filter.Node == nil && filter.Type == nil {
		sqlStmt = stmt(tx, instanceObjectsByID)
		args = []any{
			filter.ID,
		}
	} else if filter.ID == nil && filter.Project == nil && filter.Name == nil && filter.Node == nil && filter.Type == nil {
		sqlStmt = stmt(tx, instanceObjects)
		args = []any{}
	} else {
		return nil, fmt.Errorf("No statement exists for the given Filter")
	}

	// Dest function for scanning a row.
	dest := func(i int) []any {
		objects = append(objects, Instance{})
		return []any{
			&objects[i].ID,
			&objects[i].Project,
			&objects[i].Name,
			&objects[i].Node,
			&objects[i].Type,
			&objects[i].Architecture,
			&objects[i].Ephemeral,
			&objects[i].CreationDate,
			&objects[i].Stateful,
			&objects[i].LastUseDate,
			&objects[i].Description,
			&objects[i].ExpiryDate,
		}
	}

	// Select.
	err = query.SelectObjects(sqlStmt, dest, args...)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from \"instances\" table: %w", err)
	}

	return objects, nil
}

// GetInstanceDevices returns all available Instance Devices
// generator: instance GetMany
func GetInstanceDevices(ctx context.Context, tx *sql.Tx, instanceID int) (map[string]Device, error) {
	instanceDevices, err := GetDevices(ctx, tx, "instance")
	if err != nil {
		return nil, err
	}

	devices := map[string]Device{}
	for _, ref := range instanceDevices[instanceID] {
		_, ok := devices[ref.Name]
		if !ok {
			devices[ref.Name] = ref
		} else {
			return nil, fmt.Errorf("Found duplicate Device with name %q", ref.Name)
		}
	}
	return devices, nil
}

// GetInstanceConfig returns all available Instance Config
// generator: instance GetMany
func GetInstanceConfig(ctx context.Context, tx *sql.Tx, instanceID int) (map[string]string, error) {
	instanceConfig, err := GetConfig(ctx, tx, "instance")
	if err != nil {
		return nil, err
	}

	config, ok := instanceConfig[instanceID]
	if !ok {
		config = map[string]string{}
	}

	return config, nil
}

// GetInstance returns the instance with the given key.
// generator: instance GetOne
func GetInstance(ctx context.Context, tx *sql.Tx, project string, name string) (*Instance, error) {
	filter := InstanceFilter{}
	filter.Project = &project
	filter.Name = &name

	objects, err := GetInstances(ctx, tx, filter)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from \"instances\" table: %w", err)
	}

	switch len(objects) {
	case 0:
		return nil, api.StatusErrorf(http.StatusNotFound, "Instance not found")
	case 1:
		return &objects[0], nil
	default:
		return nil, fmt.Errorf("More than one \"instances\" entry matches")
	}
}

// GetInstanceID return the ID of the instance with the given key.
// generator: instance ID
func GetInstanceID(ctx context.Context, tx *sql.Tx, project string, name string) (int64, error) {
	stmt := stmt(tx, instanceID)
	rows, err := stmt.Query(project, name)
	if err != nil {
		return -1, fmt.Errorf("Failed to get \"instances\" ID: %w", err)
	}

	defer func() { _ = rows.Close() }()

	// Ensure we read one and only one row.
	if !rows.Next() {
		return -1, api.StatusErrorf(http.StatusNotFound, "Instance not found")
	}

	var id int64
	err = rows.Scan(&id)
	if err != nil {
		return -1, fmt.Errorf("Failed to scan ID: %w", err)
	}

	if rows.Next() {
		return -1, fmt.Errorf("More than one row returned")
	}

	err = rows.Err()
	if err != nil {
		return -1, fmt.Errorf("Result set failure: %w", err)
	}

	return id, nil
}

// InstanceExists checks if a instance with the given key exists.
// generator: instance Exists
func InstanceExists(ctx context.Context, tx *sql.Tx, project string, name string) (bool, error) {
	_, err := GetInstanceID(ctx, tx, project, name)
	if err != nil {
		if api.StatusErrorCheck(err, http.StatusNotFound) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

// CreateInstance adds a new instance to the database.
// generator: instance Create
func CreateInstance(ctx context.Context, tx *sql.Tx, object Instance) (int64, error) {
	// Check if a instance with the same key exists.
	exists, err := InstanceExists(ctx, tx, object.Project, object.Name)
	if err != nil {
		return -1, fmt.Errorf("Failed to check for duplicates: %w", err)
	}

	if exists {
		return -1, api.StatusErrorf(http.StatusConflict, "This \"instances\" entry already exists")
	}

	args := make([]any, 11)

	// Populate the statement arguments.
	args[0] = object.Project
	args[1] = object.Name
	args[2] = object.Node
	args[3] = object.Type
	args[4] = object.Architecture
	args[5] = object.Ephemeral
	args[6] = object.CreationDate
	args[7] = object.Stateful
	args[8] = object.LastUseDate
	args[9] = object.Description
	args[10] = object.ExpiryDate

	// Prepared statement to use.
	stmt := stmt(tx, instanceCreate)

	// Execute the statement.
	result, err := stmt.Exec(args...)
	if err != nil {
		return -1, fmt.Errorf("Failed to create \"instances\" entry: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("Failed to fetch \"instances\" entry ID: %w", err)
	}

	return id, nil
}

// CreateInstanceDevices adds new instance Devices to the database.
// generator: instance Create
func CreateInstanceDevices(ctx context.Context, tx *sql.Tx, instanceID int64, devices map[string]Device) error {
	for key, device := range devices {
		device.ReferenceID = int(instanceID)
		devices[key] = device
	}

	err := CreateDevices(ctx, tx, "instance", devices)
	if err != nil {
		return fmt.Errorf("Insert Device failed for Instance: %w", err)
	}

	return nil
}

// CreateInstanceConfig adds new instance Config to the database.
// generator: instance Create
func CreateInstanceConfig(ctx context.Context, tx *sql.Tx, instanceID int64, config map[string]string) error {
	referenceID := int(instanceID)
	for key, value := range config {
		insert := Config{
			ReferenceID: referenceID,
			Key:         key,
			Value:       value,
		}

		err := CreateConfig(ctx, tx, "instance", insert)
		if err != nil {
			return fmt.Errorf("Insert Config failed for Instance: %w", err)
		}

	}

	return nil
}

// RenameInstance renames the instance matching the given key parameters.
// generator: instance Rename
func RenameInstance(ctx context.Context, tx *sql.Tx, project string, name string, to string) error {
	stmt := stmt(tx, instanceRename)
	result, err := stmt.Exec(to, project, name)
	if err != nil {
		return fmt.Errorf("Rename Instance failed: %w", err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Fetch affected rows failed: %w", err)
	}

	if n != 1 {
		return fmt.Errorf("Query affected %d rows instead of 1", n)
	}

	return nil
}

// DeleteInstance deletes the instance matching the given key parameters.
// generator: instance DeleteOne-by-Project-and-Name
func DeleteInstance(ctx context.Context, tx *sql.Tx, project string, name string) error {
	stmt := stmt(tx, instanceDeleteByProjectAndName)
	result, err := stmt.Exec(project, name)
	if err != nil {
		return fmt.Errorf("Delete \"instances\": %w", err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Fetch affected rows: %w", err)
	}

	if n == 0 {
		return api.StatusErrorf(http.StatusNotFound, "Instance not found")
	} else if n > 1 {
		return fmt.Errorf("Query deleted %d Instance rows instead of 1", n)
	}

	return nil
}

// UpdateInstance updates the instance matching the given key parameters.
// generator: instance Update
func UpdateInstance(ctx context.Context, tx *sql.Tx, project string, name string, object Instance) error {
	id, err := GetInstanceID(ctx, tx, project, name)
	if err != nil {
		return err
	}

	stmt := stmt(tx, instanceUpdate)
	result, err := stmt.Exec(object.Project, object.Name, object.Node, object.Type, object.Architecture, object.Ephemeral, object.CreationDate, object.Stateful, object.LastUseDate, object.Description, object.ExpiryDate, id)
	if err != nil {
		return fmt.Errorf("Update \"instances\" entry failed: %w", err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Fetch affected rows: %w", err)
	}

	if n != 1 {
		return fmt.Errorf("Query updated %d rows instead of 1", n)
	}

	return nil
}

// UpdateInstanceDevices updates the instance Device matching the given key parameters.
// generator: instance Update
func UpdateInstanceDevices(ctx context.Context, tx *sql.Tx, instanceID int64, devices map[string]Device) error {
	err := UpdateDevices(ctx, tx, "instance", int(instanceID), devices)
	if err != nil {
		return fmt.Errorf("Replace Device for Instance failed: %w", err)
	}

	return nil
}

// UpdateInstanceConfig updates the instance Config matching the given key parameters.
// generator: instance Update
func UpdateInstanceConfig(ctx context.Context, tx *sql.Tx, instanceID int64, config map[string]string) error {
	err := UpdateConfig(ctx, tx, "instance", int(instanceID), config)
	if err != nil {
		return fmt.Errorf("Replace Config for Instance failed: %w", err)
	}

	return nil
}