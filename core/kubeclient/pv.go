package kubeclient

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Clients) CreatePersistentVolume(pv *corev1.PersistentVolume, opts metav1.CreateOptions) *corev1.PersistentVolume {

	newpv, err := c.KubeClient.CoreV1().PersistentVolumes().Create(ctx, pv, opts)
	if err != nil {
		logger.Error(err.Error())
	}
	logger.Info("Created PersistentVolume %q \n", newpv.GetObjectMeta().GetName())
	return newpv
}

func (c *Clients) DeletePersistentVolume(pv *corev1.PersistentVolume, ops metav1.DeleteOptions) error {

	err := c.KubeClient.CoreV1().PersistentVolumes().Delete(ctx, pv.Name, ops)
	if err != nil {
		logger.Error(err.Error())
	}
	logger.Info("Delete PersistentVolume %q \n", pv.GetObjectMeta().GetName())
	return err
}
