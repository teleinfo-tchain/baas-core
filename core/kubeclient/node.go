package kubeclient

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Clients) GetNodeList(ops metav1.ListOptions) *corev1.NodeList {

	ns, err := c.KubeClient.CoreV1().Nodes().List(ctx, ops)
	if err != nil {
		logger.Error(err.Error())
	}
	for _, n := range ns.Items {
		logger.Info("Node：", n.Name, n.Status.Addresses)
	}
	return ns
}
