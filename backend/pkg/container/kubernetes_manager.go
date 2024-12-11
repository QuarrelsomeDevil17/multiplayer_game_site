package container

import (
	"context"
	"fmt"

	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// KubernetesManager manages Kubernetes pods and deployments
type KubernetesManager struct {
	Clientset *kubernetes.Clientset
}

// NewKubernetesManager initializes a KubernetesManager
func NewKubernetesManager() (*KubernetesManager, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		config, err = rest.InClusterConfig()
		if err != nil {
			return nil, fmt.Errorf("failed to initialize Kubernetes config: %v", err)
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kubernetes client: %v", err)
	}
	return &KubernetesManager{Clientset: clientset}, nil
}

// CreateDeployment creates a new deployment in Kubernetes
func (k *KubernetesManager) CreateDeployment(namespace, deploymentName, imageName string, replicas int32) error {
	deployment := &v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: deploymentName,
		},
		Spec: v1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"app": deploymentName},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{"app": deploymentName},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  deploymentName,
							Image: imageName,
						},
					},
				},
			},
		},
	}

	_, err := k.Clientset.AppsV1().Deployments(namespace).Create(context.Background(), deployment, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create deployment: %v", err)
	}
	return nil
}

// DeleteDeployment deletes a deployment in Kubernetes
func (k *KubernetesManager) DeleteDeployment(namespace, deploymentName string) error {
	err := k.Clientset.AppsV1().Deployments(namespace).Delete(context.Background(), deploymentName, metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete deployment: %v", err)
	}
	return nil
}

// ListPods lists all pods in a namespace
func (k *KubernetesManager) ListPods(namespace string) ([]corev1.Pod, error) {
	pods, err := k.Clientset.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to list pods: %v", err)
	}
	return pods.Items, nil
}
