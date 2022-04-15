package kubeclient

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/api/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Clients) CreateIngress(ingress *v1beta1.Ingress, opts metav1.CreateOptions) (*v1beta1.Ingress, error) {

	if ingress.Namespace == "" {
		ingress.Namespace = corev1.NamespaceDefault
	}
	ingClient := c.KubeClient.NetworkingV1beta1().Ingresses(ingress.Namespace)

	newIng, err := ingClient.Create(ctx, ingress, opts)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	logger.Info("Created Ingress %q \n", newIng.GetObjectMeta().GetName())
	return newIng, nil
}

func (c *Clients) DeleteIngress(ingress *v1beta1.Ingress, ops metav1.DeleteOptions) error {
	if ingress.Namespace == "" {
		ingress.Namespace = corev1.NamespaceDefault
	}
	ingClient := c.KubeClient.ExtensionsV1beta1().Ingresses(ingress.Namespace)
	err := ingClient.Delete(ctx, ingress.Name, ops)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	logger.Info("Created deployment %q \n", ingress.GetObjectMeta().GetName())
	return err
}

func (c *Clients) GetIngressList(ns string, ops metav1.ListOptions) *v1beta1.IngressList {
	list, err := c.KubeClient.NetworkingV1beta1().Ingresses(ns).List(ctx, ops)
	if err != nil {
		logger.Error(err.Error())
	}
	for _, d := range list.Items {
		logger.Info("Ingress ：", d.Name, d.Spec.Rules)
	}
	return list
}

func (c *Clients) GetIngress(ingress *v1beta1.Ingress, ops metav1.GetOptions) (*v1beta1.Ingress, error) {
	ingClient := c.KubeClient.NetworkingV1beta1().Ingresses(ingress.Namespace)
	ing, err := ingClient.Get(ctx, ingress.Name, ops)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	logger.Info("Get deployment %q \n", ing.GetObjectMeta().GetName())
	return ing, nil
}

func (c *Clients) UpdateIngress(ingress *v1beta1.Ingress, ops metav1.UpdateOptions) (*v1beta1.Ingress, error) {
	ingClient := c.KubeClient.NetworkingV1beta1().Ingresses(ingress.Namespace)
	nweIng, err := ingClient.Update(ctx, ingress, ops)
	if err != nil {
		logger.Error(err.Error())
	}

	logger.Info("Updated deployment %q \n", nweIng.GetObjectMeta().GetName())
	return nweIng, err
}
