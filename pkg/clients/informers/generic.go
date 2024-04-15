// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package informers

import (
	"fmt"

	v1alpha1 "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	workflowv1alpha1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	walrusv1 "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	walruscorev1 "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	v1 "k8s.io/api/admissionregistration/v1"
	appsv1 "k8s.io/api/apps/v1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	v2 "k8s.io/api/autoscaling/v2"
	batchv1 "k8s.io/api/batch/v1"
	certificatesv1 "k8s.io/api/certificates/v1"
	coordinationv1 "k8s.io/api/coordination/v1"
	corev1 "k8s.io/api/core/v1"
	discoveryv1 "k8s.io/api/discovery/v1"
	eventsv1 "k8s.io/api/events/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	schedulingv1 "k8s.io/api/scheduling/v1"
	storagev1 "k8s.io/api/storage/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
	apiregistrationv1 "k8s.io/kube-aggregator/pkg/apis/apiregistration/v1"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=admissionregistration.k8s.io, Version=v1
	case v1.SchemeGroupVersion.WithResource("mutatingwebhookconfigurations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Admissionregistration().V1().MutatingWebhookConfigurations().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("validatingwebhookconfigurations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Admissionregistration().V1().ValidatingWebhookConfigurations().Informer()}, nil

		// Group=apiextensions.k8s.io, Version=v1
	case apiextensionsv1.SchemeGroupVersion.WithResource("customresourcedefinitions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apiextensions().V1().CustomResourceDefinitions().Informer()}, nil

		// Group=apiregistration.k8s.io, Version=v1
	case apiregistrationv1.SchemeGroupVersion.WithResource("apiservices"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apiregistration().V1().APIServices().Informer()}, nil

		// Group=apps, Version=v1
	case appsv1.SchemeGroupVersion.WithResource("controllerrevisions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1().ControllerRevisions().Informer()}, nil
	case appsv1.SchemeGroupVersion.WithResource("daemonsets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1().DaemonSets().Informer()}, nil
	case appsv1.SchemeGroupVersion.WithResource("deployments"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1().Deployments().Informer()}, nil
	case appsv1.SchemeGroupVersion.WithResource("replicasets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1().ReplicaSets().Informer()}, nil
	case appsv1.SchemeGroupVersion.WithResource("statefulsets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1().StatefulSets().Informer()}, nil

		// Group=argoproj.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("appprojects"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Argoprojapplication().V1alpha1().AppProjects().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("applications"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Argoprojapplication().V1alpha1().Applications().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("applicationsets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Argoprojapplication().V1alpha1().ApplicationSets().Informer()}, nil

		// Group=argoproj.io, Version=v1alpha1
	case workflowv1alpha1.SchemeGroupVersion.WithResource("clusterworkflowtemplates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Argoprojworkflow().V1alpha1().ClusterWorkflowTemplates().Informer()}, nil
	case workflowv1alpha1.SchemeGroupVersion.WithResource("cronworkflows"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Argoprojworkflow().V1alpha1().CronWorkflows().Informer()}, nil
	case workflowv1alpha1.SchemeGroupVersion.WithResource("workflows"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Argoprojworkflow().V1alpha1().Workflows().Informer()}, nil
	case workflowv1alpha1.SchemeGroupVersion.WithResource("workflowartifactgctasks"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Argoprojworkflow().V1alpha1().WorkflowArtifactGCTasks().Informer()}, nil
	case workflowv1alpha1.SchemeGroupVersion.WithResource("workfloweventbindings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Argoprojworkflow().V1alpha1().WorkflowEventBindings().Informer()}, nil
	case workflowv1alpha1.SchemeGroupVersion.WithResource("workflowtaskresults"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Argoprojworkflow().V1alpha1().WorkflowTaskResults().Informer()}, nil
	case workflowv1alpha1.SchemeGroupVersion.WithResource("workflowtasksets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Argoprojworkflow().V1alpha1().WorkflowTaskSets().Informer()}, nil
	case workflowv1alpha1.SchemeGroupVersion.WithResource("workflowtemplates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Argoprojworkflow().V1alpha1().WorkflowTemplates().Informer()}, nil

		// Group=autoscaling, Version=v1
	case autoscalingv1.SchemeGroupVersion.WithResource("horizontalpodautoscalers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Autoscaling().V1().HorizontalPodAutoscalers().Informer()}, nil

		// Group=autoscaling, Version=v2
	case v2.SchemeGroupVersion.WithResource("horizontalpodautoscalers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Autoscaling().V2().HorizontalPodAutoscalers().Informer()}, nil

		// Group=batch, Version=v1
	case batchv1.SchemeGroupVersion.WithResource("cronjobs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Batch().V1().CronJobs().Informer()}, nil
	case batchv1.SchemeGroupVersion.WithResource("jobs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Batch().V1().Jobs().Informer()}, nil

		// Group=certificates.k8s.io, Version=v1
	case certificatesv1.SchemeGroupVersion.WithResource("certificatesigningrequests"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Certificates().V1().CertificateSigningRequests().Informer()}, nil

		// Group=coordination.k8s.io, Version=v1
	case coordinationv1.SchemeGroupVersion.WithResource("leases"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Coordination().V1().Leases().Informer()}, nil

		// Group=core, Version=v1
	case corev1.SchemeGroupVersion.WithResource("componentstatuses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().ComponentStatuses().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("configmaps"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().ConfigMaps().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("endpoints"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().Endpoints().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("events"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().Events().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("limitranges"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().LimitRanges().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("namespaces"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().Namespaces().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("nodes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().Nodes().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("persistentvolumes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().PersistentVolumes().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("persistentvolumeclaims"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().PersistentVolumeClaims().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("pods"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().Pods().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("podtemplates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().PodTemplates().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("replicationcontrollers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().ReplicationControllers().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("resourcequotas"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().ResourceQuotas().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("secrets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().Secrets().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("services"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().Services().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("serviceaccounts"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().ServiceAccounts().Informer()}, nil

		// Group=discovery.k8s.io, Version=v1
	case discoveryv1.SchemeGroupVersion.WithResource("endpointslices"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Discovery().V1().EndpointSlices().Informer()}, nil

		// Group=events.k8s.io, Version=v1
	case eventsv1.SchemeGroupVersion.WithResource("events"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Events().V1().Events().Informer()}, nil

		// Group=rbac.authorization.k8s.io, Version=v1
	case rbacv1.SchemeGroupVersion.WithResource("clusterroles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rbac().V1().ClusterRoles().Informer()}, nil
	case rbacv1.SchemeGroupVersion.WithResource("clusterrolebindings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rbac().V1().ClusterRoleBindings().Informer()}, nil
	case rbacv1.SchemeGroupVersion.WithResource("roles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rbac().V1().Roles().Informer()}, nil
	case rbacv1.SchemeGroupVersion.WithResource("rolebindings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rbac().V1().RoleBindings().Informer()}, nil

		// Group=scheduling.k8s.io, Version=v1
	case schedulingv1.SchemeGroupVersion.WithResource("priorityclasses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Scheduling().V1().PriorityClasses().Informer()}, nil

		// Group=storage.k8s.io, Version=v1
	case storagev1.SchemeGroupVersion.WithResource("csidrivers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1().CSIDrivers().Informer()}, nil
	case storagev1.SchemeGroupVersion.WithResource("csinodes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1().CSINodes().Informer()}, nil
	case storagev1.SchemeGroupVersion.WithResource("csistoragecapacities"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1().CSIStorageCapacities().Informer()}, nil
	case storagev1.SchemeGroupVersion.WithResource("storageclasses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1().StorageClasses().Informer()}, nil
	case storagev1.SchemeGroupVersion.WithResource("volumeattachments"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1().VolumeAttachments().Informer()}, nil

		// Group=walrus.seal.io, Version=v1
	case walrusv1.SchemeGroupVersion.WithResource("catalogs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Walrus().V1().Catalogs().Informer()}, nil
	case walrusv1.SchemeGroupVersion.WithResource("connectors"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Walrus().V1().Connectors().Informer()}, nil
	case walrusv1.SchemeGroupVersion.WithResource("environments"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Walrus().V1().Environments().Informer()}, nil
	case walrusv1.SchemeGroupVersion.WithResource("projects"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Walrus().V1().Projects().Informer()}, nil
	case walrusv1.SchemeGroupVersion.WithResource("resources"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Walrus().V1().Resources().Informer()}, nil
	case walrusv1.SchemeGroupVersion.WithResource("resourcedefinitions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Walrus().V1().ResourceDefinitions().Informer()}, nil
	case walrusv1.SchemeGroupVersion.WithResource("resourceruns"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Walrus().V1().ResourceRuns().Informer()}, nil
	case walrusv1.SchemeGroupVersion.WithResource("schemas"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Walrus().V1().Schemas().Informer()}, nil
	case walrusv1.SchemeGroupVersion.WithResource("settings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Walrus().V1().Settings().Informer()}, nil
	case walrusv1.SchemeGroupVersion.WithResource("subjects"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Walrus().V1().Subjects().Informer()}, nil
	case walrusv1.SchemeGroupVersion.WithResource("subjectproviders"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Walrus().V1().SubjectProviders().Informer()}, nil
	case walrusv1.SchemeGroupVersion.WithResource("templates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Walrus().V1().Templates().Informer()}, nil
	case walrusv1.SchemeGroupVersion.WithResource("variables"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Walrus().V1().Variables().Informer()}, nil

		// Group=walruscore.seal.io, Version=v1
	case walruscorev1.SchemeGroupVersion.WithResource("catalogs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Walruscore().V1().Catalogs().Informer()}, nil
	case walruscorev1.SchemeGroupVersion.WithResource("connectors"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Walruscore().V1().Connectors().Informer()}, nil
	case walruscorev1.SchemeGroupVersion.WithResource("resources"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Walruscore().V1().Resources().Informer()}, nil
	case walruscorev1.SchemeGroupVersion.WithResource("resourcedefinitions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Walruscore().V1().ResourceDefinitions().Informer()}, nil
	case walruscorev1.SchemeGroupVersion.WithResource("resourceruns"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Walruscore().V1().ResourceRuns().Informer()}, nil
	case walruscorev1.SchemeGroupVersion.WithResource("schemas"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Walruscore().V1().Schemas().Informer()}, nil
	case walruscorev1.SchemeGroupVersion.WithResource("templates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Walruscore().V1().Templates().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
