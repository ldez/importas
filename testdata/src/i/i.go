package d

import (
	autoscalingv1alpha1 "knative.dev/serving/pkg/apis/autoscaling/v1alpha1"
	special "knative.dev/serving/pkg/apis/serving/v1"
)

func foo() {
	autoscalingv1alpha1.Resource("")
	special.Resource("")
}
