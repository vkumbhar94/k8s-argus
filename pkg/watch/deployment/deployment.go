// Package deployment provides the logic for mapping a Kubernetes deployment to a
// LogicMonitor w.
// nolint: dupl
package deployment

import (
	"fmt"
	"strconv"

	"github.com/logicmonitor/k8s-argus/pkg/constants"
	lmjaeger "github.com/logicmonitor/k8s-argus/pkg/jaeger"
	"github.com/logicmonitor/k8s-argus/pkg/lmctx"
	lmlog "github.com/logicmonitor/k8s-argus/pkg/log"
	"github.com/logicmonitor/k8s-argus/pkg/permission"
	"github.com/logicmonitor/k8s-argus/pkg/types"
	"github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
)

const (
	resource = "deployments"
)

// Watcher represents a watcher type that watches deployments.
type Watcher struct {
	types.DeviceManager
	*types.WConfig
}

// APIVersion is a function that implements the Watcher interface.
func (w *Watcher) APIVersion() string {
	return constants.K8sAPIVersionAppsV1
}

// Enabled is a function that check the resource can watch.
func (w *Watcher) Enabled() bool {
	return permission.HasDeploymentPermissions()
}

// Resource is a function that implements the Watcher interface.
func (w *Watcher) Resource() string {
	return resource
}

// ObjType is a function that implements the Watcher interface.
func (w *Watcher) ObjType() runtime.Object {
	return &appsv1.Deployment{}
}

// AddFunc is a function that implements the Watcher interface.
func (w *Watcher) AddFunc() func(obj interface{}) {
	return func(obj interface{}) {
		deployment := obj.(*appsv1.Deployment)
		lctx := lmlog.NewLMContextWith(logrus.WithFields(logrus.Fields{"device_id": resource + "-" + deployment.Name}))
		log := lmlog.Logger(lctx)
		span, resetter := lmjaeger.StartSpan(lctx, "newDeployment")
		defer span.Finish()
		defer resetter()
		//span.K8sObject("pod", pod)
		span.SetTag("name", deployment.Name)
		log.Infof("Handling add deployment event: %s", deployment.Name)
		w.add(lctx, deployment)
	}
}

// UpdateFunc is a function that implements the Watcher interface.
func (w *Watcher) UpdateFunc() func(oldObj, newObj interface{}) {
	return func(oldObj, newObj interface{}) {
		old := oldObj.(*appsv1.Deployment)
		new := newObj.(*appsv1.Deployment)

		lctx := lmlog.NewLMContextWith(logrus.WithFields(logrus.Fields{"device_id": resource + "-" + old.Name}))
		span, resetter := lmjaeger.StartSpan(lctx, "updateDeployment")
		defer resetter()
		defer span.Finish()
		//span.K8sObject("oldPod", old)
		//span.K8sObject("newPod", new)
		span.SetTag("name", old.Name)
		span.Info("event", old.Name+" replaced with "+new.Name)
		w.update(lctx, old, new)
	}
}

// DeleteFunc is a function that implements the Watcher interface.
func (w *Watcher) DeleteFunc() func(obj interface{}) {
	return func(obj interface{}) {
		deployment := obj.(*appsv1.Deployment)
		lctx := lmlog.NewLMContextWith(logrus.WithFields(logrus.Fields{"device_id": resource + "-" + deployment.Name}))
		log := lmlog.Logger(lctx)
		span, resetter := lmjaeger.StartSpan(lctx, "deleteDeployment")
		defer resetter()
		defer span.Finish()
		//span.K8sObject("pod", pod)
		span.SetTag("name", deployment.Name)
		log.Debugf("Handling delete deployment event: %s", deployment.Name)
		// Delete the deployment.
		if w.Config().DeleteDevices {
			span.SetTag("lm.action", "permanent delete")
			if err := w.DeleteByDisplayName(lctx, w.Resource(), fmtDeploymentDisplayName(deployment)); err != nil {
				span.Error("event", "Failed", "message", err.Error())

				log.Errorf("Failed to delete deployment: %v", err)
				return
			}
			span.Info("event", "success")
			log.Infof("Deleted deployment %s", deployment.Name)
			return
		}

		span.SetTag("lm.action", "move")
		// Move the deployment.
		w.move(lctx, deployment)
	}
}

// nolint: dupl
func (w *Watcher) add(lctx *lmctx.LMContext, deployment *appsv1.Deployment) {
	log := lmlog.Logger(lctx)
	span := lmjaeger.Span(lctx)
	d, err := w.Add(lctx, w.Resource(), deployment.Labels,
		w.args(deployment, constants.DeploymentCategory)...,
	)
	if err != nil {
		span.Error("event", "Failed", "message", err.Error())
		log.Errorf("Failed to add deployment %q: %v", fmtDeploymentDisplayName(deployment), err)
		return
	}
	span.Info("event", "created device")
	span.SetTag("lm.device.id", d.ID)
	log.Infof("Added deployment %q", fmtDeploymentDisplayName(deployment))
}

func (w *Watcher) update(lctx *lmctx.LMContext, old, new *appsv1.Deployment) {
	log := lmlog.Logger(lctx)
	span := lmjaeger.Span(lctx)
	d, err := w.UpdateAndReplaceByDisplayName(lctx, "deployments",
		fmtDeploymentDisplayName(old), nil, new.Labels,
		w.args(new, constants.DeploymentCategory)...,
	)
	if err != nil {
		span.Error("event", "Failed", "message", err.Error())
		log.Errorf("Failed to update deployment %q: %v", fmtDeploymentDisplayName(new), err)
		return
	}
	span.SetTag("lm.device.id", d.ID)
	log.Infof("Updated deployment %q", fmtDeploymentDisplayName(old))
}

// nolint: dupl
func (w *Watcher) move(lctx *lmctx.LMContext, deployment *appsv1.Deployment) {
	log := lmlog.Logger(lctx)
	span := lmjaeger.Span(lctx)
	d, err := w.UpdateAndReplaceFieldByDisplayName(lctx, w.Resource(), fmtDeploymentDisplayName(deployment), constants.CustomPropertiesFieldName, w.args(deployment, constants.DeploymentDeletedCategory)...)
	if err != nil {
		span.Error("event", "Failed to move", "message", err.Error())
		log.Errorf("Failed to move deployment %q: %v", fmtDeploymentDisplayName(deployment), err)
		return
	}
	span.SetTag("lm.device.id", d.ID)
	log.Infof("Moved deployment %q", fmtDeploymentDisplayName(deployment))
}

func (w *Watcher) args(deployment *appsv1.Deployment, category string) []types.DeviceOption {
	return []types.DeviceOption{
		w.Name(deployment.Name),
		w.ResourceLabels(deployment.Labels),
		w.DisplayName(fmtDeploymentDisplayName(deployment)),
		w.SystemCategories(category),
		w.Auto("name", deployment.Name),
		w.Auto("namespace", deployment.Namespace),
		w.Auto("selflink", deployment.SelfLink),
		w.Auto("uid", string(deployment.UID)),
		w.Custom(constants.K8sResourceCreatedOnPropertyKey, strconv.FormatInt(deployment.CreationTimestamp.Unix(), 10)),
		w.Custom(constants.K8sResourceNamePropertyKey, fmtDeploymentDisplayName(deployment)),
	}
}

// fmtDeploymentDisplayName implements the conversion for the deployment display name
func fmtDeploymentDisplayName(deployment *appsv1.Deployment) string {
	return fmt.Sprintf("%s.%s.deploy-%s", deployment.Name, deployment.Namespace, string(deployment.UID))
}

// GetDeploymentsMap implements the getting deployments map info from k8s
func GetDeploymentsMap(lctx *lmctx.LMContext, k8sClient kubernetes.Interface, namespace string) (map[string]string, error) {
	log := lmlog.Logger(lctx)
	deploymentsMap := make(map[string]string)
	deploymentList, err := k8sClient.AppsV1().Deployments(namespace).List(v1.ListOptions{})
	if err != nil || deploymentList == nil {
		log.Warnf("Failed to get the deployments from k8s")
		return nil, err
	}
	for i := range deploymentList.Items {
		deploymentsMap[fmtDeploymentDisplayName(&deploymentList.Items[i])] = deploymentList.Items[i].Name
	}

	return deploymentsMap, nil
}
