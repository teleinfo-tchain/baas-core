package kubeclient

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Clients) CreateStatefulSet(sfs *appsv1.StatefulSet, opts metav1.CreateOptions) (*appsv1.StatefulSet, error) {
	if sfs.Namespace == "" {
		sfs.Namespace = corev1.NamespaceDefault
	}
	sfsClient := c.KubeClient.AppsV1().StatefulSets(sfs.Namespace)
	newSfs, err := sfsClient.Create(ctx, sfs, opts)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	logger.Info("Created deployment %q \n", newSfs.GetObjectMeta().GetName())
	return newSfs, nil
}

func (c *Clients) DeleteStatefulSet(sfs *appsv1.StatefulSet, ops metav1.DeleteOptions) error {
	if sfs.Namespace == "" {
		sfs.Namespace = corev1.NamespaceDefault
	}
	sfsClient := c.KubeClient.AppsV1().StatefulSets(sfs.Namespace)
	err := sfsClient.Delete(ctx, sfs.Name, ops)
	if err != nil {
		logger.Error(err.Error())
	}
	logger.Info("Created deployment %q \n", sfs.GetObjectMeta().GetName())
	return err
}
