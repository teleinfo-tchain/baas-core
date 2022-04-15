package kubeclient

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
	"strconv"
)

func (c *Clients) GetPodResourceByMetrics(configPath, namespace, podName string) (string, string, error) {
	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		return "", "", err
	}

	mc, err := metrics.NewForConfig(config)
	if err != nil {
		return "", "", err
	}
	logger.Info("catch mc")

	pod, err := mc.MetricsV1beta1().PodMetricses(namespace).Get(ctx, podName, metav1.GetOptions{})
	if err != nil {
		return "", "", err
	}

	var (
		cpuTotal    float64
		memoryTotal int64
	)
	for _, p := range pod.Containers {
		c, err := strconv.ParseFloat(p.Usage.Cpu().AsDec().String(), 64)
		if err != nil {
			logger.Error(err)
		}
		//统计cpu
		cpuTotal += float64(c)
		m, _ := p.Usage.Memory().AsInt64()
		//统计memory
		memoryTotal += m >> 20

	}
	cpu := fmt.Sprintf("%.5f", cpuTotal)
	memory := fmt.Sprintf("%d", memoryTotal)
	logger.Info("catch pod CPU", cpu)
	logger.Info("catch pod CPU", memory)
	return cpu, memory, err
}
