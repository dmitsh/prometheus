// Copyright 2015 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package azure

import (
	"reflect"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
)

func TestMapFromVMWithEmptyTags(t *testing.T) {
	id := "test"
	name := "name"
	vmType := "type"
	location := "westeurope"
	networkProfile := compute.NetworkProfile{
		NetworkInterfaces: &[]compute.NetworkInterfaceReference{},
	}
	provisioningStatusCode := "ProvisioningState/succeeded"
	provisionDisplayStatus := "Provisioning succeeded"
	powerStatusCode := "PowerState/running"
	powerDisplayStatus := "VM running"
	properties := &compute.VirtualMachineProperties{
		StorageProfile: &compute.StorageProfile{
			OsDisk: &compute.OSDisk{
				OsType: "Linux",
			},
		},
		NetworkProfile: &networkProfile,
		InstanceView: &compute.VirtualMachineInstanceView{
			Statuses: &[]compute.InstanceViewStatus{
				{
					Code:          &provisioningStatusCode,
					Level:         "Info",
					DisplayStatus: &provisionDisplayStatus,
				},
				{
					Code:          &powerStatusCode,
					Level:         "Info",
					DisplayStatus: &powerDisplayStatus,
				},
			},
		},
	}

	testVM := compute.VirtualMachine{
		ID:                       &id,
		Name:                     &name,
		Type:                     &vmType,
		Location:                 &location,
		Tags:                     nil,
		VirtualMachineProperties: properties,
	}

	expectedVM := virtualMachine{
		ID:                id,
		Name:              name,
		Type:              vmType,
		Location:          location,
		OsType:            "Linux",
		Tags:              map[string]*string{},
		NetworkInterfaces: []string{},
		PowerState:        "running",
	}

	actualVM := mapFromVM(testVM)

	if !reflect.DeepEqual(expectedVM, actualVM) {
		t.Errorf("Expected %v got %v", expectedVM, actualVM)
	}
}

func TestMapFromVMWithTags(t *testing.T) {
	id := "test"
	name := "name"
	vmType := "type"
	location := "westeurope"
	tags := map[string]*string{
		"prometheus": new(string),
	}
	provisioningStatusCode := "ProvisioningState/succeeded"
	provisionDisplayStatus := "Provisioning succeeded"
	powerStatusCode := "PowerState/running"
	powerDisplayStatus := "VM running"
	networkProfile := compute.NetworkProfile{
		NetworkInterfaces: &[]compute.NetworkInterfaceReference{},
	}
	properties := &compute.VirtualMachineProperties{
		StorageProfile: &compute.StorageProfile{
			OsDisk: &compute.OSDisk{
				OsType: "Linux",
			},
		},
		NetworkProfile: &networkProfile,
		InstanceView: &compute.VirtualMachineInstanceView{
			Statuses: &[]compute.InstanceViewStatus{
				{
					Code:          &provisioningStatusCode,
					Level:         "Info",
					DisplayStatus: &provisionDisplayStatus,
				},
				{
					Code:          &powerStatusCode,
					Level:         "Info",
					DisplayStatus: &powerDisplayStatus,
				},
			},
		},
	}

	testVM := compute.VirtualMachine{
		ID:                       &id,
		Name:                     &name,
		Type:                     &vmType,
		Location:                 &location,
		Tags:                     tags,
		VirtualMachineProperties: properties,
	}

	expectedVM := virtualMachine{
		ID:                id,
		Name:              name,
		Type:              vmType,
		Location:          location,
		OsType:            "Linux",
		Tags:              tags,
		NetworkInterfaces: []string{},
		PowerState:        "running",
	}

	actualVM := mapFromVM(testVM)

	if !reflect.DeepEqual(expectedVM, actualVM) {
		t.Errorf("Expected %v got %v", expectedVM, actualVM)
	}
}

func TestMapFromVMScaleSetVMWithEmptyTags(t *testing.T) {
	id := "test"
	name := "name"
	vmType := "type"
	location := "westeurope"
	networkProfile := compute.NetworkProfile{
		NetworkInterfaces: &[]compute.NetworkInterfaceReference{},
	}
	provisioningStatusCode := "ProvisioningState/succeeded"
	provisionDisplayStatus := "Provisioning succeeded"
	powerStatusCode := "PowerState/running"
	powerDisplayStatus := "VM running"
	properties := &compute.VirtualMachineScaleSetVMProperties{
		StorageProfile: &compute.StorageProfile{
			OsDisk: &compute.OSDisk{
				OsType: "Linux",
			},
		},
		NetworkProfile: &networkProfile,
		InstanceView: &compute.VirtualMachineScaleSetVMInstanceView{
			Statuses: &[]compute.InstanceViewStatus{
				{
					Code:          &provisioningStatusCode,
					Level:         "Info",
					DisplayStatus: &provisionDisplayStatus,
				},
				{
					Code:          &powerStatusCode,
					Level:         "Info",
					DisplayStatus: &powerDisplayStatus,
				},
			},
		},
	}

	testVM := compute.VirtualMachineScaleSetVM{
		ID:                                 &id,
		Name:                               &name,
		Type:                               &vmType,
		Location:                           &location,
		Tags:                               nil,
		VirtualMachineScaleSetVMProperties: properties,
	}

	scaleSet := "testSet"
	expectedVM := virtualMachine{
		ID:                id,
		Name:              name,
		Type:              vmType,
		Location:          location,
		OsType:            "Linux",
		Tags:              map[string]*string{},
		NetworkInterfaces: []string{},
		ScaleSet:          scaleSet,
		PowerState:        "running",
	}

	actualVM := mapFromVMScaleSetVM(testVM, scaleSet)

	if !reflect.DeepEqual(expectedVM, actualVM) {
		t.Errorf("Expected %v got %v", expectedVM, actualVM)
	}
}

func TestMapFromVMScaleSetVMWithTags(t *testing.T) {
	id := "test"
	name := "name"
	vmType := "type"
	location := "westeurope"
	tags := map[string]*string{
		"prometheus": new(string),
	}
	networkProfile := compute.NetworkProfile{
		NetworkInterfaces: &[]compute.NetworkInterfaceReference{},
	}
	provisioningStatusCode := "ProvisioningState/succeeded"
	provisionDisplayStatus := "Provisioning succeeded"
	powerStatusCode := "PowerState/running"
	powerDisplayStatus := "VM running"
	properties := &compute.VirtualMachineScaleSetVMProperties{
		StorageProfile: &compute.StorageProfile{
			OsDisk: &compute.OSDisk{
				OsType: "Linux",
			},
		},
		NetworkProfile: &networkProfile,
		InstanceView: &compute.VirtualMachineScaleSetVMInstanceView{
			Statuses: &[]compute.InstanceViewStatus{
				{
					Code:          &provisioningStatusCode,
					Level:         "Info",
					DisplayStatus: &provisionDisplayStatus,
				},
				{
					Code:          &powerStatusCode,
					Level:         "Info",
					DisplayStatus: &powerDisplayStatus,
				},
			},
		},
	}

	testVM := compute.VirtualMachineScaleSetVM{
		ID:                                 &id,
		Name:                               &name,
		Type:                               &vmType,
		Location:                           &location,
		Tags:                               tags,
		VirtualMachineScaleSetVMProperties: properties,
	}

	scaleSet := "testSet"
	expectedVM := virtualMachine{
		ID:                id,
		Name:              name,
		Type:              vmType,
		Location:          location,
		OsType:            "Linux",
		Tags:              tags,
		NetworkInterfaces: []string{},
		ScaleSet:          scaleSet,
		PowerState:        "running",
	}

	actualVM := mapFromVMScaleSetVM(testVM, scaleSet)

	if !reflect.DeepEqual(expectedVM, actualVM) {
		t.Errorf("Expected %v got %v", expectedVM, actualVM)
	}
}

func TestGetPowerStatusFromVM(t *testing.T) {
	provisioningStatusCode := "ProvisioningState/succeeded"
	provisionDisplayStatus := "Provisioning succeeded"
	powerStatusCode := "PowerState/running"
	powerDisplayStatus := "VM running"
	properties := &compute.VirtualMachineScaleSetVMProperties{
		StorageProfile: &compute.StorageProfile{
			OsDisk: &compute.OSDisk{
				OsType: "Linux",
			},
		},
		InstanceView: &compute.VirtualMachineScaleSetVMInstanceView{
			Statuses: &[]compute.InstanceViewStatus{
				{
					Code:          &provisioningStatusCode,
					Level:         "Info",
					DisplayStatus: &provisionDisplayStatus,
				},
				{
					Code:          &powerStatusCode,
					Level:         "Info",
					DisplayStatus: &powerDisplayStatus,
				},
			},
		},
	}

	testVM := compute.VirtualMachineScaleSetVM{
		VirtualMachineScaleSetVMProperties: properties,
	}

	actual := getPowerState(testVM.VirtualMachineScaleSetVMProperties.InstanceView.Statuses)
	expected := "running"
	if actual != expected {
		t.Errorf("expected powerStatus %s, but got %s instead", expected, actual)
	}

	// Noq we test a virtualMachine with an empty InstanceView struct.
	testVM.VirtualMachineScaleSetVMProperties.InstanceView = &compute.VirtualMachineScaleSetVMInstanceView{}

	actual = getPowerState(testVM.VirtualMachineScaleSetVMProperties.InstanceView.Statuses)
	expected = powerStateUnknown
	if actual != expected {
		t.Errorf("expected powerStatus %s, but got %s instead", expected, actual)
	}
}
