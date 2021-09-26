package kubeclient

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Clients) GetServiceList(ns string, ops metav1.ListOptions) *corev1.ServiceList {

	services, err := c.KubeClient.CoreV1().Services(ns).List(ctx, ops)
	if err != nil {
		logger.Error(err.Error())
	}
	for _, service := range services.Items {
		logger.Info("Service：", service.Name, service.GetUID())
	}
	return services
}

func (c *Clients) CreateService(service *corev1.Service, opts metav1.CreateOptions) *corev1.Service {
	newservice, err := c.KubeClient.CoreV1().Services(service.Namespace).Create(ctx, service, opts)
	if err != nil {
		logger.Error(err.Error())
	}
	logger.Info("Created Service %q \n", newservice.GetObjectMeta().GetName())
	return newservice
}

func (c *Clients) DeleteService(service *corev1.Service, ops metav1.DeleteOptions) error {
	err := c.KubeClient.CoreV1().Services(service.Namespace).Delete(ctx, service.Name, ops)
	if err != nil {
		logger.Error(err.Error())
	}
	logger.Info("Delete Service %q \n", service.GetObjectMeta().GetName())
	return err
}
