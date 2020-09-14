package v7action

import (
	"code.cloudfoundry.org/cli/actor/actionerror"
	"code.cloudfoundry.org/cli/resources"
)

// GetRevisionsByApplicationNameAndSpace returns revisions for application.
func (actor *Actor) GetRevisionsByApplicationNameAndSpace(appName string, spaceGUID string) ([]resources.Revision, Warnings, error) {
	app, warnings, appErr := actor.GetApplicationByNameAndSpace(appName, spaceGUID)
	if appErr != nil {
		return []resources.Revision{}, warnings, appErr
	}

	revisions, v3Warnings, apiErr := actor.CloudControllerClient.GetApplicationRevisions(app.GUID)
	warnings = append(warnings, v3Warnings...)

	return revisions, warnings, apiErr
}

func (actor Actor) GetRevisionByApplicationAndVersion(appGUID string, revisionVersion int) (resources.Revision, Warnings, error) {
	revisions, warnings, apiErr := actor.CloudControllerClient.GetApplicationRevisions(appGUID)
	if apiErr != nil {
		return resources.Revision{}, Warnings(warnings), apiErr
	}

	for _, revision := range revisions {
		if revision.Version == revisionVersion {
			return revision, Warnings(warnings), nil
		}
	}

	return resources.Revision{}, Warnings(warnings), actionerror.RevisionNotFoundError{Version: revisionVersion}
}

func (actor Actor) GetRevisionByApplicationNameAndSpaceAndVersion(appGUID string, spaceGUID string, revisionVersion int) (resources.Revision, Warnings, error) {
	return resources.Revision{}, nil, nil
}
