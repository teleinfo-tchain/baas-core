package kubeclient

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"context"
)

// TODO 上下文判断需要明确
var ctx = context.TODO()

func (c *Clients) CreateDeployment(dep *appsv1.Deployment, opts metav1.CreateOptions) *appsv1.Deployment {
	if dep.Namespace == "" {
		dep.Namespace = corev1.NamespaceDefault
	}
	deploymentsClient := c.KubeClient.AppsV1().Deployments(dep.Namespace)
	newDep, err := deploymentsClient.Create(ctx, dep, opts)
	if err != nil {
		logger.Error(err.Error())
	}
	logger.Info("Created deployment %q \n", newDep.GetObjectMeta().GetName())
	return newDep
}

func (c *Clients) GetDeployment(dep *appsv1.Deployment, ops metav1.GetOptions) *appsv1.Deployment {
	deploymentsClient := c.KubeClient.AppsV1().Deployments(dep.Namespace)
	redep, err := deploymentsClient.Get(ctx, dep.Name, ops)
	if err != nil {
		logger.Error(err.Error())
	}
	logger.Info("Get deployment %q \n", dep.GetObjectMeta().GetName())
	return redep
}

func (c *Clients) GetDeploymentList(dep *appsv1.Deployment, ops metav1.ListOptions) *appsv1.DeploymentList {
	deploymentsClient := c.KubeClient.AppsV1().Deployments(dep.Namespace)
	list, err := deploymentsClient.List(ctx, ops)
	if err != nil {
		logger.Error(err.Error())
	}
	for _, d := range list.Items {
		logger.Info("Deployment ：", d.Name, d.Spec.Replicas)
	}
	return list
}

func (c *Clients) DeleteDeployment(dep *appsv1.Deployment, ops metav1.DeleteOptions) error {
	deploymentsClient := c.KubeClient.AppsV1().Deployments(dep.Namespace)
	err := deploymentsClient.Delete(ctx, dep.Name, ops)
	if err != nil {
		logger.Error(err.Error())
	}
	logger.Info("Delete deployment %q \n", dep.Name)
	return err
}

func (c *Clients) UpdateDeployment(dep *appsv1.Deployment, opts metav1.UpdateOptions) *appsv1.Deployment {
	deploymentsClient := c.KubeClient.AppsV1().Deployments(dep.Namespace)

	newDep, err := deploymentsClient.Update(ctx, dep, opts)
	if err != nil {
		logger.Error(err.Error())
	}

	logger.Info("Updated deployment %q \n", newDep.GetObjectMeta().GetName())
	return newDep
}
