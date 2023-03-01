// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package clustercost

const (
	// Label holds the string label denoting the clustercost type in the database.
	Label = "cluster_cost"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStartTime holds the string denoting the starttime field in the database.
	FieldStartTime = "start_time"
	// FieldEndTime holds the string denoting the endtime field in the database.
	FieldEndTime = "end_time"
	// FieldMinutes holds the string denoting the minutes field in the database.
	FieldMinutes = "minutes"
	// FieldConnectorID holds the string denoting the connectorid field in the database.
	FieldConnectorID = "connector_id"
	// FieldClusterName holds the string denoting the clustername field in the database.
	FieldClusterName = "cluster_name"
	// FieldTotalCost holds the string denoting the totalcost field in the database.
	FieldTotalCost = "total_cost"
	// FieldCurrency holds the string denoting the currency field in the database.
	FieldCurrency = "currency"
	// FieldCpuCost holds the string denoting the cpucost field in the database.
	FieldCpuCost = "cpu_cost"
	// FieldGpuCost holds the string denoting the gpucost field in the database.
	FieldGpuCost = "gpu_cost"
	// FieldRamCost holds the string denoting the ramcost field in the database.
	FieldRamCost = "ram_cost"
	// FieldStorageCost holds the string denoting the storagecost field in the database.
	FieldStorageCost = "storage_cost"
	// FieldAllocationCost holds the string denoting the allocationcost field in the database.
	FieldAllocationCost = "allocation_cost"
	// FieldIdleCost holds the string denoting the idlecost field in the database.
	FieldIdleCost = "idle_cost"
	// FieldManagementCost holds the string denoting the managementcost field in the database.
	FieldManagementCost = "management_cost"
	// EdgeConnector holds the string denoting the connector edge name in mutations.
	EdgeConnector = "connector"
	// Table holds the table name of the clustercost in the database.
	Table = "cluster_costs"
	// ConnectorTable is the table that holds the connector relation/edge.
	ConnectorTable = "cluster_costs"
	// ConnectorInverseTable is the table name for the Connector entity.
	// It exists in this package in order to avoid circular dependency with the "connector" package.
	ConnectorInverseTable = "connectors"
	// ConnectorColumn is the table column denoting the connector relation/edge.
	ConnectorColumn = "connector_id"
)

// Columns holds all SQL columns for clustercost fields.
var Columns = []string{
	FieldID,
	FieldStartTime,
	FieldEndTime,
	FieldMinutes,
	FieldConnectorID,
	FieldClusterName,
	FieldTotalCost,
	FieldCurrency,
	FieldCpuCost,
	FieldGpuCost,
	FieldRamCost,
	FieldStorageCost,
	FieldAllocationCost,
	FieldIdleCost,
	FieldManagementCost,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// ConnectorIDValidator is a validator for the "connectorID" field. It is called by the builders before save.
	ConnectorIDValidator func(string) error
	// ClusterNameValidator is a validator for the "clusterName" field. It is called by the builders before save.
	ClusterNameValidator func(string) error
	// DefaultTotalCost holds the default value on creation for the "totalCost" field.
	DefaultTotalCost float64
	// TotalCostValidator is a validator for the "totalCost" field. It is called by the builders before save.
	TotalCostValidator func(float64) error
	// DefaultCpuCost holds the default value on creation for the "cpuCost" field.
	DefaultCpuCost float64
	// CpuCostValidator is a validator for the "cpuCost" field. It is called by the builders before save.
	CpuCostValidator func(float64) error
	// DefaultGpuCost holds the default value on creation for the "gpuCost" field.
	DefaultGpuCost float64
	// GpuCostValidator is a validator for the "gpuCost" field. It is called by the builders before save.
	GpuCostValidator func(float64) error
	// DefaultRamCost holds the default value on creation for the "ramCost" field.
	DefaultRamCost float64
	// RamCostValidator is a validator for the "ramCost" field. It is called by the builders before save.
	RamCostValidator func(float64) error
	// DefaultStorageCost holds the default value on creation for the "storageCost" field.
	DefaultStorageCost float64
	// StorageCostValidator is a validator for the "storageCost" field. It is called by the builders before save.
	StorageCostValidator func(float64) error
	// DefaultAllocationCost holds the default value on creation for the "allocationCost" field.
	DefaultAllocationCost float64
	// AllocationCostValidator is a validator for the "allocationCost" field. It is called by the builders before save.
	AllocationCostValidator func(float64) error
	// DefaultIdleCost holds the default value on creation for the "idleCost" field.
	DefaultIdleCost float64
	// IdleCostValidator is a validator for the "idleCost" field. It is called by the builders before save.
	IdleCostValidator func(float64) error
	// DefaultManagementCost holds the default value on creation for the "managementCost" field.
	DefaultManagementCost float64
	// ManagementCostValidator is a validator for the "managementCost" field. It is called by the builders before save.
	ManagementCostValidator func(float64) error
)

// WithoutFields returns the fields ignored the given list.
func WithoutFields(ignores ...string) []string {
	if len(ignores) == 0 {
		return Columns
	}

	var s = make(map[string]bool, len(ignores))
	for i := range ignores {
		s[ignores[i]] = true
	}

	var r = make([]string, 0, len(Columns)-len(s))
	for i := range Columns {
		if s[Columns[i]] {
			continue
		}
		r = append(r, Columns[i])
	}
	return r
}