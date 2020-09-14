package ccv3

import (
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/internal"
	"code.cloudfoundry.org/cli/resources"
)

func (client *Client) GetApplicationRevisions(appGUID string, query ...Query) ([]resources.Revision, Warnings, error) {
	var revisions []resources.Revision

	_, warnings, err := client.MakeListRequest(RequestParams{
		RequestName:  internal.GetApplicationRevisionsRequest,
		URIParams:    internal.Params{"app_guid": appGUID},
		Query:        query,
		ResponseBody: resources.Revision{},
		AppendToList: func(item interface{}) error {
			revisions = append(revisions, item.(resources.Revision))
			return nil
		},
	})
	return revisions, warnings, err
}
