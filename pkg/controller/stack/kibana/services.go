package kibana

import (
	deploymentsv1alpha1 "github.com/elastic/stack-operators/pkg/apis/deployments/v1alpha1"
	"github.com/elastic/stack-operators/pkg/controller/stack/common"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ServiceName(stackName string) string {
	return stackName + "-kb"
}

func NewService(s deploymentsv1alpha1.Stack) *corev1.Service {
	stackID := common.StackID(s)
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: s.Namespace,
			Name:      ServiceName(s.Name),
			Labels:    NewLabelsWithStackID(stackID),
		},
		Spec: corev1.ServiceSpec{
			Selector: NewLabelsWithStackID(stackID),
			Ports: []corev1.ServicePort{
				corev1.ServicePort{
					Protocol: corev1.ProtocolTCP,
					Port:     HTTPPort,
				},
			},
			SessionAffinity: corev1.ServiceAffinityNone,
			// For now, expose the service as node port to ease development
			// TODO: proper ingress forwarding
			Type: corev1.ServiceTypeNodePort,
			ExternalTrafficPolicy: corev1.ServiceExternalTrafficPolicyTypeCluster,
		},
	}

}