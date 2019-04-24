package main

import (
	"fmt"
	"github.com/shinytang6/openpaas-security/client"

	svc "github.com/shinytang6/openpaas-security/service"
)

func main() {
	clientset := client.InitClient()
	//deploymentClient := clientset.AppsV1beta1().Deployments("default")
	//deployment := utils.GenerateDeployment("./examples/nginx.yaml")
	//deploy.CreateDeployment(deploymentClient, deployment)



	serviceClient := clientset.CoreV1().Services("default")
	svc, _ := svc.GetService(serviceClient, "vnc-ead-0")
	fmt.Println(svc.Spec.Ports[0].NodePort)
}

