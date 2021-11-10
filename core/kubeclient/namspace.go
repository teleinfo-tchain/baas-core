package kubeclient

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Clients) GetNamespaceList(ops metav1.ListOptions) *corev1.NamespaceList {

	nss, err := c.KubeClient.CoreV1().Namespaces().List(ctx, ops)
	if err != nil {
		logger.Error(err.Error())
	}
	for _, ns := range nss.Items {
		logger.Info("Namespace：", ns.Name, ns.Status.Phase)
	}
	return nss
}

func (c *Clients) CreateNameSpace(ns *corev1.Namespace, ops metav1.CreateOptions) (*corev1.Namespace, error) {
	nameSpace, err := c.KubeClient.CoreV1().Namespaces().Create(ctx, ns, ops)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	logger.Info("Created namesapce %q \n", nameSpace.GetObjectMeta().GetName())
	return nameSpace, nil
}

func (c *Clients) DeleteNameSpace(ns *corev1.Namespace, ops metav1.DeleteOptions) error {
	err := c.KubeClient.CoreV1().Namespaces().Delete(ctx, ns.Name, ops)
	if err != nil {
		logger.Error(err.Error())
	}
	logger.Info("Delete namesapce %q \n", ns.GetObjectMeta().GetName())
	return err
}
