package deployment

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	apps_v1beta1 "k8s.io/api/apps/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateDeployment(deploymentClient v1beta1.DeploymentInterface)  {
	var (
		deployYaml []byte
		deployJson []byte
		deployment = apps_v1beta1.Deployment{}
		err error
	)
	if deployYaml, err = ioutil.ReadFile("./nginx.yaml"); err != nil {
		panic(err.Error())
	}

	if deployJson, err = yaml.ToJSON(deployYaml); err != nil {
		panic(err.Error())
	}
	// JSON转struct
	if err = json.Unmarshal(deployJson, &deployment); err != nil {
		panic(err.Error())
	}

	if _, err = GetDeployment(deploymentClient, deployment.Name); err != nil {
		if !errors.IsNotFound(err) {
			panic(err.Error())
			return
		}
		fmt.Println("Start creating deployment...")
		result, err := deploymentClient.Create(&deployment)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("Deployment %s created!\n", result.GetObjectMeta().GetName())
	} else {
		fmt.Printf("deployments.apps %s already exists\n", deployment.Name)
		fmt.Println("Start updating deployment...")
		UpdateDeployment(deploymentClient, deployment)
		fmt.Printf("Deployment %s updated!\n", deployment.Name)
	}
}

func GetDeployment(deploymentClient v1beta1.DeploymentInterface, deploymentName string) (*apps_v1beta1.Deployment, error) {
	var (
		deployment  = &apps_v1beta1.Deployment{}
		err error
	)
	deployment, err = deploymentClient.Get(deploymentName, meta_v1.GetOptions{})
	return deployment, err
}

func UpdateDeployment(deploymentClient v1beta1.DeploymentInterface, deployment apps_v1beta1.Deployment) (*apps_v1beta1.Deployment, error) {
	var (
		new_deployment  = &apps_v1beta1.Deployment{}
		err error
	)
	new_deployment, err = deploymentClient.Update(&deployment)
	return new_deployment, err
}