package v7

import (
	"context"
	"io"
	"time"

	"code.cloudfoundry.org/cli/actor/sharedaction"
	"code.cloudfoundry.org/cli/actor/v7action"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/constant"
	uaa "code.cloudfoundry.org/cli/api/uaa/constant"
	"code.cloudfoundry.org/cli/cf/configuration/coreconfig"
	"code.cloudfoundry.org/cli/resources"
	"code.cloudfoundry.org/cli/types"
	"github.com/SermoDigital/jose/jwt"
)

//go:generate counterfeiter . Actor

type Actor interface {
	ApplyOrganizationQuotaByName(quotaName string, orgGUID string) (v7action.Warnings, error)
	ApplySpaceQuotaByName(quotaName string, spaceGUID string, orgGUID string) (v7action.Warnings, error)
	AssignIsolationSegmentToSpaceByNameAndSpace(isolationSegmentName string, spaceGUID string) (v7action.Warnings, error)
	Authenticate(credentials map[string]string, origin string, grantType uaa.GrantType) error
	BindSecurityGroupToSpaces(securityGroupGUID string, spaces []resources.Space, lifecycle constant.SecurityGroupLifecycle) (v7action.Warnings, error)
	CancelDeployment(deploymentGUID string) (v7action.Warnings, error)
	CheckRoute(domainName string, hostname string, path string, port int) (bool, v7action.Warnings, error)
	ClearTarget()
	CloudControllerAPIVersion() string
	CopyPackage(sourceApp resources.Application, targetApp resources.Application) (v7action.Package, v7action.Warnings, error)
	CreateAndUploadBitsPackageByApplicationNameAndSpace(appName string, spaceGUID string, bitsPath string) (v7action.Package, v7action.Warnings, error)
	CreateApplicationDroplet(appGUID string) (resources.Droplet, v7action.Warnings, error)
	CreateApplicationInSpace(app resources.Application, spaceGUID string) (resources.Application, v7action.Warnings, error)
	CreateBitsPackageByApplication(appGUID string) (v7action.Package, v7action.Warnings, error)
	CreateBuildpack(buildpack v7action.Buildpack) (v7action.Buildpack, v7action.Warnings, error)
	CreateDeployment(appGUID string, dropletGUID string) (string, v7action.Warnings, error)
	CreateDockerPackageByApplication(appGUID string, dockerImageCredentials v7action.DockerImageCredentials) (v7action.Package, v7action.Warnings, error)
	CreateDockerPackageByApplicationNameAndSpace(appName string, spaceGUID string, dockerImageCredentials v7action.DockerImageCredentials) (v7action.Package, v7action.Warnings, error)
	CreateIsolationSegmentByName(isolationSegment v7action.IsolationSegment) (v7action.Warnings, error)
	CreateOrgRole(roleType constant.RoleType, orgGUID string, userNameOrGUID string, userOrigin string, isClient bool) (v7action.Warnings, error)
	CreateOrganization(orgName string) (resources.Organization, v7action.Warnings, error)
	CreateOrganizationQuota(name string, limits v7action.QuotaLimits) (v7action.Warnings, error)
	CreatePrivateDomain(domainName string, orgName string) (v7action.Warnings, error)
	CreateRoute(spaceGUID, domainName, hostname, path string, port int) (resources.Route, v7action.Warnings, error)
	CreateSecurityGroup(name, filePath string) (v7action.Warnings, error)
	CreateServiceBroker(model v7action.ServiceBrokerModel) (v7action.Warnings, error)
	CreateSharedDomain(domainName string, internal bool, routerGroupName string) (v7action.Warnings, error)
	CreateSpace(spaceName, orgGUID string) (resources.Space, v7action.Warnings, error)
	CreateSpaceQuota(spaceQuotaName string, orgGuid string, limits v7action.QuotaLimits) (v7action.Warnings, error)
	CreateSpaceRole(roleType constant.RoleType, orgGUID string, spaceGUID string, userNameOrGUID string, userOrigin string, isClient bool) (v7action.Warnings, error)
	CreateUser(username string, password string, origin string) (resources.User, v7action.Warnings, error)
	CreateUserProvidedServiceInstance(instance resources.ServiceInstance) (v7action.Warnings, error)
	DeleteApplicationByNameAndSpace(name, spaceGUID string, deleteRoutes bool) (v7action.Warnings, error)
	DeleteBuildpackByNameAndStack(buildpackName string, buildpackStack string) (v7action.Warnings, error)
	DeleteDomain(domain resources.Domain) (v7action.Warnings, error)
	DeleteInstanceByApplicationNameSpaceProcessTypeAndIndex(appName string, spaceGUID string, processType string, instanceIndex int) (v7action.Warnings, error)
	DeleteOrgRole(roleType constant.RoleType, orgGUID string, userNameOrGUID string, userOrigin string, isClient bool) (v7action.Warnings, error)
	DeleteOrganization(orgName string) (v7action.Warnings, error)
	DeleteOrganizationQuota(quotaName string) (v7action.Warnings, error)
	DeleteOrphanedRoutes(spaceGUID string) (v7action.Warnings, error)
	DeleteRoute(domainName, hostname, path string, port int) (v7action.Warnings, error)
	DeleteSecurityGroup(securityGroupName string) (v7action.Warnings, error)
	DeleteServiceBroker(serviceBrokerGUID string) (v7action.Warnings, error)
	DeleteSpaceByNameAndOrganizationName(spaceName string, orgName string) (v7action.Warnings, error)
	DeleteSpaceQuotaByName(quotaName string, orgGUID string) (v7action.Warnings, error)
	DeleteSpaceRole(roleType constant.RoleType, spaceGUID string, userNameOrGUID string, userOrigin string, isClient bool) (v7action.Warnings, error)
	DeleteUser(userGuid string) (v7action.Warnings, error)
	DeleteIsolationSegmentByName(name string) (v7action.Warnings, error)
	DeleteIsolationSegmentOrganizationByName(isolationSegmentName string, orgName string) (v7action.Warnings, error)
	DisableFeatureFlag(flagName string) (v7action.Warnings, error)
	DisableServiceAccess(offeringName, brokerName, orgName, planName string) (v7action.SkippedPlans, v7action.Warnings, error)
	DownloadCurrentDropletByAppName(appName string, spaceGUID string) ([]byte, string, v7action.Warnings, error)
	DownloadDropletByGUIDAndAppName(dropletGUID string, appName string, spaceGUID string) ([]byte, v7action.Warnings, error)
	EnableFeatureFlag(flagName string) (v7action.Warnings, error)
	EnableServiceAccess(offeringName, brokerName, orgName, planName string) (v7action.SkippedPlans, v7action.Warnings, error)
	EntitleIsolationSegmentToOrganizationByName(isolationSegmentName string, orgName string) (v7action.Warnings, error)
	GetAppFeature(appGUID string, featureName string) (ccv3.ApplicationFeature, v7action.Warnings, error)
	GetAppSummariesForSpace(spaceGUID string, labels string) ([]v7action.ApplicationSummary, v7action.Warnings, error)
	GetApplicationByNameAndSpace(appName string, spaceGUID string) (resources.Application, v7action.Warnings, error)
	GetApplicationDroplets(appName string, spaceGUID string) ([]resources.Droplet, v7action.Warnings, error)
	GetApplicationLabels(appName string, spaceGUID string) (map[string]types.NullString, v7action.Warnings, error)
	GetApplicationPackages(appName string, spaceGUID string) ([]v7action.Package, v7action.Warnings, error)
	GetApplicationProcessHealthChecksByNameAndSpace(appName string, spaceGUID string) ([]v7action.ProcessHealthCheck, v7action.Warnings, error)
	GetApplicationRoutes(appGUID string) ([]resources.Route, v7action.Warnings, error)
	GetApplicationTasks(appName string, sortOrder v7action.SortOrder) ([]v7action.Task, v7action.Warnings, error)
	GetApplicationsByNamesAndSpace(appNames []string, spaceGUID string) ([]resources.Application, v7action.Warnings, error)
	GetBuildpackLabels(buildpackName string, buildpackStack string) (map[string]types.NullString, v7action.Warnings, error)
	GetBuildpacks(labelSelector string) ([]v7action.Buildpack, v7action.Warnings, error)
	GetDefaultDomain(orgGUID string) (resources.Domain, v7action.Warnings, error)
	GetDetailedAppSummary(appName string, spaceGUID string, withObfuscatedValues bool) (v7action.DetailedApplicationSummary, v7action.Warnings, error)
	GetDomain(domainGUID string) (resources.Domain, v7action.Warnings, error)
	GetDomainByName(domainName string) (resources.Domain, v7action.Warnings, error)
	GetDomainLabels(domainName string) (map[string]types.NullString, v7action.Warnings, error)
	GetEffectiveIsolationSegmentBySpace(spaceGUID string, orgDefaultIsolationSegmentGUID string) (v7action.IsolationSegment, v7action.Warnings, error)
	GetEnvironmentVariableGroup(group constant.EnvironmentVariableGroupName) (v7action.EnvironmentVariableGroup, v7action.Warnings, error)
	GetEnvironmentVariablesByApplicationNameAndSpace(appName string, spaceGUID string) (v7action.EnvironmentVariableGroups, v7action.Warnings, error)
	GetFeatureFlagByName(featureFlagName string) (v7action.FeatureFlag, v7action.Warnings, error)
	GetFeatureFlags() ([]v7action.FeatureFlag, v7action.Warnings, error)
	GetGlobalRunningSecurityGroups() ([]resources.SecurityGroup, v7action.Warnings, error)
	GetGlobalStagingSecurityGroups() ([]resources.SecurityGroup, v7action.Warnings, error)
	GetIsolationSegmentsByOrganization(orgName string) ([]v7action.IsolationSegment, v7action.Warnings, error)
	GetIsolationSegmentByName(isoSegmentName string) (v7action.IsolationSegment, v7action.Warnings, error)
	GetIsolationSegmentSummaries() ([]v7action.IsolationSegmentSummary, v7action.Warnings, error)
	GetLatestActiveDeploymentForApp(appGUID string) (v7action.Deployment, v7action.Warnings, error)
	GetLogCacheEndpoint() (string, v7action.Warnings, error)
	GetLoginPrompts() map[string]coreconfig.AuthPrompt
	GetNewestReadyPackageForApplication(app resources.Application) (v7action.Package, v7action.Warnings, error)
	GetOrgUsersByRoleType(orgGUID string) (map[constant.RoleType][]resources.User, v7action.Warnings, error)
	GetOrganizationByName(orgName string) (resources.Organization, v7action.Warnings, error)
	GetOrganizationDomains(string, string) ([]resources.Domain, v7action.Warnings, error)
	GetOrganizationLabels(orgName string) (map[string]types.NullString, v7action.Warnings, error)
	GetOrganizationQuotaByName(orgQuotaName string) (resources.OrganizationQuota, v7action.Warnings, error)
	GetOrganizationQuotas() ([]resources.OrganizationQuota, v7action.Warnings, error)
	GetOrganizationSpaces(orgGUID string) ([]resources.Space, v7action.Warnings, error)
	GetOrganizationSpacesWithLabelSelector(orgGUID string, labelSelector string) ([]resources.Space, v7action.Warnings, error)
	GetOrganizationSummaryByName(orgName string) (v7action.OrganizationSummary, v7action.Warnings, error)
	GetOrganizations(labelSelector string) ([]resources.Organization, v7action.Warnings, error)
	GetProcessByTypeAndApplication(processType string, appGUID string) (v7action.Process, v7action.Warnings, error)
	GetRawApplicationManifestByNameAndSpace(appName string, spaceGUID string) ([]byte, v7action.Warnings, error)
	GetRecentEventsByApplicationNameAndSpace(appName string, spaceGUID string) ([]v7action.Event, v7action.Warnings, error)
	GetRecentLogsForApplicationByNameAndSpace(appName string, spaceGUID string, client sharedaction.LogCacheClient) ([]sharedaction.LogMessage, v7action.Warnings, error)
	GetRevisionsByApplicationNameAndSpace(appName string, spaceGUID string) (v7action.Revisions, v7action.Warnings, error)
	GetRouteByAttributes(domain resources.Domain, hostname string, path string, port int) (resources.Route, v7action.Warnings, error)
	GetRouteDestinationByAppGUID(route resources.Route, appGUID string) (resources.RouteDestination, error)
	GetRouteLabels(routeName string, spaceGUID string) (map[string]types.NullString, v7action.Warnings, error)
	GetRouterGroups() ([]v7action.RouterGroup, error)
	GetRouteSummaries([]resources.Route) ([]v7action.RouteSummary, v7action.Warnings, error)
	GetRoutesByOrg(orgGUID string, labels string) ([]resources.Route, v7action.Warnings, error)
	GetRoutesBySpace(spaceGUID string, labels string) ([]resources.Route, v7action.Warnings, error)
	GetSSHEnabled(appGUID string) (ccv3.SSHEnabled, v7action.Warnings, error)
	GetSSHEnabledByAppName(appName string, spaceGUID string) (ccv3.SSHEnabled, v7action.Warnings, error)
	GetSSHPasscode() (string, error)
	GetSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndex(appName string, spaceGUID string, processType string, processIndex uint) (v7action.SSHAuthentication, v7action.Warnings, error)
	GetSecurityGroup(securityGroupName string) (resources.SecurityGroup, v7action.Warnings, error)
	GetSecurityGroupSummary(securityGroupName string) (v7action.SecurityGroupSummary, v7action.Warnings, error)
	GetSecurityGroups() ([]v7action.SecurityGroupSummary, v7action.Warnings, error)
	GetServiceAccess(offeringName, brokerName, orgName string) ([]v7action.ServicePlanAccess, v7action.Warnings, error)
	GetServiceBrokerByName(serviceBrokerName string) (v7action.ServiceBroker, v7action.Warnings, error)
	GetServiceBrokerLabels(serviceBrokerName string) (map[string]types.NullString, v7action.Warnings, error)
	GetServiceBrokers() ([]v7action.ServiceBroker, v7action.Warnings, error)
	GetServiceInstanceByNameAndSpace(serviceInstanceName, spaceGUID string) (resources.ServiceInstance, v7action.Warnings, error)
	GetServiceInstanceByNameAndSpaceWithRelationships(serviceInstanceName, spaceGUID string) (v7action.ServiceInstanceWithRelationships, v7action.Warnings, error)
	GetServiceOfferingLabels(serviceOfferingName, serviceBrokerName string) (map[string]types.NullString, v7action.Warnings, error)
	GetServicePlanLabels(servicePlanName, serviceOfferingName, serviceBrokerName string) (map[string]types.NullString, v7action.Warnings, error)
	GetSpaceByNameAndOrganization(spaceName string, orgGUID string) (resources.Space, v7action.Warnings, error)
	GetSpaceFeature(spaceName string, orgGUID string, feature string) (bool, v7action.Warnings, error)
	GetSpaceLabels(spaceName string, orgGUID string) (map[string]types.NullString, v7action.Warnings, error)
	GetSpaceQuotaByName(spaceQuotaName string, orgGUID string) (resources.SpaceQuota, v7action.Warnings, error)
	GetSpaceQuotasByOrgGUID(orgGUID string) ([]resources.SpaceQuota, v7action.Warnings, error)
	GetSpaceSummaryByNameAndOrganization(spaceName string, orgGUID string) (v7action.SpaceSummary, v7action.Warnings, error)
	GetSpaceUsersByRoleType(spaceGuid string) (map[constant.RoleType][]resources.User, v7action.Warnings, error)
	GetStackByName(stackName string) (v7action.Stack, v7action.Warnings, error)
	GetStackLabels(stackName string) (map[string]types.NullString, v7action.Warnings, error)
	GetStacks(string) ([]v7action.Stack, v7action.Warnings, error)
	GetStreamingLogsForApplicationByNameAndSpace(appName string, spaceGUID string, client sharedaction.LogCacheClient) (<-chan sharedaction.LogMessage, <-chan error, context.CancelFunc, v7action.Warnings, error)
	GetTaskBySequenceIDAndApplication(sequenceID int, appGUID string) (v7action.Task, v7action.Warnings, error)
	GetUnstagedNewestPackageGUID(appGuid string) (string, v7action.Warnings, error)
	GetUser(username, origin string) (resources.User, error)
	MapRoute(routeGUID string, appGUID string) (v7action.Warnings, error)
	Marketplace(filter v7action.MarketplaceFilter) ([]v7action.ServiceOfferingWithPlans, v7action.Warnings, error)
	ParseAccessToken(accessToken string) (jwt.JWT, error)
	PollBuild(buildGUID string, appName string) (resources.Droplet, v7action.Warnings, error)
	PollPackage(pkg v7action.Package) (v7action.Package, v7action.Warnings, error)
	PollStart(app resources.Application, noWait bool, handleProcessStats func(string)) (v7action.Warnings, error)
	PollStartForRolling(app resources.Application, deploymentGUID string, noWait bool, handleProcessStats func(string)) (v7action.Warnings, error)
	PollUploadBuildpackJob(jobURL ccv3.JobURL) (v7action.Warnings, error)
	PrepareBuildpackBits(inputPath string, tmpDirPath string, downloader v7action.Downloader) (string, error)
	PurgeServiceOfferingByNameAndBroker(serviceOfferingName, serviceBrokerName string) (v7action.Warnings, error)
	RefreshAccessToken() (string, error)
	RenameApplicationByNameAndSpaceGUID(oldAppName, newAppName, spaceGUID string) (resources.Application, v7action.Warnings, error)
	RenameOrganization(oldOrgName, newOrgName string) (resources.Organization, v7action.Warnings, error)
	RenameSpaceByNameAndOrganizationGUID(oldSpaceName, newSpaceName, orgGUID string) (resources.Space, v7action.Warnings, error)
	ResetOrganizationDefaultIsolationSegment(orgGUID string) (v7action.Warnings, error)
	ResetSpaceIsolationSegment(orgGUID string, spaceGUID string) (string, v7action.Warnings, error)
	ResourceMatch(resources []sharedaction.V3Resource) ([]sharedaction.V3Resource, v7action.Warnings, error)
	RestartApplication(appGUID string, noWait bool) (v7action.Warnings, error)
	RunTask(appGUID string, task v7action.Task) (v7action.Task, v7action.Warnings, error)
	ScaleProcessByApplication(appGUID string, process v7action.Process) (v7action.Warnings, error)
	ScheduleTokenRefresh(func(time.Duration) <-chan time.Time, chan struct{}, chan struct{}) (<-chan error, error)
	SetApplicationDroplet(appGUID string, dropletGUID string) (v7action.Warnings, error)
	SetApplicationDropletByApplicationNameAndSpace(appName string, spaceGUID string, dropletGUID string) (v7action.Warnings, error)
	SetApplicationManifest(appGUID string, rawManifest []byte) (v7action.Warnings, error)
	SetApplicationProcessHealthCheckTypeByNameAndSpace(appName string, spaceGUID string, healthCheckType constant.HealthCheckType, httpEndpoint string, processType string, invocationTimeout int64) (resources.Application, v7action.Warnings, error)
	SetEnvironmentVariableByApplicationNameAndSpace(appName string, spaceGUID string, envPair v7action.EnvironmentVariablePair) (v7action.Warnings, error)
	SetEnvironmentVariableGroup(group constant.EnvironmentVariableGroupName, envVars ccv3.EnvironmentVariables) (v7action.Warnings, error)
	SetOrganizationDefaultIsolationSegment(orgGUID string, isoSegGUID string) (v7action.Warnings, error)
	SetSpaceManifest(spaceGUID string, rawManifest []byte) (v7action.Warnings, error)
	SetTarget(settings v7action.TargetSettings) (v7action.Warnings, error)
	SharePrivateDomain(domainName string, orgName string) (v7action.Warnings, error)
	StageApplicationPackage(pkgGUID string) (v7action.Build, v7action.Warnings, error)
	StagePackage(packageGUID, appName, spaceGUID string) (<-chan resources.Droplet, <-chan v7action.Warnings, <-chan error)
	StartApplication(appGUID string) (v7action.Warnings, error)
	StopApplication(appGUID string) (v7action.Warnings, error)
	TerminateTask(taskGUID string) (v7action.Task, v7action.Warnings, error)
	UAAAPIVersion() string
	UnbindSecurityGroup(securityGroupName string, orgGUID string, spaceGUID string, lifecycle constant.SecurityGroupLifecycle) (v7action.Warnings, error)
	UnmapRoute(routeGUID string, destinationGUID string) (v7action.Warnings, error)
	UnsetEnvironmentVariableByApplicationNameAndSpace(appName string, spaceGUID string, EnvironmentVariableName string) (v7action.Warnings, error)
	UnsetSpaceQuota(spaceQuotaName, spaceName, orgGUID string) (v7action.Warnings, error)
	UnsharePrivateDomain(domainName string, orgName string) (v7action.Warnings, error)
	UpdateAppFeature(app resources.Application, enabled bool, featureName string) (v7action.Warnings, error)
	UpdateApplication(app resources.Application) (resources.Application, v7action.Warnings, error)
	UpdateApplicationLabelsByApplicationName(string, string, map[string]types.NullString) (v7action.Warnings, error)
	UpdateBuildpackByNameAndStack(buildpackName string, buildpackStack string, buildpack v7action.Buildpack) (v7action.Buildpack, v7action.Warnings, error)
	UpdateBuildpackLabelsByBuildpackNameAndStack(string, string, map[string]types.NullString) (v7action.Warnings, error)
	UpdateDomainLabelsByDomainName(string, map[string]types.NullString) (v7action.Warnings, error)
	UpdateOrganizationLabelsByOrganizationName(string, map[string]types.NullString) (v7action.Warnings, error)
	UpdateOrganizationQuota(quotaName string, newName string, limits v7action.QuotaLimits) (v7action.Warnings, error)
	UpdateProcessByTypeAndApplication(processType string, appGUID string, updatedProcess v7action.Process) (v7action.Warnings, error)
	UpdateRouteLabels(string, string, map[string]types.NullString) (v7action.Warnings, error)
	UpdateSecurityGroup(name, filePath string) (v7action.Warnings, error)
	UpdateSecurityGroupGloballyEnabled(securityGroupName string, lifecycle constant.SecurityGroupLifecycle, enabled bool) (v7action.Warnings, error)
	UpdateServiceBroker(serviceBrokerGUID string, model v7action.ServiceBrokerModel) (v7action.Warnings, error)
	UpdateServiceBrokerLabelsByServiceBrokerName(string, map[string]types.NullString) (v7action.Warnings, error)
	UpdateServiceOfferingLabels(serviceOfferingName string, serviceBrokerName string, labels map[string]types.NullString) (v7action.Warnings, error)
	UpdateServicePlanLabels(servicePlanName string, serviceOfferingName string, serviceBrokerName string, labels map[string]types.NullString) (v7action.Warnings, error)
	UpdateSpaceFeature(spaceName string, orgGUID string, enableds bool, feature string) (v7action.Warnings, error)
	UpdateSpaceLabelsBySpaceName(string, string, map[string]types.NullString) (v7action.Warnings, error)
	UpdateSpaceQuota(quotaName, orgGUID, newName string, limits v7action.QuotaLimits) (v7action.Warnings, error)
	UpdateStackLabelsByStackName(string, map[string]types.NullString) (v7action.Warnings, error)
	UpdateUserPassword(userGUID string, oldPassword string, newPassword string) error
	UpdateUserProvidedServiceInstance(serviceInstanceName, spaceGUID string, serviceInstanceUpdates resources.ServiceInstance) (v7action.Warnings, error)
	UploadBitsPackage(pkg v7action.Package, matchedResources []sharedaction.V3Resource, newResources io.Reader, newResourcesLength int64) (v7action.Package, v7action.Warnings, error)
	UploadBuildpack(guid string, pathToBuildpackBits string, progressBar v7action.SimpleProgressBar) (ccv3.JobURL, v7action.Warnings, error)
	UploadDroplet(dropletGUID string, dropletPath string, progressReader io.Reader, fileSize int64) (v7action.Warnings, error)
}
