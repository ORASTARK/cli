package ccv3

import (
	"bytes"
	"encoding/json"

	"code.cloudfoundry.org/cli/api/cloudcontroller"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccerror"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/internal"
	"code.cloudfoundry.org/cli/types"
)

// OrganizationQuota represents a Cloud Controller organization quota.
type OrganizationQuota struct {
	// GUID is the unique ID of the organization quota.
	GUID string `json:"guid,omitempty"`
	// Name is the name of the organization quota
	Name string `json:"name"`
	// Apps contain the various limits that are associated with applications
	Apps AppLimit `json:"apps"`
	// Services contain the various limits that are associated with services
	Services ServiceLimit `json:"services"`
	// Routes contain the various limits that are associated with routes
	Routes RouteLimit `json:"routes"`
}

type AppLimit struct {
	TotalMemory       types.NullInt `json:"total_memory_in_mb"`
	InstanceMemory    types.NullInt `json:"per_process_memory_in_mb"`
	TotalAppInstances types.NullInt `json:"total_instances"`
}

type ServiceLimit struct {
	TotalServiceInstances types.NullInt `json:"total_service_instances"`
	PaidServicePlans      bool          `json:"paid_services_allowed"`
}

type RouteLimit struct {
	TotalRoutes        types.NullInt `json:"total_routes"`
	TotalReservedPorts types.NullInt `json:"total_reserved_ports"`
}

func (client *Client) CreateOrganizationQuota(orgQuota OrganizationQuota) (OrganizationQuota, Warnings, error) {
	quotaBytes, err := json.Marshal(orgQuota)
	if err != nil {
		return OrganizationQuota{}, nil, err
	}
	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.PostOrganizationQuotaRequest,
		Body:        bytes.NewReader(quotaBytes),
	})

	if err != nil {
		return OrganizationQuota{}, nil, err
	}
	var responseOrgQuota OrganizationQuota
	response := cloudcontroller.Response{
		DecodeJSONResponseInto: &responseOrgQuota,
	}
	err = client.connection.Make(request, &response)
	if err != nil {
		return OrganizationQuota{}, response.Warnings, err
	}
	return responseOrgQuota, response.Warnings, nil
}

func (client *Client) GetOrganizationQuotas(query ...Query) ([]OrganizationQuota, Warnings, error) {
	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.GetOrganizationQuotasRequest,
		Query:       query,
	})
	if err != nil {
		return []OrganizationQuota{}, nil, err
	}

	var orgQuotasList []OrganizationQuota
	warnings, err := client.paginate(request, OrganizationQuota{}, func(item interface{}) error {
		if orgQuota, ok := item.(OrganizationQuota); ok {
			orgQuotasList = append(orgQuotasList, orgQuota)
		} else {
			return ccerror.UnknownObjectInListError{
				Expected:   OrganizationQuota{},
				Unexpected: item,
			}
		}
		return nil
	})

	return orgQuotasList, warnings, err
}