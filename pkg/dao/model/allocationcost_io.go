// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package model

import (
	"time"

	"github.com/seal-io/seal/pkg/dao/types"
)

// AllocationCostQueryInput is the input for the AllocationCost query.
type AllocationCostQueryInput struct {
	// ID holds the value of the "id" field.
	ID int `uri:"id,omitempty" json:"id,omitempty"`
}

// Model converts the AllocationCostQueryInput to AllocationCost.
func (in AllocationCostQueryInput) Model() *AllocationCost {
	return &AllocationCost{
		ID: in.ID,
	}
}

// AllocationCostCreateInput is the input for the AllocationCost creation.
type AllocationCostCreateInput struct {
	// Usage start time for current cost
	StartTime time.Time `json:"startTime,omitempty"`
	// Usage end time for current cost
	EndTime time.Time `json:"endTime,omitempty"`
	// Usage minutes from start time to end time
	Minutes float64 `json:"minutes,omitempty"`
	// Resource name for current cost, could be __unmounted__
	Name string `json:"name,omitempty"`
	// String generated from resource properties, used to identify this cost
	Fingerprint string `json:"fingerprint,omitempty"`
	// Cluster name for current cost
	ClusterName string `json:"clusterName,omitempty"`
	// Namespace for current cost
	Namespace string `json:"namespace,omitempty"`
	// Node for current cost
	Node string `json:"node,omitempty"`
	// Controller name for the cost linked resource
	Controller string `json:"controller,omitempty"`
	// Controller kind for the cost linked resource, deployment, statefulSet etc.
	ControllerKind string `json:"controllerKind,omitempty"`
	// Pod name for current cost
	Pod string `json:"pod,omitempty"`
	// Container name for current cost
	Container string `json:"container,omitempty"`
	// PV list for current cost linked
	Pvs map[string]types.PVCost `json:"pvs,omitempty"`
	// Labels for the cost linked resource
	Labels map[string]string `json:"labels,omitempty"`
	// Cost number
	TotalCost float64 `json:"totalCost,omitempty"`
	// Cost currency
	Currency int `json:"currency,omitempty"`
	// Cpu cost for current cost
	CpuCost float64 `json:"cpuCost,omitempty"`
	// Cpu core requested
	CpuCoreRequest float64 `json:"cpuCoreRequest,omitempty"`
	// GPU cost for current cost
	GpuCost float64 `json:"gpuCost,omitempty"`
	// GPU core count
	GpuCount float64 `json:"gpuCount,omitempty"`
	// Ram cost for current cost
	RamCost float64 `json:"ramCost,omitempty"`
	// Ram requested in byte
	RamByteRequest float64 `json:"ramByteRequest,omitempty"`
	// PV cost for current cost linked
	PvCost float64 `json:"pvCost,omitempty"`
	// PV bytes for current cost linked
	PvBytes float64 `json:"pvBytes,omitempty"`
	// CPU core average usage
	CpuCoreUsageAverage float64 `json:"cpuCoreUsageAverage,omitempty"`
	// CPU core max usage
	CpuCoreUsageMax float64 `json:"cpuCoreUsageMax,omitempty"`
	// Ram average usage in byte
	RamByteUsageAverage float64 `json:"ramByteUsageAverage,omitempty"`
	// Ram max usage in byte
	RamByteUsageMax float64 `json:"ramByteUsageMax,omitempty"`
	// Connector current cost linked
	Connector ConnectorQueryInput `json:"connector"`
}

// Model converts the AllocationCostCreateInput to AllocationCost.
func (in AllocationCostCreateInput) Model() *AllocationCost {
	var entity = &AllocationCost{
		StartTime:           in.StartTime,
		EndTime:             in.EndTime,
		Minutes:             in.Minutes,
		Name:                in.Name,
		Fingerprint:         in.Fingerprint,
		ClusterName:         in.ClusterName,
		Namespace:           in.Namespace,
		Node:                in.Node,
		Controller:          in.Controller,
		ControllerKind:      in.ControllerKind,
		Pod:                 in.Pod,
		Container:           in.Container,
		Pvs:                 in.Pvs,
		Labels:              in.Labels,
		TotalCost:           in.TotalCost,
		Currency:            in.Currency,
		CpuCost:             in.CpuCost,
		CpuCoreRequest:      in.CpuCoreRequest,
		GpuCost:             in.GpuCost,
		GpuCount:            in.GpuCount,
		RamCost:             in.RamCost,
		RamByteRequest:      in.RamByteRequest,
		PvCost:              in.PvCost,
		PvBytes:             in.PvBytes,
		CpuCoreUsageAverage: in.CpuCoreUsageAverage,
		CpuCoreUsageMax:     in.CpuCoreUsageMax,
		RamByteUsageAverage: in.RamByteUsageAverage,
		RamByteUsageMax:     in.RamByteUsageMax,
	}
	entity.ConnectorID = in.Connector.ID
	return entity
}

// AllocationCostUpdateInput is the input for the AllocationCost modification.
type AllocationCostUpdateInput struct {
	// ID holds the value of the "id" field.
	ID int `uri:"id" json:"-"`
	// Cost number
	TotalCost float64 `json:"totalCost,omitempty"`
	// Cost currency
	Currency int `json:"currency,omitempty"`
	// Cpu cost for current cost
	CpuCost float64 `json:"cpuCost,omitempty"`
	// GPU cost for current cost
	GpuCost float64 `json:"gpuCost,omitempty"`
	// Ram cost for current cost
	RamCost float64 `json:"ramCost,omitempty"`
	// PV cost for current cost linked
	PvCost float64 `json:"pvCost,omitempty"`
	// PV bytes for current cost linked
	PvBytes float64 `json:"pvBytes,omitempty"`
	// CPU core average usage
	CpuCoreUsageAverage float64 `json:"cpuCoreUsageAverage,omitempty"`
	// CPU core max usage
	CpuCoreUsageMax float64 `json:"cpuCoreUsageMax,omitempty"`
	// Ram average usage in byte
	RamByteUsageAverage float64 `json:"ramByteUsageAverage,omitempty"`
	// Ram max usage in byte
	RamByteUsageMax float64 `json:"ramByteUsageMax,omitempty"`
}

// Model converts the AllocationCostUpdateInput to AllocationCost.
func (in AllocationCostUpdateInput) Model() *AllocationCost {
	var entity = &AllocationCost{
		ID:                  in.ID,
		TotalCost:           in.TotalCost,
		Currency:            in.Currency,
		CpuCost:             in.CpuCost,
		GpuCost:             in.GpuCost,
		RamCost:             in.RamCost,
		PvCost:              in.PvCost,
		PvBytes:             in.PvBytes,
		CpuCoreUsageAverage: in.CpuCoreUsageAverage,
		CpuCoreUsageMax:     in.CpuCoreUsageMax,
		RamByteUsageAverage: in.RamByteUsageAverage,
		RamByteUsageMax:     in.RamByteUsageMax,
	}
	return entity
}

// AllocationCostOutput is the output for the AllocationCost.
type AllocationCostOutput struct {
	// ID holds the value of the "id" field.
	ID int `json:"id,omitempty"`
	// Usage start time for current cost
	StartTime time.Time `json:"startTime,omitempty"`
	// Usage end time for current cost
	EndTime time.Time `json:"endTime,omitempty"`
	// Usage minutes from start time to end time
	Minutes float64 `json:"minutes,omitempty"`
	// Resource name for current cost, could be __unmounted__
	Name string `json:"name,omitempty"`
	// String generated from resource properties, used to identify this cost
	Fingerprint string `json:"fingerprint,omitempty"`
	// Cluster name for current cost
	ClusterName string `json:"clusterName,omitempty"`
	// Namespace for current cost
	Namespace string `json:"namespace,omitempty"`
	// Node for current cost
	Node string `json:"node,omitempty"`
	// Controller name for the cost linked resource
	Controller string `json:"controller,omitempty"`
	// Controller kind for the cost linked resource, deployment, statefulSet etc.
	ControllerKind string `json:"controllerKind,omitempty"`
	// Pod name for current cost
	Pod string `json:"pod,omitempty"`
	// Container name for current cost
	Container string `json:"container,omitempty"`
	// PV list for current cost linked
	Pvs map[string]types.PVCost `json:"pvs,omitempty"`
	// Labels for the cost linked resource
	Labels map[string]string `json:"labels,omitempty"`
	// Cost number
	TotalCost float64 `json:"totalCost,omitempty"`
	// Cost currency
	Currency int `json:"currency,omitempty"`
	// Cpu cost for current cost
	CpuCost float64 `json:"cpuCost,omitempty"`
	// Cpu core requested
	CpuCoreRequest float64 `json:"cpuCoreRequest,omitempty"`
	// GPU cost for current cost
	GpuCost float64 `json:"gpuCost,omitempty"`
	// GPU core count
	GpuCount float64 `json:"gpuCount,omitempty"`
	// Ram cost for current cost
	RamCost float64 `json:"ramCost,omitempty"`
	// Ram requested in byte
	RamByteRequest float64 `json:"ramByteRequest,omitempty"`
	// PV cost for current cost linked
	PvCost float64 `json:"pvCost,omitempty"`
	// PV bytes for current cost linked
	PvBytes float64 `json:"pvBytes,omitempty"`
	// CPU core average usage
	CpuCoreUsageAverage float64 `json:"cpuCoreUsageAverage,omitempty"`
	// CPU core max usage
	CpuCoreUsageMax float64 `json:"cpuCoreUsageMax,omitempty"`
	// Ram average usage in byte
	RamByteUsageAverage float64 `json:"ramByteUsageAverage,omitempty"`
	// Ram max usage in byte
	RamByteUsageMax float64 `json:"ramByteUsageMax,omitempty"`
	// Connector current cost linked
	Connector *ConnectorOutput `json:"connector,omitempty"`
}

// ExposeAllocationCost converts the AllocationCost to AllocationCostOutput.
func ExposeAllocationCost(in *AllocationCost) *AllocationCostOutput {
	if in == nil {
		return nil
	}
	var entity = &AllocationCostOutput{
		ID:                  in.ID,
		StartTime:           in.StartTime,
		EndTime:             in.EndTime,
		Minutes:             in.Minutes,
		Name:                in.Name,
		Fingerprint:         in.Fingerprint,
		ClusterName:         in.ClusterName,
		Namespace:           in.Namespace,
		Node:                in.Node,
		Controller:          in.Controller,
		ControllerKind:      in.ControllerKind,
		Pod:                 in.Pod,
		Container:           in.Container,
		Pvs:                 in.Pvs,
		Labels:              in.Labels,
		TotalCost:           in.TotalCost,
		Currency:            in.Currency,
		CpuCost:             in.CpuCost,
		CpuCoreRequest:      in.CpuCoreRequest,
		GpuCost:             in.GpuCost,
		GpuCount:            in.GpuCount,
		RamCost:             in.RamCost,
		RamByteRequest:      in.RamByteRequest,
		PvCost:              in.PvCost,
		PvBytes:             in.PvBytes,
		CpuCoreUsageAverage: in.CpuCoreUsageAverage,
		CpuCoreUsageMax:     in.CpuCoreUsageMax,
		RamByteUsageAverage: in.RamByteUsageAverage,
		RamByteUsageMax:     in.RamByteUsageMax,
		Connector:           ExposeConnector(in.Edges.Connector),
	}
	if entity.Connector == nil {
		entity.Connector = &ConnectorOutput{}
	}
	entity.Connector.ID = in.ConnectorID
	return entity
}

// ExposeAllocationCosts converts the AllocationCost slice to AllocationCostOutput pointer slice.
func ExposeAllocationCosts(in []*AllocationCost) []*AllocationCostOutput {
	var out = make([]*AllocationCostOutput, 0, len(in))
	for i := 0; i < len(in); i++ {
		var o = ExposeAllocationCost(in[i])
		if o == nil {
			continue
		}
		out = append(out, o)
	}
	if len(out) == 0 {
		return nil
	}
	return out
}