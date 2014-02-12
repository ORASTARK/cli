package api

import (
	"cf/models"
	"fmt"
)

type PaginatedServiceOfferingResources struct {
	Resources []ServiceOfferingResource
}

type ServiceOfferingResource struct {
	Metadata Metadata
	Entity   ServiceOfferingEntity
}

func (resource ServiceOfferingResource) ToFields() (fields models.ServiceOfferingFields) {
	fields.Label = resource.Entity.Label
	fields.Version = resource.Entity.Version
	fields.Provider = resource.Entity.Provider
	fields.Description = resource.Entity.Description
	fields.Guid = resource.Metadata.Guid
	fields.DocumentationUrl = resource.Entity.DocumentationUrl
	return
}

func (resource ServiceOfferingResource) ToModel() (offering models.ServiceOffering) {
	offering.ServiceOfferingFields = resource.ToFields()
	for _, p := range resource.Entity.ServicePlans {
		servicePlan := models.ServicePlanFields{}
		servicePlan.Name = p.Entity.Name
		servicePlan.Guid = p.Metadata.Guid
		offering.Plans = append(offering.Plans, servicePlan)
	}
	return offering
}

type ServiceOfferingEntity struct {
	Label            string
	Version          string
	Description      string
	DocumentationUrl string `json:"documentation_url"`
	Provider         string
	ServicePlans     []ServicePlanResource `json:"service_plans"`
}

type PaginatedServicePlanResources struct {
	Resources []ServicePlanResource
}

type ServicePlanResource struct {
	Metadata Metadata
	Entity   ServicePlanEntity
}

func (resource ServicePlanResource) ToFields() (fields models.ServicePlanFields) {
	fields.Guid = resource.Metadata.Guid
	fields.Name = resource.Entity.Name
	return
}

type ServicePlanEntity struct {
	Name            string
	ServiceOffering ServiceOfferingResource `json:"service"`
}

type PaginatedServiceInstanceResources struct {
	TotalResults int `json:"total_results"`
	Resources    []ServiceInstanceResource
}

type ServiceInstanceResource struct {
	Metadata Metadata
	Entity   ServiceInstanceEntity
}

func (resource ServiceInstanceResource) ToFields() (fields models.ServiceInstanceFields) {
	fields.Guid = resource.Metadata.Guid
	fields.Name = resource.Entity.Name
	return
}

func (resource ServiceInstanceResource) ToModel() (instance models.ServiceInstance) {
	instance.ServiceInstanceFields = resource.ToFields()
	instance.ServicePlan = resource.Entity.ServicePlan.ToFields()
	instance.ServiceOffering = resource.Entity.ServicePlan.Entity.ServiceOffering.ToFields()

	instance.ServiceBindings = []models.ServiceBindingFields{}
	for _, bindingResource := range resource.Entity.ServiceBindings {
		instance.ServiceBindings = append(instance.ServiceBindings, bindingResource.ToFields())
	}
	return
}

type ServiceInstanceEntity struct {
	Name            string
	ServiceBindings []ServiceBindingResource `json:"service_bindings"`
	ServicePlan     ServicePlanResource      `json:"service_plan"`
}

type ServiceBindingResource struct {
	Metadata Metadata
	Entity   ServiceBindingEntity
}

func (resource ServiceBindingResource) ToFields() (fields models.ServiceBindingFields) {
	fields.Url = resource.Metadata.Url
	fields.Guid = resource.Metadata.Guid
	fields.AppGuid = resource.Entity.AppGuid
	return
}

type ServiceBindingEntity struct {
	AppGuid string `json:"app_guid"`
}

type V1ServicePlanDescription struct {
	ServiceName     string
	ServicePlanName string
	ServiceProvider string
}

func (v1PlanDesc V1ServicePlanDescription) String() string {
	return fmt.Sprintf("%s %s %s", v1PlanDesc.ServiceName, v1PlanDesc.ServiceProvider, v1PlanDesc.ServicePlanName)
}

type V2ServicePlanDescription struct {
	ServiceName     string
	ServicePlanName string
}

func (v2PlanDesc V2ServicePlanDescription) String() string {
	return fmt.Sprintf("%s %s", v2PlanDesc.ServiceName, v2PlanDesc.ServicePlanName)
}
