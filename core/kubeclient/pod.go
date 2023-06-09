package kubeclient

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"bytes"
	"io"
)

func (c *Clients) GetPodList(ns string, ops metav1.ListOptions) *corev1.PodList {

	pods, err := c.KubeClient.CoreV1().Pods(ns).List(ctx, ops)
	if err != nil {
		logger.Error(err.Error())
	}
	for _, pod := range pods.Items {
		logger.Info("Pod：", pod.Name, pod.Status.PodIP)
	}
	return pods
}

func (c *Clients) GetPodListByNodeName(nodeName string) *corev1.PodList {

	pods, err := c.KubeClient.CoreV1().Pods("").List(
		ctx, metav1.ListOptions{FieldSelector: "spec.nodeName=" + nodeName})
	if err != nil {
		logger.Error(err.Error())
	}
	for _, pod := range pods.Items {
		logger.Info("Pod：", pod.Name, pod.Status.PodIP)
	}
	return pods
}

func (c *Clients) CreatePod(pod *corev1.Pod, opts metav1.CreateOptions) *corev1.Pod {

	newPod, err := c.KubeClient.CoreV1().Pods(pod.Namespace).Create(ctx, pod, opts)
	if err != nil {
		logger.Error(err.Error())
	}
	logger.Info("Created pod %q \n", newPod.GetObjectMeta().GetName())
	return newPod
}

func (c *Clients) DeletePod(pod *corev1.Pod, ops metav1.DeleteOptions) {
	err := c.KubeClient.CoreV1().Pods(pod.Namespace).Delete(ctx, pod.Name, ops)
	if err != nil {
		logger.Error(err.Error())
	}
	logger.Info("Delete pod %q \n", pod.GetObjectMeta().GetName())
}

func (c *Clients) PrintPodLogs(pod corev1.Pod) {
	podLogOpts := corev1.PodLogOptions{}

	req := c.KubeClient.CoreV1().Pods(pod.Namespace).GetLogs(pod.Name, &podLogOpts)
	podLogs, err := req.Stream(ctx)
	if err != nil {
		logger.Error("error in opening stream")
	}
	if podLogs == nil {
		logger.Error("error in opening stream")
	}
	defer podLogs.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		logger.Error("error in copy information from podLogs to buf")
	}
	str := buf.String()

	logger.Info("Pod loggers :", str)
}
