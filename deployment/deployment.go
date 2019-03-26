package deployment

import (
	"fmt"

	apps_v1beta1 "k8s.io/api/apps/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/typed/apps/v1beta1"
)

func CreateDeployment(deploymentClient v1beta1.DeploymentInterface, deployment apps_v1beta1.Deployment)  {
	if _, err := GetDeployment(deploymentClient, deployment.Name); err != nil {
		if !errors.IsNotFound(err) {
			panic(err.Error())
			return
		}
		fmt.Printf("Start creating deployment...\n")
		result, err := deploymentClient.Create(&deployment)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("Deployment %s created!\n", result.GetObjectMeta().GetName())
	} else {
		UpdateDeployment(deploymentClient, deployment)
	}
}

func GetDeployment(deploymentClient v1beta1.DeploymentInterface, deploymentName string) (*apps_v1beta1.Deployment, error) {
	deployment, err := deploymentClient.Get(deploymentName, meta_v1.GetOptions{})
	return deployment, err
}

func UpdateDeployment(deploymentClient v1beta1.DeploymentInterface, deployment apps_v1beta1.Deployment) (*apps_v1beta1.Deployment, error) {
	fmt.Printf("Start updating deployment...\n")
	new_deployment, err := deploymentClient.Update(&deployment)
	fmt.Printf("Deployment %s updated!\n", deployment.Name)
	return new_deployment, err
}

func DeleteDeployment(deploymentClient v1beta1.DeploymentInterface, deploymentName string) error {
	fmt.Printf("Start deleting deployment %s...\n", deploymentName)
	err := deploymentClient.Delete(deploymentName, &meta_v1.DeleteOptions{})
	return err
}