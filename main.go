package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/shinytang6/openpaas-security/client"
	deploy "github.com/shinytang6/openpaas-security/deployment"
	svc "github.com/shinytang6/openpaas-security/service"
	apps_v1beta1 "k8s.io/api/apps/v1beta1"
	core_v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
)

func main() {
	var (
		deployYaml []byte
		deployJson []byte
		deployment = apps_v1beta1.Deployment{}
		serviceYaml []byte
		serviceJson []byte
		service = core_v1.Service{}
		err error
	)
	clientset := client.InitClient()
	deploymentClient := clientset.AppsV1beta1().Deployments("default")
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

	deploy.CreateDeployment(deploymentClient, deployment)



	serviceClient := clientset.CoreV1().Services("default")
	if serviceYaml, err = ioutil.ReadFile("./nginx_svc.yaml"); err != nil {
		panic(err.Error())
	}

	if serviceJson, err = yaml.ToJSON(serviceYaml); err != nil {
		panic(err.Error())
	}
	// JSON转struct
	if err = json.Unmarshal(serviceJson, &service); err != nil {
		panic(err.Error())
	}

	svc.CreateService(serviceClient, service)
}